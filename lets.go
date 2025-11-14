package main

import (
	"fmt"
	"math/rand"
)

// Define an interface
type Speaker interface {
	Speak() string
}

// Any type with a Speak() method automatically implements Speaker
type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Human struct{}

func (h Human) Speak() string { return "Hello!" }

// This function works with ANY Speaker
func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	fmt.Println("Hello, VS Code and Go!")

	// Create instances and pass them to MakeSound
	dog := Dog{}
	human := Human{}

	// 	pulls = ["anita max wynn", "max win please", "unreal", "shameless", "don't get it twisted", "twisted", "unbelievable", "...", "we're so due", "cmon bitch", "let us in bitch", "bad seed", "please", "BRO", "heart heart heart", "HIT HIT HTI", "retrig man",  "just do it", "just get us in", "please bro", "fr", "wtf", "cmooon", "please", "no dude", "just lock in", "rolling", "pls", "next one"]
	pulls := []string{"anita max wynn", "max win please", "unreal", "shameless", "don't get it twisted", "twisted", "unbelievable", "...", "we're so due", "cmon bitch", "let us in bitch", "bad seed", "please", "BRO", "heart heart heart", "HIT HIT HTI", "retrig man", "just do it", "just get us in", "please bro", "fr", "wtf", "cmooon", "please", "no dude", "just lock in", "rolling", "pls", "next one"}

	// creating and assigning values to an array without var keyword
	// studentsAge := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	MakeSound(dog)
	MakeSound(human)

	r := rand.Perm(10)
	randInt := r[0]
	fmt.Println(pulls[randInt])

	// The tabwriter here helps us generate aligned output.
	// looks like some aids wy of formatting output using fprint (maybe fString equivalent)
	// w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	// defer w.Flush()
	// show := func(name string, v1, v2, v3 any) {
	// 	fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	// }

	// Float32 and Float64 values are in [0, 1).
	// show("Float32", r.Float32(), r.Float32(), r.Float32())
	// show("Float64", r.Float64(), r.Float64(), r.Float64())

}
