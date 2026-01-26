# go-aml

Go client for the TDCC AML (Anti-Money Laundering) APIs.

## Features

- Name check (single query)
- Query remaining/over limit count
- Batch upload/download endpoints (planned)

## Installation

```bash
go get github.com/flc1125/go-aml
```

## Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/flc1125/go-aml"
)

func main() {
	client, err := aml.NewClient("ACCOUNT", "PASSWORD")
	if err != nil {
		panic(err)
	}

	// Single name check
	resp, _, err := client.CheckName(context.Background(), &aml.CheckNameRequest{
		EnglishName: aml.Ptr("John Doe"),
		Nationality: aml.Ptr("US"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("RCScore:", resp.RCScore)

	// Query remaining / over-limit count
	overResp, _, err := client.QueryOver(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Remain:", overResp.Remain, "Over:", overResp.Over)
}
```

## Documentation

See `docs/api.md` for the API spec summary.

## License

MIT. See `LICENSE`.
