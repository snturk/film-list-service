package service

import (
	"context"
	"filmListService/models"
	"filmListService/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var filmCollection = repository.GetCollection(repository.DB, "films")

func GetAllFilms(ctx context.Context) []models.Film {
	var films []models.Film

	cursor, err := filmCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var film models.Film
		err := cursor.Decode(&film)
		if err != nil {
			panic(err)
		}
		films = append(films, film)
	}

	return films
}

func InsertFilm(ctx context.Context, film models.Film) models.Film {
	_, err := filmCollection.InsertOne(ctx, film)
	if err != nil {
		panic(err)
	}

	return film
}

func DeleteFilm(ctx context.Context, film models.Film) {
	_, err := filmCollection.DeleteOne(ctx, film)
	if err != nil {
		panic(err)
	}
}

