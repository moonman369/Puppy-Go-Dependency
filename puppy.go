package puppy

import dog "github.com/moonman369/Dog-Go-Depedency"

var PuppyName string = "Doge"
var PuppyAge int8 = 0

func SetPuppyName(newName string) {
	PuppyName = newName
}

func SetPuppyAge(newAge int8) {
	PuppyAge = newAge
}

func Bark() string {
	return "Woof!"
}

func Barks(n int) string {
	i, res := 0, ""
	for i < n {
		res += Bark() + " "
		i++
	}
	return res
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks(n int) string {
	return dog.WhenGrownUp(Barks(n))
}
