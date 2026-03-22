package puppy

import "github.com/ParuyrBaghyan/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarkss() string {
	return dog.WhenGrowUp(Barks())
}
