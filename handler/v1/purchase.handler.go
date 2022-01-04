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

type PurchaseHandler interface {
	All(ctx *gin.Context)
	CreatePurchase(ctx *gin.Context)
	UpdatePurchase(ctx *gin.Context)
	FindOnePurchaseByID(ctx *gin.Context)
}

type purchaseHandler struct {
	purchaseService service.PurchaseService
	jwtService     service.JWTService
}

func NewPurchaseHandler(purchaseService service.PurchaseService, jwtService service.JWTService) PurchaseHandler {
	return &purchaseHandler{
		purchaseService: purchaseService,
		jwtService:     jwtService,
	}
}


// @Summary Get All Purchase by User Id
// @Description Get All Purchase by User Id
// @ID Get All Purchase
// @Param Authorization header string true "Token"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/purchase/all [get]
func (c *purchaseHandler) All(ctx *gin.Context) {
	log.Println("<< PURCHASE HANDLER - GET ALL PURCHASE BY USER ID")
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	log.Println("UserId : ", userID)
	purchases, err := c.purchaseService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", purchases)
	responseByte, _ := json.Marshal(response)
	log.Println("Response : ", string(responseByte))
	ctx.JSON(http.StatusOK, response)
}

// @Summary Create Purchse
// @Description Create Purchase
// @ID Create Purchase
// @Param Authorization header string true "Token"
// @Param body body dto.CreatePurchaseRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/purchase [post]
func (c *purchaseHandler) CreatePurchase(ctx *gin.Context) {
	log.Println("<< PURCHASE HANDLER - CREATE PURCHASE")
	var createPurchaseReq dto.CreatePurchaseRequest
	err := ctx.ShouldBind(&createPurchaseReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	reqByte, _ := json.Marshal(createPurchaseReq)
	log.Println("Request : ", string(reqByte))
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	log.Println("User Id : ", userID)
	res, err := c.purchaseService.CreatePurchase(createPurchaseReq, userID)
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

// @Summary Find Purchase By ID
// @Description Find Purchase By ID
// @ID Find Purchase By ID
// @Param Authorization header string true "Token"
// @Param id path string true "ID of the purchase to be find"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/purchase/{id} [get]
func (c *purchaseHandler) FindOnePurchaseByID(ctx *gin.Context) {
	log.Println("<< PURCHASE HANDLER - GET PURCHASE BY PURCHASE ID")
	id := ctx.Param("id")
	log.Println("Purchase Id : ", id)
	res, err := c.purchaseService.FindOnePurchaseByID(id)
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

// @Summary Update Purchase By ID
// @Description Update Purchase By ID
// @ID Update Purchase By ID
// @Param Authorization header string true "Token"
// @Param id path string true "ID of the purchase to be updated"
// @Param body body dto.UpdatePurchaseRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/purchase/{id} [put]
func (c *purchaseHandler) UpdatePurchase(ctx *gin.Context) {
	log.Println("<< PURCHASE HANDLER - UPDATE PURCHASE BY ID")
	updatePurchaseRequest := dto.UpdatePurchaseRequest{}
	err := ctx.ShouldBind(&updatePurchaseRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	reqByte, _ := json.Marshal(updatePurchaseRequest)
	log.Println("Request : ", string(reqByte))

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	log.Println("User Id : ", userID)
	
	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updatePurchaseRequest.ID = id
	purchase, err := c.purchaseService.UpdatePurchase(updatePurchaseRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		responseByte, _ := json.Marshal(response)
		log.Println("Response : ", string(responseByte))
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", purchase)
	responseByte, _ := json.Marshal(response)
	log.Println("Response : ", string(responseByte))
	ctx.JSON(http.StatusOK, response)

}
