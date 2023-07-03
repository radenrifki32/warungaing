package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/app"
	controllerOrder "github.com/rifki321/warungku/app/order/controller_order"
	"github.com/rifki321/warungku/app/order/repo"
	"github.com/rifki321/warungku/app/order/serviceOrder"
	"github.com/rifki321/warungku/categories/controller"
	"github.com/rifki321/warungku/categories/repository"
	"github.com/rifki321/warungku/categories/service"
	"github.com/rifki321/warungku/messages/controllermessage"
	"github.com/rifki321/warungku/messages/repositorymessage"
	"github.com/rifki321/warungku/messages/servicemessage"
	"github.com/rifki321/warungku/middleware"
	"github.com/rifki321/warungku/product"
	"github.com/rifki321/warungku/user"
	"github.com/rifki321/warungku/user/web"
)

func main() {
	db := app.GetConnectionDb()
	NewRepo := user.NewUserRepoRepository()
	NewService := user.NewUserService(NewRepo, db)
	NewController := user.NewUserController(NewService)
	NewRepoProduct := product.NewProductRepo()
	NewServiceProduct := product.NewProductService(NewRepoProduct, db)
	NewControllerProduct := product.NewProductController(NewServiceProduct)
	NewRepoCategory := repository.NewCategory()
	NewServiceCategory := service.NewServiceCategory(NewRepoCategory, db)
	NewControllerCategory := controller.NewControllerCategory(NewServiceCategory)
	NewRepoOrder := repo.NewRepoOrder()
	NewServiceOrder := serviceOrder.NewOrderService(db, NewRepoOrder)
	NewControllerOrder := controllerOrder.NewControllerOrder(NewServiceOrder)
	NewRepoMessage := repositorymessage.NewMessageRepository()
	NewServiceMessage := servicemessage.NewServiceMessage(db, NewRepoMessage)
	NewControllerMessage := controllermessage.NewControllerMessage(NewServiceMessage)

	router := httprouter.New()
	handler := middleware.CorsMiddleware(router)
	router.POST("/user/register", NewController.Register)
	router.POST("/user/login", NewController.Login)
	router.GET("/products", middleware.TokenAuthMiddleware(NewControllerProduct.GetAllProduct))
	router.GET("/product", NewControllerProduct.GetProductByCategories)
	router.GET("/product/:categoryid", NewControllerProduct.GetProductByCategoriesId)
	router.GET("/categories", NewControllerCategory.GetCategory)
	router.GET("/products/:productid", NewControllerProduct.GetProductById)
	router.POST("/products/create", NewControllerProduct.PostProduct)
	router.POST("/order", NewControllerOrder.OrderProduct)
	router.DELETE("/products/:productid", NewControllerProduct.DeleteProduct)
	router.PUT("/products/:productid", NewControllerProduct.UpdateProduct)
	router.POST("/message", NewControllerMessage.SendMessage)
	router.GET("/allmessage", NewControllerMessage.GetMessageByUsername)
	router.GET("/message/:messageId", NewControllerMessage.GetMessageByIdMessage)

	router.PanicHandler = web.PanicHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	server.ListenAndServe()

}
