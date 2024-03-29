package controllers

import (
	"net/http"

	"github.com/counterapi/api/pkg/repositories"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CounterController is controller for count operations.
type CounterController struct {
	Repository repositories.CounterRepository
}

// UpQuery is query for Counter up params.
type UpQuery struct{}

// DownQuery is query for Counter down params.
type DownQuery struct{}

// SetQuery is query for Counter set params.
type SetQuery struct {
	Count uint `form:"count" json:"count" binding:"required,numeric"`
}

// GetCountsQuery is query for Count params.
type GetCountsQuery struct {
	GroupBy string `form:"group_by" json:"group_by" binding:"required,oneof=hour day week month year"`
	OrderBy string `form:"order_by" json:"order_by" binding:""`
}

// Up increases Counter.
func (c CounterController) Up(ctx *gin.Context) {
	var query UpQuery

	if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	counter, err := c.Repository.IncreaseByName(ctx.Param("namespace"), ctx.Param("counter"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, counter)
}

// Down decreases Counter.
func (c CounterController) Down(ctx *gin.Context) {
	var query DownQuery

	if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	counter, err := c.Repository.DecreaseByName(ctx.Param("namespace"), ctx.Param("counter"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, counter)
}

// Get gets Counter.
func (c CounterController) Get(ctx *gin.Context) {
	var query DownQuery

	if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	counter, err := c.Repository.GetByName(ctx.Param("namespace"), ctx.Param("counter"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, counter)
}

// Set sets Counter.
func (c CounterController) Set(ctx *gin.Context) {
	var query SetQuery

	if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	counter, err := c.Repository.SetByName(ctx.Param("namespace"), ctx.Param("counter"), query.Count)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, counter)
}

// GetCounts gets counts for a counter.
func (c CounterController) GetCounts(ctx *gin.Context) {
	var query GetCountsQuery

	query.GroupBy = "day"

	if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	counts, _ := c.Repository.GroupByCounterNameAndTimeInterval(
		ctx.Param("namespace"),
		ctx.Param("counter"),
		query.GroupBy, query.OrderBy,
	)

	ctx.JSON(http.StatusOK, counts)
}
