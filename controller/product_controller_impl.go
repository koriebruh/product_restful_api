package controller

import (
	"github.com/gin-gonic/gin"
	"korie/api-product/helper"
	"korie/api-product/model/web"
	"korie/api-product/service"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller ProductControllerImpl) Create(c *gin.Context) {
	/// SAVE HASIL DECODE
	productCreateRequest := web.ProductCreateRequest{}
	err := c.ShouldBindJSON(&productCreateRequest) // DECODE JSON FROM REQ BODY
	helper.IfErrNotNul(err)

	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT
	ProductResponse := controller.ProductService.Create(c.Request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	/// GIVE OUTPUT TO CLIENT
	c.JSON(200, webResponse)
}

func (controller ProductControllerImpl) Update(c *gin.Context) {
	/// SAVE HASIL DECODE
	productUpdateRequest := web.ProductUpdateRequest{}
	err := c.ShouldBindJSON(&productUpdateRequest) // DECODE JSON FROM REQ BODY
	helper.IfErrNotNul(err)

	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT )
	productId := c.Param("id") /// MENGAMBIL PARAM
	id, err := strconv.Atoi(productId)
	helper.IfErrNotNul(err)
	productUpdateRequest.Id = id

	ProductResponse := controller.ProductService.Update(c.Request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	/// GIVE OUTPUT TO CLIENT
	c.JSON(200, webResponse)

}

func (controller ProductControllerImpl) Delete(c *gin.Context) {
	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT )
	productId := c.Param("id") /// MENGAMBIL PARAM
	id, err := strconv.Atoi(productId)
	helper.IfErrNotNul(err)

	controller.ProductService.Delete(c.Request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	/// GIVE OUTPUT TO CLIENT
	c.JSON(200, webResponse)

}

func (controller ProductControllerImpl) FindById(c *gin.Context) {
	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT )
	productId := c.Param("id") /// MENGAMBIL PARAM
	id, err := strconv.Atoi(productId)
	helper.IfErrNotNul(err)

	productResponse := controller.ProductService.FindById(c.Request.Context(), id)
	//if productResponse == (web.ProductResponse{}) {
	//	return
	//}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	/// GIVE OUTPUT TO CLIENT
	c.JSON(200, webResponse)

}

func (controller ProductControllerImpl) FindAll(c *gin.Context) {
	productResponses := controller.ProductService.FindAll(c.Request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	/// GIVE OUTPUT TO CLIENT
	c.JSON(200, webResponse)

}
