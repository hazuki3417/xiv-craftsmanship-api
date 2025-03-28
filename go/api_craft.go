// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * xiv-craftsmanship-api
 *
 * xiv-craftsmanship-api
 *
 * API version: 1.0.0
 * Contact: hazuki3417@gmail.com
 */

package openapi

import (
	"net/http"
	"strings"
)

// CraftAPIController binds http requests to an api service and writes the service results to the http response
type CraftAPIController struct {
	service CraftAPIServicer
	errorHandler ErrorHandler
}

// CraftAPIOption for how the controller is set up.
type CraftAPIOption func(*CraftAPIController)

// WithCraftAPIErrorHandler inject ErrorHandler into controller
func WithCraftAPIErrorHandler(h ErrorHandler) CraftAPIOption {
	return func(c *CraftAPIController) {
		c.errorHandler = h
	}
}

// NewCraftAPIController creates a default api controller
func NewCraftAPIController(s CraftAPIServicer, opts ...CraftAPIOption) *CraftAPIController {
	controller := &CraftAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CraftAPIController
func (c *CraftAPIController) Routes() Routes {
	return Routes{
		"GetCraft": Route{
			strings.ToUpper("Get"),
			"/craft",
			c.GetCraft,
		},
	}
}

// GetCraft - craft
func (c *CraftAPIController) GetCraft(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var nameParam string
	if query.Has("name") {
		param := query.Get("name")

		nameParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "name"}, nil)
		return
	}
	result, err := c.service.GetCraft(r.Context(), nameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
