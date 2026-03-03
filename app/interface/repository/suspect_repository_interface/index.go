package suspect_repository_interface

import "databse-cluster-master-slave-architecture-golang/app/models"

type Suspect_Repository_Interface interface {
	Create(suspect *models.Suspects) error
	GetAll(ID_Case string) ([]models.Suspects, error)
	GetById(ID string, ID_Case string) (*models.Suspects, error)
	Update(ID string, ID_Case string, suspect *models.Suspects) error
	Delete(ID string, ID_Case string) error
}
