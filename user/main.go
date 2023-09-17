package main

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"user/domain/repository"
	service2 "user/domain/service"
	"user/handler"
	"user/proto/user"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	srv.Init()

	//创建数据库连接
	db, err := gorm.Open("mysql", "xzx:xzx527416@tcp(111.229.202.109:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//只执行一次
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}
	// Register handler

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
