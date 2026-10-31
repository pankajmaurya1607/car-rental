package domain

import "strings"

type Customer struct {
	ID        string
	Name      string
	Email     string
	Phone     string
	LicenseNo string
}

func NewCustomer(
	id string,
	name string,
	email string,
	phone string,
	licensrNo string,
) (*Customer, error) {

	if strings.TrimSpace(id) == "" ||
		strings.TrimSpace(name) == "" ||
		strings.TrimSpace(email) == "" ||
		strings.TrimSpace(phone) == "" ||
		strings.TrimSpace(licensrNo) == "" {
		return nil, ErrInvalidInput
	}

	return &Customer{
		ID:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		LicenseNo: licensrNo,
	}, nil
}