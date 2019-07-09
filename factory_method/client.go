package factory_method

import "fmt"

func run() {

	phpTeam := NewPHPEngineerTeam()
	phpTeam.CreateHelloWorld()
	fmt.Printf("%s\n", phpTeam.TestHelloWorld())

	goTeam := NewGoEngineerTeam()
	goTeam.CreateHelloWorld()
	fmt.Printf("%s\n", goTeam.TestHelloWorld())
}
