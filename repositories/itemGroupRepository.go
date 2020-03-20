package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type ItemGroupRepository struct {
	*mongo.Collection
}

// NewItemGroupRepositoryRepository creates a new ItemGroupRepository
func NewItemGroupRepository(db *mongo.Database) (*ItemGroupRepository, error) {
	_, err := db.Collection("itemGroups").Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"name", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		return nil, err
	}

	return &ItemGroupRepository{
		Collection: db.Collection("itemGroups"),
	}, nil
}
