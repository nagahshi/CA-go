package main

import (
	"fmt"
	"net/http"

	"./controller"
	"./repository"
	"./router"
	"./service"
)

const port string = ":8002"

var (
	httpRouter     router.Router             = router.NewChiRouter()
	postRepository repository.PostRepository = repository.NewFireStoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	httpRouter.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "hello world")
	})

	httpRouter.Get("/posts", postController.GetPosts)
	httpRouter.Post("/posts", postController.AddPost)

	httpRouter.Serve(port)
}
