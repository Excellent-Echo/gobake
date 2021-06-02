package route

import (
	"go-bake/handler"
	"go-bake/product"

	"github.com/gin-gonic/gin"
)

var (
	productRepository = product.NewRepository(DB)
	productService    = product.NewService(productRepository)
	productHandler    = handler.NewProductHandler(productService)
)

func ProductRouter(r *gin.Engine) {
	r.GET("/products", productHandler.ShowProductHandler)
	r.GET("/products/:product_id", productHandler.GetProductByIDHandler)
	r.POST("products/", productHandler.CreateProductHandler)
	r.PUT("products/:product_id", productHandler.UpdateProductByIDHandler)
	r.DELETE("products/:product_id", productHandler.DeleteProductByIDHandler)
}
