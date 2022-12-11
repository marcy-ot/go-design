package main

import (
	"develop/gotraining/factory_pattern/framework"
	"develop/gotraining/factory_pattern/idcard"
)

// factory pattern
func main() {
	var factory framework.Factory = idcard.NewIdCardFactory()

	var card1 framework.Product = factory.Create("test1")
	var card2 framework.Product = factory.Create("test2")
	var card3 framework.Product = factory.Create("test3")

	card1.Use()
	card2.Use()
	card3.Use()
}
