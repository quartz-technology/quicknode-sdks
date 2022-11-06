import pytest


@pytest.mark.asyncio
async def test_one(client):
    await client.connect()
    try:
        result = await client.get_nft_by_wallet_address(
            '0xd8da6bf26964af9d7eed9e03e53415d37aa96045',
            first=5
        )
    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {
        'wallet': {
            'ensName': 'vitalik.eth',
            'address': '0xd8da6bf26964af9d7eed9e03e53415d37aa96045',
            'tokens': {
                'pageInfo': {
                    'hasNextPage': True,
                    'endCursor': 'YXJyYXljb25uZWN0aW9uOjQ='
                },
                'edges': [{
                              'node': {
                                  'tokenId': '3',
                                  'images': [{
                                                 'url': 'https://images.icytools.workers.dev/xs/collections/0x4cd3abbe500487b737bc0c9ec9cd7da84f0274df/tokens/c8f8e67636f90a711d346d3a9a4853bb'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/sm/collections/0x4cd3abbe500487b737bc0c9ec9cd7da84f0274df/tokens/c8f8e67636f90a711d346d3a9a4853bb'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/md/collections/0x4cd3abbe500487b737bc0c9ec9cd7da84f0274df/tokens/c8f8e67636f90a711d346d3a9a4853bb'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/lg/collections/0x4cd3abbe500487b737bc0c9ec9cd7da84f0274df/tokens/c8f8e67636f90a711d346d3a9a4853bb'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/xl/collections/0x4cd3abbe500487b737bc0c9ec9cd7da84f0274df/tokens/c8f8e67636f90a711d346d3a9a4853bb'
                                             }],
                                  'contract': {
                                      'address': '0x4cd3abbe500487b737bc0c9ec9cd7da84f0274df',
                                      'symbol': 'THRILL',
                                      'name': 'Mija - The Thrill'
                                  }
                              }
                          }, {
                              'node': {
                                  'tokenId': '4861',
                                  'images': [],
                                  'contract': {
                                      'address': '0x942bc2d3e7a589fe5bd4a5c6ef9727dfd82f5c8a',
                                      'symbol': 'EXPLORE',
                                      'name': 'Friendship Bracelets by Alexis André'
                                  }
                              }
                          }, {
                              'node': {
                                  'tokenId': '1238012972454248237435767387143779415173800484933',
                                  'images': [],
                                  'contract': {
                                      'address': '0x7e902c638db299307565062dc7cd0397431bcb11',
                                      'symbol': 'BLOCKY',
                                      'name': 'On-chain Blockies'
                                  }
                              }
                          }, {
                              'node': {
                                  'tokenId': '5',
                                  'images': [{
                                                 'url': 'https://images.icytools.workers.dev/xs/collections/0xa451f58deca2b78ec16d82da174e10366557e1a5/tokens/a069e9a8cd6cfac42c5af45c375c06d6'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/sm/collections/0xa451f58deca2b78ec16d82da174e10366557e1a5/tokens/a069e9a8cd6cfac42c5af45c375c06d6'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/md/collections/0xa451f58deca2b78ec16d82da174e10366557e1a5/tokens/a069e9a8cd6cfac42c5af45c375c06d6'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/lg/collections/0xa451f58deca2b78ec16d82da174e10366557e1a5/tokens/a069e9a8cd6cfac42c5af45c375c06d6'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/xl/collections/0xa451f58deca2b78ec16d82da174e10366557e1a5/tokens/a069e9a8cd6cfac42c5af45c375c06d6'
                                             }],
                                  'contract': {
                                      'address': '0xa451f58deca2b78ec16d82da174e10366557e1a5',
                                      'symbol': 'DELAY',
                                      'name': 'Mija - The Connection is Delayed'
                                  }
                              }
                          }, {
                              'node': {
                                  'tokenId': '4360',
                                  'images': [{
                                                 'url': 'https://images.icytools.workers.dev/xs/collections/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f/tokens/8f2e60ad7980a8a8cb4335688831ab5c'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/sm/collections/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f/tokens/8f2e60ad7980a8a8cb4335688831ab5c'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/md/collections/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f/tokens/8f2e60ad7980a8a8cb4335688831ab5c'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/lg/collections/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f/tokens/8f2e60ad7980a8a8cb4335688831ab5c'
                                             }, {
                                                 'url': 'https://images.icytools.workers.dev/xl/collections/0xe588feda21d17bd8f7b0532d3e439ab245c2f67f/tokens/8f2e60ad7980a8a8cb4335688831ab5c'
                                             }],
                                  'contract': {
                                      'address': '0xe588feda21d17bd8f7b0532d3e439ab245c2f67f',
                                      'symbol': 'BRITISH',
                                      'name': 'The British'
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
        result = await client.get_nft_by_wallet_address(
            '0x11111111111110thisisnotanaddress01111111',
            first=5
        )
    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result["wallet"] is None

    await client.close()
