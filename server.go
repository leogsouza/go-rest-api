package main

import (
	"fmt"
	"net/http"

	router "github.com/leogsouza/go-rest-api/http"
	"github.com/leogsouza/go-rest-api/repository"
	"github.com/leogsouza/go-rest-api/service"

	"github.com/leogsouza/go-rest-api/controller"
)

var (
	httpRouter router.Router = router.NewChiRouter()
)

func main() {
	const port string = ":8085"
	repo := repository.NewSQLiteRepository()
	serv := service.NewPostService(repo)
	carService := service.NewCarDetailsService()

	postController := controller.NewPostController(serv)
	carController := controller.NewCarController(carService)

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running....")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.GET("/car-details", carController.GetCarDetails)

	httpRouter.SERVE(port)
}
