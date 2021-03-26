package controllers

import (
	"github.com/counterapi/counter/pkg/repositories"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type CounterController struct {
	Repository repositories.CounterRepository
}

type UpQuery struct {
	Name string `form:"name" json:"name" binding:"required,alphanum,max=100"`
}

type DownQuery struct {
	Name string `form:"name" json:"name" binding:"required,alphanum"`
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

	err := c.Repository.IncreaseByName(query.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "increased",
	})

	return
}

// Down decreases Counter.
func (c CounterController) Down(ctx *gin.Context) {
	var query DownQuery

	if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err := c.Repository.DecreaseByName(query.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "decreased",
	})

	return
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

	counter, err := c.Repository.GetByName(query.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, counter)

	return
}
