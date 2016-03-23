package main

import (
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

type Reviewers struct {
	possible []string
}

func contains(r Reviewers, s string) bool {
	for _, i := range r.possible {
		if i == s {
			return true
		}
	}
	return false
}

func removeMeFromReviewers(r *Reviewers, name string) {
	for i, p := range r.possible {
		match, _ := regexp.MatchString(p, strings.ToLower(name))

		if match {
			r.possible = append(r.possible[:i], r.possible[i+1:]...)
		}
	}
}

func pickReviewers() {
	phabName := os.Getenv("PHABRICATOR_USERNAME")
	reviewers := Reviewers{possible}

	if phabName == "" {
		color.Red("You must set ENV['PHABRICATOR_USERNAME'] somewhere! (╯°□°）╯︵ ┻━┻")
		return
	}

	if !contains(reviewers, phabName) {
		color.Red("%v is not in the eligible reviewers list, please update the possibleReviewers file", phabName)
		return
	}

	removeMeFromReviewers(&reviewers, phabName)
	randomlySelectFromPossible(&reviewers)
}

func randomlySelectFromPossible(r *Reviewers) {
	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(r.possible))

	for _, v := range randomize[:4] {
		color.Red(r.possible[v])
	}

}

func main() {
	pickReviewers()
}
