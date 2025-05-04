package product

import (
	"github.com/brewbits-co/releasedesk/internal/services/platform"
	"github.com/brewbits-co/releasedesk/internal/services/product"
	"net/http"
)

type ProductController interface {
	HandleCreateProduct(w http.ResponseWriter, r *http.Request)
	HandleProductSetupGuide(w http.ResponseWriter, r *http.Request)
	RenderDashboard(w http.ResponseWriter, r *http.Request)
}

// NewProductController creates a new instance of productController with the provided dependencies.
func NewProductController(service product.Service, appService platform.Service) ProductController {
	return &productController{
		service:    service,
		appService: appService,
	}
}

// productController implements the product.ProductController.
type productController struct {
	service    product.Service
	appService platform.Service
}
