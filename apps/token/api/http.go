package api

/*
gin常用方式

gin.context是结构体，封装了http的request和response
和context.Context()不一样

1. 先用bindjson把body数据转换成对象取出到结构体
2. 调用controller的方法执行
3. 调用JON方法，返回status code即可


路由：
使用gin.Default对象调用restapi的方法
把httpmethod，http路径和handler注册在一起

*/

import (
	"fmt"

	"github.com/ebisu2018/xBlog/apps/token"
	"github.com/ebisu2018/xBlog/common"
	"github.com/ebisu2018/xBlog/ioc"
	"github.com/gin-gonic/gin"
)


func init()  {
	fmt.Println("loading api package")
	ioc.ApiHandler().Registry(&HttpApiHandler{})
}


func (a *HttpApiHandler) Init()  {
	a.tkSvc = ioc.Container().Get(token.AppName).(token.TokenService)
}


func (a *HttpApiHandler) Name() string {
	return token.AppName
}


func NewHttpApiHander(svc token.TokenService) *HttpApiHandler {
	return &HttpApiHandler{
		tkSvc: svc,
	}
}

type HttpApiHandler struct {
	tkSvc token.TokenService
}


func (a *HttpApiHandler) Register(r gin.IRouter) {
	v1 := r.Group("v1")
	v1.POST("/tokens", a.ApiLogin)
	// v1.DELETE("/tokens", a.ApiLogout)
}

func (a *HttpApiHandler) ApiLogin(c *gin.Context) {

	req := token.NewLoginRequest()

	// 1. 相当于json.Unmarshal，body数据转换到对象上
	err := c.BindJSON(req)
	if err != nil {
		common.Failed(c, err)
		return
	}

	// 2. 把http请求转换成controller的请求
	tk, err := a.tkSvc.Login(c.Request.Context(), req)
	if err != nil {
		common.Failed(c, err)
		return
	}

	// 3. 返回status code结果
	common.Success(c, tk)
}

func (a *HttpApiHandler) ApiLogout(c *gin.Context) {
}