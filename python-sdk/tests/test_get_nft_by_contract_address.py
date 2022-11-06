import pytest

@pytest.mark.asyncio
async def test_one(client):
    await client.connect()
    try:
        result = await client.get_nft_by_contract_address(
            '0x2106C00Ac7dA0A3430aE667879139E832307AeAa',
            first=5
        )

    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {
        'contract': {
            'address': '0x2106c00ac7da0a3430ae667879139e832307aeaa',
            'isVerified': False,
            'tokenStandard': 'ERC721',
            'tokens': {
                'pageInfo': {
                    'hasNextPage': True,
                    'endCursor': 'YXJyYXljb25uZWN0aW9uOjQ='
                },
                'edges': [{
                    'node': {
                        'tokenId': '0',
                        'images': [{
                            'url': 'https://images.icytools.workers.dev/xs/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/cde9e40e50b3116d97f157a4772f470b'
                        }, {
                            'url': 'https://images.icytools.workers.dev/sm/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/cde9e40e50b3116d97f157a4772f470b'
                        }, {
                            'url': 'https://images.icytools.workers.dev/md/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/cde9e40e50b3116d97f157a4772f470b'
                        }, {
                            'url': 'https://images.icytools.workers.dev/lg/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/cde9e40e50b3116d97f157a4772f470b'
                        }, {
                            'url': 'https://images.icytools.workers.dev/xl/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/cde9e40e50b3116d97f157a4772f470b'
                        }],
                        'contract': {
                            'address': '0x2106c00ac7da0a3430ae667879139e832307aeaa'
                        }
                    }
                }, {
                    'node': {
                        'tokenId': '1',
                        'images': [{
                            'url': 'https://images.icytools.workers.dev/xs/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/0e948cdb828f70766c55131598a3f83d'
                        }, {
                            'url': 'https://images.icytools.workers.dev/sm/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/0e948cdb828f70766c55131598a3f83d'
                        }, {
                            'url': 'https://images.icytools.workers.dev/md/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/0e948cdb828f70766c55131598a3f83d'
                        }, {
                            'url': 'https://images.icytools.workers.dev/lg/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/0e948cdb828f70766c55131598a3f83d'
                        }, {
                            'url': 'https://images.icytools.workers.dev/xl/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/0e948cdb828f70766c55131598a3f83d'
                        }],
                        'contract': {
                            'address': '0x2106c00ac7da0a3430ae667879139e832307aeaa'
                        }
                    }
                }, {
                    'node': {
                        'tokenId': '2',
                        'images': [{
                            'url': 'https://images.icytools.workers.dev/xs/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/2'
                        }, {
                            'url': 'https://images.icytools.workers.dev/sm/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/2'
                        }, {
                            'url': 'https://images.icytools.workers.dev/md/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/2'
                        }, {
                            'url': 'https://images.icytools.workers.dev/lg/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/2'
                        }, {
                            'url': 'https://images.icytools.workers.dev/xl/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/2'
                        }],
                        'contract': {
                            'address': '0x2106c00ac7da0a3430ae667879139e832307aeaa'
                        }
                    }
                }, {
                    'node': {
                        'tokenId': '3',
                        'images': [{
                            'url': 'https://images.icytools.workers.dev/xs/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/3'
                        }, {
                            'url': 'https://images.icytools.workers.dev/sm/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/3'
                        }, {
                            'url': 'https://images.icytools.workers.dev/md/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/3'
                        }, {
                            'url': 'https://images.icytools.workers.dev/lg/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/3'
                        }, {
                            'url': 'https://images.icytools.workers.dev/xl/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/3'
                        }],
                        'contract': {
                            'address': '0x2106c00ac7da0a3430ae667879139e832307aeaa'
                        }
                    }
                }, {
                    'node': {
                        'tokenId': '4',
                        'images': [{
                            'url': 'https://images.icytools.workers.dev/xs/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/4'
                        }, {
                            'url': 'https://images.icytools.workers.dev/sm/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/4'
                        }, {
                            'url': 'https://images.icytools.workers.dev/md/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/4'
                        }, {
                            'url': 'https://images.icytools.workers.dev/lg/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/4'
                        }, {
                            'url': 'https://images.icytools.workers.dev/xl/collections/0x2106c00ac7da0a3430ae667879139e832307aeaa/tokens/4'
                        }],
                        'contract': {
                            'address': '0x2106c00ac7da0a3430ae667879139e832307aeaa'
                        }
                    }
                }]
            },
            'circulatingSupply': 9999,
            'name': 'Loopy Donuts',
            'symbol': 'DONUT'
        }
    }

    await client.close()

@pytest.mark.asyncio
async def test_null_response(client):
    await client.connect()
    try:
        result = await client.get_nft_by_contract_address(
            '0x11111111111110thisisnotanaddress01111111',
            first=5
        )
    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result["contract"] is None

    await client.close()
