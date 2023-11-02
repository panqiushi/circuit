package models

import "gorm.io/gorm"

type VersionRelease struct {
	gorm.Model
	ID          uint   `gorm:"primary_key;auto_increment;column:id" json:"id"`
	VersionName string `gorm:"unique;column:version_name" json:"version_name" form:"version_name"`
	ReleaseDate string `gorm:"column:release_date" json:"release_date" form:"release_date"`
	ProjectID   uint   `gorm:"foreignkey:ProjectID;column:project_id" json:"project_id"`
	CommitID    string `gorm:"unique;column:commit_id" json:"commit_id"`
	Branch      string `gorm:"column:branch" json:"branch" form:"branch"`
	TagName     string `gorm:"unique;column:tag_name" json:"tag_name" form:"tag_name"`
	UserId      uint   `gorm:"foreignkey:UserId;column:user_id" json:"user_id"`
	ReleaseNote string `gorm:"column:release_note" json:"release_note" form:"release_note"`
	Artifactory string `gorm:"column:artifactory" json:"artifactory" form:"artifactory"`
	Project     Project
	User        User
}
