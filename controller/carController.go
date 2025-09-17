package controller

import (
	"net/http"

	"github.com/car-api/services"
	"github.com/gin-gonic/gin"
)

// Helloworld hello world
//
//	@Summary		hello world
//	@Description	printing hello world
//	@Tags			Helloworld
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @Summary		User Registration
// @Description	Register the User into DB
// @Tags			SignUP
// @Accept			json
// @Produce		json
// @Param			request	body		model.Users				true	"User details for creation"
// @Success		200		{object}	map[string]interface{}	"Signup"
// @Router			/signup [post]
func SignUp(g *gin.Context) {

	services.SignUp(g)

}
