package factory_method

// GoCode
type GoCode struct {
	Code
}

func NewGoCode() GoCode {
	g := GoCode{}
	g.printString = "hello go world!!"
	return g
}

// PHPCode
type PHPCode struct {
	Code
}

func NewPHPCode() PHPCode {
	p := PHPCode{}
	p.printString = "hello php world!!"
	return p
}
