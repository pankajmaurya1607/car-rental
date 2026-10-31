package domain

import "errors"

var (
	ErrVehicleNotFound     = errors.New("vehicle not found")
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrReservationNotFound = errors.New("reservation not found")
	ErrPaymentNotFound     = errors.New("payment not found")

	ErrVehicleNotAvailable    = errors.New("vehicle not available")
	ErrReservationConflict    = errors.New("reservation time slot conflict")
	ErrInvalidRentalPeriod    = errors.New("invalid rental period")
	ErrInvalidStateTransition = errors.New("invalid state transation")
	ErrReservationCancelled   = errors.New("resarvation is cancelled")
	ErrPaymentFailed          = errors.New("payment failed")

	ErrInvalidInput = errors.New("invalid input")
)