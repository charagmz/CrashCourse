package main

import (
	"os"

	"github.com/charagmz/CrashCourse/cache"
	"github.com/charagmz/CrashCourse/controller"
	router "github.com/charagmz/CrashCourse/http"
	"github.com/charagmz/CrashCourse/repository"
	"github.com/charagmz/CrashCourse/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {

	//const port string = ":8000"
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostById)
	httpRouter.POST("/posts", postController.AddPost)

	//httpRouter.SERVE(port)
	httpRouter.SERVE(os.Getenv("PORT"))
}
