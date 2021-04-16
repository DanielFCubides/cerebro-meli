package adapters

import (
	"cerebro/infrastructure"
	"cerebro/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	err := infrastructure.Injector.Provide(NewRestAdapter)
	if err != nil {
		log.Println("Error providing RestAdapter instance:", err)
		panic(err)
	}
}

type RestAdapter struct {
	mutantVerifier usecase.MutantVerifier
}

func NewRestAdapter(mv usecase.MutantVerifier) *RestAdapter {
	return &RestAdapter{
		mutantVerifier: mv,
	}
}

func (a *RestAdapter) MutantVerifier(c *gin.Context) {
	var body map[string][]string
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if _, ok := body["dna"]; !ok {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if a.mutantVerifier.IsMutant(body["dna"]) {
		c.AbortWithStatus(http.StatusOK)
	}
	c.AbortWithStatus(http.StatusForbidden)

}

func (a *RestAdapter) GetStats(c *gin.Context) {
	stats := a.mutantVerifier.GetStats()
	c.JSON(200, gin.H{
		"count_mutant_dna": stats.CountMutantDna,
		"count_human_dna":  stats.CountHumanDna,
		"ratio":            stats.Ratio,
	})
}
