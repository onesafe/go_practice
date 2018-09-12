package main

import (
	"github.com/onesafe/go_practice/oop/employee"
)

func main() {
	e := employee.New("bran", "wang", 300, 30)
	e.LeavesRemaining()
}
