package db

import (
	"gitlab/nefco/accessControl/app/models"
)

func (dataBase SqlDB) CreatePolicy(policy *models.Policy) error {
	err := dataBase.db.Create(policy).Error
	return err
}

func (dataBase SqlDB) GetPolicies() (*[]models.Policy, error) {
	policies := new([]models.Policy)
	err := dataBase.db.Find(policies).Error
	return policies, err
}

func (dataBase SqlDB) GetPolicy(id int) (*models.Policy, error) {
	policy := new(models.Policy)
	err := dataBase.db.Where(id).Find(policy).Error
	return policy, err
}

func (dataBase SqlDB) UpdatePolicy(policy *models.Policy) error {
	err := dataBase.db.Save(policy).Error
	return err
}

func (dataBase SqlDB) DeletePolicy(policy *models.Policy) error {
	err := dataBase.db.Where("id = ?", policy.ID).Delete(policy).Error
	return err
}
