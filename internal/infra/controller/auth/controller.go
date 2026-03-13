package authcontroller

import (
	"fmt"
	"net/http"

	_http "github.com/EduCoelhoTs/base-hex-arq-api/internal/adapter/http"
	loginusecase "github.com/EduCoelhoTs/base-hex-arq-api/internal/application/usecase/auth/login"
	infraerrors "github.com/EduCoelhoTs/base-hex-arq-api/internal/infra/error"
	"github.com/EduCoelhoTs/base-hex-arq-api/pkg/xjson"
	"github.com/EduCoelhoTs/base-hex-arq-api/pkg/xvalidator"
)

type Controller interface {
	GetRoutes() _http.Routes
	Login(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	LoginUseCase loginusecase.LoginUseCase
}

func NewController(loginUseCase loginusecase.LoginUseCase) *controller {
	return &controller{LoginUseCase: loginUseCase}
}

func (c *controller) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var body LoginRequestDTO

	if err := xjson.Decode(r.Body, &body); err != nil {
		fmt.Printf("logincontroller.create decode body error: %v", err)
		xjson.ResponseHttpError(w, http.StatusBadRequest, infraerrors.ERR_INTERNAL_SERVER_ERROR)
	}

	if err := xvalidator.ValidateStruct(body); err != nil {
		fmt.Printf("logincontroller.create validate body error: %v", err)
		xjson.ResponseHttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input := loginusecase.LoginInput{
		Email:    body.Email,
		Password: body.Password,
	}

	token, err := c.LoginUseCase.Execute(r.Context(), input)
	if err != nil {
		fmt.Printf("logincontroller.create execute usecase error: %v", err)
		xjson.ResponseHttpError(w, http.StatusInternalServerError, infraerrors.ERR_INTERNAL_SERVER_ERROR)
		return
	}

	xjson.ReponseHttp(w, http.StatusOK, LoginResponseDTO{Token: token})

}
