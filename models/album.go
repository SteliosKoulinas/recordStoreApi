package models

import "gorm.io/gorm"

type Album struct {
    gorm.Model
    Artist      string `json:"artist"`
    Title       string `json:"title"`
    Year        int    `json:"year"`
}