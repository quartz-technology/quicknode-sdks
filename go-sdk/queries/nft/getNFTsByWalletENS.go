package nft

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/quartz-technology/quicknode-sdks/go-sdk/spec"
)

const getNFTsByWalletENSRawQuery = `query WalletNFTs($ensName: String, $first: Int, $after: String) {
    wallet(ensName: $ensName) {
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

type GetNFTsByWalletENSFilterInput struct {
	spec.PaginationArgs
}

type GetNFTsByWalletENSOutput struct {
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

func GetNFTsByWalletENS(ctx context.Context, client graphql.Client, ensName string, filterOpt GetNFTsByWalletENSFilterInput) (*GetNFTsByWalletENSOutput, error) {
	res := &graphql.Response{
		Data: GetNFTsByWalletENSOutput{},
	}

	variable := map[string]interface{}{
		"ensName": ensName,
		"first":   filterOpt.First,
		"after":   filterOpt.After,
	}

	req := &graphql.Request{
		Query:     getNFTsByWalletENSRawQuery,
		Variables: variable,
		OpName:    "WalletNFTs",
	}

	if err := client.MakeRequest(ctx, req, res); err != nil {
		return nil, err
	}

	var data GetNFTsByWalletENSOutput
	if err := mapstructure.Decode(res.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
