package services

import (
	"gitlab/nefco/access-management-system/src/storage"
)

type ServicesService interface {
	GetAllServices() (*items, error)
	GetActionsByService(int) (*items, error)
}

type servicesService struct {
	storage storage.DB
}

func NewServicesService(
	storage storage.DB,
) ServicesService {
	return &servicesService{
		storage,
	}
}

func (service *servicesService) GetAllServices() (*items, error) {
	services, err := service.storage.GetAllServices()
	if err != nil {
		return nil, err
	}

	return &items{services}, nil
}

func (service *servicesService) GetActionsByService(
	servicesId int,
) (*items, error) {
	serv, err := service.storage.GetService(servicesId)
	if err != nil {
		return nil, NewEntityNotFoundError("service", servicesId)
	}

	actions, err := service.storage.GetActionsByService(serv)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	return &items{actions}, nil
}
