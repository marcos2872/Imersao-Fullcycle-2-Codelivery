package main

import (
	"fmt"
	"github.com/codmarcos/imersaofsfc2-simulator/application/route"
)

func main() {
	route2 := route.Route{
		ID: "1",
		ClientID: "1",
	}
	route2.LoadPositions()
	stringjson, err := route2.ExportJsonPositions()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stringjson[9])
}