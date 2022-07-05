package handler

import (
	"github.com/gin-gonic/gin"
	"go-dummy-clean-arch/api/presenter"
	"go-dummy-clean-arch/entity/exceptions"
	"go-dummy-clean-arch/usecase/artist"
	"net/http"
)

type artistHandler struct {
	useCase artist.UseCase
}

func MapArtistRoutes(useCase artist.UseCase, r *gin.Engine) {
	handler := artistHandler{useCase: useCase}
	r.GET("/artists/:id", handler.GetByID)
	r.POST("/artists", handler.Save)
}

func (h *artistHandler) GetByID(c *gin.Context) {
	var input struct {
		ArtistID int64 `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		err = exceptions.IllegalArguments(err)
		_ = c.Error(err)
		return
	}

	artist, err := h.useCase.GetByID(c.Request.Context(), input.ArtistID)
	if err != nil {
		err = exceptions.Unknown(err)
		_ = c.Error(err)
		return
	}

	var json = presenter.Artist{}
	json.FromEntity(&artist)
	c.JSON(http.StatusOK, json)
}

func (h *artistHandler) Save(c *gin.Context) {
	var input presenter.Artist // TODO: i don't think using the presenter here is the best idea
	err := c.ShouldBindJSON(&input)
	if err != nil {
		err = exceptions.IllegalArguments(err)
		_ = c.Error(err)
		return
	}
	artist := input.ToEntity()

	err = h.useCase.Save(c.Request.Context(), *artist)
	if err != nil {
		err = exceptions.Unknown(err)
		_ = c.Error(err)
		return
	}
}
