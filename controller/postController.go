package controller

import (
	"encoding/json"
	"net/http"

	"../errors"

	"../entity"
	"../service"
)

type controller struct{}

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
}

var (
	postService service.PostService
)

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "error on get posts"})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func (*controller) AddPost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post
	res.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "error a unmarshal body data"})
		return
	}

	err = postService.Validate(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	postCreated, errToCreate := postService.Create(&post)
	if errToCreate != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error on save post"})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(postCreated)
}

func (*controller) DelePost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post
	res.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "error a unmarshal body data"})
		return
	}

	err = postService.Validate(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	postCreated, errToCreate := postService.Create(&post)
	if errToCreate != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error on save post"})
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(postCreated)
}
