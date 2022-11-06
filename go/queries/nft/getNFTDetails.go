package nft

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/mitchellh/mapstructure"
)

const getNFTDetailsRawQuery = `
  query Token($address: String!, $tokenId: String!) {
    token(contractAddress: $address, tokenId: $tokenId) {
      ... on ERC721Token {
        tokenId
        attributes {
          name
          value
        }
        contract {
          address
          isVerified
          tokenStandard
          ... on ERC721Contract {
            name
          }
        }
        images {
          height
          mimeType
          url
          width
        }
        name
        symbol
        metadata {
          animation_url
          background_color
          description
          external_url
          image
          image_data
          name
          youtube_url
        }
      }
    }
  }
`

type GetNFTDetailsOutput struct {
	Token struct {
		// on ERC721Token
		TokenId    string
		Attributes []struct {
			Name  string
			Value string
		}
		Contract struct {
			Address       string
			IsVerified    bool
			TokenStandard string
			// On ERC 721 Contract
			Name string
		}
		Images []struct {
			Height   int
			MimeType string
			Url      string
			Width    int
		}
		Name     string
		Symbol   string
		Metadata struct {
			AnimationUrl    string
			BackgroundColor string
			Description     string
			ExternalUrl     string
			Image           string
			ImageData       string
			Name            string
			YoutubeUrl      string
		}
	}
}

// GetNFTDetails retrieves information of a given NFT.
func GetNFTDetails(ctx context.Context, client graphql.Client, address, tokenID string) (*GetNFTDetailsOutput, error) {
	res := &graphql.Response{
		Data: GetNFTDetailsOutput{},
	}

	req := &graphql.Request{
		Query: getNFTDetailsRawQuery,
		Variables: map[string]interface{}{
			"address": address,
			"tokenId": tokenID,
		},
		OpName: "Token",
	}

	if err := client.MakeRequest(ctx, req, res); err != nil {
		return nil, err
	}

	var data GetNFTDetailsOutput
	if err := mapstructure.Decode(res.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
