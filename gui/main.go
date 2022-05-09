package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/chzyer/readline"
	_color "github.com/gookit/color"
	"go.pkg/nchoquet/config_files"
	"go.pkg/nchoquet/files"
	. "go.pkg/nchoquet/helpers"
	"go.pkg/nchoquet/technos"
	"image/color"
	"log"
)

type FormItemArray = []*widget.FormItem

type Data struct {
	techno      string
	projectPath string
}

func NewData() Data {
	return Data{}
}

func buildButtons(cb func(_ string)) FormItemArray {
	return FormItemArray{
		{Widget: widget.NewButton(technos.JavaScript, func() {
			log.Println(technos.JavaScript)
			cb(technos.JavaScript)
		})},
		{Widget: widget.NewButton(technos.React15, func() {
			log.Println(technos.React15)
			cb(technos.React15)
		})},
		{Widget: widget.NewButton(technos.React16, func() {
			log.Println(technos.React16)
			cb(technos.React16)
		})},
		{Widget: widget.NewButton(technos.React17, func() {
			log.Println(technos.React17)
			cb(technos.React17)
		})},
		{Widget: widget.NewButton(technos.React18, func() {
			log.Println(technos.React18)
			cb(technos.React18)
		})},
		{Widget: widget.NewButton(technos.Vue2, func() {
			log.Println(technos.Vue2)
			cb(technos.Vue2)
		})},
		{Widget: widget.NewButton(technos.Vue3, func() {
			log.Println(technos.Vue3)
			cb(technos.Vue3)
		})},
		{Widget: widget.NewButton(technos.TypeScript, func() {
			log.Println(technos.TypeScript)
			cb(technos.TypeScript)
		})},
		{Widget: widget.NewButton(technos.Angular, func() {
			log.Println(technos.Angular)
			cb(technos.Angular)
		})},
		{Widget: widget.NewButton(technos.PHP, func() {
			log.Println(technos.PHP)
			cb(technos.PHP)
		})},
		{Widget: widget.NewButton(technos.Laravel, func() {
			log.Println(technos.Laravel)
			cb(technos.Laravel)
		})},
		{Widget: widget.NewButton(technos.Symfony, func() {
			log.Println(technos.Symfony)
			cb(technos.Symfony)
		})},
		{Widget: widget.NewButton(technos.Go, func() {
			log.Println(technos.Go)
			cb(technos.Go)
		})},
		{Widget: widget.NewButton(technos.Java, func() {
			log.Println(technos.Java)
			cb(technos.Java)
		})},
		{Widget: widget.NewButton(technos.Spring, func() {
			log.Println(technos.Spring)
			cb(technos.Spring)
		})},
	}
}

func createTechnologiesTab(onValid func(techno string)) *container.TabItem {
	buttons := buildButtons(onValid)

	content := &widget.Form{
		Items: buttons,
	}

	return container.NewTabItem("Technologie", content)
}

func createProjectPathTab(onValid func(projectPath string)) *container.TabItem {
	input := widget.NewEntry()
	input.SetPlaceHolder("ex: C:\\Users\\workspaces... || /home/workspaces...")

	button := widget.NewButton("Validate", func() {
		onValid(input.Text)
	})

	content := &widget.Form{
		Items: FormItemArray{
			{Widget: input, HintText: "Chemin du projet"},
			{Widget: button},
		},
	}

	return container.NewTabItem("Chemin du projet", content)
}

func createAlertWindow(text string, error bool, a *fyne.App) (w fyne.Window, err error) {
	var title = "Success"
	if error {
		title = "Error"
	}
	w = (*a).NewWindow(title)
	icon, err := fyne.LoadResourceFromPath("./assets/images/logo.png")
	if err != nil {
		println(fmt.Errorf("an error occured when load logo"))
		return w, fmt.Errorf("an error occured when load logo")
	}

	var _color = color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 1,
	}
	if error {
		_color = color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 1,
		}
	}

	w.SetIcon(icon)
	w.SetFixedSize(true)
	w.SetContent(canvas.NewText(text, _color))

	return w, nil
}

