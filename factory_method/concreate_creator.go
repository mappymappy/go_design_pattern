package factory_method

// Programmer

type Programmer interface {
	CreateHelloWorld() HelloWorldInterface
}

// Gopher Team
type GoEngineerTeam struct {
	EngineerTeam
}

func NewGoEngineerTeam() GoEngineerTeam {
	return GoEngineerTeam{}
}

func (g *GoEngineerTeam) CreateHelloWorld() {
	g.engineer = Gopher{}
	g.result = g.engineer.CreateHelloWorld()
}

type Gopher struct{}

func (gop Gopher) CreateHelloWorld() HelloWorldInterface {
	return NewGoCode()
}

// PHPer Team
type PHPEngineerTeam struct {
	EngineerTeam
}

func NewPHPEngineerTeam() PHPEngineerTeam {
	return PHPEngineerTeam{}
}

func (p *PHPEngineerTeam) CreateHelloWorld() {
	p.engineer = PHPer{}
	p.result = p.engineer.CreateHelloWorld()
}

type PHPer struct {
}

func (p PHPer) CreateHelloWorld() HelloWorldInterface {
	return NewPHPCode()
}
