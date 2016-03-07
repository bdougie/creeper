package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

var possible = []string{
	"afogel",
	"ambriz",
	"brett",
	"brian",
	"christian",
	"jared",
	"jason",
	"megan",
	"Patrick",
	"joe",
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

func removeFromReviewers(r *Reviewers, name string) {
	for i, p := range r.possible {
		match, _ := regexp.MatchString(p, strings.ToLower(name))

		if match {
			r.possible = append(r.possible[:i], r.possible[i+1:]...)
			fmt.Println(r.possible)
		}
	}
}

func pickReviewers() {
	phabName := os.Getenv("PHABRICATOR_USERNAME")
	reviewers := Reviewers{possible}

	removeFromReviewers(&reviewers, phabName)

	if phabName == "" {
		color.Red("You must set ENV['PHABRICATOR_USERNAME'] somewhere! (╯°□°）╯︵ ┻━┻")
		return
	}

	if !contains(reviewers, phabName) {
		color.Red("%v is not in the eligible reviewers list, please update the possibleReviewers file", phabName)
		return
	}

	randomlySelectFromPossible(&reviewers)
}

func randomlySelectFromPossible(r *Reviewers) {
	firstThree := rand.Perm(len(r.possible))
	for _, v := range firstThree {
		fmt.Println(r.possible[v])
	}

}

func main() {
	pickReviewers()
}
