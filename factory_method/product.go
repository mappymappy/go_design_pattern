package factory_method

// Code
type Code struct {
	printString string
}

func (c Code) HelloWorld() string {
	return c.printString
}
