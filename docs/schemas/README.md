# Schemas

## Overview

REST APIs for managing Schema entities

### Available Operations

* [DeleteSchema](#deleteschema) - Delete a particular schema revision for an Api.
* [DownloadSchema](#downloadschema) - Download the latest schema for a particular apiID.
* [DownloadSchemaRevision](#downloadschemarevision) - Download a particular schema revision for an Api.
* [GetSchema](#getschema) - Get information about the latest schema.
* [GetSchemaDiff](#getschemadiff) - Get a diff of two schema revisions for an Api.
* [GetSchemaRevision](#getschemarevision) - Get information about a particular schema revision for an Api.
* [GetSchemas](#getschemas) - Get information about all schemas associated with a particular apiID.
* [RegisterSchema](#registerschema) - Register a schema.

## DeleteSchema

Delete a particular schema revision for an Api.

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
    res, err := s.Schemas.DeleteSchema(ctx, operations.DeleteSchemaRequest{
        APIID: "ipsa",
        RevisionID: "omnis",
        VersionID: "voluptate",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## DownloadSchema

Download the latest schema for a particular apiID.

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
    res, err := s.Schemas.DownloadSchema(ctx, operations.DownloadSchemaRequest{
        APIID: "cum",
        VersionID: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

## DownloadSchemaRevision

Download a particular schema revision for an Api.

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
    res, err := s.Schemas.DownloadSchemaRevision(ctx, operations.DownloadSchemaRevisionRequest{
        APIID: "doloremque",
        RevisionID: "reprehenderit",
        VersionID: "ut",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

## GetSchema

Returns information about the last uploaded schema for a particular API version. 
This won't include the schema itself, that can be retrieved via the downloadSchema operation.

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
    res, err := s.Schemas.GetSchema(ctx, operations.GetSchemaRequest{
        APIID: "maiores",
        VersionID: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

## GetSchemaDiff

Get a diff of two schema revisions for an Api.

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
    res, err := s.Schemas.GetSchemaDiff(ctx, operations.GetSchemaDiffRequest{
        APIID: "corporis",
        BaseRevisionID: "dolore",
        TargetRevisionID: "iusto",
        VersionID: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaDiff != nil {
        // handle response
    }
}
```

## GetSchemaRevision

Returns information about the last uploaded schema for a particular schema revision. 
This won't include the schema itself, that can be retrieved via the downloadSchema operation.

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
    res, err := s.Schemas.GetSchemaRevision(ctx, operations.GetSchemaRevisionRequest{
        APIID: "harum",
        RevisionID: "enim",
        VersionID: "accusamus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

## GetSchemas

Returns information the schemas associated with a particular apiID. 
This won't include the schemas themselves, they can be retrieved via the downloadSchema operation.

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
    res, err := s.Schemas.GetSchemas(ctx, operations.GetSchemasRequest{
        APIID: "commodi",
        VersionID: "repudiandae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Schemata != nil {
        // handle response
    }
}
```

## RegisterSchema

Allows uploading a schema for a particular API version.
This will be used to populate ApiEndpoints and used as a base for any schema generation if present.

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
    res, err := s.Schemas.RegisterSchema(ctx, operations.RegisterSchemaRequest{
        RequestBody: operations.RegisterSchemaRequestBody{
            File: operations.RegisterSchemaRequestBodyFile{
                Content: []byte("quae"),
                File: "ipsum",
            },
        },
        APIID: "quidem",
        VersionID: "molestias",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
