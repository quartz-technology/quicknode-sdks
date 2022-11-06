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

func TestGetContractEventLog(t *testing.T) {
	c := client.New()

	type TestCase map[string]struct {
		address string
		filter  nft.GetContractEventLogFilterInput
		expect  string
	}

	testCases := TestCase{
		"executes query": {
			address: "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
			filter: nft.GetContractEventLogFilterInput{
				Types: []spec.LogType{spec.TRANSFER, spec.ORDER, spec.MINT},
				PaginationArgs: spec.PaginationArgs{
					After: "YXJyYXljb25uZWN0aW9uOjA=",
					First: 1,
				},
			},
			expect: `{
        "Contract": {
            "Address": "0x60e4d786628fea6478f785a6d7e704777c86a7c6",
            "Logs": {
                "PageInfo": {
                    "HasNextPage": true,
                    "EndCursor": "YXJyYXljb25uZWN0aW9uOjE="
                },
                "Edges": [
                    {
                        "Node": {
                            "BlockNumber": 15906937,
                            "Type": "ORDER",
                            "FromAddress": "0x05aee7caea14ae3c18aa415e6b779c0ca90f07dd",
                            "ToAddress": "0x32500659b6eaf3a0f03fffb5b8da0eca3fc80735",
                            "EstimatedConfirmedAt": "2022-11-05T22:56:35.000Z",
                            "TransactionHash": "0x80685cb411671ec612b12dcc76cc3e4b9439046e7d162ea5661bbf54670b8f72",
                            "Token": {
                                "Contract": {
                                    "Address": "0x60e4d786628fea6478f785a6d7e704777c86a7c6"
                                },
                                "TokenId": "489"
                            },
                            "Marketplace": "OPENSEA",
                            "PriceInEth": 12.25
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
			res, err := c.GetContractEventLog(context.TODO(), tt.address, tt.filter)

			assert.Nil(t, err)
			data, err := json.Marshal(res)
			assert.Nil(t, err)
			assert.JSONEq(t, tt.expect, string(data))
		})
	}
}
