package service

import (
	"testing"

	"github.com/charagmz/CrashCourse/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Create an structure "MockRepository", this mock repository structure is going to implement te post repository interface
type MockRepository struct {
	mock.Mock //create an structure provided by the Testify framework
}

// Implement the interface "PostReposity" for our "MockRepository"
func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	//Stub the function returning the arguments that we receive
	args := mock.Called()                           //this is going to get the arguments by the mock
	resultPost := args.Get(0)                       //the first argument is the post
	return resultPost.(*entity.Post), args.Error(1) //Do type assertion for the post, get the args parameter for error
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	//Stub the function returning the arguments that we receive
	args := mock.Called()                             //this is going to get the arguments by the mock
	resultPosts := args.Get(0)                        //the first argument is the post's array
	return resultPosts.([]entity.Post), args.Error(1) //Do type assertion for the array, get the args parameter for error
}

func (mock *MockRepository) Delete(post *entity.Post) error {
	//Stub the function returning the arguments that we receive
	args := mock.Called() //this is going to get the arguments by the mock
	//resultPost := args.Get(0)                         //the first argument is the post
	return args.Error(1) //Get the args parameter for error
}

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

// To call FindAll(), create an expectation on the mock repository
func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository) //Create a new reference to the mock repository

	var identifier int64 = 1

	post := entity.Post{ID: identifier, Title: "A", Text: "B"}

	//setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil) //when the FindAll methos is invoked on this mock repo is going to return an array, and nil as the error

	//create a service with the mockRepo
	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	//create an assertion on the expectation
	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository) //Create a new reference to the mock repository

	post := entity.Post{Title: "A", Text: "B"} //Not assign the identifier because it is generated in the Create method

	mockRepo.On("Save").Return(&post, nil)

	//create a service with the mockRepo
	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	//create an assertion on the expectation
	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}
