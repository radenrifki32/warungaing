package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/app"
	"github.com/rifki321/warungku/product"
	"github.com/rifki321/warungku/user"
	"github.com/rifki321/warungku/user/web"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func main() {
	db := app.GetConnectionDb()
	NewRepo := user.NewUserRepoRepository()
	NewService := user.NewUserService(NewRepo, db)
	NewController := user.NewUserController(NewService)
	NewRepoProduct := product.NewProductRepo()
	NewServiceProduct := product.NewProductService(NewRepoProduct, db)
	NewControllerProduct := product.NewProductController(NewServiceProduct)
	router := httprouter.New()
	handler := corsMiddleware(router)
	router.PanicHandler = web.PanicHandler
	router.POST("/user/register", NewController.Register)
	router.POST("/user/login", NewController.Login)
	router.GET("/products", NewControllerProduct.GetAllProduct)
	router.GET("/product", NewControllerProduct.GetProductByCategories)

	router.GET("/product/:categoryid", NewControllerProduct.GetProductByCategoriesId)
	router.GET("/products/:productid", NewControllerProduct.GetProductById)
	router.POST("/products/create", NewControllerProduct.PostProduct)
	router.DELETE("/products/:productid", NewControllerProduct.DeleteProduct)
	router.PUT("/products/:productid", NewControllerProduct.UpdateProduct)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	server.ListenAndServe()

}
