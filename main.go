package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/itgamerskyHuang/category/common"
	"github.com/itgamerskyHuang/category/handler"
	pb "github.com/itgamerskyHuang/category/proto/category"
	"github.com/jinzhu/gorm"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("192.168.0.201:8500", "/micro/config")
	if err != nil {
		log.Println(err)
	}
	// 注册中心
	consulregistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.0.201:8500",
		}
	})

	srv := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address(":8998"),
		// 添加consul作为注册中心
		micro.Registry(consulregistry),
	)

	// 获取mysql的配置,路径中不带前缀
	mysqlconf := common.NewMysqlConfig(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlconf.User+":"+mysqlconf.Pwd+"@tcp("+mysqlconf.Host+")/"+mysqlconf.Database+"?charset=utf8")

	// db, err := gorm.Open("mysql", "root:123456@tcp(192.168.0.201:3306)/micro?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	defer db.Close()
	db.SingularTable(true)
	// 创建数据库表
	// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Category{})

	// Register handler
	pb.RegisterCategoryHandler(srv.Server(), handler.NewCategory(db))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
