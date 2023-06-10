# Apis

## Overview

REST APIs for managing Api entities

### Available Operations

* [DeleteAPI](#deleteapi) - Delete an Api.
* [GenerateOpenAPISpec](#generateopenapispec) - Generate an OpenAPI specification for a particular Api.
* [GeneratePostmanCollection](#generatepostmancollection) - Generate a Postman collection for a particular Api.
* [GetAllAPIVersions](#getallapiversions) - Get all Api versions for a particular ApiEndpoint.
* [GetApis](#getapis) - Get a list of Apis for a given workspace
* [UpsertAPI](#upsertapi) - Upsert an Api

## DeleteAPI

Delete a particular version of an Api. The will also delete all associated ApiEndpoints, Metadata, Schemas & Request Logs (if using a Postgres datastore).

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Speakeasy-API"
	"Speakeasy-API/pkg/models/operations"
)

func main() {
    s := speakeasyapi.New(
        speakeasyapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apis.DeleteAPI(ctx, operations.DeleteAPIRequest{
        APIID: "quod",
        VersionID: "esse",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DeleteAPIRequest](../../models/operations/deleteapirequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.DeleteAPIResponse](../../models/operations/deleteapiresponse.md), error**


## GenerateOpenAPISpec

This endpoint will generate any missing operations in any registered OpenAPI document if the operation does not already exist in the document.
Returns the original document and the newly generated document allowing a diff to be performed to see what has changed.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Speakeasy-API"
	"Speakeasy-API/pkg/models/operations"
)

func main() {
    s := speakeasyapi.New(
        speakeasyapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apis.GenerateOpenAPISpec(ctx, operations.GenerateOpenAPISpecRequest{
        APIID: "totam",
        VersionID: "porro",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerateOpenAPISpecDiff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GenerateOpenAPISpecRequest](../../models/operations/generateopenapispecrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GenerateOpenAPISpecResponse](../../models/operations/generateopenapispecresponse.md), error**


## GeneratePostmanCollection

Generates a postman collection containing all endpoints for a particular API. Includes variables produced for any path/query/header parameters included in the OpenAPI document.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Speakeasy-API"
	"Speakeasy-API/pkg/models/operations"
)

func main() {
    s := speakeasyapi.New(
        speakeasyapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apis.GeneratePostmanCollection(ctx, operations.GeneratePostmanCollectionRequest{
        APIID: "dolorum",
        VersionID: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostmanCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GeneratePostmanCollectionRequest](../../models/operations/generatepostmancollectionrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GeneratePostmanCollectionResponse](../../models/operations/generatepostmancollectionresponse.md), error**


## GetAllAPIVersions

Get all Api versions for a particular ApiEndpoint.
Supports filtering the versions based on metadata attributes.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Speakeasy-API"
	"Speakeasy-API/pkg/models/operations"
)

func main() {
    s := speakeasyapi.New(
        speakeasyapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apis.GetAllAPIVersions(ctx, operations.GetAllAPIVersionsRequest{
        APIID: "nam",
        Metadata: map[string][]string{
            "occaecati": []string{
                "deleniti",
            },
            "hic": []string{
                "totam",
                "beatae",
                "commodi",
                "molestiae",
            },
            "modi": []string{
                "impedit",
            },
        },
        Op: &operations.GetAllAPIVersionsOp{
            And: false,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Apis != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetAllAPIVersionsRequest](../../models/operations/getallapiversionsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetAllAPIVersionsResponse](../../models/operations/getallapiversionsresponse.md), error**


## GetApis

Get a list of all Apis and their versions for a given workspace.
Supports filtering the APIs based on metadata attributes.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"Speakeasy-API"
	"Speakeasy-API/pkg/models/operations"
)

func main() {
    s := speakeasyapi.New(
        speakeasyapi.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{
        Metadata: map[string][]string{
            "esse": []string{
                "excepturi",
            },
            "aspernatur": []string{
                "ad",
            },
            "natus": []string{
                "iste",
            },
        },
        Op: &operations.GetApisOp{
            And: false,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Apis != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetApisRequest](../../models/operations/getapisrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetApisResponse](../../models/operations/getapisresponse.md), error**


## UpsertAPI

Upsert an Api. If the Api does not exist, it will be created.
If the Api exists, it will be updated.

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
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Apis.UpsertAPI(ctx, operations.UpsertAPIRequest{
        APIInput: shared.APIInput{
            APIID: "dolor",
            Description: "natus",
            MetaData: map[string][]string{
                "hic": []string{
                    "fuga",
                    "in",
                    "corporis",
                    "iste",
                },
                "iure": []string{
                    "quidem",
                    "architecto",
                    "ipsa",
                    "reiciendis",
                },
            },
            VersionID: "est",
        },
        APIID: "mollitia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.API != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.UpsertAPIRequest](../../models/operations/upsertapirequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.UpsertAPIResponse](../../models/operations/upsertapiresponse.md), error**

