package store

import (
	"context"
	"mangafox/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ChapterCollection = "chapter"

func (store Store) CreateChapter(chapter models.Chapter) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	result, err := store.db.Collection(ChapterCollection).InsertOne(ctx, chapter)
	if err != nil {
		return primitive.NewObjectID(), err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return objectID, err
	}

	return primitive.NewObjectID(), err
}

func (store Store) FindChapterBySourceAndNumber(source string, number float64, language string) (models.Chapter, error) {
	ctx, cancel := context.WithTimeout(store.context, 30*time.Second)
	defer cancel()

	var result models.Chapter
	filter := bson.M{
		"source":   source,
		"number":   number,
		"language": language,
	}
	err := store.db.Collection(ChapterCollection).FindOne(ctx, filter).Decode(&result)
	return result, err
}
