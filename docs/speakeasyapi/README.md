# SpeakeasyAPI SDK

## Overview

The Speakeasy API allows teams to manage common operations with their APIs

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
            APIKey: "YOUR_API_KEY_HERE",
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
