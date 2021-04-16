package adapters

import (
	"cerebro/infrastructure"
	_ "cerebro/infrastructure"
	_ "cerebro/repository/implementations"
	_ "cerebro/usecase/implementations"
	"github.com/gin-gonic/gin"
	"log"
)

func ServerSetup() *gin.Engine {
	router := gin.Default()
	//migrate models
	//Inject adapters
	var adapter *RestAdapter
	invokeFunc := func(a *RestAdapter) { adapter = a }
	err := infrastructure.Injector.Invoke(invokeFunc)
	if err != nil {
		log.Println("Error providing RestAdapter instance:", err)
		panic(err)
	}
	//define routes
	router.POST("/mutant/", adapter.MutantVerifier)
	router.GET("/stats", adapter.GetStats)

	return router

}
