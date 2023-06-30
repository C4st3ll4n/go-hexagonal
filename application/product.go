package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

func init(){
	govalidator.SetFieldsREquiredByDefault()
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)

type Product struct {
	ID string `valid:"uuidv4"`
	Name string `valid:required`
	Price float64 `valid:float,optional`
	Status string `valid:required`
}

func(p *Product) IsValid() (bool, error){
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("Invalid status")
	}

	if p.Price < 0 {
		return false, errors.New("Invalid price")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func(p *Product) Enable() error{
	if p.Price > 0 {
		p,Status = ENABLED
		return nil
	}
	return errors.New("Invalid price; Must be greather than zero")
}

func(p *Product) Disable() {
	if p.Price == 0 {
		p,Status = DISABLED
		return nil
	}
	return errors.New("Invalid price; Must be zero")
}

func(p *Product) GetID() string {
	return p.ID
}

func(p *Product) GetName() string {
	return p.Name
}

func(p *Product) GetStatus() string {
	return p.Status
}

func(p *Product) GetPrice() float64 {
	return p.Price
}