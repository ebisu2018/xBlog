package common

import (
	"net/http"

	"github.com/ebisu2018/xBlog/exception"
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Failed(c *gin.Context, err error)  {


	var e *exception.ApiException
	if v, ok := err.(*exception.ApiException); ok {
		e = v
	} else {
		e = exception.NewApiException(http.StatusInternalServerError, err.Error())
		e.HttpCode = http.StatusInternalServerError
	}

	c.JSON(http.StatusBadRequest, e)
}