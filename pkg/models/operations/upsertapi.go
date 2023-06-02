// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Speakeasy-API/pkg/models/shared"
	"net/http"
)

type UpsertAPIRequest struct {
	// A JSON representation of the Api to upsert
	APIInput shared.APIInput `request:"mediaType=application/json"`
	// The ID of the Api to upsert.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type UpsertAPIResponse struct {
	// OK
	API         *shared.API
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
