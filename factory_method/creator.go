package factory_method

type EngineerTeam struct {
	result   HelloWorldInterface
	engineer Programmer
}

type HelloWorldInterface interface {
	HelloWorld() string
}

func (e *EngineerTeam) CreateHelloWorld() {
	//override
}

func (e *EngineerTeam) TestHelloWorld() string {
	return e.result.HelloWorld()
}
