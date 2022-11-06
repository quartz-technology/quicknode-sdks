import pytest


@pytest.mark.asyncio
async def test_one(client):
    await client.connect()
    try:
        result = await client.get_nft_details(
            '0x23581767a106ae21c074b2276D25e5C3e136a68b',
            '400'
        )
    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result == {
        'token': {
            'tokenId': '400',
            'attributes': [{
                               'name': 'Background',
                               'value': 'Yellow'
                           }, {
                               'name': 'Beak',
                               'value': 'Small'
                           }, {
                               'name': 'Body',
                               'value': 'Guardian'
                           }, {
                               'name': 'Eyes',
                               'value': 'Discerning'
                           }, {
                               'name': 'Feathers',
                               'value': 'Purple'
                           }, {
                               'name': 'Headwear',
                               'value': 'Flower'
                           }],
            'contract': {
                'address': '0x23581767a106ae21c074b2276d25e5c3e136a68b',
                'isVerified': True,
                'tokenStandard': 'ERC721',
                'name': 'Moonbirds'
            },
            'images': [{
                           'height': 100,
                           'mimeType': 'image/png',
                           'url': 'https://images.icytools.workers.dev/xs/collections/0x23581767a106ae21c074b2276d25e5c3e136a68b/tokens/bde0877c73293c2d23da83c8e823a3fe',
                           'width': 100
                       }, {
                           'height': 200,
                           'mimeType': 'image/png',
                           'url': 'https://images.icytools.workers.dev/sm/collections/0x23581767a106ae21c074b2276d25e5c3e136a68b/tokens/bde0877c73293c2d23da83c8e823a3fe',
                           'width': 200
                       }, {
                           'height': 400,
                           'mimeType': 'image/png',
                           'url': 'https://images.icytools.workers.dev/md/collections/0x23581767a106ae21c074b2276d25e5c3e136a68b/tokens/bde0877c73293c2d23da83c8e823a3fe',
                           'width': 400
                       }, {
                           'height': 800,
                           'mimeType': 'image/png',
                           'url': 'https://images.icytools.workers.dev/lg/collections/0x23581767a106ae21c074b2276d25e5c3e136a68b/tokens/bde0877c73293c2d23da83c8e823a3fe',
                           'width': 800
                       }, {
                           'height': 1200,
                           'mimeType': 'image/png',
                           'url': 'https://images.icytools.workers.dev/xl/collections/0x23581767a106ae21c074b2276d25e5c3e136a68b/tokens/bde0877c73293c2d23da83c8e823a3fe',
                           'width': 1200
                       }],
            'name': '#400',
            'symbol': None,
            'metadata': {
                'animation_url': None,
                'background_color': None,
                'description': None,
                'external_url': 'https://moonbirds.xyz/',
                'image': 'https://live---metadata-5covpqijaa-uc.a.run.app/images/400',
                'image_data': None,
                'name': '#400',
                'youtube_url': None
            }
        }
    }

    await client.close()


@pytest.mark.asyncio
async def test_null_response(client):
    await client.connect()
    try:
        result = await client.get_nft_details(
            '0x23581767a106ae21c074b2276D25e5C3e136a68b',
            '444444444'
        )
    except Exception as exc:
        print(f"Received exception {exc} ")
    assert result['token'] is None
    await client.close()
