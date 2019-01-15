package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/toferc/foundations/models"
)

// InitDB initializes the DB Schema
func InitDB(db *pg.DB) error {
	err := createSchema(db)
	if err != nil {
		panic(err)
	}
	return err
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		(*models.Episode)(nil),
		(*models.Image)(nil),
		(*models.Video)(nil),
		(*models.Podcast)(nil),
		(*models.Practice)(nil),
		(*models.Experience)(nil),
		(*models.InterestMap)(nil),
		(*models.Skill)(nil),
		(*models.Stream)(nil),
		(*models.LearningPoint)(nil),
		(*models.KnowledgePoint)(nil),
		(*models.LearningResource)(nil),
		(*models.User)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
