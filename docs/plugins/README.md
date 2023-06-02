# Plugins

## Overview

REST APIs for managing and running plugins

### Available Operations

* [GetPlugins](#getplugins) - Get all plugins for the current workspace.
* [RunPlugin](#runplugin) - Run a plugin
* [UpsertPlugin](#upsertplugin) - Upsert a plugin

## GetPlugins

Get all plugins for the current workspace.

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
    res, err := s.Plugins.GetPlugins(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plugins != nil {
        // handle response
    }
}
```

## RunPlugin

Run a plugin

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Speakeasy-API"
	"Speakeasy-API/pkg/models/operations"
	"Speakeasy-API/pkg/models/shared"
)

func main() {
    s := speakeasyapi.New(
        speakeasyapi.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Plugins.RunPlugin(ctx, operations.RunPluginRequest{
        Filters: &shared.Filters{
            Filters: []shared.Filter{
                shared.Filter{
                    Key: "repellat",
                    Operator: "mollitia",
                    Value: "occaecati",
                },
            },
            Limit: 253291,
            Offset: 414369,
            Operator: "quam",
        },
        PluginID: "molestiae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BoundedRequests != nil {
        // handle response
    }
}
```

## UpsertPlugin

Upsert a plugin

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Speakeasy-API"
	"Speakeasy-API/pkg/models/shared"
	"Speakeasy-API/pkg/types"
)

func main() {
    s := speakeasyapi.New(
        speakeasyapi.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Plugins.UpsertPlugin(ctx, shared.Plugin{
        Code: "velit",
        CreatedAt: types.MustTimeFromString("2022-09-06T22:51:09.401Z"),
        EvalHash: speakeasyapi.String("quis"),
        PluginID: "vitae",
        Title: "Miss",
        UpdatedAt: types.MustTimeFromString("2022-05-14T10:37:30.872Z"),
        WorkspaceID: "odit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Plugin != nil {
        // handle response
    }
}
```
