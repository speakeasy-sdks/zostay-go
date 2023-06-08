<!-- Start SDK Example Usage -->
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
            "provident": []string{
                "quibusdam",
                "unde",
                "nulla",
            },
            "corrupti": []string{
                "vel",
                "error",
                "deserunt",
                "suscipit",
            },
            "iure": []string{
                "debitis",
                "ipsa",
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
<!-- End SDK Example Usage -->