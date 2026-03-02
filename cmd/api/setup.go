package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	_http "github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/http"
	_chi "github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/http/chi"
	"github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/repository/postgres"
	"github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/repository/postgres/sqlc"
	"github.com/EduCoelhoTs/base-hex-arq-api/internal/application/service"
	userusecase "github.com/EduCoelhoTs/base-hex-arq-api/internal/application/usecase/user"
	"github.com/EduCoelhoTs/base-hex-arq-api/internal/infra/config"
	usercontroller "github.com/EduCoelhoTs/base-hex-arq-api/internal/infra/controller/user"
)

func startServer() error {
	//use when running in local environment, in production the env vars will be set by the hosting provider
	loadFile := true
	config, err := config.LoadConfig(loadFile)
	if err != nil {
		return fmt.Errorf("error to load enviroments variables? %w", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	conn, err := postgres.NewPostgres(ctx, config)

	db := sqlc.New(conn.Db)

	userRepository := postgres.NewUserRepository(db)

	userService := service.NewUserService(ctx, userRepository)

	userUseCase := userusecase.NewCreateUserUseCase(userService)

	userController := usercontroller.NewController(userUseCase)

	handler := _chi.NewChiHandler()

	httpServer := _http.NewHttpServer(handler, *userController.GetRoutes(), &config.Port)

	if err := httpServer.Start(); err != nil {
		return err
	}

	slog.Info(fmt.Sprintf("Server is running on %s\n", config.Port))
	return nil
}
