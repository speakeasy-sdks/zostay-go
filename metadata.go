// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package speakeasyapi

import (
	"Speakeasy-API/pkg/models/operations"
	"Speakeasy-API/pkg/models/shared"
	"Speakeasy-API/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// metadata - REST APIs for managing Version Metadata entities
type metadata struct {
	sdkConfiguration sdkConfiguration
}

func newMetadata(sdkConfig sdkConfiguration) *metadata {
	return &metadata{
		sdkConfiguration: sdkConfig,
	}
}

// DeleteVersionMetadata - Delete metadata for a particular apiID and versionID.
func (s *metadata) DeleteVersionMetadata(ctx context.Context, request operations.DeleteVersionMetadataRequest) (*operations.DeleteVersionMetadataResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/metadata/{metaKey}/{metaValue}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteVersionMetadataResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// GetVersionMetadata - Get all metadata for a particular apiID and versionID.
func (s *metadata) GetVersionMetadata(ctx context.Context, request operations.GetVersionMetadataRequest) (*operations.GetVersionMetadataResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/metadata", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetVersionMetadataResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.VersionMetadata
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.VersionMetadata = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}

// InsertVersionMetadata - Insert metadata for a particular apiID and versionID.
func (s *metadata) InsertVersionMetadata(ctx context.Context, request operations.InsertVersionMetadataRequest) (*operations.InsertVersionMetadataResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/metadata", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "VersionMetadataInput", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.InsertVersionMetadataResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.VersionMetadata
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.VersionMetadata = out
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Error = out
		}
	}

	return res, nil
}
