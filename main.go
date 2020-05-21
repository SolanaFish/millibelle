package main

import (
	"github.com/solanafish/millibelle/internal/ledScreen"
	"github.com/solanafish/millibelle/internal/screenController"
)

func main() {
	ledScreen.Init()
	screenController.Init()
	ledScreen.Finish()
}
