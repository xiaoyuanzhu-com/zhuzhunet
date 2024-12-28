package enum

type APIStatus string

const (
	APIStatusDefault    APIStatus = "default"
	APIStatusInProgress APIStatus = "in-progress"
	APIStatusSuccess    APIStatus = "success"
	APIStatusError      APIStatus = "error"
)
