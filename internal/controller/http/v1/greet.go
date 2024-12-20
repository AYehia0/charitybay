package v1

import (
	"charitybay/internal/entity"
	"charitybay/internal/usecase"
	"charitybay/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type greetsRoutes struct {
	t usecase.Greet
	l logger.LoggerI
}

func newGreetingRoutes(handler *gin.RouterGroup, t usecase.Greet, l logger.LoggerI) {
	r := &greetsRoutes{t, l}

	h := handler.Group("/greetings")
	{
		h.GET("/", r.getGreets)
		h.POST("/", r.createGreet)
	}
}

type greetResponse struct {
	Greet []entity.Greet `json:"greets"`
}

// @Summary     Show all greets
// @Description Show all greets
// @ID          get-greets
// @Tags	    greet
// @Produce     json
// @Success     200 {object} greetResponse
// @Failure     500 {object} response
// @Router      /greets [get]
func (r *greetsRoutes) getGreets(c *gin.Context) {
	greets, err := r.t.GetGreets(c.Request.Context())
	if err != nil {
		r.l.Errorf("http - v1 - getGreets: %v", err)
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	c.JSON(http.StatusOK, greetResponse{greets})
}

// @Summary     Create a greet
// @Description Create a greet
// @ID          create-greet
// @Tags	    greet
// @Accept      json
// @Produce     json
// @Param       greet body entity.Greet true "Greet object"
// @Success     201 {object} response
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /greets [post]
func (r *greetsRoutes) createGreet(c *gin.Context) {
	var greet entity.Greet
	if err := c.ShouldBindJSON(&greet); err != nil {
		r.l.Errorf("http - v1 - createGreet: %v", err)
		errorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	err := r.t.CreateGreet(c.Request.Context(), greet)
	if err != nil {
		r.l.Errorf("http - v1 - createGreet: %v", err)
		errorResponse(c, http.StatusInternalServerError, "database problems")
		return
	}

	c.JSON(http.StatusCreated, response{})
}
