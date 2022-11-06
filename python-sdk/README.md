# QuickNode Python-SDK

This is the SDK to access QuickNode directly from Python.

## Requirements
- [Python >= 3.10](https://www.python.org/downloads/release/python-3100/)

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