func main() {
	data := NewData()

	a := app.New()
	window := a.NewWindow("Norsys | Project Base Generator")
	icon, err := fyne.LoadResourceFromPath("./assets/images/logo.png")
	if err != nil {
		println(fmt.Errorf("an error occured when load logo"))
		println("")
		return
	}

	window.SetIcon(icon)
	window.Resize(fyne.Size{
		Width:  600,
		Height: 300,
	})
	window.SetFixedSize(false)

	tabs := container.NewAppTabs()
	projectPathTab := createProjectPathTab(func(projectPath string) {
		if projectPath != "" {
			data.projectPath = projectPath

			var w fyne.Window
			var wErr, fErr error

			for path, content := range config_files.ConfigFiles[data.techno] {
				fErr = files.Create(data.projectPath+Slash()+path, content)
				if fErr != nil {
					w, wErr = createAlertWindow(fErr.Error(), true, &a)
					if wErr != nil {
						_color.Red.Println(wErr.Error())
					} else {
						w.Show()
					}
					_color.Red.Println(fErr.Error())
					break
				}
			}

			if wErr == nil && fErr == nil {
				w, wErr = createAlertWindow(
					"le projet "+data.techno+" à bien été généré dans le répertoire "+data.projectPath+" !",
					false,
					&a,
				)

				if wErr != nil {
					_color.Red.Println(wErr.Error())
				} else {
					w.Show()
				}
			}
		}
	})
	technosTab := createTechnologiesTab(func(techno string) {
		if techno != "" {
			data.techno = techno
			tabs.Select(projectPathTab)
		}
	})

	tabs.Append(technosTab)
	tabs.Append(projectPathTab)

	tabs.SetTabLocation(container.TabLocationTop)

	window.SetContent(tabs)
	window.ShowAndRun()

	readline.New("Tap on touch to quit")
	// C:\Users\Nicolas\Documents\workspaces\golang-workspace\project-base-generator\test\test2\test3\test4
}

/*package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window

func makeBasicControlsPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	hbox.Append(ui.NewButton("Button"), false)
	hbox.Append(ui.NewCheckbox("Checkbox"), false)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("Entries")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	entryForm.Append("Entry", ui.NewEntry(), false)
	entryForm.Append("Password Entry", ui.NewPasswordEntry(), false)
	entryForm.Append("Search Entry", ui.NewSearchEntry(), false)
	entryForm.Append("Multiline Entry", ui.NewMultilineEntry(), true)
	entryForm.Append("Multiline Entry No Wrap", ui.NewNonWrappingMultilineEntry(), true)

	return vbox
}

func makeNumbersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)
	pbar := ui.NewProgressBar()
	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(spinbox, false)
	vbox.Append(slider, false)
	vbox.Append(pbar, false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	vbox.Append(ip, false)

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	vbox.Append(cbox, false)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("Editable Item 1")
	ecbox.Append("Editable Item 2")
	ecbox.Append("Editable Item 3")
	vbox.Append(ecbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}

func makeDataChoosersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewDatePicker(), false)
	vbox.Append(ui.NewTimePicker(), false)
	vbox.Append(ui.NewDateTimePicker(), false)
	vbox.Append(ui.NewFontButton(), false)
	vbox.Append(ui.NewColorButton(), false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	button = ui.NewButton("Save File")
	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry2.SetText(filename)
	})
	grid.Append(button,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry2,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	msggrid := ui.NewGrid()
	msggrid.SetPadded(true)
	grid.Append(msggrid,
		0, 2, 2, 1,
		false, ui.AlignCenter, false, ui.AlignStart)

	button = ui.NewButton("Message Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBox(mainwin,
			"This is a normal message box.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	button = ui.NewButton("Error Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(mainwin,
			"This message box describes an error.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

func setupUI() {
	mainwin = ui.NewWindow("libui Control Gallery", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("Basic Controls", makeBasicControlsPage())
	tab.SetMargined(0, true)

	tab.Append("Numbers and Lists", makeNumbersPage())
	tab.SetMargined(1, true)

	tab.Append("Data Choosers", makeDataChoosersPage())
	tab.SetMargined(2, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}*/
