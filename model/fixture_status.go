package model

const (
	StatusTypeInPlay = "IP"
	StatusTypeFinished = "FI"
	StatusTypeScheduled = "SC"
)

type FixtureStatus struct {
	Id string `gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
}