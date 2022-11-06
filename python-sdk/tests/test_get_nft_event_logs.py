import pytest


@pytest.mark.asyncio
async def test_all_types(client):
    await client.connect()
    try:
        result = await client.get_nft_event_logs(
            '0x60E4d786628Fea6478F785A6d7e704777c86a7c6',
            '100',
            first=2,
            filter=['MINT', 'ORDER', 'TRANSFER']
        )
    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {
        'token': {
            'tokenId': '100',
            'contract': {
                'address': '0x60e4d786628fea6478f785a6d7e704777c86a7c6'
            },
            'logs': {
                'edges': [{
                              'node': {
                                  'blockNumber': 15409149,
                                  'type': 'ORDER',
                                  'fromAddress': '0x3452072f31bf2243f1293884a6c5d8df2726a7e8',
                                  'toAddress': '0x2ae7668a2e0b2aa0377288a017558ed12d259335',
                                  'estimatedConfirmedAt': '2022-08-25T12:41:08.000Z',
                                  'transactionHash': '0x24841172fafb9f9b07fba0b4e4e5b51e0424f22871f5d2f44f55190452ed8936',
                                  'marketplace': 'OPENSEA',
                                  'priceInEth': 14.69
                              }
                          }, {
                              'node': {
                                  'blockNumber': 15171194,
                                  'type': 'TRANSFER',
                                  'fromAddress': '0x828b292daefaa119c021af635879d9629177391b',
                                  'toAddress': '0x3452072f31bf2243f1293884a6c5d8df2726a7e8',
                                  'estimatedConfirmedAt': '2022-07-19T05:32:18.000Z',
                                  'transactionHash': '0x8e16a2be5bb08dc1a5c78edbbed89739aaa81f5d459b4ed5e6b5a6eac1756fd1'
                              }
                          }],
                'pageInfo': {
                    'hasNextPage': True,
                    'endCursor': 'YXJyYXljb25uZWN0aW9uOjE='
                }
            }
        }
    }

    await client.close()


@pytest.mark.asyncio
async def test_one_type(client):
    await client.connect()
    try:
        result = await client.get_nft_event_logs(
            '0x60E4d786628Fea6478F785A6d7e704777c86a7c6',
            '100',
            first=2,
            filter=['MINT']
        )
    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {
        'token': {
            'tokenId': '100',
            'contract': {
                'address': '0x60e4d786628fea6478f785a6d7e704777c86a7c6'
            },
            'logs': {
                'edges': [{
                              'node': {
                                  'blockNumber': 13117233,
                                  'type': 'MINT',
                                  'fromAddress': '0x0000000000000000000000000000000000000000',
                                  'toAddress': '0x5e92a45225c3925b6180b8672a3cf5dd75bcc4a1',
                                  'estimatedConfirmedAt': '2021-08-29T00:45:27.000Z',
                                  'transactionHash': '0xf93a29f670b97370b6e180c958da47b96ba3455ceb479d9d803919405c18ef1f'
                              }
                          }],
                'pageInfo': {
                    'hasNextPage': False,
                    'endCursor': 'YXJyYXljb25uZWN0aW9uOjE='
                }
            }
        }
    }

    await client.close()


@pytest.mark.asyncio
async def test_null_response(client):
    await client.connect()
    try:
        result = await client.get_nft_event_logs(
            '0x11111111111110thisisnotanaddress01111111',
            '100',
            first=2,
        )
    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result['token'] is None

    await client.close()
