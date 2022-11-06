import { providers } from 'ethers';
import {
    NFTCollectionDetails,
    NFTDetails,
    NFTsInCollection,
    NFTsOwner,
    NFTTransfersForCollection,
    TokenMetadata,
    TokenMetadataPaginated,
    TokenTransactions,
    WalletTokenBalance,
} from './types';

export class QuickNodeProvider {
    private parentProvider: providers.JsonRpcProvider;

    constructor(parentProvider: providers.JsonRpcProvider) {
        this.parentProvider = parentProvider;
    }

    public async broadcastRawTransaction(signedTx: string): Promise<string> {
        return (await this.parentProvider.send('qn_broadcastRawTransaction', [
            signedTx,
        ])) as string;
    }

    public async fetchNFTCollectionDetails(
        contractsAddresses: string[]
    ): Promise<NFTCollectionDetails[]> {
        return (await this.parentProvider.send('qn_fetchNFTCollectionDetails', {
            // @ts-ignore
            contracts: contractsAddresses,
        })) as NFTCollectionDetails[];
    }

    public async fetchNFTs(
        walletAddress: string,
        nftContractAddresses?: string[],
        omitFields?: string[],
        page?: number,
        perPage?: number
    ): Promise<NFTDetails> {
        return (await this.parentProvider.send('qn_fetchNFTs', {
            // @ts-ignore
            wallet: walletAddress,
            contracts: nftContractAddresses,
            omitFields: omitFields,
            page: page,
            perPage: perPage,
        })) as NFTDetails;
    }

    public async fetchNFTsByCollection(
        collectionAddress: string,
        tokensIDs?: string[],
        omitFields?: string[],
        page?: number,
        perPage?: number
    ) {
        return (await this.parentProvider.send('qn_fetchNFTsByCollection', {
            // @ts-ignore
            collection: collectionAddress,
            tokens: tokensIDs,
            omitFields: omitFields,
            page: page,
            perPage: perPage,
        })) as NFTsInCollection;
    }

    public async getTokenMetadataByContractAddress(
        contractAddress: string
    ): Promise<TokenMetadata | null> {
        const res = await this.parentProvider.send(
            'qn_getTokenMetadataByContractAddress',
            {
                // @ts-ignore
                contract: contractAddress,
            }
        );

        return res.contract as TokenMetadata | null;
    }

    public async getTokenMetadataBySymbol(
        symbol: string,
        page?: number,
        perPage?: number
    ): Promise<TokenMetadataPaginated> {
        return (await this.parentProvider.send('qn_getTokenMetadataBySymbol', {
            // @ts-ignore
            symbol: symbol,
            page: page,
            perPage: perPage,
        })) as TokenMetadataPaginated;
    }

    public async getTransfersByNFT(
        collectionAddress: string,
        collectionTokenId: string,
        page?: number,
        perPage?: number
    ) {
        return (await this.parentProvider.send('qn_getTransfersByNFT', {
            // @ts-ignore
            collection: collectionAddress,
            collectionTokenId: collectionTokenId,
            page: page,
            perPage: perPage,
        })) as NFTTransfersForCollection;
    }

    public async getWalletTokenBalance(
        wallet: string,
        contractAddresses?: string[]
    ): Promise<WalletTokenBalance> {
        const res = (await this.parentProvider.send(
            'qn_getWalletTokenBalance',
            {
                // @ts-ignore
                wallet: wallet,
                contracts: contractAddresses,
            }
        )) as WalletTokenBalance;
        res.owner = res.owner.toLowerCase();

        return res;
    }

    // TODO: Does not seem to work right now. Inform Quick Node about this.
    public async getWalletTokenTransactions(
        walletAddress: string,
        tokenAddress: string
    ): Promise<TokenTransactions> {
        return (await this.parentProvider.send(
            'qn_getWalletTokenTransactions',
            {
                // @ts-ignore
                address: walletAddress,
                contract: tokenAddress,
            }
        )) as TokenTransactions;
    }

    public async verifyNFTsOwner(
        walletAddress: string,
        identifiedContractAddresses: string[]
    ): Promise<NFTsOwner> {
        const res = await this.parentProvider.send('qn_verifyNFTsOwner', [
            walletAddress,
            identifiedContractAddresses,
        ]);

        return res as NFTsOwner;
    }
}
