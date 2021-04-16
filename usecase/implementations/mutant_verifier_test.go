package implementations

import (
	"cerebro/repository/mocks"
	"cerebro/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MutantVerifierSuite struct {
	suite.Suite
	verifier usecase.MutantVerifier
	repo     *mocks.DNARepo
}

func TestManagerService(t *testing.T) {
	suite.Run(t, &MutantVerifierSuite{})
}

func (s *MutantVerifierSuite) SetupTest() {
	s.repo = &mocks.DNARepo{}
	s.verifier = NewMutantVerifierImpl(s.repo)
	s.repo.On("Save", mock.Anything, mock.Anything).Return(nil)
	s.repo.On("GetStats").Return(40, 100, 0.4)
}

func (s *MutantVerifierSuite) TestForMutantDNA() {
	isMutant := s.verifier.IsMutant([]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"})
	s.True(isMutant)
}

func (s *MutantVerifierSuite) TestForMutantDNAForAliens() {
	isMutant := s.verifier.IsMutant([]string{"XTGCGX", "CXGTGC", "TTXTGT", "XGXXGG", "CCCXTX", "TCXCTG"})
	s.False(isMutant)
}

func (s *MutantVerifierSuite) TestForMutantDNAMinimumMatrix() {
	isMutant := s.verifier.IsMutant([]string{"ATGC", "CAGT", "TTAT", "AAAA"})
	s.True(isMutant)
}

func (s *MutantVerifierSuite) TestForMutantDNAWhenEmpty() {
	isMutant := s.verifier.IsMutant([]string{})
	s.False(isMutant)
}

func (s *MutantVerifierSuite) TestForHumanDNA() {
	isMutant := s.verifier.IsMutant([]string{"ATGCGA", "CCGTCC", "TTATGT", "AGAAGG", "ACCCTA", "TCACTG"})
	s.False(isMutant)
}

func (s *MutantVerifierSuite) TestForHumanDNAMinimumWhitOneGeneMatrix() {
	isMutant := s.verifier.IsMutant([]string{"ATGC", "CCGT", "TTAT", "AAAA"})
	s.False(isMutant)
}

func (s *MutantVerifierSuite) TestForHumanDNAMinimumWhitZeroGeneMatrix() {
	isMutant := s.verifier.IsMutant([]string{"ATGC", "CAGT", "TTAT", "AATA"})
	s.False(isMutant)
}

func (s *MutantVerifierSuite) TestInvalidMatrix() {
	isMutant := s.verifier.IsMutant([]string{"ATGC", "CAGT", "TTT", "AATA"})
	s.False(isMutant)
}
