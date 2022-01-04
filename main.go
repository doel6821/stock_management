package main

import (
	"time"

	"stock_management/config"
	"stock_management/docs"
	v1 "stock_management/handler/v1"
	"stock_management/middleware"
	"stock_management/repo"
	"stock_management/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	db             *gorm.DB                = config.SetupDatabaseConnection()
	userRepo       repo.UserRepository     = repo.NewUserRepo(db)
	productRepo    repo.ProductRepository  = repo.NewProductRepo(db)
	cartRepo       repo.CartRepository     = repo.NewCartRepo(db)
	purchaseRepo   repo.PurchaseRepository = repo.NewPurchaseRepo(db)
	historyRepo    repo.HistoryRepository  = repo.NewHistoryRepo(db)
	authService    service.AuthService     = service.NewAuthService(userRepo)
	jwtService     service.JWTService      = service.NewJWTService()
	userService    service.UserService     = service.NewUserService(userRepo)
	productService service.ProductService  = service.NewProductService(productRepo)
	cartService    service.CartService     = service.NewCartService(cartRepo, productRepo, historyRepo)
	purchaseService service.PurchaseService    = service.NewPurchaseService(purchaseRepo, productRepo, historyRepo)
	historyService  service.HistoryService = service.NewHistoryService(productRepo,historyRepo)
	authHandler    v1.AuthHandler          = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler    v1.UserHandler          = v1.NewUserHandler(userService, jwtService)
	productHandler v1.ProductHandler       = v1.NewProductHandler(productService, jwtService)
	cartHandler    v1.CartHandler          = v1.NewCartHandler(cartService, jwtService)
	purchaseHandler v1.PurchaseHandler    = v1.NewPurchaseHandler(purchaseService, jwtService)
	historyHandler  v1.HistoryHandler    = v1.NewHistoryHandler(historyService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
		
	}))
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "Stock Management"
	docs.SwaggerInfo.Description = "Simple Ecommerce Stock Management"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = "localhost:8080"

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	productRoutes := server.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/all", productHandler.All)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	cartRouters := server.Group("api/cart", middleware.AuthorizeJWT(jwtService))
	{
		cartRouters.GET("/all", cartHandler.All)
		cartRouters.POST("/", cartHandler.CreateCart)
		cartRouters.GET("/:id", cartHandler.FindOneCartByID)
	}

	purchaseRouter := server.Group("api/purchase", middleware.AuthorizeJWT(jwtService))
	{
		purchaseRouter.GET("/all", purchaseHandler.All)
		purchaseRouter.POST("/", purchaseHandler.CreatePurchase)
		purchaseRouter.GET("/:id", purchaseHandler.FindOnePurchaseByID)
		purchaseRouter.PUT("/:id", purchaseHandler.UpdatePurchase)
	}

	historyRouter := server.Group("api/history", middleware.AuthorizeJWT(jwtService))
	{
		historyRouter.GET("/:productId", historyHandler.FindHistoryByProductID)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
