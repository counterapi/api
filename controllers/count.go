package controllers

import (
	"github.com/counterapi/counter/repositories"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type CountController struct {
	Repository repositories.CountRepository
}

type GetCountsQuery struct {
	Name string `form:"name" json:"name" binding:"required,alphanum,max=100"`
}

// Up increases Counter.
func (c CountController) GetCounts(ctx *gin.Context) {
	var query GetCountsQuery

	if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	counts, _ := c.Repository.ListByCounterName(query.Name)

	ctx.JSON(http.StatusOK, counts)

	return
}
