package utils

const (
	// user roles
	UserRoleID   = 1
	WasherRoleID = 2
	AdminRoleID  = 3
	UserRole     = "user"
	WasherRole   = "washer"
	AdminRole    = "admin"

	// washer status
	AvailableWasherStatusID = 1
	WashingWasherStatusID   = 2
	InActiveWasherStatusID  = 3
	AvailableStatus         = "available"
	WashingStatus           = "washing"
	InActiveStatus          = "inactive"

	// washer order status
	PreparingOrderWasherStatusID = 1
	OngoingOrderWasherStatusID   = 2
	ArrivedOrderWasherStatusID   = 3
	WashingOrderWasherStatusID   = 4
	FinishedOrderWasherStatusID  = 5
)
