package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cocoasterr/net_http/app/controllers"
	PGConfig "github.com/cocoasterr/net_http/infra/db/postgres"
	PGRepository "github.com/cocoasterr/net_http/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/net_http/services"
)

func main() {
	db, err := PGConfig.ConnectPG()
	if err != nil {
		log.Fatal("Failed to Connect DB")
	}
	AuthRepo := PGRepository.NewUserRepository(db)
	AuthService := services.NewAuthService(AuthRepo.Repository)
	AuthController := controllers.NewAuthController(*AuthService)
	http.HandleFunc("/api/auth/register", AuthController.Register)
	http.HandleFunc("/api/auth/login", AuthController.Login)

	ProductRepo := PGRepository.NewProductRepository(db)
	ProductService := services.NewProductService(ProductRepo.Repository)
	Productcontroller := controllers.NewProductController(*ProductService)

	http.HandleFunc("/api/create-product", Productcontroller.CreateProductController)
	http.HandleFunc("/api/index-product", Productcontroller.IndexProdcuctController)
	http.HandleFunc("/api/find-product/", Productcontroller.FindProduct)
	http.HandleFunc("/api/update-product/", Productcontroller.UpdateProduct)
	http.HandleFunc("/api/delete-product/", Productcontroller.DeleteProduct)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{"Message": "page not found!"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("Running on port :8080")
	http.ListenAndServe(":8080", nil)
}
