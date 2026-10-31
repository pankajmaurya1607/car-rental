package domain

import (
	"errors"
	"strings"
)

type Vehicle struct {
	ID             string
	RegistrationNo string
	Brand          string
	Model          string
	Type           VehicleType
	Year           int64
	PricePerHour   float64
	Status         VehicleStatus
}

func NewVehicle(
	id string,
	registrationNo string,
	brand string,
	model string,
	vehicleType VehicleType,
	year int64,
	pricePerHour float64,
) (*Vehicle, error) {

	if strings.TrimSpace(id) == "" ||
		strings.TrimSpace(registrationNo) == "" ||
		strings.TrimSpace(brand) == "" ||
		strings.TrimSpace(model) == "" {
		return nil, ErrInvalidInput
	}

	if year <= 0 || pricePerHour <= 0 {
		return nil, ErrInvalidInput
	}

	if !isValidVehicleType(vehicleType) {
		return nil, ErrInvalidInput
	}

	return &Vehicle{
		ID:             id,
		RegistrationNo: registrationNo,
		Brand:          brand,
		Model:          model,
		Type:           vehicleType,
		Year:           year,
		PricePerHour:   pricePerHour,
		Status:         VehicleStatusAvailable,
	}, nil
}

func isValidVehicleType(vehicleType VehicleType) bool {
	switch vehicleType {
	case VehicleTypeHatchback,
		VehicleTypeSedan,
		VehicleTypeSUV,
		VehicleTypeLuxury:
		return true
	default:
		return false
	}
}

func (v *Vehicle) MarkAsRented() error {
	if v.Status != VehicleStatusAvailable {
		return ErrVehicleNotAvailable
	}

	v.Status = VehicleStatusRented

	return nil
}

func (v *Vehicle) MarkAsAvailable() error {
	if v.Status != VehicleStatusRented {
		return ErrInvalidStateTransition
	}

	v.Status = VehicleStatusAvailable

	return nil
}

func (v *Vehicle) MarkAsMaintenance() error {
	if v.Status == VehicleStatusRented {
		return errors.New("cannot send rented vehicle to maintenance")
	}

	v.Status = VehicleStatusMaintenance

	return nil
}
