package service

import (
	"CrashCourse/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil) //in this case repo is nill because isn't necesary

	err := testService.Validate(nil) //test with an empty post

	assert.NotNil(t, err)                             //Should exist an error
	assert.Equal(t, "The post is empty", err.Error()) //compare the error message
}

func TestValidateEmptyTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: ""}
	testService := NewPostService(nil) //in this case repo is nill because isn't necesary

	err := testService.Validate(&post) //test with an empty post

	assert.NotNil(t, err)                                   //Should exist an error
	assert.Equal(t, "The post title is empty", err.Error()) //compare the error message
}

func TestValidateValidPost(t *testing.T) {
	post := entity.Post{ID: 1, Title: "Title 1", Text: "Text 1"}
	testService := NewPostService(nil) //in this case repo is nill because isn't necesary

	err := testService.Validate(&post) //test with an empty post

	assert.Nil(t, err) //Shouldn't exist an error
}
