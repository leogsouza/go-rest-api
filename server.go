package main

import (
	"fmt"
	"net/http"

	router "github.com/leogsouza/go-rest-api/http"

	"github.com/leogsouza/go-rest-api/controller"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8085"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running....")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
