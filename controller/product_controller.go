package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController interface {
	Create(writer http.ResponseWriter, request *http.Request, c *gin.Context)
	Update(writer http.ResponseWriter, request *http.Request, c *gin.Context)
	Delete(writer http.ResponseWriter, request *http.Request, c *gin.Context)
	FindById(writer http.ResponseWriter, request *http.Request, c *gin.Context)
	FindAll(writer http.ResponseWriter, request *http.Request, c *gin.Context)
}
