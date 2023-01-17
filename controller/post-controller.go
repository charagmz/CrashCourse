package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/charagmz/CrashCourse/cache"
	"github.com/charagmz/CrashCourse/entity"
	"github.com/charagmz/CrashCourse/errors"
	"github.com/charagmz/CrashCourse/service"
)

type controller struct{}

var (
	postService service.PostService
	postCache   cache.PostCache
)

type PostController interface {
	GetPostById(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
	return &controller{}
}

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// Get parameter
	postID := strings.Split(r.URL.Path, "/")[2]

	// Get the post from redis
	var post *entity.Post = postCache.Get(postID)
	var err error
	// If the post isn't in redis
	if post == nil {
		// Get the post from service
		post, err = postService.FindByID(postID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errors.ServiceError{Message: "No post found"})
		}
		//log.Println("from repo get", post.ID)
		// Store the post in redis
		postCache.Set(postID, post)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	result, err := postService.Create(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
