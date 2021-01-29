package main

import (
	"github.com/itgamerskyHuang/category/handler"

	pb "github.com/itgamerskyHuang/category/proto/category"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	// Create service
	// srv := service.New(
	// 	service.Name("category"),
	// 	service.Version("latest"),
	// )
	srv := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Address(":8998"),
	)

	// Register handler
	pb.RegisterCategoryHandler(srv.Server(), new(handler.Category))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
