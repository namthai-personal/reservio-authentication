package main

import (
	"context"
	"log"
	"os"
	"reservio-be/infrastructure/database"
	"reservio-be/routes"

	"reservio-be/docs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var ginLambda *ginadapter.GinLambda

const (
	PORT = ":3000"
)

// @title			Auth Service
// @version		1.0.0
// @description	Reservio 2.0 Auth Service
// @BasePath		/
func init() {
	docs.SwaggerInfo.Title = "Auth Service"
	database.OpenDbConnection()

	router := routes.SetupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	env := os.Getenv("ENV")

	if env == "local" {
		router.Run(PORT)
	} else {
		ginLambda = ginadapter.New(router)
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
