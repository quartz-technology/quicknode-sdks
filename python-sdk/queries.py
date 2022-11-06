from dataclasses import dataclass


GET_COLLECTION_DETAILS = """
query CollectionDetails($address: String!) {
    contract(address: $address) {
      ... on ERC721Contract {
        address
        isVerified
        tokenStandard
        attributes {
          name
          rarity
          value
          valueCount
        }
        circulatingSupply
        name
        stats {
          average
          ceiling
          floor
          totalSales
          volume
        }
        symbol
      }
    }
  }
"""

GET_CONTRACT_EVENT_LOGS = """
query ContractEvents(
    $address: String!
    $filter: LogsFilterInputType
    $first: Int
    $after: String
  ) {
    contract(address: $address) {
      address
      logs(filter: $filter, first: $first, after: $after) {
        pageInfo {
          hasNextPage
          endCursor
        }
        edges {
          node {
            blockNumber
            type
            fromAddress
            toAddress
            estimatedConfirmedAt
            transactionHash
            token {
              contract {
                address
              }
              ... on ERC721Token {
                tokenId
              }
            }
            ... on OrderLog {
              marketplace
              priceInEth
            }
          }
        }
      }
    }
  }
"""

GET_NFT_DETAILS = """
query Token($contractAddress: String!, $tokenId: String!) {
    token(contractAddress: $contractAddress, tokenId: $tokenId) {
      ... on ERC721Token {
        tokenId
        attributes {
          name
          value
        }
        contract {
          address
          isVerified
          tokenStandard
          ... on ERC721Contract {
            name
          }
        }
        images {
          height
          mimeType
          url
          width
        }
        name
        symbol
        metadata {
          animation_url
          background_color
          description
          external_url
          image
          image_data
          name
          youtube_url
        }
      }
    }
  }
"""

GET_NFT_EVENT_LOGS = """
query NFTEvents(
    $address: String!
    $tokenId: String!
    $filter: LogsFilterInputType
    $first: Int
    $after: String
  ) {
    token(contractAddress: $address, tokenId: $tokenId) {
      ... on ERC721Token {
        tokenId
        contract {
          address # Included key field for caching
        }
        logs(filter: $filter, first: $first, after: $after) {
          edges {
            node {
              blockNumber
              type
              fromAddress
              toAddress
              estimatedConfirmedAt
              transactionHash
              ... on OrderLog {
                marketplace
                priceInEth
              }
            }
          }
          pageInfo {
            hasNextPage
            endCursor
          }
        }
      }
    }
  }
"""

GET_NFT_BY_CONTRACT_ADDRESS = """
query ContractNFTs($address: String!, $first: Int, $after: String) {
    contract(address: $address) {
      address
      isVerified
      tokenStandard
      tokens(first: $first, after: $after) {
        pageInfo {
          hasNextPage
          endCursor
        }
        edges {
          node {
            tokenId
            images {
              url
            }
            ... on ERC721Token {
              contract {
                address # Included key field for caching
              }
            }
          }
        }
      }
      ... on ERC721Contract {
        circulatingSupply
        name
        symbol
      }
    }
  }
"""

GET_NFT_BY_WALLET_ADDRESS = """
query WalletNFTs($address: String!, $first: Int, $after: String) {
  wallet(address: $address) {
    ensName
    address
    tokens(first: $first, after: $after) {
      pageInfo {
        hasNextPage
        endCursor
      }
      edges {
        node {
          tokenId
          images {
            url
          }
          ... on ERC721Token {
            contract {
              address
              ... on ERC721Contract {
                symbol
                name
              }
            }
          }
        }
      }
    }
  }
}
"""

GET_NFT_BY_WALLET_AND_CONTRACT = """
query NFTsWalletAndContract(
    $filter: TokensFilterInputType
    $address: String
    $first: Int
    $after: String
  ) {
    wallet(address: $address) {
      ensName
      address
      tokens(filter: $filter, first: $first, after: $after) {
        pageInfo {
          hasNextPage
          endCursor
        }
        edges {
          node {
            tokenId
            images {
              url
            }
            ... on ERC721Token {
              contract {
                address
                ... on ERC721Contract {
                  symbol
                  name
                }
              }
            }
          }
        }
      }
    }
  }
"""

GET_NFT_BY_WALLET_ENS = """
query WalletNFTs($ensName: String, $first: Int, $after: String) {
    wallet(ensName: $ensName) {
      ensName
      address
      tokens(first: $first, after: $after) {
        pageInfo {
          hasNextPage
          endCursor
        }
        edges {
          node {
            tokenId
            images {
              url
            }
            ... on ERC721Token {
              contract {
                address
                ... on ERC721Contract {
                  symbol
                  name
                }
              }
            }
          }
        }
      }
    }
  }
"""
