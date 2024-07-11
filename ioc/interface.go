package ioc

import "github.com/gin-gonic/gin"

type IocInf interface {
	Init()
	Name() string
}

type GinInf interface {
	Register(r gin.IRouter)
}