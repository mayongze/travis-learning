package main

import (
	"fmt"
	"log"
)

var (
	// GitSHA 自动构建
	GitSHA string = ""
)

func main() {
	log.Println("GitSHA=", GitSHA)
	str := "Hello World!"
	result := Reverse(str)
	log.Println(result)
	fmtFormat()
}

func fmtFormat() {
	type Student struct {
		Name  string
		Class int
	}
	var joy = Student{Name: "joy", Class: 5}

	placeholders := []string{"%v", "%+v", "%#v", "%T", "%%", "%q"}
	for _, v := range placeholders {
		fmt.Printf("Printf(\"%s\", joy)\t", v)
		fmt.Printf(v, joy)
		fmt.Println()
	}
}
