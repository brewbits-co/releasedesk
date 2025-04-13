package product

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
)

func (s *service) CreateProduct(info product.BasicInfo) (product.Product, error) {
	info.SetupGuideCompleted = false
	info.Slug.Format()

	newProduct := product.NewProduct(info)

	err := newProduct.IsValid()
	if err != nil {
		return product.Product{}, err
	}

	err = s.productRepo.Save(&newProduct)
	if err != nil {
		return product.Product{}, err
	}

	return newProduct, nil
}
