package apps

import (
	"fmt"

	_ "github.com/ebisu2018/xBlog/apps/token/impl"
	_ "github.com/ebisu2018/xBlog/apps/user/impl"
	_ "github.com/ebisu2018/xBlog/apps/token/api"
)

func init()  {
	fmt.Println("loading app package")
}