# QuickNode Golang SDK

:bulb: This implementation is based on the [official Typescript SDK](https://github.com/quiknode-labs/qn-oss)
made by QuickNode.

## Description

This SDK aims to extend the current QN SDK to another language so QN can be integrated
to more projects.

The project is split in two parts : a [client](./client) and [queries](./queries).

Client is a simplistic initialization of a GraphQL client with the possibility
to add IcyTools API Key in order to increase rate limit.
This client wrap all implemented queries to mimic the same behavior as the
official SDK.

Queries are pure functions that execute GraphQL query, thanks [Golang interface
system](https://gobyexample.com/interfaces), you can either use the SDK client
or provide your own with personalised configuration.
Each query is implemented using the same pattern:

1. The raw GraphQL query is defined.
2. Input and Output types are declared.
3. A pure function insert arguments, execute the query and extract outputs.

## Quickstart

```golang
package main

import (
  "context"
  "fmt"
  "log"

  "github.com/quartz-technology/quicknode-sdks/go-sdk/client"
)

func main() {
  c := client.New()

  collection, err := c.GetCollectionDetails(context.TODO(), "0x60e4d786628fea6478f785a6d7e704777c86a7c6")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%#v\n", collection)
}
```

:bulb: You can also check [tests files](./queries/nft) to see usage for each queries.

## Supported Queries

| Methods                                                                     | Status |
|-----------------------------------------------------------------------------|--------|
| [getCollectionDetails](./queries/nft/getCollectionDetails.go)               | TESTED |
| [getContractEventLogs](./queries/nft/getContractEventLogs.go)               | TESTED |
| [getNFTDetails](./queries/nft/getNFTDetails.go)                             | TESTED |
| [getNFTEventLogs](./queries/nft/getNFTEventLogs.go)                         | TESTED |
| [getNFTsByContractAddress](./queries/nft/getNFTsByContractAddress.go)       | TESTED |
| [getNFTsByWalletAddress](./queries/nft/getNFTsByWalletAddress.go)           | TESTED |
| [getNFTsByWalletAndContracts](./queries/nft/getNFTsByWalletAndContracts.go) | TESTED |
| [getNFTsByWalletENS](./queries/nft/getNFTsByWalletENS.go)                   | TESTED |

## Issues

Even if every query has tests, the data used to compare with function's value are static. However
the client will returns different results as things happens in the Chain.

## Further improvement

- Implement an effective generator system based on GQL query, [Genqlient](https://github.com/Khan/genqlient)
  is a good solution but the generated code is too verbose to be maintainable. A lighter homemade solution
  may be better.
- Implements new useful queries and more tests.
- Implement a generic generator (think about gRPC generator) to generate clients in any languages.
  This is the best way to easily maintains SDKs.
- Improve the test system to not use static data.