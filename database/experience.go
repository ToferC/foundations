package database

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/toferc/foundations/models"
)

// SaveExperience saves a Experience to the DB
func SaveExperience(db *pg.DB, e *models.Experience) error {

	// Save experience in Database
	_, err := db.Model(e).
		OnConflict("(id) DO UPDATE").
		Set("verb = ?verb").
		Insert(e)
	if err != nil {
		panic(err)
	}
	return err
}

// UpdateExperience updates a runequest experience
func UpdateExperience(db *pg.DB, e *models.Experience) error {

	err := db.Update(e)
	if err != nil {
		panic(err)
	}
	return err
}

// ListExperiences queries Experience names and add to slice
func ListExperiences(db *pg.DB) ([]*models.Experience, error) {
	var es []*models.Experience

	_, err := db.Query(&es, `SELECT * FROM experiences`)

	if err != nil {
		panic(err)
	}

	// Print names and PK
	for i, e := range es {

		fmt.Println(i, e)
	}
	return es, nil
}

// ListUserExperiences queries Character names and add to slice
func ListUserExperiences(db *pg.DB, username string) ([]*models.Experience, error) {
	var exs []*models.Experience

	_, err := db.Query(&exs, `SELECT * FROM experiences WHERE user_name = ?`, username)

	if err != nil {
		return []*models.Experience{}, err
	}

	// Print names and PK
	for i, ex := range exs {

		fmt.Println(i, ex)
	}
	return exs, nil
}

// PKLoadExperience loads a single experience from the DB by pk
func PKLoadExperience(db *pg.DB, pk int64) (*models.Experience, error) {
	// Select user by Primary Key
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
