package main

import (
	"category/handler"
	pb "category/proto"

	"github.com/micro/go-micro/v2"
)

func main() {
	// Create service
	// srv := service.New(
	// 	service.Name("category"),
	// 	service.Version("latest"),
	// )
	srv := micro.

		// Register handler
		pb.RegisterCategoryHandler(srv.Server(), new(handler.Category))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
