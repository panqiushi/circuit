package repository

import (
	"circuit.io/circuit/internal/models"
	"circuit.io/circuit/internal/models/db"
)

func CreateProject(project *models.Project) (*models.Project, error) {
	if err := db.DB().Create(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func FindProjectById(id uint) (*models.Project, error) {
	var project models.Project
	if err := db.DB().First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func FindProjectByName(name string) (*models.Project, error) {
	var project models.Project
	if err := db.DB().Where("name= ?", name).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func FindAllProject() ([]models.Project, error) {
	var projects []models.Project
	if err := db.DB().Preload("User").Preload("Users").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func AddMemberToProject(project *models.Project, user *models.User) (*models.Project, error) {
	if err := db.DB().Model(project).Association("Users").Append(user); err != nil {
		return nil, err
	}
	return project, nil
}

func RemoveMemberFromProject(project *models.Project, user *models.User) (*models.Project, error) {
	if err := db.DB().Model(project).Association("Users").Delete(user); err != nil {
		return nil, err
	}
	return project, nil
}

func AddVersionReleaseToProject(project *models.Project, version_release *models.VersionRelease) (*models.Project, error) {
	if err := db.DB().Model(project).Association("VersionReleases").Append(version_release); err != nil {
		return nil, err
	}
	return project, nil
}

func RemoveVersionReleaseFromProject(project *models.Project, version_release *models.VersionRelease) (*models.Project, error) {
	if err := db.DB().Model(project).Association("VersionReleases").Delete(version_release); err != nil {
		return nil, err
	}
	return project, nil
}
