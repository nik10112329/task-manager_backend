// package rest
// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type MyRouter struct {
// }

// func (MyRouter) Getting() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"method": "GET"})
// 	}
// }
// func (MyRouter) GetUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Query("displayName")
// 		c.Query("email")
// 		c.Query("")
// 	}
// }
// func MyRouterConstructor() MyRouter {
// 	return MyRouter{}
// }
//