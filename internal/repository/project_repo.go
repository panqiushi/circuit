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

func FindProjecptById(id uint) (*models.Project, error) {
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

func FindAll() ([]models.Project, error) {
	var projects []models.Project
	if err := db.DB().Preload("User").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
