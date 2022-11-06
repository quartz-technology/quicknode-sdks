from gql import Client, gql
from gql.transport.aiohttp import AIOHTTPTransport

from queries import GET_COLLECTION_DETAILS, GET_CONTRACT_EVENT_LOGS, GET_NFT_DETAILS, GET_NFT_EVENT_LOGS, GET_NFT_BY_CONTRACT_ADDRESS, GET_NFT_BY_WALLET_ADDRESS, GET_NFT_BY_WALLET_AND_CONTRACT, GET_NFT_BY_WALLET_ENS

class QuickNodeSDK:
    def __init__(self, apiKey: str = ""):
        self._client = Client(
            transport = AIOHTTPTransport(
                url="https://graphql.icy.tools/graphql",
                headers={'x-api-key': apiKey})
        )
        self._session = None

    async def connect(self):
        self._session = await self._client.connect_async(reconnecting=True)

    async def close(self):
        await self._client.close_async()
    
    async def get_collection_details(self, address: str):
        params = {"address": address}

        return await self._session.execute(
            gql(GET_COLLECTION_DETAILS),
            variable_values=params
        )

    async def get_contract_event_logs(self, address: str, filter: list[str] = None, first: int = None, after: str = None):
        params = {
            "address": address,
            "filter": {"typeIn": filter},
            "first": first,
            "after": after
        }

        return await self._session.execute(
            gql(GET_CONTRACT_EVENT_LOGS),
            variable_values=params
        )

    async def get_nft_details(self, contractAddress: str, tokenID: str):
        params = {
            "contractAddress": contractAddress,
            "tokenId": tokenID,
        }

        return await self._session.execute(
            gql(GET_NFT_DETAILS), 
            variable_values=params
        )

    async def get_nft_event_logs(self, address: str, tokenID: str, filter: list[str] = None, first: int = None, after: str = None):
        params = {
            "address": address,
            "tokenId": tokenID,
            "filter": {"typeIn": filter},
            "first": first,
            "after": after
        }

        return await self._session.execute(
            gql(GET_NFT_EVENT_LOGS), 
            variable_values=params
        )

    async def get_nft_by_contract_address(self, address: str, first: int = None, after: str = None):
        params = {
            "address": address,
            "first": first,
            "after": after
            }

        return await self._session.execute(
            gql(GET_NFT_BY_CONTRACT_ADDRESS),
            variable_values=params
        )

    async def get_nft_by_wallet_address(self, address: str, first: int = None, after: str = None):
        params = {
            "address": address,
            "first": first,
            "after": after
            }

        return await self._session.execute(
            gql(GET_NFT_BY_WALLET_ADDRESS),
            variable_values=params
        )

    async def get_nft_by_wallet_and_contracts(self, address: str, filter: list[str] = None, first: int = None, after: str = None):
        params = {
            "address": address,
            "filter": {"contractAddressIn": filter},
            "first": first,
            "after": after
        }

        return await self._session.execute(
            gql(GET_NFT_BY_WALLET_AND_CONTRACT),
            variable_values=params
        )

    async def get_nft_by_wallet_ens(self, ensName: str, first: int = None, after: str = None):
        params = {
            "ensName": ensName,
            "first": first,
            "after": after
        }

        return await self._session.execute(
            gql(GET_NFT_BY_WALLET_ENS),
            variable_values=params
        )