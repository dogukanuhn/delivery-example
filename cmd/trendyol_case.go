package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dogukanuhn/delivery-system/cfg"
	_ "github.com/dogukanuhn/delivery-system/docs"
	"github.com/dogukanuhn/delivery-system/domain/dto"
	"github.com/dogukanuhn/delivery-system/domain/interfaces"
	"github.com/dogukanuhn/delivery-system/internal/logger"
	"github.com/dogukanuhn/delivery-system/internal/repositories/packageRepository"
	"github.com/dogukanuhn/delivery-system/internal/repositories/sackRepository"
	"github.com/dogukanuhn/delivery-system/internal/response"
	"github.com/dogukanuhn/delivery-system/internal/services/deliveryService"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Execute(env string) {

	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg.MockDatabase()

	echoInstance := echo.New()

	// Middleware
	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.Recover())
	echoInstance.Use(middleware.CORS())

	// Routes
	echoInstance.GET("/", HealthCheck)
	echoInstance.POST("/deliver", DeliverRoute)

	echoInstance.GET("/swagger/*", echoSwagger.WrapHandler)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	echoInstance.Logger.Fatal(echoInstance.Start(port))

}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

// Delivery Endpoint
// @Summary Post request for handle delivery status
// @Description If package or sacks in correct delivery point, change status to unloaded. If not, set status to loaded.
// @Description Branch Point can only receive packages, Distribution Point can receive everything, Transfer Point can receive only sacks and packages in sacks
// @Tags root
// @Accept */*
// @Param Routes body dto.DeliverDTO true "Request struct"
// @Produce json
// @Success 200 {object} dto.DeliverDTO
// @Failure  400 {object} response.ErrorResponse
// @Router / [post]
func DeliverRoute(c echo.Context) error {
	deliveryRoutes := new(dto.DeliverDTO)

	logger := logger.WithCollection(cfg.GetDatabase().Collection("logs"))

	if err := c.Bind(deliveryRoutes); err != nil {
		logger.Fatal(err)
		return c.JSON(400, response.ErrorResponse{
			Message: "Request Parse Error",
			Code:    400,
		})
	}

	var packageRepo interfaces.IPackageRepository
	var sackRepo interfaces.ISackRepository

	packageRepo = packageRepository.NewRepository(cfg.GetDatabase().Collection("packages"))
	sackRepo = sackRepository.NewRepository(cfg.GetDatabase().Collection("sacks"))

	deliveryHandler := deliveryService.NewHandler(packageRepo, sackRepo, logger)

	deliveryHandler.Deliver(*deliveryRoutes)

	return c.JSON(http.StatusOK, deliveryRoutes)
}
