package cases_repository

import (
	"databse-cluster-master-slave-architecture-golang/app/database"
	"databse-cluster-master-slave-architecture-golang/app/models"

	"gorm.io/gorm"
)

type Cases_Repository struct {
	master *gorm.DB
	slave  *gorm.DB
}

func NewCasesRepositoryRegistry() *Cases_Repository {
	return &Cases_Repository{
		master: database.Connect().Master,
		slave:  database.Connect().SlaveCases,
	}
}

func (repo *Cases_Repository) Create(cases *models.Cases) error {

	errCreate := repo.master.Table("cases").Create(cases).Error

	return errCreate

}

func (repo *Cases_Repository) GetAll() ([]models.Cases, error) {

	var cases []models.Cases

	errGet := repo.slave.Table("cases").Preload("Suspects").Find(&cases).Error

	return cases, errGet

}

func (repo *Cases_Repository) GetById(ID string) (*models.Cases, error) {

	var cases *models.Cases

	errGet := repo.slave.Table("cases").Preload("Suspects").Where("id = ?", ID).First(&cases).Error

	return cases, errGet

}

func (repo *Cases_Repository) GetByCaseNumber(case_number string) (*models.Cases, error) {

	var cases *models.Cases

	errGet := repo.slave.Table("cases").Preload("Suspects").Where("case_number = ?", case_number).First(&cases).Error

	return cases, errGet

}

func (repo *Cases_Repository) Update(ID string, cases *models.Cases) error {

	errUpdate := repo.master.Table("cases").Where("id = ?", ID).Updates(cases).Error

	return errUpdate

}

func (repo *Cases_Repository) Delete(ID string) error {

	var cases *models.Cases

	errDelete := repo.master.Table("cases").Unscoped().Where("id = ?", ID).Delete(&cases).Error

	return errDelete

}
