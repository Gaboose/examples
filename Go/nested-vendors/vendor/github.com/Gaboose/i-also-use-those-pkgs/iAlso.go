package iAlso

import (
	"fmt"

	"github.com/Gaboose/color"
	"github.com/Gaboose/flavor"
	"github.com/Gaboose/fruit"
)

func WhatDoYouSee() string {
	return fmt.Sprintf("I see %s - it's %s and it tastes %s",
		fruit.GetFruit(), color.GetColor(), flavor.GetFlavor())
}
