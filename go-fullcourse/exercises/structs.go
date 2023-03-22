package exercises

import "fmt"

type Doctor struct {
	number int
	actorName string
	episodes []string
	companions []string
}

func Struct_Declaration() (Doctor) {
	
	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string {"Liz Shaw", "Jo Grant", "Sarah Jane Smith"},
	}
	// or

	bDoctor := struct{name string}{name: "John Pertwee"}
	fmt.Println(bDoctor)

	return aDoctor
}

func Struct_Manipulating() (int, string, []string){
	// The other declaration form of a Struct is using "Positional Syntax", so, it isn't necessary to pass the name before the value like above.
	aDoctor := Doctor {
		number: 3,
		actorName: "Leonardo DiCaprio",
		companions: []string {"Liz Shaw", "Jo Grant", "Sarah Jane Smith"},
	}
	// This format isn't recommended, also is possible to complete ignore some fields of structs, but not recommended.

	fmt.Println(aDoctor.companions[1]) // It's possible to get the value of a specific index in a Slice just as before.

	fmt.Println()
	fmt.Println("Name transformation")
	fmt.Println()

	bDoctor := struct{name string}{name: "Victor Teixeira"}
	cDoctor := bDoctor
	cDoctor.name = "Tom Baker"

	fmt.Println(bDoctor)
	fmt.Println(cDoctor)

	return aDoctor.number, aDoctor.actorName, aDoctor.companions
}

func Struct_Composition() {
	type Animal struct {
		Name string
		Origin string
	}

	type Bird struct {
		Animal
		SpeedKPH float32
		CanFly bool
	}

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	
	// or

	c := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 30,
		CanFly: false,
	}

	fmt.Println(b)
	fmt.Println()
	fmt.Println(c)
}