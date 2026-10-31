package domain

import "time"

type Reservation struct {
	ID         string
	CustomerID string
	VehicleID  string

	StartTime time.Time
	EndTime   time.Time

	ActualPichupTime *time.Time
	ActualReturnTime *time.Time

	QuotedAmount float64
	Status       ReservationStatus

	CreatedAt time.Time
}

func NewReservation(
	id string,
	customerID string,
	vehicleID string,
	startTime time.Time,
	endTime time.Time,
	quotedAmount float64,
) (*Reservation, error) {

	if id == "" || customerID == "" || vehicleID == "" {
		return nil, ErrInvalidInput
	}

	if !endTime.After(startTime) {
		return nil, ErrInvalidRentalPeriod
	}

	if quotedAmount <= 0 {
		return nil, ErrInvalidInput
	}

	return &Reservation{
		ID:           id,
		CustomerID:   customerID,
		VehicleID:    vehicleID,
		StartTime:    startTime,
		EndTime:      endTime,
		QuotedAmount: quotedAmount,
		Status:       ReservationStatusConfirmed,
		CreatedAt:    time.Now(),
	}, nil
}

func (r *Reservation) Activate(pickupTime time.Time) error {
	if r.Status != ReservationStatusConfirmed {
		return ErrInvalidStateTransition
	}

	if !pickupTime.Before(r.EndTime) {
		return ErrInvalidRentalPeriod
	}

	r.ActualPichupTime = &pickupTime
	r.Status = ReservationStatusActive

	return nil
}

func (r *Reservation) Completed(returnTime time.Time) error {
	if r.Status != ReservationStatusActive {
		return ErrInvalidStateTransition
	}

	if r.ActualPichupTime == nil || !returnTime.After(*r.ActualPichupTime) {
		return ErrInvalidRentalPeriod
	}

	r.ActualReturnTime = &returnTime
	r.Status = ReservationStatusCompleted

	return nil
}

func (r *Reservation) Cancel() error {
	if r.Status != ReservationStatusConfirmed {
		return ErrInvalidStateTransition
	}

	r.Status = ReservationStatusCancelled

	return nil
}

func (r *Reservation) HasOverlap(startTime, endTime time.Time) bool {
	return startTime.Before(r.EndTime) && r.StartTime.Before(endTime)
}	