package main

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/soap", handleRequest)

	router.Run(":7000")
}

func handleRequest(ctx *gin.Context) {
	var request Request
	err := ctx.ShouldBindXML(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	resp, err := createResponse()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.Data(http.StatusOK, "Content-Type: text/html; charset=utf-8", resp)
}

func createResponse() ([]byte, error) {
	var resp bytes.Buffer

	template, err := template.ParseFiles("response.html")
	if err != nil {
		return resp.Bytes(), err
	}

	// hard coded responce
	template.Execute(&resp, GetPriceResponse{
		Price: 1.6,
	})

	return resp.Bytes(), nil
}

type Request struct {
	Envalop  xml.Name `xml:"Envelope"`
	GetPrice GetPrice `xml:"Body>GetPrice"`
}

type GetPrice struct {
	Item string `xml:"Item"`
}

type GetPriceResponse struct {
	Price float64 `xml:"Price"`
}
