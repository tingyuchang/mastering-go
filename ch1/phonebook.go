package ch1

import "fmt"

type Entity struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entity{}

// init initials default data
func init() {
	data = append(data, Entity{
		Name:    "Matt",
		Surname: "Chang",
		Tel:     "12345",
	})
	data = append(data, Entity{"Nancy", "Docy", "0987234"})
	data = append(data, Entity{"Tom", "Jackson", "2356803"})
	data = append(data, Entity{"Jeff", "Hamilton", "34683"})
	data = append(data, Entity{"Lilly", "Back", "345331"})
	data = append(data, Entity{"George", "Carloze", "742367"})
	data = append(data, Entity{"Mike", "Yang", "95w211"})
}

// Search does linear search now, if the data become bigger & bigger, change to other ways
func Search(surname string) *Entity {
	for _, v := range data {
		if v.Surname == surname {
			// check the difference between return &v and &data[i]
			return &v
		}
	}
	return nil
}

// List just shows all data
func List() {
	for _, v := range data {
		fmt.Println(v)
	}
}
