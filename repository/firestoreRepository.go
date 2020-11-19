package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"../entity"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type repo struct{}

func NewFireStoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "gocrashcourseapi"
	collectionName string = "posts"
	path           string = "/home/billy/FIREBASE/gocrashcourseapi-firebase-adminsdk-j9i76-bb2c1356c0.json"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile(path))
	if err != nil {
		log.Fatalf("Fail to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	docRef, result, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Fail to add new post: %v", err)
		return nil, err
	}

	fmt.Println(docRef, result, err)

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile(path))
	if err != nil {
		log.Fatalf("Fail to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post

	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			log.Fatalf("Fail to iterate lists posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			Id:    doc.Data()["Id"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (*repo) Delete(post *entity.Post) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile(path))
	if err != nil {
		log.Fatalf("Fail to create a firestore client: %v", err)
		return err
	}
	defer client.Close()

	var id string = strconv.FormatInt(post.Id, 10)

	if _, err = client.Collection(collectionName).Doc(id).Delete(ctx); err != nil {
		log.Fatalf("Fail to create a firestore client: %v", err)
		return err
	}

	return nil
}
