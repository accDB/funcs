package funcs

type In interface {
	From([]string) error
	To() ([]string, error)
}

type Out interface {
	From([]byte) error
	To() ([]byte, error)
	ReUse()
}
