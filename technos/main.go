package technos

import "reflect"

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

func inArray(val interface{}, array interface{}) (exists bool /*, index int*/) {
	exists = false
	//index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				//index = i
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
