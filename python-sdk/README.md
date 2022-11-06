# QuickNode Python-SDK

This is the SDK to access QuickNode directly from Python.

## Requirements
- [Python >= 3.10](https://www.python.org/downloads/release/python-3100/)

## Quickstart

```sh
python3 -m pip install -r requirements.txt
```

## Usage

```py
import asyncio
from client import QuickNodeSDK

async def main():
    client = QuickNodeSDK()

    await client.connect()
    try:
        result = await client.get_collection_details('0x2106C00Ac7dA0A3430aE667879139E832307AeAa')

        print(result)
    except Exception as exc:
        print(f"Received exception {exc}")

    await client.close()

asyncio.run(main())
```
:bulb: You can also check [tests files](./tests/) to see usage for each queries.

## Supported Queries
| Methods                                                                     | Status |
|-----------------------------------------------------------------------------|--------|
| [get_collection_details](./client.py#L21)                                   | TESTED |
| [get_contract_event_logs](./client.py#L29)                                  | TESTED |
| [get_nft_details](./client.py#L42)                                          | TESTED |
| [get_nft_event_logs](./client.py#L53)                                       | TESTED |
| [get_nft_by_contract_address](./client.py#L67)                              | TESTED |
| [get_nft_by_wallet_address](./client.py#L79)                                | TESTED |
| [get_nft_by_wallet_and_contracts](./client.py#L91)                          | TESTED |
| [get_nft_by_wallet_ens](./client.py#L104)                                   | TESTED |

## Issues

Even if every query has tests, the data used to compare with function's value are static. However the client will returns different results as things happens in the Chain.