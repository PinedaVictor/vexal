// Package commands handles internal logic for cli workflow
package commands

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var s = spinner.New(spinner.CharSets[33], 100*time.Millisecond)
var c = color.New(color.FgCyan, color.Bold)

// StartSpinner sets starts spinner animation and sets config
func StartSpinner(prefix string) {
	s.Start() // Start the spinner
	s.FinalMSG = c.Sprint("Complete! \n")
	s.Color("green", "bold")
	s.Prefix = c.Sprint(prefix) // Prefix text before the spinner
}

// StopSpinner stops the s spinner
func StopSpinner() {
	s.Stop()
}
