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

func TestGetNFTsByWalletENS(t *testing.T) {
	c := client.New()

	type TestCase map[string]struct {
		ensName string
		filter  nft.GetNFTsByWalletENSFilterInput
		expect  string
	}

	testCases := TestCase{
		"executes query": {
			ensName: "vitalik.eth",
			filter: nft.GetNFTsByWalletENSFilterInput{
				PaginationArgs: spec.PaginationArgs{
					First: 1,
					After: "YXJyYXljb25uZWN0aW9uOjM=",
				},
			},
			expect: `{
    "Wallet": {
        "EnsName": "vitalik.eth",
        "Address": "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
        "Tokens": {
            "PageInfo": {
                "HasNextPage": true,
                "EndCursor": "YXJyYXljb25uZWN0aW9uOjQ="
            },
            "Edges": [
                {
                    "Node": {
                        "TokenId": "4360",
                        "Images": [
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/xs\/collections\/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f\/tokens\/8f2e60ad7980a8a8cb4335688831ab5c"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/sm\/collections\/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f\/tokens\/8f2e60ad7980a8a8cb4335688831ab5c"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/md\/collections\/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f\/tokens\/8f2e60ad7980a8a8cb4335688831ab5c"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/lg\/collections\/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f\/tokens\/8f2e60ad7980a8a8cb4335688831ab5c"
                            },
                            {
                                "Url": "https:\/\/images.icytools.workers.dev\/xl\/collections\/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f\/tokens\/8f2e60ad7980a8a8cb4335688831ab5c"
                            }
                        ],
                        "Contract": {
                            "Address": "0xe588feda21d17bd8f7b0532d3e439ab245c2f67f",
                            "Symbol": "BRITISH",
                            "Name": "The British"
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
			res, err := c.GetNFTsByWalletENS(context.TODO(), tt.ensName, tt.filter)

			assert.Nil(t, err)
			data, err := json.Marshal(res)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.expect, string(data))
		})
	}
}
