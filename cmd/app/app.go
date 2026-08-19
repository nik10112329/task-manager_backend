package main

import (
	// "learning-project/internal/transport/rest"

	// "github.com/gin-gonic/gin"
	"fmt"
	"learning-project/internal/config"
	"learning-project/internal/user"
	"learning-project/pkg/logging"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

//	func posting(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"method": "POST"})
//	}
//
//	func putting(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"method": "PUT"})
//	}
//
//	func deleting(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
//	}
//
//	func patching(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"method": "PATCH"})
//	}
//
//	func head(c *gin.Context) {
//		c.Status(http.StatusOK)
//	}
//
//	func options(c *gin.Context) {
//		c.Status(http.StatusOK)
//	}
//
//	func mine(transferPoint chan int, number int) {
//		time.Sleep(1 * time.Second)
//		transferPoint <- 10
//		fmt.Println(number)
//	}
func appStart(router *httprouter.Router, cfg *config.Config) {

	logger := logging.GetLogger()
	listenType := cfg.Listen.Type

	var listener net.Listener
	var listenErr error

	if listenType == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("Create socket")

		socketPath := path.Join(appDir, "app.sock")

		logger.Debugf("Socket path: %s", socketPath)

		logger.Info("Listen unix socket")

		listener, listenErr = net.Listen("unix", socketPath)

		logger.Debugf("Socket path: %s", socketPath)
	} else {
		logger.Info("Listen TCP")
		listener, listenErr = net.Listen("tcp",
			fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		log.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	server.Serve(listener)

}
func main() {
	//close(transferPoint)
	logger := logging.GetLogger()

	config := config.GetConfig()

	logger.Info("create router")

	router := httprouter.New()

	logger.Info("register user handler")

	handler := user.NewHandler(logger)

	handler.Register(router)

	appStart(router, config)
	// router, myRouter := gin.Default(), rest.MyRouter{}
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	//
	// router.GET("/someGet", myRouter.Getting())
	//
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
	//router.GET("/welcome", func(c *gin.Context) {
	//	firstname := c.DefaultQuery("firstname", "Guest")
	//	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	//	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//})
	//if err := router.Run(); err != nil {
	//	return
	//}
}
