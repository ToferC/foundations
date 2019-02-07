package database

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/toferc/foundations/models"
)

// SaveExperience saves a Experience to the DB
func SaveExperience(db *pg.DB, e *models.Experience) (int64, error) {

	// Save experience in Database
	_, err := db.Model(e).
		OnConflict("(id) DO UPDATE").
		Set("verb = ?verb").
		Insert(e)
	if err != nil {
		panic(err)
	}
	return e.ID, err
}

// UpdateExperience updates a runequest experience
func UpdateExperience(db *pg.DB, e *models.Experience) (int64, error) {

	err := db.Update(e)
	if err != nil {
		panic(err)
	}
	return e.ID, err
}

// LoadStreamExperiences queries Experience names and add to slice
func LoadStreamExperiences(db *pg.DB, name string) ([]*models.Experience, error) {
	var temp, exs []*models.Experience

	err := db.Model(&temp).
		Column("experience.*", "LearningResource").
		Select()

	if err != nil {
		panic(err)
	}

	// Add stream experiences to Slice -- Need DB work here
	for i, e := range temp {
		if e.Stream.Name == name {
			fmt.Println(i, e)
			exs = append(exs, e)
		}
	}
	return exs, nil
}

// ListUserExperiences queries Character names and add to slice
func ListUserExperiences(db *pg.DB, username string) ([]*models.Experience, error) {
	var exs []*models.Experience

	err := db.Model(&exs).
		Where("user_name = ?", username).
		Column("experience.*", "LearningResource").
		Select()

	if err != nil {
		return []*models.Experience{}, err
	}

	// Print names and PK

	return exs, nil
}

// ListExperiences queries Character names and add to slice
func ListExperiences(db *pg.DB) ([]*models.Experience, error) {
	var exs []*models.Experience

	err := db.Model(&exs).
		Column("experience.*", "LearningResource").
		Select()

	if err != nil {
		return []*models.Experience{}, err
	}

	// Print names and PK

	return exs, nil
}

// PKLoadExperience loads a single experience from the DB by pk
func PKLoadExperience(db *pg.DB, pk int64) (*models.Experience, error) {
	// Select experience by Primary Key
	e := &models.Experience{ID: pk}
	err := db.Select(e)

	if err != nil {
		fmt.Println(err)
		return &models.Experience{}, err
	}

	fmt.Println("Experience loaded From DB")
	return e, nil
}

// DeleteExperience deletes a single experience from DB by ID
func DeleteExperience(db *pg.DB, pk int64) error {

	e := models.Experience{ID: pk}

	fmt.Println("Deleting experience...")

	err := db.Delete(&e)

	return err
}
