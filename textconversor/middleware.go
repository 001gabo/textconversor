package main

import (
	"fmt"
	"github.com/textconversor/entities"
)

type middleware struct {
	conv entities.Conversor
}

func newMiddleware(conv entities.Conversor) *middleware {
	return &middleware{
		conv: conv,
	}
}

func (mw *middleware) parse(i entities.Input) (response string, err error) {

	fmt.Println("Texto recibido", i.Texto)

	//llamando al metodo de parseo dependiendo del tipo
	response, err = mw.conv.Parse(i)

	if err != nil {
		fmt.Println("Hubo un error al momento de hacer el parseo en la implementacion concreta", err)
		return "", err
	}

	return response, nil
}
