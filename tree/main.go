package tree

type ITreeElement interface {
	IsDir() bool
	IsFile() bool
}

type TreeElement struct {
	ITreeElement

	Type     string
	Name     string
	Children []TreeElement
}

func (e TreeElement) IsFile() bool {
	return e.Type == "File"
}

func (e TreeElement) IsDir() bool {
	return e.Type == "Dir"
}

func File(name string) TreeElement {
	return TreeElement{
		Type: "File",
		Name: name,
	}
}

func Dir(name string, children []TreeElement) TreeElement {
	return TreeElement{
		Type:     "Dir",
		Name:     name,
		Children: children,
	}
}
