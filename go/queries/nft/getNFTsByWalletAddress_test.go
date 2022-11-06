package nft_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/quartz-technology/qn-sdk/client"
	"github.com/quartz-technology/qn-sdk/queries/nft"
	"github.com/quartz-technology/qn-sdk/spec"
	"github.com/stretchr/testify/assert"
)

func TestGetNFTsByWalletAddress(t *testing.T) {
	c := client.New()

	type TestCase map[string]struct {
		address string
		filter  nft.GetNFTsByWalletAddressFilterInput
		expect  string
	}

	testCases := TestCase{
		"executes query": {
			address: "0x13928eb9a86c8278a45b6ff2935c7730b58ac675",
			filter: nft.GetNFTsByWalletAddressFilterInput{
				PaginationArgs: spec.PaginationArgs{
					First: 1,
					After: "YXJyYXljb25uZWN0aW9uOjA=",
				},
			},
			expect: `{
    "Wallet": {
        "EnsName": "bitcoinpirate.eth",
        "Address": "0x13928eb9a86c8278a45b6ff2935c7730b58ac675",
        "Tokens": {
            "PageInfo": {
                "HasNextPage": true,
                "EndCursor": "YXJyYXljb25uZWN0aW9uOjE="
            },
            "Edges": [
                {
                    "Node": {
                        "TokenId": "5131",
                        "Images": [
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/xs\/collections\/0x0f78c6eee3c89ff37fd9ef96bd685830993636f2\/tokens\/4fc875c9740fc8be8f99ca858189c5c1"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/sm\/collections\/0x0f78c6eee3c89ff37fd9ef96bd685830993636f2\/tokens\/4fc875c9740fc8be8f99ca858189c5c1"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/md\/collections\/0x0f78c6eee3c89ff37fd9ef96bd685830993636f2\/tokens\/4fc875c9740fc8be8f99ca858189c5c1"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/lg\/collections\/0x0f78c6eee3c89ff37fd9ef96bd685830993636f2\/tokens\/4fc875c9740fc8be8f99ca858189c5c1"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/xl\/collections\/0x0f78c6eee3c89ff37fd9ef96bd685830993636f2\/tokens\/4fc875c9740fc8be8f99ca858189c5c1"
                            }
                        ],
                        "Contract": {
                            "Address": "0x0f78c6eee3c89ff37fd9ef96bd685830993636f2",
                            "Symbol": "Nuclear Nerds",
                            "Name": "Nuclear Nerds of the Accidental Apocalypse"
                        }
                    }
                }
            ]
        }
    }
}`,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetNFTsByWalletAddress(context.TODO(), tt.address, tt.filter)

			assert.Nil(t, err)
			data, err := json.Marshal(res)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.expect, string(data))
		})
	}
}
