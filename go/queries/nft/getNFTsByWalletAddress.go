package nft

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/quartz-technology/qn-sdk/spec"
)

const getNFTsByWalletAddressRawQuery = `query WalletNFTs($address: String!, $first: Int, $after: String) {
  wallet(address: $address) {
    ensName
    address
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

type GetNFTsByWalletAddressFilterInput struct {
	spec.PaginationArgs
}

type GetNFTsByWalletAddressOutput struct {
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

func GetNFTsByWalletAddress(ctx context.Context, client graphql.Client, address string, filterOpt GetNFTsByWalletAddressFilterInput) (*GetNFTsByWalletAddressOutput, error) {
	res := &graphql.Response{
		Data: GetNFTsByWalletAddressOutput{},
	}

	variable := map[string]interface{}{
		"address": address,
		"first":   filterOpt.First,
		"after":   filterOpt.After,
	}

	req := &graphql.Request{
		Query:     getNFTsByWalletAddressRawQuery,
		Variables: variable,
		OpName:    "WalletNFTs",
	}

	if err := client.MakeRequest(ctx, req, res); err != nil {
		return nil, err
	}

	var data GetNFTsByWalletAddressOutput
	if err := mapstructure.Decode(res.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
