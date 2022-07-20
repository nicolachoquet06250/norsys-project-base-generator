package technos

import (
	"fmt"
	. "test_go_webserver/helpers"
)

const (
	JavaScript = "JavaScript Vanilla"

	React15 = "JavaScript React v15"
	React16 = "JavaScript React v16"
	React17 = "JavaScript React v17"
	React18 = "JavaScript React v18"

	Vue2 = "JavaScript Vue v2"
	Vue3 = "JavaScript Vue v3"

	TypeScript = "TypeScript Vanilla"
	Angular    = "TypeScript Angular"

	PHP     = "PHP From Scratch"
	Laravel = "PHP Laravel"
	Symfony = "PHP Symfony"

	Go = "Go"

	Java   = "Java From Scratch"
	Spring = "Java Spring"
)

func IsTechno(techno string) bool {
	exists, _ := InArray(techno, []string{
		JavaScript,
		React15, React16, React17, React18,
		Vue2, Vue3,
		TypeScript, Angular,
		PHP, Laravel, Symfony,
		Go,
		Java, Spring,
	})

	return exists
}

type Techno struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func All() []Techno {
	return []Techno{
		{
			Name:  JavaScript,
			Value: "javascript",
		},

		{
			Name:  React15,
			Value: "react-15",
		},
		{
			Name:  React16,
			Value: "react-16",
		},
		{
			Name:  React17,
			Value: "react-17",
		},
		{
			Name:  React18,
			Value: "react-18",
		},

		{
			Name:  Vue2,
			Value: "vue-2",
		},
		{
			Name:  Vue3,
			Value: "vue-3",
		},

		{
			Name:  TypeScript,
			Value: "typescript",
		},
		{
			Name:  Angular,
			Value: "angular",
		},

		{
			Name:  PHP,
			Value: "php-from-scratch",
		},
		{
			Name:  Laravel,
			Value: "laravel",
		},
		{
			Name:  Symfony,
			Value: "symfony",
		},

		{
			Name:  Go,
			Value: "go",
		},

		{
			Name:  Java,
			Value: "java",
		},
		{
			Name:  Spring,
			Value: "spring",
		},
	}
}

func FromValue(value string) (Techno, error) {
	for _, techno := range All() {
		if techno.Value == value {
			return techno, nil
		}
	}

	return Techno{}, fmt.Errorf("Aucune techno du nom de '%s' n'a été trouvée", value)
}
