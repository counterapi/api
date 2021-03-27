package controllers

import (
	"net/http"

	"github.com/counterapi/counter/pkg/repositories"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CountController is controller for count operations.
type CountController struct {
	Repository repositories.CountRepository
}

// GetCountsQuery is query for Count params.
type GetCountsQuery struct {
	Name    string `form:"name" json:"name" binding:"required,alphanum,max=100"`
	GroupBy string `form:"group_by" json:"group_by" binding:"required"`
	OrderBy string `form:"order_by" json:"order_by" binding:""`
}

// GetCounts gets counts for a counter.
func (c CountController) GetCounts(ctx *gin.Context) {
	var query GetCountsQuery

	if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	counts, _ := c.Repository.GroupByCounterNameAndTimeInterval(query.Name, query.GroupBy, query.OrderBy)

	ctx.JSON(http.StatusOK, counts)
}
