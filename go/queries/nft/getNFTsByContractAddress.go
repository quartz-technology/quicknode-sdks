package nft

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/quartz-technology/qn-sdk/spec"
)

const getNFTsByContractAddressRawQuery = `
  query ContractNFTs($address: String!, $first: Int, $after: String) {
    contract(address: $address) {
      address
      isVerified
      tokenStandard
      tokens(first: $first, after: $after) {
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
                address # Included key field for caching
              }
            }
          }
        }
      }
      ... on ERC721Contract {
        circulatingSupply
        name
        symbol
      }
    }
  }`

type GetNFTsByContractAddressFilterInput struct {
	spec.PaginationArgs
}

type GetNFTsByContractAddressOutput struct {
	Contract struct {
		Address       string
		IsVerified    bool
		TokenStandard string
		Tokens        struct {
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
					}
				}
			}
		}
		CirculatingSupply int
		Name              string
		Symbol            string
	}
}

func GetNFTsByContractAddress(ctx context.Context, client graphql.Client, address string, filterOpt GetNFTsByContractAddressFilterInput) (*GetNFTsByContractAddressOutput, error) {
	res := &graphql.Response{
		Data: GetNFTsByContractAddressOutput{},
	}

	variable := map[string]interface{}{
		"address": address,
		"first":   filterOpt.First,
		"after":   filterOpt.After,
	}

	req := &graphql.Request{
		Query:     getNFTsByContractAddressRawQuery,
		Variables: variable,
		OpName:    "ContractNFTs",
	}

	if err := client.MakeRequest(ctx, req, res); err != nil {
		return nil, err
	}

	var data GetNFTsByContractAddressOutput
	if err := mapstructure.Decode(res.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
