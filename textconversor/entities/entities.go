package entities

type Input struct {
	Texto           string
	OriginFormat    string
	DestinityFormat string
}

type Conversor interface {
	Parse(i Input) (response string, err error)
}
