package webmaticlib

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Headless    bool   `json:"headless"`
	XMLData     string `json:"xml_data"`
}
