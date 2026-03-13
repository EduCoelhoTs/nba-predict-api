package bootstrap

import (
	"context"

	"github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/repository/postgres"
	authcontroller "github.com/EduCoelhoTs/base-hex-arq-api/internal/infra/controller/auth"
	usercontroller "github.com/EduCoelhoTs/base-hex-arq-api/internal/infra/controller/user"
)

type App struct {
	UserController usercontroller.Controller
	AuthController authcontroller.Controller
}

func NewApp(ctx context.Context, db postgres.QueriesPort) *App {
	return &App{
		UserController: NewUserModule(ctx, db),
		A
	}
}
