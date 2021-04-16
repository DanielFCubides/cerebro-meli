package usecase

import "cerebro/domain"

type MutantVerifier interface {
	IsMutant(dna []string) bool
	GetStats() domain.Stats
}
