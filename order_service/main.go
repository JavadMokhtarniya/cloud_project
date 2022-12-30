package main
import (
	"flag"
	"github.com/gin-gonic/gin"
)
func main(){
	authaddr:=flag.string("authAddr","localhost:50051","Authentication service address (GRPC)")
	prodaddr:=flag.string("prodAddr","localhost:50052","Product service address (GRPC)")

	flag.Parse()
	svc:=NewInmemservice(*prodaddr)
	handler:=NewHandler(svc)
	authMW:=NewAuthMiddleware(*authaddr)

	router:=gin.Default()

	router.GET("/basket",authMW.hasAccess,handler.getBasket)
	router.POST("/basket/:pid",authMW.hasAccess,handler.addProductToBasket)
	router.PUT("/basket/:pid",authMW.hasAccess,handler.modifyBasket)

	router.Run(":3000")
}