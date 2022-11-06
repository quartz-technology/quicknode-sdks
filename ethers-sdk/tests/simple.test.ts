import { describe, it } from 'mocha';
import { ethers } from 'ethers';
import { QuickNodeProvider } from '../src';
import { expect } from 'chai';
import { TokenMetadata } from '../src/types';

describe('QuickNode SDK', () => {
    const nullAddress = '0x0000000000000000000000000000000000000000';
    const vitalikAddress = '0xd8da6bf26964af9d7eed9e03e53415d37aa96045';
    // const phaAddress = '0x8Db2568E2E5697CC21606291b2b846D36a25f49c';
    const usdcTokenAddress = '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48';

    const parentProvider = new ethers.providers.JsonRpcProvider(
        'https://fittest-sleek-log.discover.quiknode.pro/YOUR_TOKEN_HERE/'
    );
    const qnProvider = new QuickNodeProvider(parentProvider);

    it('should fetch NFT collection details', async () => {
        const res = await qnProvider.fetchNFTCollectionDetails([
            '0x2106C00Ac7dA0A3430aE667879139E832307AeAa',
        ]);

        console.log(res)

        expect(res).not.to.be.empty;
    });

    it('should fail to fetch NFT collection details', async () => {
        const res = await qnProvider.fetchNFTCollectionDetails([nullAddress]);

        expect(res).to.be.empty;
    });

    it('should fetch NFTs owned by wallet', async () => {
        const res = await qnProvider.fetchNFTs(
            '0xf1BCf736a46D41f8a9d210777B3d75090860a665'
        );

        expect(res.assets).not.to.be.empty;
    });

    it('should get all the NFTs in a collection', async () => {
        const res = await qnProvider.fetchNFTsByCollection(
            '0x60E4d786628Fea6478F785A6d7e704777c86a7c6'
        );

        expect(res.tokens).not.to.be.empty;
        expect(res.totalItems).to.be.equal(19427);
    });

    it('should get token metadata by symbol', async () => {
        const res = await qnProvider.getTokenMetadataBySymbol('USDC');

        expect(res.pageNumber).to.be.equal(1);
        expect(res.tokens).to.be.length(20);
    });

    it('should get token metadata by symbol using pagination filters', async () => {
        const res = await qnProvider.getTokenMetadataBySymbol('USDC', 2, 23);

        expect(res.pageNumber).to.be.equal(2);
        expect(res.tokens).to.be.length(23);
    });

    it('should get transfers for NFT in collection', async () => {
        const res = await qnProvider.getTransfersByNFT(
            '0x60E4d786628Fea6478F785A6d7e704777c86a7c6',
            '1'
        );

        expect(res.transfers).not.to.be.empty;
    });

    it('should get token metadata by contract address', async () => {
        const res = await qnProvider.getTokenMetadataByContractAddress(
            '0x4d224452801ACEd8B2F0aebE155379bb5D594381'
        );

        expect(res).not.to.be.null;
    });

    it('should fail to get token metadata by contract address', async () => {
        const res = await qnProvider.getTokenMetadataByContractAddress(
            nullAddress
        );

        expect(res).to.be.null;
    });

    it('should verify NFTs Owner', async () => {
        const identifiedContractAddresses = [
            '0x2106c00ac7da0a3430ae667879139e832307aeaa:3643',
            '0xd07dc4262bcdbf85190c01c996b4c06a461d2430:133803',
        ];
        const res = await qnProvider.verifyNFTsOwner(
            '0x91b51c173a4bdaa1a60e234fc3f705a16d228740',
            identifiedContractAddresses
        );

        expect(res.assets).to.be.eql(identifiedContractAddresses);
    });

    it('should fail to verify NFTs Owner', async () => {
        const identifiedContractAddresses = [
            '0x2106c00ac7da0a3430ae667879139e832307aeaa:3643',
            '0xd07dc4262bcdbf85190c01c996b4c06a461d2430:133803',
        ];
        const res = await qnProvider.verifyNFTsOwner(
            nullAddress,
            identifiedContractAddresses
        );

        expect(res.assets).to.be.empty;
    });

    it('should get the wallet token transactions (USDC Token)', async () => {
        const res = await qnProvider.getWalletTokenTransactions(
            vitalikAddress,
            usdcTokenAddress
        );
        const usdcMetadata =
            (await qnProvider.getTokenMetadataByContractAddress(
                usdcTokenAddress
            )) as TokenMetadata;

        expect(res.token.genesisTransaction).to.be.equal(
            usdcMetadata.genesisTransaction
        );
    });

    it('should get the wallet token balance', async () => {
        const owner = vitalikAddress;
        const res = await qnProvider.getWalletTokenBalance(owner);

        expect(res.owner).to.be.equal(owner);
        expect(res.assets).not.to.be.empty;
    });

    it('should get the wallet token balance with filter (USDC only)', async () => {
        const owner = vitalikAddress;
        const res = await qnProvider.getWalletTokenBalance(owner, [
            usdcTokenAddress,
        ]);

        expect(res.owner).to.be.equal(owner);
        expect(res.assets).to.be.length(1);
    });
});
