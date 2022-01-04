package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"stock_management/common/obj"
	"stock_management/common/response"
	"stock_management/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type HistoryHandler interface {
	FindHistoryByProductID(ctx *gin.Context)
}

type historyHandler struct {
	historyService service.HistoryService
	jwtService  service.JWTService
}

func NewHistoryHandler(historyService service.HistoryService, jwtService service.JWTService) HistoryHandler {
	return &historyHandler{
		historyService: historyService,
		jwtService:  jwtService,
	}
}

// @Summary Find History By Product ID
// @Description Find History By Product ID
// @ID Find History By Product ID
// @Param Authorization header string true "Token"
// @Param productId path string true "ID of the product to be find"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/history/{productId} [get]
func (c *historyHandler) FindHistoryByProductID(ctx *gin.Context) {
	log.Println("<< HISTORY HANDLER - FIND BY PRODUCT ID >>")
	id := ctx.Param("productId")
	log.Println("ProductId : ", id)
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	log.Println("UserId : ", userID)
	res, err := c.historyService.FindHistoryByProductID(id, userID)
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
