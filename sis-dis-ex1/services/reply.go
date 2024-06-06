package services

type Reply struct{}

func (t *Reply) Echo(msg string) string {
	return msg
}

func (t *Reply) Hello() string {
	return "Hello"
}
