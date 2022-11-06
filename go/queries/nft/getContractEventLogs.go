package nft

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/quartz-technology/qn-sdk/spec"
)

const getContractEventLogRawQuery = `
  query ContractEvents(
    $address: String!
    $filter: LogsFilterInputType
    $first: Int
    $after: String
  ) {
    contract(address: $address) {
      address
      logs(filter: $filter, first: $first, after: $after) {
        pageInfo {
          hasNextPage
          endCursor
        }
        edges {
          node {
            blockNumber
            type
            fromAddress
            toAddress
            estimatedConfirmedAt
            transactionHash
            token {
              contract {
                address
              }
              ... on ERC721Token {
                tokenId
              }
            }
            ... on OrderLog {
              marketplace
              priceInEth
            }
          }
        }
      }
    }
  }
`

type GetContractEventLogFilterInput struct {
	spec.PaginationArgs
	Types []spec.LogType
}

type GetContractEventOutput struct {
	Contract struct {
		Address string
		Logs    struct {
			PageInfo struct {
				HasNextPage bool
				EndCursor   string
			}
			Edges []struct {
				Node struct {
					BlockNumber          int
					Type                 string
					FromAddress          string
					ToAddress            string
					EstimatedConfirmedAt string
					TransactionHash      string
					Token                struct {
						Contract struct {
							Address string
						}
						// On ERC721Token
						TokenId string
					}
					// On OrderLog
					Marketplace string
					PriceInEth  float64
				}
			}
		}
	}
}

// GetContractEventLog retrieves all events of a contract.
func GetContractEventLog(ctx context.Context, client graphql.Client, address string, filterOpt GetContractEventLogFilterInput) (*GetContractEventOutput, error) {
	res := &graphql.Response{
		Data: GetContractEventOutput{},
	}

	variable := map[string]interface{}{
		"address": address,
		"types": struct {
			TypeIn []spec.LogType
		}{
			TypeIn: filterOpt.Types,
		},
		"first": filterOpt.First,
		"after": filterOpt.After,
	}

	req := &graphql.Request{
		Query:     getContractEventLogRawQuery,
		Variables: variable,
		OpName:    "ContractEvents",
	}

	if err := client.MakeRequest(ctx, req, res); err != nil {
		return nil, err
	}

	var data GetContractEventOutput
	if err := mapstructure.Decode(res.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
