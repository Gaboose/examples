package main

import (
	"fmt"

	"github.com/Gaboose/color"
	"github.com/Gaboose/fruit"
	"github.com/Gaboose/i-also-use-those-pkgs"

	// 'go build' does not search vendor dirs recursively
	// it won't find flavor from this package
	// "github.com/Gaboose/flavor"
)

func main() {
	fmt.Println("From root:")
	fmt.Printf("I see %s - it's %s and it tastes %s\n",
		fruit.GetFruit(), color.GetColor(), "(can't import flavor)")

	fmt.Println("\nFrom 'iAlso':")
	fmt.Println(iAlso.WhatDoYouSee())
}
