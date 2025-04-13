package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/schemas"
	"github.com/brewbits-co/releasedesk/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func (c *productController) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	var productCreationRequest product.BasicInfo

	err = utils.NewDecoder().Decode(&productCreationRequest, r.Form)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	newProduct, err := c.service.CreateProduct(productCreationRequest)
	if err != nil {
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, schemas.NewErrorResponse("A product with the same identifier already exists.", []string{
			"Please use a unique product name and slug.",
			"Check the existing products before creating a new one.",
		}))
		return
	}

	log.Println(newProduct.ID)

	render.Status(r, http.StatusCreated)
}

func (c *productController) HandleProductSetupGuide(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}
	slug := chi.URLParam(r, "slug")

	err = c.service.ApplyProductSetupGuide(
		values.Slug(slug),
		values.VersionFormat(r.Form.Get("VersionFormat")),
		product.SetupChannelsOption(r.Form.Get("Channels")),
		r.Form["CustomChannels"],
	)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, schemas.NewErrorResponse(err.Error(), nil))
		return
	}
}
