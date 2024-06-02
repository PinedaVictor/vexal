// Package commands handles internal logic for cli workflow
package internal

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var s = spinner.New(spinner.CharSets[33], 100*time.Millisecond)
var c = color.New(color.FgCyan, color.Bold)

// StartSpinner sets starts spinner animation and sets config
func StartSpinner(prefix string) {
	s.Start() // Start the spinner
	s.Color("green", "bold")
	s.Prefix = c.Sprint(prefix) // Prefix text before the spinner
}

// StopSpinner stops the s spinner
func StopSpinner(userFeedback string) {
	s.FinalMSG = c.Sprint(fmt.Sprintf("%s \n", userFeedback))
	s.Stop()
}
