package infrastructure

import (
	"go.uber.org/dig"
	"log"
)

const (
	moduleName = "injection"
)

var Injector = dig.New(dig.DeferAcyclicVerification())

func CheckInjection(err error, instanceName string) {

	if err == nil {
		return
	}

	log.Printf("error on dependency injection %s", err.Error())

	panic(err)

}
