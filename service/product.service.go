package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"stock_management/dto"
	"stock_management/entity"
	"stock_management/repo"

	"github.com/mashingan/smapping"

	_product "stock_management/service/product"
)

type ProductService interface {
	All() (*[]_product.ProductResponse, error)
	CreateProduct(productRequest dto.CreateProductRequest, userID string) (*_product.ProductResponse, error)
	UpdateProduct(updateProductRequest dto.UpdateProductRequest, userID string) (*_product.ProductResponse, error)
	FindOneProductByID(productID string) (*_product.ProductResponse, error)
	DeleteProduct(productID string, userID string) error
}

type productService struct {
	productRepo repo.ProductRepository
}

func NewProductService(productRepo repo.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (c *productService) All() (*[]_product.ProductResponse, error) {
	products, err := c.productRepo.All()
	if err != nil {
		return nil, err
	}

	prods := _product.NewProductArrayResponse(products)
	return &prods, nil
}

func (c *productService) CreateProduct(productRequest dto.CreateProductRequest, userID string) (*_product.ProductResponse, error) {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&productRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	product.UserID = id
	product.CreatedAt =time.Now()
	p, err := c.productRepo.InsertProduct(product)
	if err != nil {
		return nil, err
	}

	res := _product.NewProductResponse(p)
	return &res, nil
}

func (c *productService) FindOneProductByID(productID string) (*_product.ProductResponse, error) {
	product, err := c.productRepo.FindOneProductByID(productID)

	if err != nil {
		return nil, err
	}

	res := _product.NewProductResponse(product)
	return &res, nil
}

func (c *productService) UpdateProduct(updateProductRequest dto.UpdateProductRequest, userID string) (*_product.ProductResponse, error) {
	product, err := c.productRepo.FindOneProductByID(fmt.Sprintf("%d", updateProductRequest.ID))
	if err != nil {
		return nil, err
	}

	uid, _ := strconv.ParseInt(userID, 0, 64)
	if product.UserID != uid {
		return nil, errors.New("produk ini bukan milik anda")
	}

	product = entity.Product{}
	err = smapping.FillStruct(&product, smapping.MapFields(&updateProductRequest))

	if err != nil {
		return nil, err
	}

	product.UserID = uid
	product.UpdatedAt = time.Now()
	product, err = c.productRepo.UpdateProduct(product)

	if err != nil {
		return nil, err
	}

	res := _product.NewProductResponse(product)
	return &res, nil
}

func (c *productService) DeleteProduct(productID string, userID string) error {
	product, err := c.productRepo.FindOneProductByID(productID)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%d", product.UserID) != userID {
		return errors.New("produk ini bukan milik anda")
	}

	c.productRepo.DeleteProduct(productID)
	return nil

}
