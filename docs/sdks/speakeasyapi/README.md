# SpeakeasyAPI SDK

## Overview

Speakeasy API: The Speakeasy API allows teams to manage common operations with their APIs

The Speakeasy Platform Documentation
<https://docs.speakeasyapi.dev>
### Available Operations

* [ValidateAPIKey](#validateapikey) - Validate the current api key.

## ValidateAPIKey

Validate the current api key.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Speakeasy-API"
)

func main() {
    s := speakeasyapi.New(
        speakeasyapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SpeakeasyAPI.ValidateAPIKey(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ValidateAPIKeyResponse](../../models/operations/validateapikeyresponse.md), error**

