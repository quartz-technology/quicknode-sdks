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

func TestGetNFTEventLogs(t *testing.T) {
	c := client.New()

	type TestCase map[string]struct {
		address string
		tokenID string
		filter  nft.GetNFTEventLogsFilterInput
		expect  string
	}

	testCases := TestCase{
		"executes query": {
			address: "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
			tokenID: "0",
			filter: nft.GetNFTEventLogsFilterInput{
				Types: []spec.LogType{spec.TRANSFER, spec.ORDER, spec.MINT},
				PaginationArgs: spec.PaginationArgs{
					First: 1,
				},
			},
			expect: `{
    "Token": {
        "TokenId": "0",
        "Contract": {
            "Address": "0x60e4d786628fea6478f785a6d7e704777c86a7c6"
        },
        "Logs": {
            "Edges": [
                {
                    "Node": {
                        "BlockNumber": 13117215,
                        "Type": "MINT",
                        "FromAddress": "0x0000000000000000000000000000000000000000",
                        "ToAddress": "0x9056d15c49b19df52ffad1e6c11627f035c0c960",
                        "EstimatedConfirmedAt": "2021-08-29T00:41:11.000Z",
                        "TransactionHash": "0x612a5190e324f5455432af3aa2c97e9d35f036929263f55e57d7975a0093bd6c",
						"Marketplace": "",
						"PriceInEth": 0
                    }
                }
            ],
            "PageInfo": {
                "HasNextPage": true,
                "EndCursor": "YXJyYXljb25uZWN0aW9uOjA="
            }
        }
    }
}`,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetNFTDetailsLogs(context.TODO(), tt.address, tt.tokenID, tt.filter)

			assert.Nil(t, err)
			data, err := json.Marshal(res)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.expect, string(data))
		})
	}
}
