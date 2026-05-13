package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	respository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		respository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.respository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.respository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.respository.GetProductById(id_product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProductById(id_product int) (*model.Product, error) {
	product, err := pu.respository.DeleteProductById(id_product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
