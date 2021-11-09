package environment

type Environment int

const (
	Local Environment = iota
	Production
)

func (d Environment) String() string {
	return [...]string{"local", "production"}[d]
}

func GetStringFrom(s string) Environment {
	if s == "production" {
		return Production
	}

	return Local
}
