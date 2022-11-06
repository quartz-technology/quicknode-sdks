package nft_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/quartz-technology/qn-sdk/client"
	"github.com/stretchr/testify/assert"
)

func TestGetNFTDetails(t *testing.T) {
	c := client.New()

	type TestCase map[string]struct {
		address string
		tokenId string
		expect  string
	}

	testCases := TestCase{
		"executes query": {
			address: "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
			tokenId: "0",
			expect: `{
        "Token": {
            "TokenId": "0",
            "Attributes": [
                {
                    "Name": "Background",
                    "Value": "M1 Purple"
                },
                {
                    "Name": "Clothes",
                    "Value": "M1 Toga"
                },
                {
                    "Name": "Eyes",
                    "Value": "M1 Scumbag"
                },
                {
                    "Name": "Fur",
                    "Value": "M1 Cheetah"
                },
                {
                    "Name": "Mouth",
                    "Value": "M1 Bored Unshaven"
                }
            ],
            "Contract": {
                "Address": "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
                "IsVerified": true,
                "TokenStandard": "ERC721",
                "Name": "Mutant Ape Yacht Club"
            },
            "Images": [
                {
                    "Height": 100,
                    "MimeType": "image\/png",
                    "Url": "https:\/\/images.icytools.workers.dev\/xs\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0",
                    "Width": 100
                },
                {
                    "Height": 200,
                    "MimeType": "image\/png",
                    "Url": "https:\/\/images.icytools.workers.dev\/sm\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0",
                    "Width": 200
                },
                {
                    "Height": 400,
                    "MimeType": "image\/png",
                    "Url": "https:\/\/images.icytools.workers.dev\/md\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0",
                    "Width": 400
                },
                {
                    "Height": 800,
                    "MimeType": "image\/png",
                    "Url": "https:\/\/images.icytools.workers.dev\/lg\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0",
                    "Width": 800
                },
                {
                    "Height": 1200,
                    "MimeType": "image\/png",
                    "Url": "https:\/\/images.icytools.workers.dev\/xl\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0",
                    "Width": 1200
                }
            ],
            "Name": "",
            "Symbol": "",
            "Metadata": {
                "AnimationUrl": "",
                "BackgroundColor": "",
                "Description": "",
                "ExternalUrl": "",
                "Image": "ipfs:\/\/QmURua8bNrAwX76Tp6G9t6Lospdxyt61JGy3UsXY7skfR1",
                "ImageData": "",
                "Name": "",
                "YoutubeUrl": ""
            }
        }
    }`,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetNFTDetails(context.TODO(), tt.address, tt.tokenId)

			assert.Nil(t, err)
			data, err := json.Marshal(res)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.expect, string(data))
		})
	}
}
