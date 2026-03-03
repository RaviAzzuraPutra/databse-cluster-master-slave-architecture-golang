package cases_registry

import (
	"databse-cluster-master-slave-architecture-golang/app/controller/cases_controller"
	"databse-cluster-master-slave-architecture-golang/app/repository/cases_repository"
	"databse-cluster-master-slave-architecture-golang/app/service/cases_service"
)

type Cases_Module struct {
	Cases_Controller *cases_controller.Cases_Controller
}

func Case_Registry() *Cases_Module {
	repository := cases_repository.NewCasesRepositoryRegistry()

	service := cases_service.NewCasesServiceRegistry(repository)

	controller := cases_controller.NewCasesControllerRegistry(service)

	return &Cases_Module{
		Cases_Controller: controller,
	}
}
