# Creeper

![Al](http://static.tvtropes.org/pmwiki/pub/images/weirdal.jpg)

Selecting people to review your code is hard. It's also hard to remember who you
picked last and who you ping too much for pull requests.
`creep` is a script that picks 4 people to review my code at random and
copys to your clipboard.

This is a Phabricator/Arcanist specific but the concept it easy to read. 


## Architecture

There is an array that holds all possible reviewers of my code:

```go
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
```

The Creeper Struct holds all the methods and possible reviewers array.

```go
type Creeper struct {
	possible []string
}
```

`removeMeFromPossible()` and `contains()` are used to remove (me/the user) from the possible reviewers array, which is why I opted to use a Struct.

`randomlySelectFromPossible()` does as titled and picks the reviewers at
random.

Everything else is self-explanatory and fired off in `main()` with the `pickReviewers()` function

If you have Bloc-Alpha access this is a port of my [script/reviewers](https://github.com/Bloc/Bloc/blob/master/script/reviewers). You see the code side by side is about the same. 

## Setup

Download the latest [Go](https://golang.org/dl/) and then:
```bash
$ cd creep
$ go install
```

Installing this will allow you to run the creep from any repo.

Make sure you set your `ENV['PHABRICATOR_USERNAME']` in your `.bashrc` or
equivalent.

## Run

`$ creep`
