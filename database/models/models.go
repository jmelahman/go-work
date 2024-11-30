package models

import (
	"time"
)

type TaskClassification int

const (
	Break TaskClassification = iota
	Chore
	Toil
	Work
)

func (tc TaskClassification) String() string {
	return [...]string{"Break", "Chore", "Toil", "Work"}[tc]
}

type Task struct {
	ID             int                `json:"id"`
	Description    string             `json:"description"`
	Classification TaskClassification `json:"classification"`
	Start          time.Time          `json:"start"`
	End            time.Time          `json:"end"`
}
