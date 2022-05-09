package autocomplete

import (
	"github.com/chzyer/readline"
	"github.com/gookit/color"
	. "go.pkg/nchoquet/helpers"
	. "go.pkg/nchoquet/tree"
	"io/ioutil"
	"os"
)

func GetPrompt(prompt ...string) string {
	var _prompt = "»"
	if len(prompt) > 0 {
		_prompt = prompt[0]
	}

	return _prompt
}

func CreateAutoCompleteLineReader(completer *readline.PrefixCompleter, prompt ...string) *readline.Instance {
	var _prompt = GetPrompt(prompt...)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:          color.Red.Render(_prompt + " "),
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		FuncFilterInputRune: func(r rune) (rune, bool) {
			switch r {
			// block CtrlZ feature
			case readline.CharCtrlZ:
				return r, false
			}
			return r, true
		},
	})
	if err != nil {
		panic(err)
	}

	return rl
}

func listFiles(path string) func(string) []string {
	return func(line string) []string {
		names := make([]string, 0)
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			names = append(names, f.Name())
		}
		return names
	}
}

func buildDynamicTree(path string) []TreeElement {
	var localTree []TreeElement

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			localTree = append(localTree, Dir(f.Name(), buildDynamicTree(path+Slash()+f.Name())))
		} else {
			localTree = append(localTree, File(f.Name()))
		}
	}

	return localTree
}

func buildPcItemDynamic(tree []TreeElement) []readline.PrefixCompleterInterface {
	var finalTree []readline.PrefixCompleterInterface

	for _, e := range tree {
		if e.IsDir() {
			finalTree = append(
				finalTree,
				readline.PcItem(
					e.Name,
					readline.PcItem(
						Slash(),
						buildPcItemDynamic(e.Children)...,
					),
				),
			)
		}
	}

	return finalTree
}

func buildDynamicPcItemR(path string, skip bool) *readline.PrefixCompleter {
	if skip {
		return readline.PcItem(
			Slash(),
			buildPcItemDynamic(buildDynamicTree(path))...,
		)
	}

	return readline.PcItem(path,
		readline.PcItem(
			Slash(),
			buildPcItemDynamic(buildDynamicTree(path))...,
		),
	)
}

func pathCompleter() *readline.PrefixCompleter {
	var pwd, _ = os.Getwd()

	return readline.NewPrefixCompleter(
		readline.PcItem(PwdVar(), buildDynamicPcItemR(pwd, true)),
		buildDynamicPcItemR(pwd, false),
	)
}

var PathCompleter = pathCompleter()

var TechnoCompleter = readline.NewPrefixCompleter(
	readline.PcItem("JavaScript",
		readline.PcItem("Vanilla"),
		readline.PcItem("React",
			readline.PcItem("v15"),
			readline.PcItem("v16"),
			readline.PcItem("v17"),
			readline.PcItem("v18"),
		),
		readline.PcItem("Vue",
			readline.PcItem("v2"),
			readline.PcItem("v3"),
		),
	),
	readline.PcItem("TypeScript",
		readline.PcItem("Angular"),
	),
	readline.PcItem("PHP",
		readline.PcItem("From Scratch"),
		readline.PcItem("Symfony"),
		readline.PcItem("Laravel"),
	),
	readline.PcItem("Go"),
	readline.PcItem("Java",
		readline.PcItem("From Scratch"),
		readline.PcItem("Spring"),
	),
)
