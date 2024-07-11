package ioc

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Ioc struct {
	container map[string]IocInf
}

func (c *Ioc) Registry(iocObj IocInf) {
	c.container[iocObj.Name()] = iocObj
}

func (c *Ioc) Get(name string) any {
	return c.container[name]
}

func (c *Ioc) InitObj() {
	fmt.Println(c.container)
	for _, obj := range c.container {
		obj.Init()
	}
}


func (c *Ioc) RouterRegistry(r gin.IRouter)  {
	for _, obj := range c.container {
		if api, ok := obj.(GinInf); ok {
			api.Registry(r)
		}
	}
}