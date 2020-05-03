package main

import (
	"fmt"

	"github.com/dema501/randomjoke/internal/pkg/request"

	"github.com/dema501/randomjoke/internal/pkg/jokemaker"
	"github.com/dema501/randomjoke/internal/pkg/namegiver"
)

func main() {
	sa := request.New()
	if j, err := MakeJoke(namegiver.New(sa), jokemaker.New(sa)); err != nil {
		fmt.Printf("[ERROR] %v", err)
	} else {
		fmt.Println(j)
	}
}

func MakeJoke(nm namegiver.Doer, jm jokemaker.Doer) (string, error) {
	if err := nm.Generate(); err != nil {
		return "", err
	}

	firstName, lastName := nm.GetName()
	jm.SetName(firstName, lastName)

	if err := jm.Generate(); err != nil {
		return "", err
	}

	return jm.GetJoke(), nil
}
