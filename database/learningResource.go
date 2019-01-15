package database

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/toferc/foundations/models"
)

// SaveLearningResource saves a LearningResource to the DB
func SaveLearningResource(db *pg.DB, lr *models.LearningResource) error {

	fmt.Println("Save learningresource in Database")
	_, err := db.Model(lr).
		OnConflict("(id) DO UPDATE").
		Set("title = ?title").
		Insert(lr)
	if err != nil {
		panic(err)
	}
	return err
}

// UpdateLearningResource updates a learningresource
func UpdateLearningResource(db *pg.DB, lr *models.LearningResource) error {

	err := db.Update(lr)
	if err != nil {
		panic(err)
	}
	return err
}

// ListLearningResources queries LearningResource names and add to slice
func ListLearningResources(db *pg.DB) ([]*models.LearningResource, error) {
	var lrs []*models.LearningResource

	_, err := db.Query(&lrs, `SELECT * FROM learningresources`)

	if err != nil {
		panic(err)
	}

	// Print names and PK
	for i, lr := range lrs {

		fmt.Println(i, lr)
	}
	return lrs, nil
}

// LearningResourceExists returns true if LR already exists in database
func LearningResourceExists(db *pg.DB, path string) bool {
	fmt.Println("Checking Learning Resource " + path)

	var lr models.LearningResource

	exists, err := db.Model(&lr).Where("path = ?", path).Exists()
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Println(err)
	}
	return exists

}

// ListUserLearningResources queries User's LR and add to slice
func ListUserLearningResources(db *pg.DB, username string) ([]*models.LearningResource, error) {
	var lrs []*models.LearningResource

	_, err := db.Query(&lrs, `SELECT * FROM learning_resources WHERE author = ?`, username)

	if err != nil {
		return []*models.LearningResource{}, err
	}

	// Print names and PK
	for i, lr := range lrs {

		fmt.Println(i, lr)
	}
	return lrs, nil
}

//LoadLearningResource will load a learningresource from path
func LoadLearningResource(db *pg.DB, path string) (*models.LearningResource, error) {
	lr := new(models.LearningResource)

	fmt.Println("Loading Learning Resource " + path)
	err := db.Model(lr).
		Where("path = ?", path).
		Limit(1).
		Select()
	if err != nil {
		fmt.Println(err)
		return new(models.LearningResource), err
	}
	return lr, nil
}

// PKLoadLearningResource loads a single learningresource from the DB by pk
func PKLoadLearningResource(db *pg.DB, pk int64) (*models.LearningResource, error) {
	// Select user by Primary Key
	lr := &models.LearningResource{ID: pk}
	err := db.Select(lr)

	if err != nil {
		fmt.Println(err)
		return &models.LearningResource{}, err
	}

	fmt.Println("LearningResource loaded From DB")
	return lr, nil
}

// DeleteLearningResource deletes a single learningresource from DB by ID
func DeleteLearningResource(db *pg.DB, pk int64) error {

	lr := models.LearningResource{ID: pk}

	fmt.Println("Deleting learningresource...")

	err := db.Delete(&lr)

	return err
}
