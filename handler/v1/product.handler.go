package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"stock_management/common/obj"
	"stock_management/common/response"
	"stock_management/dto"
	"stock_management/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	All(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	FindOneProductByID(ctx *gin.Context)
}

type productHandler struct {
	productService service.ProductService
	jwtService     service.JWTService
}

func NewProductHandler(productService service.ProductService, jwtService service.JWTService) ProductHandler {
	return &productHandler{
		productService: productService,
		jwtService:     jwtService,
	}
}


// @Summary Get All Product
// @Description Get All Product
// @ID Get All Product
// @Param Authorization header string true "Token"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product/all [get]
func (c *productHandler) All(ctx *gin.Context) {
	log.Println("<< PRODUCT HANDLER - FIND ALL PRODUCT")
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	log.Println("User Id : ", userID)
	products, err := c.productService.All()
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", products)
	responseByte, _ := json.Marshal(response)
	log.Println("Response : ", string(responseByte))
	ctx.JSON(http.StatusOK, response)
}

// @Summary Create Product
// @Description Create Product
// @ID Create Product
// @Param Authorization header string true "Token"
// @Param body body dto.CreateProductRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product [post]
func (c *productHandler) CreateProduct(ctx *gin.Context) {
	log.Println("<< PRODUCT HANDLER - CREATE PRODUCT")
	var createProductReq dto.CreateProductRequest
	err := ctx.ShouldBind(&createProductReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		return
	}
	reqByte, _ := json.Marshal(createProductReq)
	log.Println("Response : ", string(reqByte))
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.productService.CreateProduct(createProductReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	responseByte, _ := json.Marshal(response)
	log.Println("Response : ", string(responseByte))
	ctx.JSON(http.StatusCreated, response)

}

// @Summary Find Product By ID
// @Description Find Product By ID
// @ID Find Product By ID
// @Param Authorization header string true "Token"
// @Param id path string true "ID of the product to be find"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product/{id} [get]
func (c *productHandler) FindOneProductByID(ctx *gin.Context) {
	log.Println("<< PRODUCT HANDLER - FIND PRODUCT BY PRODUCT ID")
	id := ctx.Param("id")
	log.Println("Product Id : " , id)
	res, err := c.productService.FindOneProductByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	responseByte, _ := json.Marshal(response)
	log.Println("Response : ", string(responseByte))
	ctx.JSON(http.StatusOK, response)
}

// @Summary Delete Product By ID
// @Description Delete Product By ID
// @ID Delete Product By ID
// @Param Authorization header string true "Token"
// @Param id path string true "ID of the product to be delete"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product/{id} [delete]
func (c *productHandler) DeleteProduct(ctx *gin.Context) {
	log.Println("<< PRODUCT HANDLER - DELETE PRODUCT BY PRODUCT ID")
	id := ctx.Param("id")
	log.Println("Product Id : ", id)
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	log.Println("User Id : ", userID)
	err := c.productService.DeleteProduct(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	responseByte, _ := json.Marshal(response)
	log.Println("Response : ", string(responseByte))
	ctx.JSON(http.StatusOK, response)
}

// @Summary Update Product By ID
// @Description Update Product By ID
// @ID Update Product By ID
// @Param Authorization header string true "Token"
// @Param id path string true "ID of the product to be updated"
// @Param body body dto.UpdateProductRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product/{id} [put]
func (c *productHandler) UpdateProduct(ctx *gin.Context) {
	log.Println("<< PRODUCT HANDLER - UPDATE PRODUCT BY PRODUCT ID")
	updateProductRequest := dto.UpdateProductRequest{}
	err := ctx.ShouldBind(&updateProductRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	reqByte, _ := json.Marshal(updateProductRequest)
	log.Println("Response : ", string(reqByte))

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	log.Println("User Id : ", userID)
	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateProductRequest.ID = id
	product, err := c.productService.UpdateProduct(updateProductRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", product)
	responseByte, _ := json.Marshal(response)
	log.Println("Response : ", string(responseByte))
	ctx.JSON(http.StatusOK, response)

}
