package technos

import (
	"fmt"
	"npbg/configFiles"
	. "npbg/helpers"
	"npbg/technos/enum"
)

func IsTechno(techno string) bool {
	exists, _ := InArray(techno, []string{
		enum.JavaScript,
		enum.React15, enum.React16, enum.React17, enum.React18,
		enum.Vue2, enum.Vue3,
		enum.TypeScript, enum.Angular,
		enum.PHP, enum.Laravel, enum.Symfony,
		enum.Go,
		enum.Java, enum.Spring,
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
			Name:  enum.JavaScript,
			Value: "javascript",
		},

		{
			Name:  enum.React15,
			Value: "react-15",
		},
		{
			Name:  enum.React16,
			Value: "react-16",
		},
		{
			Name:  enum.React17,
			Value: "react-17",
		},
		{
			Name:  enum.React18,
			Value: "react-18",
		},

		{
			Name:  enum.Vue2,
			Value: "vue-2",
		},
		{
			Name:  enum.Vue3,
			Value: "vue-3",
		},

		{
			Name:  enum.TypeScript,
			Value: "typescript",
		},
		{
			Name:  enum.Angular,
			Value: "angular",
		},

		{
			Name:  enum.PHP,
			Value: "php-from-scratch",
		},
		{
			Name:  enum.Laravel,
			Value: "laravel",
		},
		{
			Name:  enum.Symfony,
			Value: "symfony",
		},

		{
			Name:  enum.Go,
			Value: "go",
		},

		{
			Name:  enum.Java,
			Value: "java",
		},
		{
			Name:  enum.Spring,
			Value: "spring",
		},
	}
}

func AllAvailable() (t []Techno) {
	all := All()

	for k, _ := range configFiles.ConfigFiles {
		present := false
		var n int

		for i, e := range all {
			if e.Name == k {
				present = true
				n = i
				break
			}
		}
		if present {
			t = append(t, all[n])
		}
	}

	return t
}

func FromValue(value string) (Techno, error) {
	for _, techno := range All() {
		if techno.Value == value {
			return techno, nil
		}
	}

	return Techno{}, fmt.Errorf("Aucune techno du nom de '%s' n'a été trouvée", value)
}
