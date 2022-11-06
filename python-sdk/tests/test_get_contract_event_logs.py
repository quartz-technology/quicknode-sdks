import pytest

"""
@pytest.mark.asyncio
async def test_all_types(client):
    await client.connect()
    try:
        result = await client.get_contract_event_logs(
            '0x60E4d786628Fea6478F785A6d7e704777c86a7c6',
            first=2,
            filter=['TRANSFER', 'ORDER', 'MINT']
        )

    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {
        'contract': {
            'address': '0x60e4d786628fea6478f785a6d7e704777c86a7c6',
            'logs': {
                'pageInfo': {
                    'hasNextPage': True,
                    'endCursor': 'YXJyYXljb25uZWN0aW9uOjE='
                },
                'edges': [{
                              'node': {
                                  'blockNumber': 15907764,
                                  'type': 'TRANSFER',
                                  'fromAddress': '0xcf5e233bcb1f1b3536e0879be2d3294049d7bfc7',
                                  'toAddress': '0xc28f7ee92cd6619e8eec6a70923079fbafb86196',
                                  'estimatedConfirmedAt': '2022-11-06T01:42:47.000Z',
                                  'transactionHash': '0x2fe02df2e88f91b206e9fa5580cf171b5c6fffcd874f5ca16bae1e29fdc61306',
                                  'token': {
                                      'contract': {
                                          'address': '0x60e4d786628fea6478f785a6d7e704777c86a7c6'
                                      },
                                      'tokenId': '29228'
                                  }
                              }
                          }, {
                              'node': {
                                  'blockNumber': 15907713,
                                  'type': 'TRANSFER',
                                  'fromAddress': '0xb4debf1547c4df92ccf15ae0c70a61ae6e5f002d',
                                  'toAddress': '0xc8f7e8566cb717c32e8727700c6d646a967deead',
                                  'estimatedConfirmedAt': '2022-11-06T01:32:35.000Z',
                                  'transactionHash': '0xf0c9cf2859b18d7ebb3bcc16ce5356b68a96ed666161be09a3e4c0ed9d93e57c',
                                  'token': {
                                      'contract': {
                                          'address': '0x60e4d786628fea6478f785a6d7e704777c86a7c6'
                                      },
                                      'tokenId': '12367'
                                  }
                              }
                          }]
            }
        }
    }

    await client.close()
"""

@pytest.mark.asyncio
async def test_one_type(client):
    await client.connect()
    try:
        result = await client.get_contract_event_logs(
            '0x60E4d786628Fea6478F785A6d7e704777c86a7c6',
            first=2,
            filter=['MINT']
        )

    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {
        'contract': {
            'address': '0x60e4d786628fea6478f785a6d7e704777c86a7c6',
            'logs': {
                'pageInfo': {
                    'hasNextPage': True,
                    'endCursor': 'YXJyYXljb25uZWN0aW9uOjE='
                },
                'edges': [{
                    'node': {
                        'blockNumber': 15782593,
                        'type': 'MINT',
                        'fromAddress': '0x0000000000000000000000000000000000000000',
                        'toAddress': '0x3c05c8c71fc3bd1e095a71f1ab58a9142b41e8ad',
                        'estimatedConfirmedAt': '2022-10-19T13:56:59.000Z',
                        'transactionHash': '0x410a3d39fc380ae9f24131a78c35d16c8a7bb7a20d8cb835871d3fdb64a7e0c4',
                        'token': {
                            'contract': {
                                'address': '0x60e4d786628fea6478f785a6d7e704777c86a7c6'
                            },
                            'tokenId': '16227'
                        }
                    }
                }, {
                    'node': {
                        'blockNumber': 15657628,
                        'type': 'MINT',
                        'fromAddress': '0x0000000000000000000000000000000000000000',
                        'toAddress': '0x9524b86e777c0c6e1e1afe0bb8bf931a5b81386c',
                        'estimatedConfirmedAt': '2022-10-02T02:59:59.000Z',
                        'transactionHash': '0x64a195a18d501f1aafd911a64f80d990f0b5311ad8bad44c41136c69641c7cc5',
                        'token': {
                            'contract': {
                                'address': '0x60e4d786628fea6478f785a6d7e704777c86a7c6'
                            },
                            'tokenId': '24337'
                        }
                    }
                }]
            }
        }
    }
    await client.close()


@pytest.mark.asyncio
async def test_null_response(client):
    await client.connect()
    try:
        result = await client.get_contract_event_logs(
            '0x11111111111110thisisnotanaddress01111111',
            first=2,
        )

    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result["contract"] is None
    await client.close()
