// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/andrewstucki/oapi-codegen version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/andrewstucki/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /json)
	JSONExample(c *gin.Context)

	// (POST /multipart)
	MultipartExample(c *gin.Context)

	// (POST /multiple)
	MultipleRequestAndResponseTypes(c *gin.Context)

	// (POST /reusable-responses)
	ReusableResponses(c *gin.Context)

	// (POST /text)
	TextExample(c *gin.Context)

	// (POST /unknown)
	UnknownExample(c *gin.Context)

	// (POST /unspecified-content-type)
	UnspecifiedContentType(c *gin.Context)

	// (POST /urlencoded)
	URLEncodedExample(c *gin.Context)

	// (POST /with-headers)
	HeadersExample(c *gin.Context, params HeadersExampleParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// JSONExample operation middleware
func (siw *ServerInterfaceWrapper) JSONExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.JSONExample(c)
}

// MultipartExample operation middleware
func (siw *ServerInterfaceWrapper) MultipartExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.MultipartExample(c)
}

// MultipleRequestAndResponseTypes operation middleware
func (siw *ServerInterfaceWrapper) MultipleRequestAndResponseTypes(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.MultipleRequestAndResponseTypes(c)
}

// ReusableResponses operation middleware
func (siw *ServerInterfaceWrapper) ReusableResponses(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.ReusableResponses(c)
}

// TextExample operation middleware
func (siw *ServerInterfaceWrapper) TextExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.TextExample(c)
}

// UnknownExample operation middleware
func (siw *ServerInterfaceWrapper) UnknownExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UnknownExample(c)
}

// UnspecifiedContentType operation middleware
func (siw *ServerInterfaceWrapper) UnspecifiedContentType(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.UnspecifiedContentType(c)
}

// URLEncodedExample operation middleware
func (siw *ServerInterfaceWrapper) URLEncodedExample(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.URLEncodedExample(c)
}

// HeadersExample operation middleware
func (siw *ServerInterfaceWrapper) HeadersExample(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params HeadersExampleParams

	headers := c.Request.Header

	// ------------- Required header parameter "header1" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header1")]; found {
		var Header1 string
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for header1, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "header1", runtime.ParamLocationHeader, valueList[0], &Header1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter header1: %s", err)})
			return
		}

		params.Header1 = Header1

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Header parameter header1 is required, but not found: %s", err)})
		return
	}

	// ------------- Optional header parameter "header2" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header2")]; found {
		var Header2 int
		n := len(valueList)
		if n != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Expected one value for header2, got %d", n)})
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "header2", runtime.ParamLocationHeader, valueList[0], &Header2)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter header2: %s", err)})
			return
		}

		params.Header2 = &Header2

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.HeadersExample(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.POST(options.BaseURL+"/json", wrapper.JSONExample)

	router.POST(options.BaseURL+"/multipart", wrapper.MultipartExample)

	router.POST(options.BaseURL+"/multiple", wrapper.MultipleRequestAndResponseTypes)

	router.POST(options.BaseURL+"/reusable-responses", wrapper.ReusableResponses)

	router.POST(options.BaseURL+"/text", wrapper.TextExample)

	router.POST(options.BaseURL+"/unknown", wrapper.UnknownExample)

	router.POST(options.BaseURL+"/unspecified-content-type", wrapper.UnspecifiedContentType)

	router.POST(options.BaseURL+"/urlencoded", wrapper.URLEncodedExample)

	router.POST(options.BaseURL+"/with-headers", wrapper.HeadersExample)

	return router
}

type BadrequestResponse struct {
}

type ReusableresponseResponseHeaders struct {
	Header1 string
	Header2 int
}
type ReusableresponseJSONResponse struct {
	Body Example

	Headers ReusableresponseResponseHeaders
}

func (t ReusableresponseJSONResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Body)
}

type JSONExampleRequestObject struct {
	Body *JSONExampleJSONRequestBody
}

type JSONExample200JSONResponse Example

func (t JSONExample200JSONResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Example)(t))
}

