package repository

import (
	"warung-makan/model"
	"warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type customerRepository struct {
	db *sqlx.DB
}

type CustomerRepository interface {
	GetAll() ([]model.Customer, error)
	GetById(id string) (model.Customer, error)
	GetByName(name string) ([]model.Customer, error)
	GetByCredentials(customername, password string) (model.Customer, error)

	Insert(customer *model.Customer) (model.Customer, error)
	Update(customer *model.Customer) (model.Customer, error)
	Delete(id string) error
}

func (p *customerRepository) GetAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := p.db.Select(&customers, utils.USER_GET_ALL+" order by id")
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (p *customerRepository) GetById(id string) (model.Customer, error) {
	var customer model.Customer
	err := p.db.Get(&customer, utils.USER_GET_BY_ID, id)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (p *customerRepository) GetByName(name string) ([]model.Customer, error) {
	var customer []model.Customer
	err := p.db.Select(&customer, utils.USER_GET_BY_NAME, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (p *customerRepository) GetByCredentials(customername, password string) (model.Customer, error) {
	var customer model.Customer
	err := p.db.Get(&customer, utils.USER_GET_BY_CREDENTIALS, customername, password)
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (p *customerRepository) Insert(newCustomer *model.Customer) (model.Customer, error) {
	_, err := p.db.NamedExec(utils.USER_INSERT, newCustomer)
	if err != nil {
		return model.Customer{}, err
	}
	customer := newCustomer
	return *customer, nil
}

func (p *customerRepository) Update(newData *model.Customer) (model.Customer, error) {
	_, err := p.db.NamedExec(utils.USER_UPDATE, newData)
	if err != nil {
		return model.Customer{}, err
	}
	return *newData, nil
}

func (p *customerRepository) Delete(id string) error {
	_, err := p.db.Exec(utils.USER_DELETE, id)
	return err
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
