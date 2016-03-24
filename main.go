package main

import (
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var possible = []string{
	"afogel",
	"ambriz",
	"brett",
	"brian",
	"jared",
	"jason",
	"megan",
	"Patrick",
	"joe",
	"stan",
	"aaron",
}

type Creeper struct {
	possible []string
}

func contains(r Creeper, s string) bool {
	for _, i := range r.possible {
		if i == s {
			return true
		}
	}
	return false
}

func removeMeFromPossible(r *Creeper, name string) {
	for i, p := range r.possible {
		match, _ := regexp.MatchString(p, strings.ToLower(name))

		if match {
			r.possible = append(r.possible[:i], r.possible[i+1:]...)
		}
	}
}

func pickReviewers() {
	phabName := os.Getenv("PHABRICATOR_USERNAME")
	reviewers := Creeper{possible}

	if phabName == "" {
		color.Red("You must set ENV['PHABRICATOR_USERNAME'] somewhere! (╯°□°）╯︵ ┻━┻")
		return
	}

	if !contains(reviewers, phabName) {
		color.Red("Weird, %v is not in the eligible reviewers list, please update the possible reviewers array", phabName)
		return
	}

	removeMeFromPossible(&reviewers, phabName)
	randomlySelectFromPossible(&reviewers)
}

func randomlySelectFromPossible(r *Creeper) {
	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(r.possible))

	reviewers := make([]string, 4)
	i := 0

	for _, v := range randomize[:4] {
		reviewers[i] = r.possible[v]
		i++
	}

	joined := strings.Join(reviewers, " ")
	color.Red(joined)

	copyToClipboard(joined)
}

func copyToClipboard(j string) {
	if err := clipboard.WriteAll(string(j)); err != nil {
		panic(err)
	}

	color.Green("┬──┬◡ﾉ(° -°ﾉ) copied to your clipboard")
}

func main() {
	pickReviewers()
}