type JSONExample400Response = BadrequestResponse

type JSONExampledefaultResponse struct {
	StatusCode int
}

type MultipartExampleRequestObject struct {
	Body *multipart.Reader
}

type MultipartExample200MultipartResponse func(writer *multipart.Writer) error

type MultipartExample400Response = BadrequestResponse

type MultipartExampledefaultResponse struct {
	StatusCode int
}

type MultipleRequestAndResponseTypesRequestObject struct {
	JSONBody      *MultipleRequestAndResponseTypesJSONRequestBody
	FormdataBody  *MultipleRequestAndResponseTypesFormdataRequestBody
	Body          io.Reader
	MultipartBody *multipart.Reader
	TextBody      *MultipleRequestAndResponseTypesTextRequestBody
}

type MultipleRequestAndResponseTypes200JSONResponse Example

func (t MultipleRequestAndResponseTypes200JSONResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((Example)(t))
}

type MultipleRequestAndResponseTypes200FormdataResponse Example

type MultipleRequestAndResponseTypes200ImagepngResponse struct {
	Body          io.Reader
	ContentLength int64
}

type MultipleRequestAndResponseTypes200MultipartResponse func(writer *multipart.Writer) error

type MultipleRequestAndResponseTypes200TextResponse string

type MultipleRequestAndResponseTypes400Response = BadrequestResponse

type ReusableResponsesRequestObject struct {
	Body *ReusableResponsesJSONRequestBody
}

type ReusableResponses200JSONResponse = ReusableresponseJSONResponse

type ReusableResponses400Response = BadrequestResponse

type ReusableResponsesdefaultResponse struct {
	StatusCode int
}

type TextExampleRequestObject struct {
	Body *TextExampleTextRequestBody
}

type TextExample200TextResponse string

type TextExample400Response = BadrequestResponse

type TextExampledefaultResponse struct {
	StatusCode int
}

type UnknownExampleRequestObject struct {
	Body io.Reader
}

type UnknownExample200Videomp4Response struct {
	Body          io.Reader
	ContentLength int64
}

type UnknownExample400Response = BadrequestResponse

type UnknownExampledefaultResponse struct {
	StatusCode int
}

type UnspecifiedContentTypeRequestObject struct {
	ContentType string
	Body        io.Reader
}

type UnspecifiedContentType200VideoResponse struct {
	Body          io.Reader
	ContentType   string
	ContentLength int64
}

type UnspecifiedContentType400Response = BadrequestResponse

type UnspecifiedContentType401Response struct {
}

type UnspecifiedContentType403Response struct {
}

type UnspecifiedContentTypedefaultResponse struct {
	StatusCode int
}

type URLEncodedExampleRequestObject struct {
	Body *URLEncodedExampleFormdataRequestBody
}

type URLEncodedExample200FormdataResponse Example

type URLEncodedExample400Response = BadrequestResponse

type URLEncodedExampledefaultResponse struct {
	StatusCode int
}

type HeadersExampleRequestObject struct {
	Params HeadersExampleParams
	Body   *HeadersExampleJSONRequestBody
}

type HeadersExample200ResponseHeaders struct {
	Header1 string
	Header2 int
}

type HeadersExample200JSONResponse struct {
	Body    Example
	Headers HeadersExample200ResponseHeaders
}

func (t HeadersExample200JSONResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Body)
}

type HeadersExample400Response = BadrequestResponse

type HeadersExampledefaultResponse struct {
	StatusCode int
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /json)
	JSONExample(ctx context.Context, request JSONExampleRequestObject) interface{}

	// (POST /multipart)
	MultipartExample(ctx context.Context, request MultipartExampleRequestObject) interface{}

	// (POST /multiple)
	MultipleRequestAndResponseTypes(ctx context.Context, request MultipleRequestAndResponseTypesRequestObject) interface{}

	// (POST /reusable-responses)
	ReusableResponses(ctx context.Context, request ReusableResponsesRequestObject) interface{}

	// (POST /text)
	TextExample(ctx context.Context, request TextExampleRequestObject) interface{}

	// (POST /unknown)
	UnknownExample(ctx context.Context, request UnknownExampleRequestObject) interface{}

	// (POST /unspecified-content-type)
	UnspecifiedContentType(ctx context.Context, request UnspecifiedContentTypeRequestObject) interface{}

	// (POST /urlencoded)
	URLEncodedExample(ctx context.Context, request URLEncodedExampleRequestObject) interface{}

	// (POST /with-headers)
	HeadersExample(ctx context.Context, request HeadersExampleRequestObject) interface{}
}

