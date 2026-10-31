package domain

import "time"

type Payment struct {
	ID            string
	ReservationID string
	Amount        int64 // Amount in paisa
	Status        PaymentStatus
	TransactionID string
	CreatedAt     time.Time
}

func NewPayment(
	id string,
	reservationID string,
	amount int64,
) (*Payment, error) {
	if id == "" || reservationID == "" || amount <= 0 {
		return nil, ErrInvalidInput
	}

	return &Payment{
		ID:            id,
		ReservationID: reservationID,
		Amount:        amount,
		Status:        PaymentSatusPending,
		CreatedAt:     time.Now(),
	}, nil
}

func (p *Payment) MarkAsSuccessful(transactionID string) error {
	if p.Status != PaymentSatusPending || transactionID == "" {
		return ErrInvalidStateTransition
	}

	p.TransactionID = transactionID
	p.Status = PaymentStatusSuccess

	return nil
}

func (p *Payment) MarkAsFailed() error {
	if p.Status != PaymentSatusPending {
		return ErrInvalidStateTransition
	}

	p.Status = PaymentStatusFailed

	return nil
}
