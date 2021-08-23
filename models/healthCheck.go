package models

import "time"

type HealthCheck struct {
	Status string
	Time   time.Time
}
