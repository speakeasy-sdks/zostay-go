// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Speakeasy-API/pkg/models/shared"
	"net/http"
)

type GenerateRequestPostmanCollectionRequest struct {
	// The ID of the request to retrieve.
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GenerateRequestPostmanCollectionResponse struct {
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	PostmanCollection []byte
	StatusCode        int
	RawResponse       *http.Response
}
