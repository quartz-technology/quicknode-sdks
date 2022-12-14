import pytest

@pytest.mark.asyncio
async def test_one(client):
    await client.connect()

    try:
        result = await client.get_collection_details('0x2106C00Ac7dA0A3430aE667879139E832307AeAa')
    except Exception as exc:
        print(f"Received exception {exc} while trying to get continent name")

        assert result == {
            'contract':
                {
                    'address':
                        '0x2106c00ac7da0a3430ae667879139e832307aeaa',
                    'isVerified': False,
                    'tokenStandard': 'ERC721',
                    'attributes': [
                        {
                            'name': 'Headgear',
                            'rarity': 0.24002400240024002,
                            'value': 'Strawberry Cream',
                            'valueCount': 24
                        },
                        {
                            'name': 'Headgear',
                            'rarity': 1.05010501050105,
                            'value': 'Pirate',
                            'valueCount': 105
                        },
                        {
                            'name': 'Eyes',
                            'rarity': 0.16001600160016002,
                            'value': 'Nerdy',
                            'valueCount': 16
                        },
                        {
                            'name': 'Accessory',
                            'rarity': 0.2000200020002,
                            'value': 'Yellow Skirt',
                            'valueCount': 20
                        },
                        {
                            'name': 'Headgear',
                            'rarity': 0.25002500250025,
                            'value': 'Silver Girlish',
                            'valueCount': 25
                        },
                        {
                            'name': 'Eyes',
                            'rarity': 1.3201320132013201,
                            'value': 'Sleeping',
                            'valueCount': 132
                        },
                        {
                            'name': 'Accessory',
                            'rarity': 0.14001400140014,
                            'value': 'Cloud',
                            'valueCount': 14
                        },
                        {
                            'name': 'Accessory',
                            'rarity': 0.07000700070007,
                            'value': 'Green Skirt',
                            'valueCount': 7
                        },
                        {
                            'name': 'Accessory',
                            'rarity': 0.020002000200020003,
                            'value': 'Skirt',
                            'valueCount': 2
                        },
                        {
                            'name': 'Base',
                            'rarity': 0.020002000200020003,
                            'value': 'Chocolate Fudge',
                            'valueCount': 2
                        },
                        {
                            'name': 'Base',
                            'rarity': 0.020002000200020003,
                            'value': 'Vanilla Candy',
                            'valueCount': 2
                        },
                        {
                            'name': 'Headgear',
                            'rarity': 1.5901590159015901,
                            'value': 'Cat',
                            'valueCount': 159
                        },
                        {
                            'name': 'Accessory',
                            'rarity': 1.7301730173017302,
                            'value': 'Donut Inflatable',
                            'valueCount': 173
                        },
                        {
                            'name': 'Headgear',
                            'rarity': 0.34003400340034,
                            'value': 'Brown Hair',
                            'valueCount': 34
                        },
                        {
                            'name': 'Mouth',
                            'rarity': 0.48004800480048004,
                            'value': 'Diamond',
                            'valueCount': 48
                        },
                        {
                            'name': 'Accessory',
                            'rarity': 0.22002200220022003,
                            'value': 'Pink Wings',
                            'valueCount': 22
                        }, {
                            'name': 'Mouth',
                            'rarity': 0.09000900090009001,
                            'value': 'Covid Mask',
                            'valueCount': 9
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.2000200020002,
                            'value': 'Japanese',
                            'valueCount': 20
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.71007100710071,
                            'value': 'Chastity Belt',
                            'valueCount': 71
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.51005100510051,
                            'value': 'Bat Wings',
                            'valueCount': 51
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.56005600560056,
                            'value': 'Bunny',
                            'valueCount': 56
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.15001500150015,
                            'value': 'Green Pouch',
                            'valueCount': 15
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.59005900590059,
                            'value': 'Donut Nest',
                            'valueCount': 59
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.71007100710071,
                            'value': 'Red Girlish',
                            'valueCount': 71
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.4000400040004,
                            'value': 'Red Pouch',
                            'valueCount': 40
                        }, {
                            'name': 'Glaze',
                            'rarity': 1.5501550155015502,
                            'value': 'Stars',
                            'valueCount': 155
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.49004900490049,
                            'value': 'Icecream',
                            'valueCount': 49
                        }, {
                            'name': 'Glaze',
                            'rarity': 0.4000400040004,
                            'value': 'Pink and Sparks',
                            'valueCount': 40
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.27002700270027,
                            'value': 'Blonde Girlish',
                            'valueCount': 27
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.19001900190019003,
                            'value': 'Red Skate',
                            'valueCount': 19
                        }, {
                            'name': 'Eyes',
                            'rarity': 0.12001200120012001,
                            'value': 'Sleepmask',
                            'valueCount': 12
                        }, {
                            'name': 'Mouth',
                            'rarity': 1.0701070107010702,
                            'value': 'Fangs',
                            'valueCount': 107
                        }, {
                            'name': 'Glaze',
                            'rarity': 1.2601260126012601,
                            'value': 'Radioactive',
                            'valueCount': 126
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.18011801180118,
                            'value': 'Green Frog',
                            'valueCount': 118
                        }, {
                            'name': 'Accessory',
                            'rarity': 1.2501250125012502,
                            'value': 'Adam',
                            'valueCount': 125
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.12001200120012001,
                            'value': 'Planets',
                            'valueCount': 12
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.22002200220022003,
                            'value': 'Green Casino Hat',
                            'valueCount': 22
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.29002900290029004,
                            'value': 'Red Skirt',
                            'valueCount': 29
                        }, {
                            'name': 'Glaze',
                            'rarity': 3.2803280328032804,
                            'value': 'Galaxy',
                            'valueCount': 328
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.3401340134013402,
                            'value': 'Goose',
                            'valueCount': 134
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.12011201120112,
                            'value': 'Red Frog',
                            'valueCount': 112
                        }, {
                            'name': 'Glaze',
                            'rarity': 0.24002400240024002,
                            'value': 'Rock',
                            'valueCount': 24
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.65006500650065,
                            'value': 'Elvis',
                            'valueCount': 65
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.8200820082008201,
                            'value': 'Pink Casino Hat',
                            'valueCount': 82
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.13001300130013002,
                            'value': 'Duck Inflatable',
                            'valueCount': 13
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.77007700770077,
                            'value': 'Blue Skate',
                            'valueCount': 77
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.21012101210121,
                            'value': 'Brain',
                            'valueCount': 121
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.7200720072007201,
                            'value': 'Blue Wings',
                            'valueCount': 72
                        }, {
                            'name': 'Eyes',
                            'rarity': 0.46004600460046,
                            'value': 'Spiral',
                            'valueCount': 46
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.47004700470047006,
                            'value': 'VS',
                            'valueCount': 47
                        }, {
                            'name': 'Mouth',
                            'rarity': 2.48024802480248,
                            'value': 'Gold Teeth',
                            'valueCount': 248
                        }, {
                            'name': 'Glaze',
                            'rarity': 0.9300930093009301,
                            'value': 'Elegant Blue',
                            'valueCount': 93
                        }, {
                            'name': 'Glaze',
                            'rarity': 1.4901490149014902,
                            'value': 'Chocolate Cookie',
                            'valueCount': 149
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.1001100110011002,
                            'value': 'Filling Syringe',
                            'valueCount': 110
                        }, {
                            'name': 'Mouth',
                            'rarity': 1.4101410141014101,
                            'value': 'Mustache',
                            'valueCount': 141
                        }, {
                            'name': 'Background',
                            'rarity': 4.5004500450045,
                            'value': 'Jade',
                            'valueCount': 450
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.6000600060006,
                            'value': 'Blue Girlish',
                            'valueCount': 60
                        }, {
                            'name': 'Eyes',
                            'rarity': 0.47004700470047006,
                            'value': 'Fox Sleepmask',
                            'valueCount': 47
                        }, {
                            'name': 'Glaze',
                            'rarity': 1.6501650165016502,
                            'value': 'Elegant Pink',
                            'valueCount': 165
                        }, {
                            'name': 'Back-accessory',
                            'rarity': 1.6001600160016,
                            'value': 'Superman',
                            'valueCount': 160
                        }, {
                            'name': 'Eyes',
                            'rarity': 2.7002700270027002,
                            'value': 'Sunglasses',
                            'valueCount': 270
                        }, {
                            'name': 'Glaze',
                            'rarity': 1.08010801080108,
                            'value': 'Elegant Fuchsia',
                            'valueCount': 108
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.74007400740074,
                            'value': 'Crown',
                            'valueCount': 74
                        }, {
                            'name': 'Mouth',
                            'rarity': 0.7800780078007801,
                            'value': 'Kinky',
                            'valueCount': 78
                        }, {
                            'name': 'Eyes',
                            'rarity': 1.39013901390139,
                            'value': 'Green Party Glasses',
                            'valueCount': 139
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.45004500450045004,
                            'value': 'Unicorn',
                            'valueCount': 45
                        }, {
                            'name': 'Eyes',
                            'rarity': 0.9900990099009901,
                            'value': 'Crossed Eye',
                            'valueCount': 99
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.8200820082008201,
                            'value': 'Fish',
                            'valueCount': 82
                        }, {
                            'name': 'Eyes',
                            'rarity': 2.33023302330233,
                            'value': 'Curious',
                            'valueCount': 233
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.65006500650065,
                            'value': 'Pink Horns',
                            'valueCount': 65
                        }, {
                            'name': 'Back-accessory',
                            'rarity': 1.51015101510151,
                            'value': 'King',
                            'valueCount': 151
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.95009500950095,
                            'value': 'Liberty',
                            'valueCount': 95
                        }, {
                            'name': 'Mouth',
                            'rarity': 0.95009500950095,
                            'value': 'Rainbow Beam',
                            'valueCount': 95
                        }, {
                            'name': 'Glaze',
                            'rarity': 2.4002400240024,
                            'value': 'Purple Candy',
                            'valueCount': 240
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.4901490149014902,
                            'value': 'Angel',
                            'valueCount': 149
                        }, {
                            'name': 'Background',
                            'rarity': 4.280428042804281,
                            'value': 'Red',
                            'valueCount': 428
                        }, {
                            'name': 'Eyes',
                            'rarity': 1.9501950195019502,
                            'value': 'Yellow Party Glasses',
                            'valueCount': 195
                        }, {
                            'name': 'Mouth',
                            'rarity': 1.5001500150015001,
                            'value': 'Donut Mask',
                            'valueCount': 150
                        }, {
                            'name': 'Glaze',
                            'rarity': 0.9000900090009001,
                            'value': 'Rainbow',
                            'valueCount': 90
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.2601260126012601,
                            'value': 'Pink Hat',
                            'valueCount': 126
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.53005300530053,
                            'value': 'Diaper',
                            'valueCount': 53
                        }, {
                            'name': 'Background',
                            'rarity': 4.2004200420042,
                            'value': 'Yellow',
                            'valueCount': 420
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.63016301630163,
                            'value': 'Red Horns',
                            'valueCount': 163
                        }, {
                            'name': 'Accessory',
                            'rarity': 0.8200820082008201,
                            'value': 'VS Red',
                            'valueCount': 82
                        }, {
                            'name': 'Eyes',
                            'rarity': 1.18011801180118,
                            'value': 'Bored',
                            'valueCount': 118
                        }, {
                            'name': 'Glaze',
                            'rarity': 2.5602560256025604,
                            'value': 'Blue Skittles',
                            'valueCount': 256
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.54005400540054,
                            'value': 'Professor',
                            'valueCount': 54
                        }, {
                            'name': 'Accessory',
                            'rarity': 1.21012101210121,
                            'value': 'Crocodile Inflatable',
                            'valueCount': 121
                        }, {
                            'name': 'Glaze',
                            'rarity': 3.7603760376037605,
                            'value': 'Caramel',
                            'valueCount': 376
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.9600960096009601,
                            'value': 'Roman',
                            'valueCount': 96
                        }, {
                            'name': 'Eyes',
                            'rarity': 2.6702670267026702,
                            'value': 'Happy',
                            'valueCount': 267
                        }, {
                            'name': 'Background',
                            'rarity': 4.5004500450045,
                            'value': 'Purple',
                            'valueCount': 450
                        }, {
                            'name': 'Eyes',
                            'rarity': 3.0003000300030003,
                            'value': 'Pretty Please',
                            'valueCount': 300
                        }, {
                            'name': 'Glaze',
                            'rarity': 0.92009200920092,
                            'value': 'Teal and Sparks',
                            'valueCount': 92
                        }, {
                            'name': 'Eyes',
                            'rarity': 1.7001700170017002,
                            'value': 'Diva',
                            'valueCount': 170
                        }, {
                            'name': 'Eyes',
                            'rarity': 2.1702170217021703,
                            'value': 'Surprised',
                            'valueCount': 217
                        }, {
                            'name': 'Accessory',
                            'rarity': 1.51015101510151,
                            'value': 'Arrow',
                            'valueCount': 151
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.17011701170117,
                            'value': 'Fruit Bowl',
                            'valueCount': 117
                        }, {
                            'name': 'Base',
                            'rarity': 29.832983298329832,
                            'value': 'Donut Body',
                            'valueCount': 2983
                        }, {
                            'name': 'Eyes',
                            'rarity': 1.48014801480148,
                            'value': 'X Eyes',
                            'valueCount': 148
                        }, {
                            'name': 'Eyes',
                            'rarity': 2.25022502250225,
                            'value': 'Purple Party Glasses',
                            'valueCount': 225
                        }, {
                            'name': 'Accessory',
                            'rarity': 1.36013601360136,
                            'value': 'Yellow Pouch',
                            'valueCount': 136
                        }, {
                            'name': 'Background',
                            'rarity': 4.09040904090409,
                            'value': 'Pink',
                            'valueCount': 409
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.2601260126012601,
                            'value': 'Orange Hat',
                            'valueCount': 126
                        }, {
                            'name': 'Accessory',
                            'rarity': 1.9001900190019,
                            'value': 'Purple Skirt',
                            'valueCount': 190
                        }, {
                            'name': 'Background',
                            'rarity': 4.14041404140414,
                            'value': 'Green',
                            'valueCount': 414
                        }, {
                            'name': 'Glaze',
                            'rarity': 3.8003800380038,
                            'value': 'Starcandy',
                            'valueCount': 380
                        }, {
                            'name': 'Headgear',
                            'rarity': 0.7000700070007001,
                            'value': 'Bouquet',
                            'valueCount': 70
                        }, {
                            'name': 'Eyes',
                            'rarity': 3.5503550355035505,
                            'value': 'Cool',
                            'valueCount': 355
                        }, {
                            'name': 'Mouth',
                            'rarity': 3.9703970397039705,
                            'value': 'La la la',
                            'valueCount': 397
                        }, {
                            'name': 'Background',
                            'rarity': 4.220422042204221,
                            'value': 'Blue',
                            'valueCount': 422
                        }, {
                            'name': 'Glaze',
                            'rarity': 3.6603660366036603,
                            'value': 'Pink Candy',
                            'valueCount': 366
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.7001700170017002,
                            'value': 'Loopies',
                            'valueCount': 170
                        }, {
                            'name': 'Accessory',
                            'rarity': 1.7401740174017402,
                            'value': 'Babypink Pouch',
                            'valueCount': 174
                        }, {
                            'name': 'Headgear',
                            'rarity': 1.0401040104010402,
                            'value': 'Quinn',
                            'valueCount': 104
                        }, {
                            'name': 'Headgear',
                            'rarity': 70.10701070107011,
                            'value': None,
                            'valueCount': 7010
                        }, {
                            'name': 'Eyes',
                            'rarity': 70.10701070107011,
                            'value': None,
                            'valueCount': 7010
                        }, {
                            'name': 'Accessory',
                            'rarity': 82.83828382838284,
                            'value': None,
                            'valueCount': 8283
                        }, {
                            'name': 'Base',
                            'rarity': 70.12701270127013,
                            'value': None,
                            'valueCount': 7012
                        }, {
                            'name': 'Mouth',
                            'rarity': 87.26872687268727,
                            'value': None,
                            'valueCount': 8726
                        }, {
                            'name': 'Glaze',
                            'rarity': 70.11701170117011,
                            'value': None,
                            'valueCount': 7011
                        }, {
                            'name': 'Background',
                            'rarity': 70.06700670067006,
                            'value': None,
                            'valueCount': 7006
                        }, {
                            'name': 'Back-accessory',
                            'rarity': 96.88968896889689,
                            'value': None,
                            'valueCount': 9688
                        }],
                    'circulatingSupply': 9999,
                    'name': 'Loopy Donuts',
                    'stats': {
                        'average': None,
                        'ceiling': None,
                        'floor': None,
                        'totalSales': 0,
                        'volume': 0
                    },
                    'symbol': 'DONUT'
                }
        }

    await client.close()


@pytest.mark.asyncio
async def test_two(client):
    await client.connect()

    try:
        result = await client.get_collection_details('0x11111111111110thisisnotanaddress01111111')
    except Exception as exc:
        print(f"Received exception {exc} ")

    assert result['contract'] is None

    await client.close()
