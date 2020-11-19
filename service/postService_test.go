package service

import (
	"testing"

	"../entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.(*entity.Post), args.Error(1)
}

func (mock *mockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mockRepository)

	var (
		id int64 = 1
	)
	post := entity.Post{
		Id:    id,
		Title: "Test Title",
		Text:  "Test Text",
	}

	// expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, result[0].Id, id)
	assert.Equal(t, result[0].Title, "Test Title")
	assert.Equal(t, result[0].Text, "Test Text")
}

func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)

	post := entity.Post{
		Title: "Test Title",
		Text:  "Test Text",
	}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.Id)
	assert.Equal(t, result.Title, "Test Title")
	assert.Equal(t, result.Text, "Test Text")
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "The post is empty")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	testService := NewPostService(nil)

	post := entity.Post{
		Id:    1,
		Title: "",
		Text:  "",
	}

	err := testService.Validate(&post)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "The post title is empty")
}
