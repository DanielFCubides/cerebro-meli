package implementations

import (
	"cerebro/infrastructure"
	"cerebro/repository"
	"github.com/jinzhu/gorm"
	"log"
)

func init() {
	err := infrastructure.Injector.Provide(NewMysqlDNARepo)
	if err != nil {
		log.Println("Error providing MutantVerifierImpl instance:", err)
		panic(err)
	}
}

type MysqlDNARepo struct {
	db *gorm.DB
}

func (repo *MysqlDNARepo) Save(dna []string, isMutant bool) error {
	dnaConcat := ""
	for _, s := range dna {
		dnaConcat += s + "|"
	}
	err := repo.db.Save(&repository.DNA{
		Dna:      dnaConcat,
		IsMutant: isMutant,
	}).Error
	return err
}

func (repo *MysqlDNARepo) GetStats() (mutants int, humans int, ratio float64) {
	err := repo.db.Model(&repository.DNA{}).Where("is_mutant = ?", true).Count(&mutants).Error
	if err != nil {
		log.Printf(err.Error())
	}
	err = repo.db.Model(&repository.DNA{}).Where("is_mutant = ? ", false).Count(&humans).Error
	if err != nil {
		log.Printf(err.Error())
	}
	if humans != 0 {
		return mutants, humans, float64(mutants) / float64(humans)
	}

	return mutants, humans, 1.0

}

func NewMysqlDNARepo(conn infrastructure.Connection) repository.DNARepo {
	return &MysqlDNARepo{db: conn.GetDatabase()}
}
