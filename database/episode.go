package database

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/toferc/foundations/models"
)

// SaveEpisode saves a Episode to the DB
func SaveEpisode(db *pg.DB, e *models.Episode) error {

	// Save episode in Database
	_, err := db.Model(e).
		OnConflict("(id) DO UPDATE").
		Set("episode = ?episode").
		Insert(e)
	if err != nil {
		panic(err)
	}
	return err
}

// UpdateEpisode updates a runequest episode
func UpdateEpisode(db *pg.DB, e *models.Episode) error {

	err := db.Update(e)
	if err != nil {
		panic(err)
	}
	return err
}

// ListEpisodes queries Episode names and add to slice
func ListEpisodes(db *pg.DB) ([]*models.Episode, error) {
	var es []*models.Episode

	_, err := db.Query(&es, `SELECT * FROM episodes`)

	if err != nil {
		panic(err)
	}

	// Print names and PK
	for i, e := range es {

		fmt.Println(i, e.Title)
	}
	return es, nil
}

// PKLoadEpisode loads a single episode from the DB by pk
func PKLoadEpisode(db *pg.DB, pk int64) (*models.Episode, error) {
	// Select user by Primary Key
	e := &models.Episode{ID: pk}
	err := db.Select(e)

	if err != nil {
		fmt.Println(err)
		return &models.Episode{}, err
	}

	fmt.Println("Episode loaded From DB")
	return e, nil
}

// DeleteEpisode deletes a single episode from DB by ID
func DeleteEpisode(db *pg.DB, pk int64) error {

	e := models.Episode{ID: pk}

	fmt.Println("Deleting episode...")

	err := db.Delete(&e)

	return err
}
