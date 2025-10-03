package model

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Status  string   `json:"status"` // export type Status = "Open" | "Blocked" | "Failed" | "Passed";
	Title   string   `json:"title"`
	Details []string `json:"details"`
}
