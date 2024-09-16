package controller

import (
	"github.com/gin-gonic/gin"
	"korie/api-product/helper"
	"korie/api-product/model/web"
	"korie/api-product/service"
	"net/http"
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

func (controller ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, c *gin.Context) {
	/// SAVE HASIL DECODE
	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)

	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT
	ProductResponse := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	/// GIVE OUTPUT TO CLIENT
	helper.WriteToResponseJson(writer, webResponse)
}

func (controller ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, c *gin.Context) {
	/// SAVE HASIL DECODE
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT )
	productId := c.Param("id") /// MENGAMBIL PARAM
	id, err := strconv.Atoi(productId)
	helper.IfErrNotNul(err)
	productUpdateRequest.Id = id

	ProductResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	/// GIVE OUTPUT TO CLIENT
	helper.WriteToResponseJson(writer, webResponse)

}

func (controller ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, c *gin.Context) {
	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT )
	productId := c.Param("id") /// MENGAMBIL PARAM
	id, err := strconv.Atoi(productId)
	helper.IfErrNotNul(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	/// GIVE OUTPUT TO CLIENT
	helper.WriteToResponseJson(writer, webResponse)

}

func (controller ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, c *gin.Context) {
	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT )
	productId := c.Param("id") /// MENGAMBIL PARAM
	id, err := strconv.Atoi(productId)
	helper.IfErrNotNul(err)

	productResponse := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	/// GIVE OUTPUT TO CLIENT
	helper.WriteToResponseJson(writer, webResponse)

}

func (controller ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, c *gin.Context) {
	productResponses := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	/// GIVE OUTPUT TO CLIENT
	helper.WriteToResponseJson(writer, webResponse)

}
