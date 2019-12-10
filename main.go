package main

import (
	"./app"
	"flag"
	"fmt"
)

func main() {
	dir := flag.String("dir", "./", "Target directory path")
	dist := flag.String("dist", "./", "Distribution directory path")
	flag.Parse()
	fmt.Println(*dir)
	fmt.Println(*dist)
	if err := app.Start(*dir, *dist); err != nil {
		//TODO Error handling
	}
}
