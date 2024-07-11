package ioc

import "github.com/gin-gonic/gin"

type IocInf interface {
	Init()
	Name() string
}

type GinInf interface {
	Registry(r gin.IRouter)
}