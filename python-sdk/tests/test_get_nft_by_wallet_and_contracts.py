import pytest

@pytest.mark.asyncio
async def test_one_contract(client):
    await client.connect()
    try:
        result = await client.get_nft_by_wallet_and_contracts(
            '0x13928eB9A86c8278a45B6fF2935c7730b58AC675',
            filter=['0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D'],
            first=2
        )

    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {'wallet': {'ensName': 'bitcoinpirate.eth', 'address': '0x13928eb9a86c8278a45b6ff2935c7730b58ac675', 'tokens': {'pageInfo': {'hasNextPage': False, 'endCursor': 'YXJyYXljb25uZWN0aW9uOjE='}, 'edges': []}}}


    await client.close()


@pytest.mark.asyncio
async def test_multiple_contracts(client):
    await client.connect()
    try:
        result = await client.get_nft_by_wallet_and_contracts(
            '0x13928eB9A86c8278a45B6fF2935c7730b58AC675',
            filter=['0xba30E5F9Bb24caa003E9f2f0497Ad287FDF95623', '0xbCe3781ae7Ca1a5e050Bd9C4c77369867eBc307e'],
            first=2
        )

    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {'wallet': {'ensName': 'bitcoinpirate.eth', 'address': '0x13928eb9a86c8278a45b6ff2935c7730b58ac675', 'tokens': {'pageInfo': {'hasNextPage': False, 'endCursor': 'YXJyYXljb25uZWN0aW9uOjE='}, 'edges': []}}}

    await client.close()


@pytest.mark.asyncio
async def test_null_response(client):
    await client.connect()
    try:
        result = await client.get_nft_by_wallet_and_contracts(
            '0x11111111111110thisisnotanaddress01111111',
            filter=['0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d'],
            first=2
        )

    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result["wallet"] is None

    await client.close()