package validator

import (
	"errors"
	"github.com/textconversor/entities"
	"github.com/textconversor/singleton"
)

type validator struct {
}

func NewValidator() *validator {
	return &validator{}
}

func (v *validator) Validate(singleton *singleton.Single, request entities.Input) error {

	originFlag := validateOriginFormat(singleton, request.OriginFormat)

	if originFlag != true {
		return errors.New("El formato de origen no es soportado")
	}

	destinyFlag := validateDestinyFormat(singleton, request.DestinityFormat)

	if destinyFlag != true {
		return errors.New("El formato de destino no es soportado")
	}
	return nil
}

func validateOriginFormat(singleton *singleton.Single, origin string) bool {

	for _, v := range singleton.SupportedFormats {
		if v == origin {
			return true
		}
	}
	return false
}

func validateDestinyFormat(singleton *singleton.Single, destiny string) bool {

	for _, v := range singleton.SupportedFormats {
		if v == destiny {
			return true
		}
	}
	return false
}
