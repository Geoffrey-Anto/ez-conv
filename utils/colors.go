package utils

type Colors struct{}

func (c Colors) Reset() string {
	return "\033[0m"
}

func (c Colors) Red() string {
	return "\033[31m"
}

func (c Colors) Green() string {
	return "\033[32m"
}

func (c Colors) Yellow() string {
	return "\033[33m"
}

func (c Colors) Blue() string {
	return "\033[34m"
}

func (c Colors) Purple() string {
	return "\033[35m"
}

func (c Colors) Cyan() string {
	return "\033[36m"
}

func (c Colors) White() string {
	return "\033[37m"
}
