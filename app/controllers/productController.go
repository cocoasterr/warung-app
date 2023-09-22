package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cocoasterr/net_http/app/controllers/helper"
	"github.com/cocoasterr/net_http/services"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (c *ProductController) CreateProductController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", http.StatusInternalServerError)
		return
	}
	tokenString := r.Header.Get("Authorization")
	helper.AuthSession(w, tokenString)

	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid Request Body!", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	if err := c.ProductService.CreateService(ctx, payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{"Message": "Success!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *ProductController) IndexProdcuctController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusInternalServerError)
		return
	}
	tokenString := r.Header.Get("Authorization")
	helper.AuthSession(w, tokenString)

	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	limit, _ := strconv.Atoi(r.URL.Query()["limit"][0])

	ctx := r.Context()
	data, total, err := c.ProductService.IndexService(ctx, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"message": "success!",
		"data":    data,
		"total":   total,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (c *ProductController) FindProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusInternalServerError)
		return
	}

	tokenString := r.Header.Get("Authorization")
	helper.AuthSession(w, tokenString)

	id := r.URL.Path[len("/api/find-product/"):]
	// id := r.URL.Query().Get("id")
	ctx := r.Context()
	data, err := c.ProductService.FindByService(ctx, "id", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{
		"message": "success!",
		"data":    data,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed!", http.StatusInternalServerError)
		return
	}
	tokenString := r.Header.Get("Authorization")
	helper.AuthSession(w, tokenString)

	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid Request Body!", http.StatusBadRequest)
		return
	}
	id := r.URL.Path[len("/api/update-product/"):]
	ctx := r.Context()
	err := c.ProductService.UpdateService(ctx, payload, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{"Message": "Success!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed!", http.StatusInternalServerError)
		return
	}
	tokenString := r.Header.Get("Authorization")
	helper.AuthSession(w, tokenString)

	id := r.URL.Path[len("/api/delete-product/"):]
	ctx := r.Context()
	err := c.ProductService.DeleteService(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{"Message": "Success!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
