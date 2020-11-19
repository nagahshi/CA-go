package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"../entity"
	"../repository"
	"../service"
	"github.com/stretchr/testify/assert"
)

const (
	ID    int64  = 123
	TITLE string = "Title 1"
	TEXT  string = "Text 1"
)

var (
	postRepositoryT repository.PostRepository = repository.NewFireStoreRepository()
	postServiceT    service.PostService       = service.NewPostService(postRepositoryT)
	postControllerT PostController            = NewPostController(postServiceT)
)

func TestAddPost(t *testing.T) {
	var jsonStr = []byte(`{"title": "Title 1", "text": "Text 1"}`)

	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))

	handler := http.HandlerFunc(postControllerT.AddPost)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	statusCode := response.Code

	if statusCode != http.StatusOK {
		t.Errorf("return wrong status code: %v want %v", statusCode, http.StatusOK)
	}

	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	assert.NotNil(t, post.Id)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	cleanUp(&post)
}

func cleanUp(post *entity.Post) {
	postRepositoryT.Delete(post)
}

func TestGetPosts(t *testing.T) {
	setup()

	req, _ := http.NewRequest("GET", "/posts", nil)

	handler := http.HandlerFunc(postControllerT.GetPosts)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	statusCode := response.Code

	if statusCode != http.StatusOK {
		t.Errorf("return wrong status code: %v want %v", statusCode, http.StatusOK)
	}

	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	assert.NotNil(t, posts[0].Id)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	cleanUp(&posts[0])
}

func setup() {
	var post = entity.Post{
		Id:    ID,
		Title: TITLE,
		Text:  TEXT,
	}

	postRepositoryT.Save(&post)
}
