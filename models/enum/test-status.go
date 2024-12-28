package enum

type TestStatus string

const (
	TestStatusDefault   TestStatus = "default"
	TestStatusPending   TestStatus = "pending"
	TestStatusRunning   TestStatus = "running"
	TestStatusCompleted TestStatus = "completed"
	TestStatusFailed    TestStatus = "failed"
)
