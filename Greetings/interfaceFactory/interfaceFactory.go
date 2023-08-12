package interfaceFactory

import "fmt"

type animal interface {
	sound() string
}

type dog struct{}

func (d *dog) sound() string {
	return "Bark"
}

type cat struct{}

func (c *cat) sound() string {
	return "Meow"
}

func GetSound(animals []string) {
	fmt.Println("=============================================================")
	fmt.Println("Factory")
	fmt.Println()

	for _, ani := range animals {
		var a animal
		var err error
		if a, err = getAnimal(ani); err != nil {
			fmt.Printf("%+v \n", err)
		} else {
			fmt.Println(a.sound())
		}
	}

	fmt.Println()
}

func getAnimal(s string) (animal, error) {
	switch s {
	case "Dog":
		return &dog{}, nil
	case "Cat":
		return &cat{}, nil
	}

	return nil, fmt.Errorf("animal: %s not defined", s)
}
