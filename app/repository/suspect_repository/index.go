package suspect_repository

import (
	"databse-cluster-master-slave-architecture-golang/app/database"
	"databse-cluster-master-slave-architecture-golang/app/models"

	"gorm.io/gorm"
)

type Suspect_Repository struct {
	master *gorm.DB
	slave2 *gorm.DB
}

func NewSuspectRepositoryRegistry() *Suspect_Repository {
	return &Suspect_Repository{
		master: database.Connect().Master,
		slave2: database.Connect().SlaveSuspects,
	}
}

func (repo *Suspect_Repository) Create(suspect *models.Suspects) error {

	errCreate := repo.master.Table("suspects").Create(suspect).Error

	return errCreate

}

func (repo *Suspect_Repository) GetAll(ID_Case string) ([]models.Suspects, error) {

	var suspects []models.Suspects

	errGet := repo.slave2.Table("suspects").Where("case_id = ?", ID_Case).Find(&suspects).Error

	return suspects, errGet

}

func (repo *Suspect_Repository) GetById(ID string, ID_Case string) (*models.Suspects, error) {

	var suspect *models.Suspects

	errGet := repo.slave2.Table("suspects").Where("id = ? AND case_id = ?", ID, ID_Case).First(&suspect).Error

	return suspect, errGet

}

func (repo *Suspect_Repository) Update(ID string, ID_Case string, suspect *models.Suspects) error {

	errUpdate := repo.master.Table("suspects").Where("id = ? AND case_id = ?", ID, ID_Case).Updates(suspect).Error

	return errUpdate

}

func (repo *Suspect_Repository) Delete(ID string, ID_Case string) error {

	var suspect *models.Suspects

	errDelete := repo.master.Table("suspects").Unscoped().Where("id = ? AND case_id = ?", ID, ID_Case).Delete(&suspect).Error

	return errDelete

}
