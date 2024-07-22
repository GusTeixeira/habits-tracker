package runner

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GusTeixeira/habits-tracker/api/responses"
	"github.com/GusTeixeira/habits-tracker/internals/middlewares"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func EnvOr(name, value string) string {
	val, exists := os.LookupEnv(name)
	if exists {
		return val
	}
	return value
}

type AdapterHandler func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func Lambda(adapter *chiadapter.ChiLambda) func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return adapter.ProxyWithContext(ctx, req)
	}
}

func Run(applicationName string) error {

	root := chi.NewRouter()
	root.Use(middlewares.Cors)
	root.Use(middleware.Logger)
	//root.Mount("/habits", routes.HabitsRoutes())
	root.Get("/docs/*", httpSwagger.WrapHandler)

	root.NotFound(func(w http.ResponseWriter, r *http.Request) {
		responses.HTTPError(w, "recurso não encontrado, verifique a documentacao", 404)
	})

	_, isDevMode := os.LookupEnv("DEV")
	if isDevMode {
		var host string = EnvOr("HOST", "0.0.0.0")
		var port string = EnvOr("PORT", "8080")
		var addr string = fmt.Sprintf("%s:%s", host, port)

		log.Printf("Iniciando servidor em: %s\n", addr)
		if err := http.ListenAndServe(":8080", root); err != nil {
			return err
		}
		return nil
	}
	adapter := chiadapter.New(root)

	lambda.StartWithOptions(Lambda(adapter), lambda.WithContext(context.Background()))
	log.Println("close connected")
	return nil
}
