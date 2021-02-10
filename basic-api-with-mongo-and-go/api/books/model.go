package books

import (
	"basic-api-with-mongo-and-go/db"
	"context"
	"fmt"
)

type Book struct {
	Title       string `json:"title"`
	Description string `json:"descrition"`
	ReleaseDate string `json:"release_date"`
}

func (b Book) Save() error {
	fmt.Printf("%v", b)
	coll := db.DB.Collection("books")
	result, err := coll.InsertOne(context.Background(), map[string]string{
		"title":        b.Title,
		"description":  b.Description,
		"release_date": b.ReleaseDate,
	})
	fmt.Println(result)
	return err
}
