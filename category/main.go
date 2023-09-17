package main

import (
	"category/common"
	"category/domain/repository"
	service2 "category/domain/service"
	"category/handler"
	"category/proto/category"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}

	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//new service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		//添加consul作为注册中心
		micro.Registry(consulRegistry),
	)

	//获取mysql的配置,路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	//禁止复表
	db.SingularTable(true)

	//执行一次
	//rp := repository.NewCategoryRepository(db)
	//rp.InitTable()

	//初始化
	service.Init()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))
	err = category.RegisterCatagoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	//运行
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
