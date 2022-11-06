#!/usr/bin/env python3

import asyncio
from client import QuickNodeSDK

async def main():
    client = QuickNodeSDK()

    await client.connect()
    try:
        #result = await client.get_collection_details('0x2106C00Ac7dA0A3430aE667879139E832307AeAa')
        result = await client.get_contract_event_logs('0x60E4d786628Fea6478F785A6d7e704777c86a7c6', first=2, filter=['TRANSFER', 'ORDER', 'MINT'])
        #result = await client.get_nft_details('0x23581767a106ae21c074b2276D25e5C3e136a68b', '400')
        #result = await client.get_nft_event_logs('0x60E4d786628Fea6478F785A6d7e704777c86a7c6', '100', first= 2, filter=['MINT'])
        #result = await client.get_nft_by_contract_address('0x2106C00Ac7dA0A3430aE667879139E832307AeAa', first=5)
        #result = await client.get_nft_by_wallet_address('0xd8da6bf26964af9d7eed9e03e53415d37aa96045', first=5)
        #result = await client.get_nft_by_wallet_and_contracts('0x13928eB9A86c8278a45B6fF2935c7730b58AC675', filter=['0xba30E5F9Bb24caa003E9f2f0497Ad287FDF95623', '0xbCe3781ae7Ca1a5e050Bd9C4c77369867eBc307e'], first=2)
        #result = await client.get_nft_by_wallet_ens('vitalik.eth', first=5)
        print(result)
    except Exception as exc:
        print(f"Received exception {exc}")

    await client.close()

asyncio.run(main())