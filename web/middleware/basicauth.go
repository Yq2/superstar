// file: middleware/basicauth.go

package middleware

import "github.com/kataras/iris/middleware/basicauth"

// BasicAuth middleware sample.
//使用基本认证中间件
//BasicAuth 是 context.Handler实现类型
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})

var BasicAuthh = basicauth.New(basicauth.Config{
	Users:map[string]string{
		"admin":"password",
	},
})