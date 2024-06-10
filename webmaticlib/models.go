package webmaticlib

import "gorm.io/gorm"

type ProjectMap struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	FileName string `json:"file_name"`
}
