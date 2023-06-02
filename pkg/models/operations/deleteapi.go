// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Speakeasy-API/pkg/models/shared"
	"net/http"
)

type DeleteAPIRequest struct {
	// The ID of the Api to delete.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to delete.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIResponse struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
