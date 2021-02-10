package main

import (
	"bufio"
	"fmt"
	"github.com/textconversor/entities"
	"github.com/textconversor/singleton"
	"github.com/textconversor/validator"
	"os"
)

func main() {

	var resp string
	var err error

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Texto a traducir: ")
	text, _ := reader.ReadString('\n')

	fmt.Print("Formato de origen: ")
	origin, _ := reader.ReadString('\n')

	fmt.Print("Formato de destino: ")
	destiny, _ := reader.ReadString('\n')

	bin := Binary{}
	mor := Morse{}

	//para binario
	mwBin := newMiddleware(bin)

	//para morse
	mwMorse := newMiddleware(mor)

	validator := validator.NewValidator()

	request := entities.Input{
		Texto:           text,
		OriginFormat:    origin,
		DestinityFormat: destiny,
	}

	//creating instance of supported formats
	singleton := singleton.GetInstances()

	//validating formats
	err = validator.Validate(singleton, request)
	if err != nil {
		fmt.Println("Hubo un error al momento de validar los formatos ingresados (origen o destino)", err)
		return
	}

	if request.DestinityFormat == "morse" {
		resp, err = mwMorse.parse(request)
	} else if request.DestinityFormat == "binary" {
		resp, err = mwBin.parse(request)
	} else {

	}

	if err != nil {
		fmt.Println("Hubo un error al momento de hacer el parseo", err)
	}

	fmt.Println("Texto convertido", resp)

}
