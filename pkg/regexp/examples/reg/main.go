package main

import (
	"fmt"
	"regexp"

	regexp2 "github.com/jackie-feng/commontools/pkg/regexp"
)

func main() {
	reg := regexp2.Compile(".*.*=.*")
	fmt.Println(reg.Match("a=cccccc"))

	reg2, _ := regexp.Compile("a.*.*c")
	fmt.Println(reg2.Match([]byte("a=c")))
}
