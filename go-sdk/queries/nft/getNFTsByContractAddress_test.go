package nft_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/quartz-technology/quicknode-sdks/go-sdk/client"
	"github.com/quartz-technology/quicknode-sdks/go-sdk/queries/nft"
	"github.com/quartz-technology/quicknode-sdks/go-sdk/spec"
	"github.com/stretchr/testify/assert"
)

func TestGetNFTsByContractAddress(t *testing.T) {
	c := client.New()

	type TestCase map[string]struct {
		address string
		filter  nft.GetNFTsByContractAddressFilterInput
		expect  string
	}

	testCases := TestCase{
		"executes query": {
			address: "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
			filter: nft.GetNFTsByContractAddressFilterInput{
				PaginationArgs: spec.PaginationArgs{
					First: 1,
					After: "YXJyYXljb25uZWN0aW9uOjA=",
				},
			},
			expect: `{
    "Contract": {
        "Address": "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
        "IsVerified": true,
        "TokenStandard": "ERC721",
        "Tokens": {
            "PageInfo": {
                "HasNextPage": true,
                "EndCursor": "YXJyYXljb25uZWN0aW9uOjE="
            },
            "Edges": [
                {
                    "Node": {
                        "TokenId": "0",
                        "Images": [
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/xs\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/sm\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/md\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/lg\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/xl\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/bdf56dec1c120d5cdbcc4816b60468a0"
                            }
                        ],
                        "Contract": {
                            "Address": "0x60e4d786628fea6478f785a6d7e704777c86a7c6"
                        }
                    }
                },
                {
                    "Node": {
                        "TokenId": "1",
                        "Images": [
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/xs\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/3ab725edf59d2a65b6542983abf3485e"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/sm\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/3ab725edf59d2a65b6542983abf3485e"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/md\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/3ab725edf59d2a65b6542983abf3485e"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/lg\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/3ab725edf59d2a65b6542983abf3485e"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/xl\/collections\/0xcb07acb32212d2a3a7fa14bf570a06af081903dd\/tokens\/3ab725edf59d2a65b6542983abf3485e"
                            }
                        ],
                        "Contract": {
                            "Address": "0x60e4d786628fea6478f785a6d7e704777c86a7c6"
                        }
                    }
                }
            ]
        },
        "CirculatingSupply": 19377,
        "Name": "Mutant Ape Yacht Club",
        "Symbol": "MAYC"
    }
}`,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetNFTsByContractAddress(context.TODO(), tt.address, tt.filter)

			assert.Nil(t, err)
			data, err := json.Marshal(res)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.expect, string(data))
		})
	}
}
