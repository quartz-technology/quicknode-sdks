package nft

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/quartz-technology/qn-sdk/spec"
)

const getNFTsByWalletAndContractsRawQuery = `query NFTsWalletAndContract(
    $filter: TokensFilterInputType
    $address: String
    $first: Int
    $after: String
  ) {
    wallet(address: $address) {
      ensName
      address
      tokens(filter: $filter, first: $first, after: $after) {
        pageInfo {
          hasNextPage
          endCursor
        }
        edges {
          node {
            tokenId
            images {
              url
            }
            ... on ERC721Token {
              contract {
                address
                ... on ERC721Contract {
                  symbol
                  name
                }
              }
            }
          }
        }
      }
    }
  }
`

type GetNFTsByWalletAndContractsFilterInput struct {
	spec.PaginationArgs
	Contracts []string
}

type GetNFTsByWalletAndContractsOutput struct {
	Wallet struct {
		EnsName string
		Address string
		Tokens  struct {
			PageInfo struct {
				HasNextPage bool
				EndCursor   string
			}
			Edges []struct {
				Node struct {
					TokenId string
					Images  []struct {
						Url string
					}
					Contract struct {
						Address string
						Symbol  string
						Name    string
					}
				}
			}
		}
	}
}

func GetNFTsByWalletAndContracts(ctx context.Context, client graphql.Client, address string, filterOpt GetNFTsByWalletAndContractsFilterInput) (*GetNFTsByWalletAndContractsOutput, error) {
	res := &graphql.Response{
		Data: GetNFTsByWalletAndContractsOutput{},
	}

	variable := map[string]interface{}{
		"address": address,
		"filter": map[string][]string{
			"contractAddressIn": filterOpt.Contracts,
		},
		"first": filterOpt.First,
		"after": filterOpt.After,
	}

	req := &graphql.Request{
		Query:     getNFTsByWalletAndContractsRawQuery,
		Variables: variable,
		OpName:    "NFTsWalletAndContract",
	}

	if err := client.MakeRequest(ctx, req, res); err != nil {
		return nil, err
	}

	var data GetNFTsByWalletAndContractsOutput
	if err := mapstructure.Decode(res.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
