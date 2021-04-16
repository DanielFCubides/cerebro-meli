package repository

type DNARepo interface {
	Save(dna []string, isMutant bool) error
	GetStats() (mutants int, humans int, ratio float64)
}
