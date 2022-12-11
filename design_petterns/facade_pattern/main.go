package main

import (
	"develop/go-design/design_petterns/facade_pattern/pagemaker"
)

func main() {
	pagemaker.PageMaker.MakeWelcomePage("test3@test.com", "welcome.html")
}
