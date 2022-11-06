package client

import (
	"context"

	"github.com/quartz-technology/qn-sdk/queries/nft"
	"github.com/quartz-technology/qn-sdk/spec"
)

func (c *Client) GetCollectionDetails(ctx context.Context, address string) (*nft.GetCollectionDetailsOutput, error) {
	return nft.GetCollectionDetails(ctx, c, address)
}

// GetContractEventLog retrieve logs of a given contract
// By default, it retrieves all types of logs.
func (c *Client) GetContractEventLog(ctx context.Context, address string, filterOpt nft.GetContractEventLogFilterInput) (*nft.GetContractEventOutput, error) {
	if len(filterOpt.Types) == 0 {
		filterOpt.Types = []spec.LogType{spec.TRANSFER, spec.MINT, spec.ORDER}
	}

	return nft.GetContractEventLog(ctx, c, address, filterOpt)
}

func (c *Client) GetNFTDetails(ctx context.Context, address, tokenID string) (*nft.GetNFTDetailsOutput, error) {
	return nft.GetNFTDetails(ctx, c, address, tokenID)
}

func (c *Client) GetNFTDetailsLogs(ctx context.Context, address string, tokenID string, filterOpt nft.GetNFTEventLogsFilterInput) (*nft.GetNFTEventLogsOutput, error) {
	if len(filterOpt.Types) == 0 {
		filterOpt.Types = []spec.LogType{spec.TRANSFER, spec.MINT, spec.ORDER}
	}

	return nft.GetNFTEventLogs(ctx, c, address, tokenID, filterOpt)
}

func (c *Client) GetNFTsByContractAddress(ctx context.Context, address string, filterOpt nft.GetNFTsByContractAddressFilterInput) (*nft.GetNFTsByContractAddressOutput, error) {
	return nft.GetNFTsByContractAddress(ctx, c, address, filterOpt)
}

func (c *Client) GetNFTsByWalletAddress(ctx context.Context, address string, filterOpt nft.GetNFTsByWalletAddressFilterInput) (*nft.GetNFTsByWalletAddressOutput, error) {
	return nft.GetNFTsByWalletAddress(ctx, c, address, filterOpt)
}

func (c *Client) GetNFTsByWalletENS(ctx context.Context, ensName string, filterOpt nft.GetNFTsByWalletENSFilterInput) (*nft.GetNFTsByWalletENSOutput, error) {
	return nft.GetNFTsByWalletENS(ctx, c, ensName, filterOpt)
}

func (c *Client) GetNFTsByWalletAndContracts(ctx context.Context, address string, filterOpt nft.GetNFTsByWalletAndContractsFilterInput) (*nft.GetNFTsByWalletAndContractsOutput, error) {
	return nft.GetNFTsByWalletAndContracts(ctx, c, address, filterOpt)
}
