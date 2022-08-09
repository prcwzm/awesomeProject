package main

type Stringer interface {
	String() string
}

var value interface{}

func PrintMethod() string {
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
	return ""
}
