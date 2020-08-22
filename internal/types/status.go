package types

type Status string

const (
	ACTIVE     Status = "ACTIVE"
	DEACTIVATE Status = "DEACTIVATE"
	REVOKED    Status = "REVOKED"
)
