package main

import (
	"fmt"
	"os"

	_ "github.com/ebisu2018/xBlog/apps"
	"github.com/ebisu2018/xBlog/apps/token/api"
	tokenImpl "github.com/ebisu2018/xBlog/apps/token/impl"
	"github.com/ebisu2018/xBlog/apps/user/impl"
	"github.com/ebisu2018/xBlog/common"
	"github.com/ebisu2018/xBlog/config"
	"github.com/ebisu2018/xBlog/ioc"
	"github.com/gin-gonic/gin"
)


func nonIocVersion()  {
	userSvc := impl.NewUserServiceImpl()
	tkSvc := tokenImpl.NewTokenServiceImpl(userSvc)
	apiHandler := api.NewHttpApiHander(tkSvc)
	httpAddr := config.ReadConfig().HttpApi.HttpEndpoint()
	
	client := gin.Default()
	apiHandler.Register(client.Group(fmt.Sprintf("%v", common.API)))
	fmt.Printf("HTTP API Addr: %v", httpAddr)
	fmt.Println(client.Run(httpAddr))
}


func iocVersion() {
	ioc.Container().InitObj()
	ioc.ApiHandler().InitObj()

	client := gin.Default()
	ioc.ApiHandler().RouterRegistry(client.Group(fmt.Sprintf("%v", common.API)))

	httpAddr := config.ReadConfig().HttpApi.HttpEndpoint()
	fmt.Printf("HTTP API Addr: %v", httpAddr)
	fmt.Println(client.Run(httpAddr))
}

func main() {
	err := config.LoadFromTomlFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// nonIocVersion()
	iocVersion()
}