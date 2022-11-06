package nft

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/quartz-technology/quicknode-sdks/go-sdk/spec"
)

const getNFTEventLogsRawQuery = `query NFTEvents(
    $address: String!
    $tokenId: String!
    $filter: LogsFilterInputType
    $first: Int
    $after: String
  ) {
    token(contractAddress: $address, tokenId: $tokenId) {
      ... on ERC721Token {
        tokenId
        contract {
          address # Included key field for caching
        }
        logs(filter: $filter, first: $first, after: $after) {
          edges {
            node {
              blockNumber
              type
              fromAddress
              toAddress
              estimatedConfirmedAt
              transactionHash
              ... on OrderLog {
                marketplace
                priceInEth
              }
            }
          }
          pageInfo {
            hasNextPage
            endCursor
          }
        }
      }
    }
  }`

type GetNFTEventLogsFilterInput struct {
	spec.PaginationArgs
	Types []spec.LogType
}

type GetNFTEventLogsOutput struct {
	Token struct {
		TokenId  string
		Contract struct {
			Address string
		}
		Logs struct {
			Edges []struct {
				Node struct {
					BlockNumber          int
					Type                 string
					FromAddress          string
					ToAddress            string
					EstimatedConfirmedAt string
					TransactionHash      string
					Marketplace          string
					PriceInEth           float64
				}
			}
			PageInfo struct {
				HasNextPage bool
				EndCursor   string
			}
		}
	}
}

func GetNFTEventLogs(ctx context.Context, client graphql.Client, address string, tokenId string, filterOpt GetNFTEventLogsFilterInput) (*GetNFTEventLogsOutput, error) {
	res := &graphql.Response{
		Data: GetNFTEventLogsOutput{},
	}

	variable := map[string]interface{}{
		"address": address,
		"tokenId": tokenId,
		"types": struct {
			TypeIn []spec.LogType
		}{
			TypeIn: filterOpt.Types,
		},
		"first": filterOpt.First,
		"after": filterOpt.After,
	}

	req := &graphql.Request{
		Query:     getNFTEventLogsRawQuery,
		Variables: variable,
		OpName:    "NFTEvents",
	}

	if err := client.MakeRequest(ctx, req, res); err != nil {
		return nil, err
	}

	var data GetNFTEventLogsOutput
	if err := mapstructure.Decode(res.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
