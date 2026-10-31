package domain

type VehicleType string

const (
	VehicleTypeHatchback VehicleType = "HATCHBACK"
	VehicleTypeSedan     VehicleType = "SEDAN"
	VehicleTypeSUV       VehicleType = "SUV"
	VehicleTypeLuxury    VehicleType = "Luxury"
)

type VehicleStatus string

const (
	VehicleStatusAvailable   VehicleStatus = "AVAILABLE"
	VehicleStatusRented      VehicleStatus = "RENTED"
	VehicleStatusMaintenance VehicleStatus = "Maintenance"
)

type ReservationStatus string

const (
	ReservationStatusConfirmed ReservationStatus = "CONFIRMED"
	ReservationStatusActive    ReservationStatus = "ACTIVE"
	ReservationStatusCompleted ReservationStatus = "COMPLETED"
	ReservationStatusCancelled ReservationStatus = "CANCELLED"
)

type PaymentStatus string

const (
	PaymentSatusPending  PaymentStatus = "PENDING"
	PaymentStatusSuccess PaymentStatus = "SUCCESS"
	PaymentStatusFailed  PaymentStatus = "FAILED"
)
