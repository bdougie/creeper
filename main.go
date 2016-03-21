package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/koding/cache"
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

	if phabName == "" {
		color.Red("You must set ENV['PHABRICATOR_USERNAME'] somewhere! (╯°□°）╯︵ ┻━┻")
		return
	}

	if !contains(reviewers, phabName) {
		color.Red("%v is not in the eligible reviewers list, please update the possibleReviewers file", phabName)
		return
	}

	removeFromReviewers(&reviewers, phabName)
	randomlySelectFromPossible(&reviewers)
	// p := reviewers.possible
	// Shuffle(p)
}

func randomlySelectFromPossible(r *Reviewers) {
	firstThree := rand.Perm(len(r.possible))

	for _, v := range firstThree[:3] {
		go fmt.Println(r.possible[v])
		// color.Red(r.possible[v])
	}

}

func Shuffle(slc []string) {
	N := len(slc)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + rand.Intn(N-i)
		slc[r], slc[i] = slc[i], slc[r]
	}
	fmt.Println(slc)
}

func main() {
	pickReviewers()
	NewMemoryWithTTL(2 * time.Second)
}
