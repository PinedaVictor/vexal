// Package internal handles CLI feedback and loading state
package internal

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#10b981")).Bold(true)
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ef4444")).Bold(true)
	warnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#f59e0b")).Bold(true)
	dimStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#1e90ff"))
	quoteStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#1e90ff")).Italic(true)
	timerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#10b981"))
)

var officeQuotes = []string{
	"Bears. Beets. Battlestar Galactica.",
	"That's what she said.",
	"I am Beyoncé, always.",
	"Identity theft is not a joke, Jim!",
	"Why are you the way that you are?",
	"I'm not superstitious, but I am a little stitious.",
	"Wikipedia is the best thing ever.",
	"I love inside jokes. I hope to be a part of one someday.",
	"Would I rather be feared or loved? Both.",
	"I am running away from my responsibilities.",
	"Fact: I am a great driver.",
	"Fool me once, strike one.",
	"I talk a lot, so I've learned to tune myself out.",
	"I don't care what they say about me.",
	"How the turntables...",
	"No. God. Please. No.",
	"I declare bankruptcy!",
	"I am faster than 80% of all snakes.",
	"Dwight, you ignorant slut.",
	"You have no idea how high I can fly.",
	"Dunder Mifflin, this is Pam.",
	"I hate so much about the things that you choose to be.",
	"Number one: how dare you.",
	"I want people to be afraid of how much they love me.",
	"Perfectenschlag.",
	"Parkour!",
}

var DftlDoneMsg = "All done"
var s = spinner.New(spinner.CharSets[14], 80*time.Millisecond)

var (
	spinnerMu    sync.Mutex
	spinnerStop  chan struct{}
	spinnerStart time.Time
	spinnerLabel string
)

func StartSpinner(prefix string) {
	quote := officeQuotes[rand.Intn(len(officeQuotes))]

	spinnerMu.Lock()
	spinnerStop = make(chan struct{})
	spinnerStart = time.Now()
	spinnerLabel = prefix
	spinnerMu.Unlock()

	updatePrefix(prefix, quote, 0)
	s.Color("fgHiYellow")
	s.Start()

	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		for {
			select {
			case <-spinnerStop:
				return
			case <-tick.C:
				spinnerMu.Lock()
				elapsed := int(time.Since(spinnerStart).Seconds())
				lbl := spinnerLabel
				spinnerMu.Unlock()
				updatePrefix(lbl, "", elapsed)
			}
		}
	}()
}

func StopSpinner(msg string) {
	spinnerMu.Lock()
	if spinnerStop != nil {
		close(spinnerStop)
		spinnerStop = nil
	}
	spinnerMu.Unlock()

		if msg == "" {
		s.FinalMSG = ""
	} else {
		s.FinalMSG = successStyle.Render("✓") + " " + dimStyle.Render(msg) + "\n"
	}
	s.Stop()
}

func updatePrefix(label, quote string, elapsed int) {
	prefix := dimStyle.Render("❯ "+label)
	if quote != "" {
		prefix += " " + quoteStyle.Render("\""+quote+"\"")
	}
	prefix += " " + timerStyle.Render(fmt.Sprintf("(%ds)", elapsed)) + " "
	s.Prefix = prefix
}

func PreFeedback(msg string) {
	fmt.Println(dimStyle.Render(msg))
}

func PostFeedback(msg string) {
	fmt.Println(successStyle.Render("✓") + " " + msg)
}

func UserFeedback(msg string) {
	fmt.Println(warnStyle.Render("⚠") + " " + msg)
}

func UserErrFeedback(msg string) {
	fmt.Println(errorStyle.Render("✗") + " " + msg)
}
