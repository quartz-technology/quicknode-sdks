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

func TestGetNFTsByWalletAndContracts(t *testing.T) {
	c := client.New()

	type TestCase map[string]struct {
		address string
		filter  nft.GetNFTsByWalletAndContractsFilterInput
		expect  string
	}

	testCases := TestCase{
		"executes query": {
			address: "0x13928eB9A86c8278a45B6fF2935c7730b58AC675",
			filter: nft.GetNFTsByWalletAndContractsFilterInput{
				Contracts: []string{
					"0xba30E5F9Bb24caa003E9f2f0497Ad287FDF95623",
					"0xbCe3781ae7Ca1a5e050Bd9C4c77369867eBc307e",
					"0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D",
				},
				PaginationArgs: spec.PaginationArgs{
					First: 1,
				},
			},
			expect: `{
    "Wallet": {
        "EnsName": "bitcoinpirate.eth",
        "Address": "0x13928eb9a86c8278a45b6ff2935c7730b58ac675",
        "Tokens": {
            "PageInfo": {
                "HasNextPage": false,
                "EndCursor": "YXJyYXljb25uZWN0aW9uOjA="
            },
            "Edges": []
        }
    }
}`,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetNFTsByWalletAndContracts(context.TODO(), tt.address, tt.filter)

			assert.Nil(t, err)
			data, err := json.Marshal(res)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.expect, string(data))
		})
	}
}
