package models

import "time"

type DataAccelerationConfig struct {
	AccelerationEnabled *string    `json:"accelerationEnabled,omitempty"`
	LastUpdatedAt       *time.Time `json:"lastUpdatedAt,omitempty"`
	AccelerationStatus  *string    `json:"accelerationStatus,omitempty"`
}
