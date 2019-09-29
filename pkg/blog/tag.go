package blog

import (
	"context"

	"github.com/nomkhonwaan/myblog/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tag is a label attached to the post for the purpose of identification
type Tag struct {
	// Identifier of the tag
	ID primitive.ObjectID `bson:"_id" json:"id" graphql:"-"`

	// Name of the tag
	Name string `bson:"name" json:"name" graphql:"name"`

	// Valid URL string composes with name and ID
	Slug string `bson:"slug" json:"slug" graphql:"slug"`
}

type TagRepository interface {
	FindAllByIDs(context.Context, []primitive.ObjectID) ([]Tag, error)
}

func NewTagRepository(col mongo.Collection) MongoTagRepository {
	return MongoTagRepository{col}
}

type MongoTagRepository struct {
	col mongo.Collection
}

func (repo MongoTagRepository) FindAllByIDs(ctx context.Context, ids []primitive.ObjectID) ([]Tag, error) {
	cur, err := repo.col.Find(ctx, bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var tags []Tag
	err = cur.Decode(&tags)

	return tags, err
}
