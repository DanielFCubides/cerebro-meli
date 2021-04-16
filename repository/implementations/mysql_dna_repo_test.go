package implementations

import (
	"cerebro/infrastructure"
	"cerebro/repository"
	"errors"
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MysqlDNARepoSuite struct {
	suite.Suite
	db   *gorm.DB
	repo repository.DNARepo
}

func (s *MysqlDNARepoSuite) SetupTest() {
	connection := infrastructure.NewMockConnection()
	s.repo = NewMysqlDNARepo(connection)
}

func TestRegisterUserHandlerInit(t *testing.T) {
	suite.Run(t, new(MysqlDNARepoSuite))
}

func (s *MysqlDNARepoSuite) TestGetStatsSuccessfully() {
	mocket.Catcher.Reset().NewMock().
		WithQuery("SELECT count(*) FROM \"dnas\"  WHERE (is_mutant = true)").
		WithReply([]map[string]interface{}{
			{
				"COUNT(*)": 40,
			},
		})
	mocket.Catcher.NewMock().
		WithQuery("SELECT count(*) FROM \"dnas\"  WHERE (is_mutant = false )").
		WithReply([]map[string]interface{}{
			{
				"COUNT(*)": 100,
			},
		})

	mutants, humans, ratio := s.repo.GetStats()
	s.Equal(40, mutants)
	s.Equal(100, humans)
	s.Equal(0.4, ratio)
}

func (s *MysqlDNARepoSuite) TestGetStatsWhenNoHumans() {
	mocket.Catcher.Reset().NewMock().
		WithQuery("SELECT count(*) FROM \"dnas\"  WHERE (is_mutant = true)").
		WithReply([]map[string]interface{}{
			{
				"COUNT(*)": 40,
			},
		})
	mocket.Catcher.NewMock().
		WithQuery("SELECT count(*) FROM \"dnas\"  WHERE (is_mutant = false )").
		WithReply([]map[string]interface{}{
			{
				"COUNT(*)": 100,
			},
		})

	mutants, humans, ratio := s.repo.GetStats()
	s.Equal(40, mutants)
	s.Equal(100, humans)
	s.Equal(0.4, ratio)
}

func (s *MysqlDNARepoSuite) TestSaveSucessfull() {
	mocket.Catcher.Reset()
	err := s.repo.Save([]string{"asdasd", "asdasd"}, true)
	s.NoError(err)
}

func (s *MysqlDNARepoSuite) TestSaveOnFailure() {
	mocket.Catcher.Reset().NewMock().
		WithQuery("INSERT INTO \"dnas\" (\"dna\",\"is_mutant\") VALUES (?,?)").
		WithError(errors.New("mock error"))
	err := s.repo.Save([]string{"asdasd", "asdasd"}, true)
	s.Error(err)
}
