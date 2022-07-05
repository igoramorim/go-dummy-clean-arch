package handler

import (
	"github.com/gin-gonic/gin"
	"go-dummy-clean-arch/entity/exceptions"
	"go-dummy-clean-arch/usecase/album"
	"net/http"
)

type albumHandler struct {
	useCase album.UseCase
}

func MapAlbumRoutes(useCase album.UseCase, r *gin.Engine) {
	handler := albumHandler{useCase: useCase}
	r.GET("/albums/:id", handler.GetByID)
}

func (h *albumHandler) GetByID(c *gin.Context) {
	var input struct {
		AlbumID int64 `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		err = exceptions.IllegalArguments(err)
		_ = c.Error(err)
		return
	}

	album, err := h.useCase.GetByID(c.Request.Context(), input.AlbumID)
	if err != nil {
		err = exceptions.Unknown(err)
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, album) // TODO return album presenter
}
