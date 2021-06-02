package handler

import (
	"go-bake/entity"
	"go-bake/helper"
	"go-bake/product"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) ShowProductHandler(c *gin.Context) {
	products, err := h.productService.GetAllProduct()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("succes get all product", 200, "status OK", products)
	c.JSON(200, response)
}

func (h *productHandler) CreateProductHandler(c *gin.Context) {

	var inputProduct entity.ProductInput

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		splitError := helper.SplitErrorInformation(err)

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
	}

	newProduct, err := h.productService.SaveNewProduct(inputProduct)
	if err != nil {

		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create product", 201, "status created", newProduct)
	c.JSON(200, response)

}

func (h *productHandler) GetProductByIDHandler(c *gin.Context) {
	id := c.Param("product_id")

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)

		return
	}

	response := helper.APIResponse("success get  product by id", 200, "succes", product)
	c.JSON(200, response)
}

func (h *productHandler) UpdateProductByIDHandler(c *gin.Context) {
	id := c.Params.ByName("product_id")

	var updateProductInput entity.UpdateProductInput

	if err := c.ShouldBindJSON(&updateProductInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{
			"errors": splitError})

		c.JSON(400, responseError)

		return
	}

	// idParam, _ := strconv.Atoi(id)

	// productData := int(c.MustGet("currentProduct").(int))

	// if idParam != productData {
	// 	responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "Product id not authorize"})

	// 	c.JSON(401, responseError)
	// 	return
	// }

	product, err := h.productService.UpdateProductByID(id, updateProductInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("update user succeed", 200, "success", product)
	c.JSON(200, response)
}

func (h *productHandler) DeleteProductByIDHandler(c *gin.Context) {
	id := c.Params.ByName("product_id")
	product, err := h.productService.DeleteProductByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)

		return
	}

	response := helper.APIResponse("success delete product by ID", 200, "succes", product)
	c.JSON(200, response)
}
