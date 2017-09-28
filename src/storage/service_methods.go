package storage

import (
	"gitlab/nefco/access-management-system/src/models"
)

func (dataBase *SqlDB) GetService(id int) (*models.Service, error) {
	service := new(models.Service)

	err := dataBase.Where(id).Find(service).Error

	return service, err
}

func (dataBase *SqlDB) GetAllServices() (*[]models.Service, error) {
	services := new([]models.Service)

	err := dataBase.Find(services).Error

	return services, err
}

func (dataBase *SqlDB) GetActionsByService(
	service *models.Service,
) (*[]models.Action, error) {
	actions := new([]models.Action)

	err := dataBase.Where("service_id = ?", service.ID).Find(actions).Error

	return actions, err
}
