package nft

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/mitchellh/mapstructure"
)

const getCollectionDetailsRawQuery = `
  query CollectionDetails($address: String!) {
    contract(address: $address) {
      ... on ERC721Contract {
        address
        isVerified
        tokenStandard
        attributes {
          name
          rarity
          value
          valueCount
        }
        circulatingSupply
        name
        stats {
          average
          ceiling
          floor
          totalSales
          volume
        }
        symbol
      }
    }
  }
`

type GetCollectionDetailsOutput struct {
	Contract struct {
		Address       string
		IsVerified    bool
		tokenStandard string
		Attributes    []struct {
			Name       string
			Rarity     float64
			Value      string
			ValueCount int
		}
		CirculatingSupply int
		Name              string
		Stats             struct {
			Average    float64
			Ceiling    float64
			Floor      float64
			TotalSales float64
			Volume     float64
		}
		Symbol        string
		TokenStandard string
	}
}

// GetCollectionDetails returns a bunch on a given NFT Collection.
func GetCollectionDetails(ctx context.Context, client graphql.Client, address string) (*GetCollectionDetailsOutput, error) {
	res := &graphql.Response{
		Data: GetCollectionDetailsOutput{},
	}

	req := &graphql.Request{
		Query: getCollectionDetailsRawQuery,
		Variables: map[string]interface{}{
			"address": address,
		},
		OpName: "CollectionDetails",
	}

	if err := client.MakeRequest(ctx, req, res); err != nil {
		return nil, err
	}

	var data GetCollectionDetailsOutput
	if err := mapstructure.Decode(res.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
