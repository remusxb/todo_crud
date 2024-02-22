package dto

type HealthCheckOutput struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
