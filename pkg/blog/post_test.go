package blog_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/nomkhonwaan/myblog/pkg/blog"
	"github.com/nomkhonwaan/myblog/pkg/mongo"
	mock_mongo "github.com/nomkhonwaan/myblog/pkg/mongo/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestPost_MarshalJSON(t *testing.T) {
	// Given
	id := primitive.NewObjectID()
	createdAt := time.Now()
	post := Post{
		ID:          id,
		Title:       "Children of Dune",
		Slug:        "children-of-dune-" + id.Hex(),
		Status:      Draft,
		Markdown:    "Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit. Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem.",
		HTML:        "Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh.",
		PublishedAt: time.Time{},
		AuthorID:    "github|c7834cb0-2b79-4d27-a817-520a6420c11b",
		Categories:  []mongo.DBRef{},
		Tags:        []mongo.DBRef{},
		CreatedAt:   createdAt,
		UpdatedAt:   time.Time{},
	}

	// When
	result, err := json.Marshal(post)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, "{\"id\":\""+id.Hex()+"\",\"title\":\"Children of Dune\",\"slug\":\"children-of-dune-"+id.Hex()+"\",\"status\":\"DRAFT\",\"markdown\":\"Integer tincidunt ante vel ipsum. Praesent blandit lacinia erat. Vestibulum sed magna at nunc commodo placerat. Praesent blandit. Nam nulla. Integer pede justo, lacinia eget, tincidunt eget, tempus vel, pede. Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem.\",\"html\":\"Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh.\",\"publishedAt\":\"0001-01-01T00:00:00Z\",\"authorId\":\"github|c7834cb0-2b79-4d27-a817-520a6420c11b\",\"createdAt\":\""+createdAt.Format(time.RFC3339Nano)+"\",\"updatedAt\":\"0001-01-01T00:00:00Z\"}", string(result))
}

func TestMongoPostRepository_Create(t *testing.T) {
	t.Run("When insert into the collection successfully", func(t *testing.T) {
		// Given
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		col := mock_mongo.NewMockCollection(ctrl)
		ctx := context.Background()
		authorID := "github|303589"

		col.EXPECT().InsertOne(ctx, gomock.Any()).Return(&mgo.InsertOneResult{}, nil)

		postRepo := NewPostRepository(col)

		// When
		result, err := postRepo.Create(ctx, authorID)

		// Then
		assert.Nil(t, err)
		assert.True(t, time.Since(result.CreatedAt) < time.Minute)
		assert.Equal(t, fmt.Sprintf("%s", result.ID.Hex()), result.Slug)
		assert.Equal(t, Draft, result.Status)
		assert.Equal(t, authorID, result.AuthorID)
	})

	t.Run("When insert into the collection un-successfully", func(t *testing.T) {
		// Given
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		col := mock_mongo.NewMockCollection(ctrl)
		ctx := context.Background()
		authorID := "github|303589"

		col.EXPECT().InsertOne(ctx, gomock.Any()).Return(&mgo.InsertOneResult{}, errors.New("something went wrong"))

		postRepo := NewPostRepository(col)

		expected := Post{}

		//result When
		result, err := postRepo.Create(ctx, authorID)

		// Then
		assert.EqualError(t, err, "something went wrong")
		assert.Equal(t, expected, result)
	})
}

func TestMongoPostRepository_FindAll(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cur := mock_mongo.NewMockCursor(ctrl)
	col := mock_mongo.NewMockCollection(ctrl)
	ctx := context.Background()

	repo := NewPostRepository(col)

	tests := map[string]struct {
		q       PostQuery
		filter  interface{}
		options func() *options.FindOptions
		err     error
	}{
		"With default query options": {
			q:      NewPostQueryBuilder().Build(),
			filter: bson.M{},
			options: func() *options.FindOptions {
				return (&options.FindOptions{}).SetSkip(0).SetLimit(5)
			},
		},
		"With specified offset and limit": {
			q:      NewPostQueryBuilder().WithOffset(10).WithLimit(5).Build(),
			filter: bson.M{},
			options: func() *options.FindOptions {
				return (&options.FindOptions{}).SetSkip(10).SetLimit(5)
			},
		},
		"With status draft": {
			q:      NewPostQueryBuilder().WithStatus(Draft).Build(),
			filter: bson.M{"status": Draft},
			options: func() *options.FindOptions {
				return (&options.FindOptions{}).SetSkip(0).SetLimit(5)
			},
		},
		"With status published": {
			q:      NewPostQueryBuilder().WithStatus(Published).Build(),
			filter: bson.M{"status": Published},
			options: func() *options.FindOptions {
				options := (&options.FindOptions{}).SetSkip(0).SetLimit(5)
				options.Sort = map[string]interface{}{
					"publishedAt": -1,
				}
				return options
			},
		},
		"When an error has occurred while finding the result": {
			q:      &MongoPostQuery{},
			filter: bson.M{},
			options: func() *options.FindOptions {
				return (&options.FindOptions{}).SetSkip(0).SetLimit(0)
			},
			err: errors.New("something went wrong"),
		},
	}

	// When
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			col.EXPECT().Find(ctx, test.filter, test.options()).Return(cur, test.err)

			if test.err == nil {
				cur.EXPECT().Close(ctx).Return(nil)
				cur.EXPECT().Decode(gomock.Any()).Return(nil)

				_, err := repo.FindAll(ctx, test.q)
				assert.Nil(t, err)
			} else {
				_, err := repo.FindAll(ctx, test.q)
				assert.EqualError(t, err, test.err.Error())
			}
		})
	}

	// Then
}

