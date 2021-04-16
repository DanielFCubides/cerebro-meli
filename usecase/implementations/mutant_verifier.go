package implementations

import (
	"cerebro/domain"
	"cerebro/infrastructure"
	"cerebro/repository"
	"cerebro/usecase"
	"fmt"
	"log"
)

func init() {
	err := infrastructure.Injector.Provide(NewMutantVerifierImpl)
	if err != nil {
		log.Println("Error providing MutantVerifierImpl instance:", err)
		panic(err)
	}
}

type MutantVerifierImpl struct {
	sequenceMap map[string]string
	repo        repository.DNARepo
}

func NewMutantVerifierImpl(r repository.DNARepo) usecase.MutantVerifier {
	return &MutantVerifierImpl{sequenceMap: map[string]string{
		"A": "AAA",
		"T": "TTT",
		"C": "CCC",
		"G": "GGG",
	},
		repo: r}
}

func (m *MutantVerifierImpl) IsMutant(dna []string) bool {
	mutantGens := 0
	valid := verifyDNA(dna)
	if !valid {
		return false
	}
	for i, DNAseq := range dna {
		for j, nucleobase := range DNAseq {
			nucleotide := string(nucleobase)
			if !m.isValidNucleotide(nucleotide) {
				return false
			}
			mutantGens += m.exploreNeighborhoods(i, j, nucleotide, dna)
			if mutantGens > 1 {
				_ = m.repo.Save(dna, true)
				return true
			}
		}
	}
	_ = m.repo.Save(dna, false)
	return false
}

func (m *MutantVerifierImpl) isValidNucleotide(nucleotide string) bool {
	if _, ok := m.sequenceMap[nucleotide]; ok {
		return true
	}
	return false
}

func verifyDNA(dna []string) bool {
	length := len(dna)
	if length < 4 {
		return false
	}
	for _, dnaSeq := range dna {
		if len(dnaSeq) != length {
			return false
		}
	}
	return true
}

func (m *MutantVerifierImpl) exploreNeighborhoods(i int, j int, nucleobase string, dna []string) int {
	mutantGenes := 0
	for _, neighbour := range m.getNeighborhoods(i, j, dna) {
		if neighbour == m.sequenceMap[nucleobase] {
			mutantGenes += 1
		}
	}
	return mutantGenes
}

func (m *MutantVerifierImpl) getNeighborhoods(i int, j int, dna []string) []string {
	length := len(dna)
	neighborhoods := make([]string, 4)

	//check right
	sequenceLenth := 3
	if j+sequenceLenth < length {
		neighborhoods[0] = dna[i][j+1 : j+4]
	}
	//check right-down
	if i+sequenceLenth < length && j+sequenceLenth < length {
		neighborhoods[1] = fmt.Sprintf("%s%s%s", string(dna[i+1][j+1]), string(dna[i+2][j+2]), string(dna[i+3][j+3]))
	}
	//check down
	if i+sequenceLenth < length {
		neighborhoods[2] = fmt.Sprintf("%s%s%s", string(dna[i+1][j]), string(dna[i+2][j]), string(dna[i+3][j]))
	}
	//check left-down
	if i+sequenceLenth < length && j-sequenceLenth >= 0 {
		neighborhoods[3] = fmt.Sprintf("%s%s%s", string(dna[i+1][j-1]), string(dna[i+2][j-2]), string(dna[i+3][j-3]))
	}
	return neighborhoods
}

func (m *MutantVerifierImpl) GetStats() domain.Stats {
	mutants, humans, ratio := m.repo.GetStats()
	return domain.Stats{
		CountMutantDna: mutants,
		CountHumanDna:  humans,
		Ratio:          ratio,
	}
}
