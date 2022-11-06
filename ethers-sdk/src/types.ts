//                                  //
//  USED FOR PAGINATED RESPONSES    //
//                                  //
export interface Pagination {
    totalItems: number;
    totalPages: number;
    pageNumber: number;
}

//                          //
//  USED BY qn_fetchNFTs    //
//                          //
export interface NFTDetails extends Pagination {
    owner: string;
    assets: NFTAssetDetailsWithHistory[];
}

export interface BaseNFTAssetDetails {
    name: string;
    collectionTokenId: string;
    collectionName: string;
    imageUrl: string;
    collectionAddress: string;
    traits: string[];
    chain: string;
    network: string;
}

export interface NFTAssetDetailsWithHistory extends BaseNFTAssetDetails {
    provenance: NFTAssetProvenance[];
}

export interface NFTAssetProvenance {
    blockNumber: number;
    date: number;
    from: string;
    to: string;
}

//                                          //
//  USED BY qn_fetchNFTCollectionDetails    //
//                                          //
export interface NFTCollectionDetails {
    name: string;
    description: string;
    address: string;
    genesisBlock: number;
    genesisTransaction: string;
    erc721: boolean;
    erc1155: boolean;
}

//                                      //
//  USED BY qn_fetchNFTsByCollection    //
//                                      //
export interface NFTsInCollection extends Pagination {
    collection: string;
    tokens: BaseNFTAssetDetails[];
}

//                                                                                  //
//  USED BY qn_getTokenMetadataByContractAddress AND qn_getTokenMetadataBySymbol    //
//                                                                                  //
export interface TokenMetadata {
    name: string;
    symbol: string;
    address: string;
    decimals: number;
    genesisBlock: number;
    genesisTransaction: string;
}

export interface TokenMetadataPaginated extends Pagination {
    tokens: TokenMetadata[];
}

//                                  //
//  USED BY qn_getTransfersByNFT    //
//                                  //
export interface NFTTransfersForCollection extends Pagination {
    collection: string;
    transfers: NFTTransfer[];
}

export interface NFTTransfer {
    blockNumber: number;
    date: number;
    from: string;
    to: string;
    txHash: string;
}

//                                      //
//  USED BY qn_getWalletTokenBalance    //
//                                      //
export interface WalletTokenBalance {
    owner: string;
    assets: TokenBalance[];
}

export interface TokenBalance {
    address: string;
    amount: number;
    chain: string;
    decimals: number;
    name: string;
    network: string;
    symbol: string;
}

//                              //
//  USED BY qn_verifyNFTsOwner  //
//                              //
export interface NFTsOwner {
    owner: string;
    assets: string[];
}

//                                          //
//  USED BY qn_getWalletTokenTransactions   //
//                                          //
export interface TokenTransactions extends Pagination {
    token: TokenMetadata;
    transfers: TokenTransfer[];
}

export interface TokenTransfer {
    amount: number;
    blockNumber: number;
    date: number;
    from: string;
    logIndex: number;
    to: string;
    txHash: string;
    valueSent: number;
}
