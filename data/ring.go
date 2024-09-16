package data

import "gorm.io/gorm"

type WebsiteType struct {
	gorm.Model
	Name string
}

type Website struct {
	gorm.Model
	Name        string
	URL         string
	Description string
	Type        []*WebsiteType
}

var Websites []*Website
