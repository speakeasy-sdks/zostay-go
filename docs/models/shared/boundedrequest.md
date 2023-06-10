# BoundedRequest

A BoundedRequest is a request that has been logged by the Speakeasy without the contents of the request.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `APIEndpointID`                                             | *string*                                                    | :heavy_check_mark:                                          | The ID of the ApiEndpoint this request was made to.         |
| `APIID`                                                     | *string*                                                    | :heavy_check_mark:                                          | The ID of the Api this request was made to.                 |
| `CreatedAt`                                                 | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | Creation timestamp.                                         |
| `CustomerID`                                                | *string*                                                    | :heavy_check_mark:                                          | The ID of the customer that made this request.              |
| `Latency`                                                   | *int64*                                                     | :heavy_check_mark:                                          | The latency of the request.                                 |
| `Metadata`                                                  | [][RequestMetadata](../../models/shared/requestmetadata.md) | :heavy_minus_sign:                                          | Metadata associated with this request                       |
| `Method`                                                    | *string*                                                    | :heavy_check_mark:                                          | HTTP verb.                                                  |
| `Path`                                                      | *string*                                                    | :heavy_check_mark:                                          | The path of the request.                                    |
| `RequestFinishTime`                                         | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | The time the request finished.                              |
| `RequestID`                                                 | *string*                                                    | :heavy_check_mark:                                          | The ID of this request.                                     |
| `RequestStartTime`                                          | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | The time the request was made.                              |
| `Status`                                                    | *int64*                                                     | :heavy_check_mark:                                          | The status code of the request.                             |
| `VersionID`                                                 | *string*                                                    | :heavy_check_mark:                                          | The version ID of the Api this request was made to.         |
| `WorkspaceID`                                               | *string*                                                    | :heavy_check_mark:                                          | The workspace ID this request was made to.                  |