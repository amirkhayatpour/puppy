package puppy

import (
	"fmt"
	"github.com/amirkhayatpour/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowsUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowsUp(Barks())
}

func VersionOne() {
	fmt.Println("puppy package version: v1.0.0")
}

func VersionOneTwo() {
	fmt.Println("puppy package version: v1.2.0")
}

func VersionOneThree() {
	fmt.Println("puppy package version: v1.3.0")
}
