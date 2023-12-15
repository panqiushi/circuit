package repository

import (
	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/models/db"
)

func CreateVersionRelease(version_release *models.VersionRelease) (*models.VersionRelease, error) {
	if err := db.DB().Create(version_release).Error; err != nil {
		return nil, err
	}
	return version_release, nil
}

func FindVersionReleaseById(id uint) (*models.VersionRelease, error) {
	var version_release models.VersionRelease
	if err := db.DB().First(&version_release, id).Error; err != nil {
		return nil, err
	}
	return &version_release, nil
}

func FindVersionReleaseByName(name string) (*models.VersionRelease, error) {
	var version_release models.VersionRelease
	if err := db.DB().Where("version_name= ?", name).First(&version_release).Error; err != nil {
		return nil, err
	}
	return &version_release, nil
}

func FindVersionReleaseByProjectId(project_id uint) ([]models.VersionRelease, error) {
	var version_releases []models.VersionRelease
	if err := db.DB().Preload("Project").Where("project_id= ?", project_id).Find(&version_releases).Error; err != nil {
		return nil, err
	}
	return version_releases, nil
}
