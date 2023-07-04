package puppy

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
