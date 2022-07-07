package technos

import (
	"fmt"
	"reflect"
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

func inArray(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}

	return
}

func IsTechno(techno string) bool {
	return inArray(techno, []string{
		JavaScript,
		React15, React16, React17, React18,
		Vue2, Vue3,
		TypeScript, Angular,
		PHP, Laravel, Symfony,
		Go,
		Java, Spring,
	})
}

type Techno struct {
	Name  string
	Value string
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