type StrictHandlerFunc func(ctx *gin.Context, args interface{}) interface{}

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// JSONExample operation middleware
func (sh *strictHandler) JSONExample(ctx *gin.Context) {
	var request JSONExampleRequestObject

	var body JSONExampleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.JSONExample(ctx, request.(JSONExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "JSONExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case JSONExample200JSONResponse:
		ctx.JSON(200, v)
	case JSONExample400Response:
		ctx.Status(400)
	case JSONExampledefaultResponse:
		ctx.Status(v.StatusCode)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// MultipartExample operation middleware
func (sh *strictHandler) MultipartExample(ctx *gin.Context) {
	var request MultipartExampleRequestObject

	if reader, err := ctx.Request.MultipartReader(); err == nil {
		request.Body = reader
	} else {
		ctx.Error(err)
		return
	}

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.MultipartExample(ctx, request.(MultipartExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipartExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case MultipartExample200MultipartResponse:
		writer := multipart.NewWriter(ctx.Writer)
		ctx.Writer.Header().Set("Content-Type", writer.FormDataContentType())
		defer writer.Close()
		if err := v(writer); err != nil {
			ctx.Error(err)
		}
	case MultipartExample400Response:
		ctx.Status(400)
	case MultipartExampledefaultResponse:
		ctx.Status(v.StatusCode)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// MultipleRequestAndResponseTypes operation middleware
func (sh *strictHandler) MultipleRequestAndResponseTypes(ctx *gin.Context) {
	var request MultipleRequestAndResponseTypesRequestObject

	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/json") {
		var body MultipleRequestAndResponseTypesJSONRequestBody
		if err := ctx.Bind(&body); err != nil {
			ctx.Error(err)
			return
		}
		request.JSONBody = &body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/x-www-form-urlencoded") {
		if err := ctx.Request.ParseForm(); err != nil {
			ctx.Error(err)
			return
		}
		var body MultipleRequestAndResponseTypesFormdataRequestBody
		if err := runtime.BindForm(&body, ctx.Request.Form, nil, nil); err != nil {
			ctx.Error(err)
			return
		}
		request.FormdataBody = &body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "image/png") {
		request.Body = ctx.Request.Body
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "multipart/form-data") {
		if reader, err := ctx.Request.MultipartReader(); err == nil {
			request.MultipartBody = reader
		} else {
			ctx.Error(err)
			return
		}
	}
	if strings.HasPrefix(ctx.GetHeader("Content-Type"), "text/plain") {
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.Error(err)
			return
		}
		body := MultipleRequestAndResponseTypesTextRequestBody(data)
		request.TextBody = &body
	}

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.MultipleRequestAndResponseTypes(ctx, request.(MultipleRequestAndResponseTypesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "MultipleRequestAndResponseTypes")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case MultipleRequestAndResponseTypes200JSONResponse:
		ctx.JSON(200, v)
	case MultipleRequestAndResponseTypes200FormdataResponse:
		if form, err := runtime.MarshalForm(v, nil); err != nil {
			ctx.Error(err)
		} else {
			ctx.Data(200, "application/x-www-form-urlencoded", []byte(form.Encode()))
		}
	case MultipleRequestAndResponseTypes200ImagepngResponse:
		if closer, ok := v.Body.(io.ReadCloser); ok {
			defer closer.Close()
		}
		ctx.DataFromReader(200, v.ContentLength, "image/png", v.Body, nil)
	case MultipleRequestAndResponseTypes200MultipartResponse:
		writer := multipart.NewWriter(ctx.Writer)
		ctx.Writer.Header().Set("Content-Type", writer.FormDataContentType())
		defer writer.Close()
		if err := v(writer); err != nil {
			ctx.Error(err)
		}
	case MultipleRequestAndResponseTypes200TextResponse:
		ctx.Data(200, "text/plain", []byte(v))
	case MultipleRequestAndResponseTypes400Response:
		ctx.Status(400)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// ReusableResponses operation middleware
func (sh *strictHandler) ReusableResponses(ctx *gin.Context) {
	var request ReusableResponsesRequestObject

	var body ReusableResponsesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.ReusableResponses(ctx, request.(ReusableResponsesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReusableResponses")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case ReusableResponses200JSONResponse:
		ctx.Header("header1", fmt.Sprint(v.Headers.Header1))
		ctx.Header("header2", fmt.Sprint(v.Headers.Header2))
		ctx.JSON(200, v)
	case ReusableResponses400Response:
		ctx.Status(400)
	case ReusableResponsesdefaultResponse:
		ctx.Status(v.StatusCode)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// TextExample operation middleware
func (sh *strictHandler) TextExample(ctx *gin.Context) {
	var request TextExampleRequestObject

	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Error(err)
		return
	}
	body := TextExampleTextRequestBody(data)
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.TextExample(ctx, request.(TextExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TextExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case TextExample200TextResponse:
		ctx.Data(200, "text/plain", []byte(v))
	case TextExample400Response:
		ctx.Status(400)
	case TextExampledefaultResponse:
		ctx.Status(v.StatusCode)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// UnknownExample operation middleware
func (sh *strictHandler) UnknownExample(ctx *gin.Context) {
	var request UnknownExampleRequestObject

	request.Body = ctx.Request.Body

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.UnknownExample(ctx, request.(UnknownExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnknownExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case UnknownExample200Videomp4Response:
		if closer, ok := v.Body.(io.ReadCloser); ok {
			defer closer.Close()
		}
		ctx.DataFromReader(200, v.ContentLength, "video/mp4", v.Body, nil)
	case UnknownExample400Response:
		ctx.Status(400)
	case UnknownExampledefaultResponse:
		ctx.Status(v.StatusCode)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// UnspecifiedContentType operation middleware
func (sh *strictHandler) UnspecifiedContentType(ctx *gin.Context) {
	var request UnspecifiedContentTypeRequestObject

	request.ContentType = ctx.ContentType()

	request.Body = ctx.Request.Body

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.UnspecifiedContentType(ctx, request.(UnspecifiedContentTypeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UnspecifiedContentType")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case UnspecifiedContentType200VideoResponse:
		if closer, ok := v.Body.(io.ReadCloser); ok {
			defer closer.Close()
		}
		ctx.DataFromReader(200, v.ContentLength, v.ContentType, v.Body, nil)
	case UnspecifiedContentType400Response:
		ctx.Status(400)
	case UnspecifiedContentType401Response:
		ctx.Status(401)
	case UnspecifiedContentType403Response:
		ctx.Status(403)
	case UnspecifiedContentTypedefaultResponse:
		ctx.Status(v.StatusCode)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// URLEncodedExample operation middleware
func (sh *strictHandler) URLEncodedExample(ctx *gin.Context) {
	var request URLEncodedExampleRequestObject

	if err := ctx.Request.ParseForm(); err != nil {
		ctx.Error(err)
		return
	}
	var body URLEncodedExampleFormdataRequestBody
	if err := runtime.BindForm(&body, ctx.Request.Form, nil, nil); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.URLEncodedExample(ctx, request.(URLEncodedExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "URLEncodedExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case URLEncodedExample200FormdataResponse:
		if form, err := runtime.MarshalForm(v, nil); err != nil {
			ctx.Error(err)
		} else {
			ctx.Data(200, "application/x-www-form-urlencoded", []byte(form.Encode()))
		}
	case URLEncodedExample400Response:
		ctx.Status(400)
	case URLEncodedExampledefaultResponse:
		ctx.Status(v.StatusCode)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// HeadersExample operation middleware
func (sh *strictHandler) HeadersExample(ctx *gin.Context, params HeadersExampleParams) {
	var request HeadersExampleRequestObject

	request.Params = params

	var body HeadersExampleJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) interface{} {
		return sh.ssi.HeadersExample(ctx, request.(HeadersExampleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HeadersExample")
	}

	response := handler(ctx, request)

	switch v := response.(type) {
	case HeadersExample200JSONResponse:
		ctx.Header("header1", fmt.Sprint(v.Headers.Header1))
		ctx.Header("header2", fmt.Sprint(v.Headers.Header2))
		ctx.JSON(200, v)
	case HeadersExample400Response:
		ctx.Status(400)
	case HeadersExampledefaultResponse:
		ctx.Status(v.StatusCode)
	case error:
		ctx.Error(v)
	case nil:
	default:
		ctx.Error(fmt.Errorf("Unexpected response type: %T", v))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xYS4/iOBD+K1btnkaB0D194rbTGmnfI9Ezp9UcirgAzya2166QRoj/vnJsaBjSCFo8",
	"pNXeEqde/qq+KsdLKExljSbNHoZLcOSt0Z7alzFKR//U5Dm8SfKFU5aV0TCEDyhH6dsqA0e1x3FJa/Ug",
	"XxjNpFtVtLZUBQbV/JsP+kvwxYwqDE8/OprAEH7IX0LJ41ef0zNWtiRYrVbZdxF8+g0ymBFKcm208fFu",
	"1zYvLMEQPDulpxCMRLH7TjGlmabkgrcgmoIIAus4hkuwzlhyrCJGcyxr6vaUVsz4GxUcd6D0xOxj+Wg0",
	"o9JeSDWZkCPNIoEngg0vfG2tcUxSjBcieChYeHJzcpABKw6BwdP2ukgBe8hgTs5HR3f9QX8Q8mUsabQK",
	"hvC+XcrAIs/aDW0SZE1X3n99+vSnUF5gzaZCVgWW5UJU6PwMy5KkUJpNiLEu2PehdeXazP8ik/rHhGUo",
	"m7aCPhi5uETFtIW5Vc/3g8GVCnOVwUN01mVjE1S+xbDWzATrsgP0L/pvbRotyDnj0s7yqi5ZWXS8naxd",
	"tP9YixwD+cZePjGu6klkvBDq5/J0U+BTM+gkydPMNF7MTCPYCElYikbxTKwVv2O30gKFV3paklgHlXVm",
	"sqTUc3/ScpT28jnYuDiXsh0rz72maXpt8mpXki6MJPk2s6rCKeVWT3fVg21kGMJ4waFs97vrmYooA6Zn",
	"zm2JSh8eHVdqJ/8jfTZiR7quzya9neR1E3dNKi8K1GIc+DjxgcRdvvZIOkqeRlsStxlxhzHaO61do2uG",
	"5L8+qT7T81FD6oxkvXY1ngpYHRdfxyxpHQPbG7l/BIpzJcnklX040fLNQPWWCjVRJHtpF70Y22st4dHo",
	"whHvDu1wAtaGxcZYOJjzjEREIBPeiIZEVXsWFr0XitsuUqp4uJe01zy+vET2GD2FyX5EVt9dKKfvbpXR",
	"h8Hd6SrvL1w3O8P3FT6Ofv8YZU79wznblD/xjHI+vzeiczhW97buALop/HMUeJnpBak5SYFaCkdcO01S",
	"zBWuf1v3uJkMvKTVosOKuPX61xLCCEkXC5CBxoo273epCJQLyLKrKTt0PXHQ1j1kh+4svv6Hf6gvedNz",
	"6TpdZRAvZWKx1K4MGWW2wzyPlzl93+B0Sq6vTI5Wwerr6t8AAAD//ygSomqZEwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
