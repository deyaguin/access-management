package services

import (
	"gitlab/nefco/access-management-system/src/storage"
)

type ServicesService interface {
	GetAllServices() (*pureItems, error)
	GetActionsByService(int) (*pureItems, error)
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

func (service *servicesService) GetAllServices() (*pureItems, error) {
	services, err := service.storage.GetAllServices()
	if err != nil {
		return nil, err
	}

	return &pureItems{services}, nil
}

func (service *servicesService) GetActionsByService(
	servicesId int,
) (*pureItems, error) {
	serv, err := service.storage.GetService(servicesId)
	if err != nil {
		return nil, NewEntityNotFoundError("service", servicesId)
	}

	actions, err := service.storage.GetActionsByService(serv)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	return &pureItems{actions}, nil
}
