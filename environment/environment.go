package environment

import "errors"

var ErrUnknownEnvironment = errors.New("unknown environment")

type Environment string

const (
	Unknown Environment = ""
	Dev     Environment = "dev"
	Stage   Environment = "stage"
	Prod    Environment = "prod"
)

func (e Environment) String() string {
	return string(e)
}

func ParseEnvironment(env string) (Environment, error) {
	switch env {
	case Dev.String():
		return Dev, nil
	case Stage.String():
		return Stage, nil
	case Prod.String():
		return Prod, nil
	}

	return Unknown, ErrUnknownEnvironment
}
