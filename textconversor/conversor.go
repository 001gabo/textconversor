package main

import (
	"fmt"
	"github.com/textconversor/entities"
)

type Binary struct {
}

type Morse struct {
}

func (b Binary) Parse(i entities.Input) (response string, err error) {
	for _, c := range i.Texto {
		response = fmt.Sprintf("%s%.8b", i.Texto, c)
	}
	return response, nil
}

func (m Morse) Parse(i entities.Input) (response string, err error) {
	return "", nil
}
