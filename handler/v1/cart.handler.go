package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"stock_management/common/obj"
	"stock_management/common/response"
	"stock_management/dto"
	"stock_management/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	All(ctx *gin.Context)
	CreateCart(ctx *gin.Context)
	FindOneCartByID(ctx *gin.Context)
}

type cartHandler struct {
	cartService service.CartService
	jwtService  service.JWTService
}

func NewCartHandler(cartService service.CartService, jwtService service.JWTService) CartHandler {
	return &cartHandler{
		cartService: cartService,
		jwtService:  jwtService,
	}
}

// @Summary Get All Cart by UserId
// @Description Get All Cart
// @ID Get All Cart
// @Param Authorization header string true "Token"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/cart/all [get]
func (c *cartHandler) All(ctx *gin.Context) {
	log.Println("<< CART HANDLER - FIND ALL BY USER ID >>")
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	log.Println("UserId : ", userID)
	carts, err := c.cartService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", carts)
	responseByte, _ := json.Marshal(response)
	log.Println("Response : ", string(responseByte))
	ctx.JSON(http.StatusOK, response)
}

// @Summary Create Cart
// @Description Create Cart
// @ID Create Cart
// @Param Authorization header string true "Token"
// @Param body body dto.CreateCartRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/cart [post]
func (c *cartHandler) CreateCart(ctx *gin.Context) {
	log.Println("<< CART HANDLER - CREATE >>")
	var createCartReq dto.CreateCartRequest
	err := ctx.ShouldBind(&createCartReq)

	reqByte, _ := json.Marshal(createCartReq)
	log.Println("Request : ", string(reqByte))
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	log.Println("userId :", userID)
	res, err := c.cartService.CreateCart(createCartReq, userID)
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

// @Summary Find Cart By ID
// @Description Find Cart By ID
// @ID Find Cart By ID
// @Param Authorization header string true "Token"
// @Param id path string true "ID of the Cart to be find"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/cart/{id} [get]
func (c *cartHandler) FindOneCartByID(ctx *gin.Context) {
	log.Println("<< CART HANDLER - FIND BY ID >>")
	id := ctx.Param("id")
	log.Println("Cart Id : ", id)
	res, err := c.cartService.FindOneCartByID(id)
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
