package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

type post struct {
	title   string
	content string
	author
}

type website struct {
	posts []post
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (p post) details() {
	fmt.Println("Titile: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}

	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}

	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}

/*
OUTPUT:
Contents of Website
Titile:  Inheritance in Go
Content:  Go supports composition instead of inheritance
Author:  Naveen Ramanathan
Bio:  Golang Enthusiast

Titile:  Struct instead of Classes in Go
Content:  Go does not support classes but methods can be added to structs
Author:  Naveen Ramanathan
Bio:  Golang Enthusiast

Titile:  Concurrency
Content:  Go is a concurrent language and not a parallel one
Author:  Naveen Ramanathan
Bio:  Golang Enthusiast
*/