func TestMongoPostRepository_FindByID(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	singleResult := mock_mongo.NewMockSingleResult(ctrl)
	col := mock_mongo.NewMockCollection(ctrl)

	ctx := context.Background()
	repo := NewPostRepository(col)

	tests := map[string]struct {
		id  interface{}
		err error
	}{
		"With existing post ID": {
			id: primitive.NewObjectID(),
		},
		"When an error has occurred while finding the result": {
			id:  primitive.NewObjectID(),
			err: errors.New("test find by ID error"),
		},
	}

	// When
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			col.EXPECT().FindOne(ctx, bson.M{"_id": test.id.(primitive.ObjectID)}, gomock.Any()).Return(singleResult)
			singleResult.EXPECT().Decode(gomock.Any()).Return(test.err)

			if test.err == nil {
				_, err := repo.FindByID(ctx, test.id)
				assert.Nil(t, err)
			} else {
				_, err := repo.FindByID(ctx, test.id)
				assert.EqualError(t, err, test.err.Error())
			}
		})
	}

	// Then
}

func TestMongoPostRepository_Save(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	singleResult := mock_mongo.NewMockSingleResult(ctrl)
	col := mock_mongo.NewMockCollection(ctrl)

	ctx := context.Background()
	repo := NewPostRepository(col)
	tagID := primitive.NewObjectID()

	tests := map[string]struct {
		q      PostQuery
		id     interface{}
		update interface{}
		err    error
	}{
		"With default query options": {
			q:      NewPostQueryBuilder().Build(),
			id:     primitive.NewObjectID(),
			update: bson.M{"$set": bson.M{}},
		},
		"When updating post's title": {
			q:      NewPostQueryBuilder().WithTitle("Test update post title").Build(),
			id:     primitive.NewObjectID(),
			update: bson.M{"$set": bson.M{"title": "Test update post title"}},
		},
		"When updating post's content": {
			q:      NewPostQueryBuilder().WithMarkdown("Test update post content").WithHTML("<p>Test update post content</p>").Build(),
			id:     primitive.NewObjectID(),
			update: bson.M{"$set": bson.M{"markdown": "Test update post content", "html": "<p>Test update post content</p>"}},
		},
		"When updating post's tags": {
			q:      NewPostQueryBuilder().WithTags([]Tag{{ID: tagID, Name: "Blog", Slug: "blog-" + tagID.Hex()}}).Build(),
			id:     primitive.NewObjectID(),
			update: bson.M{"$set": bson.M{"tags": bson.A{mongo.DBRef{Ref: "tags", ID: tagID}}}},
		},
		"When an error has occurred while updating the post": {
			q:      NewPostQueryBuilder().Build(),
			id:     primitive.NewObjectID(),
			update: bson.M{"$set": bson.M{}},
			err:    errors.New("something went wrong"),
		},
	}

	// When
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			col.EXPECT().UpdateOne(ctx, bson.M{"_id": test.id.(primitive.ObjectID)}, test.update).Return(nil, test.err)

			if test.err == nil {
				col.EXPECT().FindOne(ctx, bson.M{"_id": test.id.(primitive.ObjectID)}).Return(singleResult)
				singleResult.EXPECT().Decode(gomock.Any()).Return(nil)

				_, err := repo.Save(ctx, test.id, test.q)
				assert.Nil(t, err)
			} else {
				_, err := repo.Save(ctx, test.id, test.q)
				assert.EqualError(t, err, test.err.Error())
			}
		})
	}
}

func TestNewPostQueryBuilder(t *testing.T) {
	// Given

	// When
	qb := NewPostQueryBuilder()

	// Then
	q := qb.Build()
	assert.Equal(t, int64(0), q.Offset())
	assert.Equal(t, int64(5), q.Limit())
}

func TestPostQueryBuilder_WithStatus(t *testing.T) {
	// Given

	// When
	qb := NewPostQueryBuilder().WithStatus(Published)

	// Then
	q := qb.Build()
	assert.Equal(t, Published, q.Status())
}

func TestPostQueryBuilder_WithOffset(t *testing.T) {
	// Given

	// When
	qb := NewPostQueryBuilder().WithOffset(99)

	// Then
	q := qb.Build()
	assert.Equal(t, int64(99), q.Offset())
}

func TestPostQueryBuilder_WithLimit(t *testing.T) {
	// Given

	// When
	qb := NewPostQueryBuilder().WithLimit(99)

	// Then
	q := qb.Build()
	assert.Equal(t, int64(99), q.Limit())
}

func TestPostQueryBuilder_Build(t *testing.T) {
	// Given
	qb := NewPostQueryBuilder()

	// When
	q := qb.Build()

	// Then
	assert.Equal(t, int64(0), q.Offset())
	assert.Equal(t, int64(5), q.Limit())
}

//func TestPostQuery_Status(t *testing.T) {
//
//}
//
//func TestPostQuery_Offset(t *testing.T) {
//
//}
//
//func TestPostQuery_Limit(t *testing.T) {
//
//}
