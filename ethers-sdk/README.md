# ethers-provider-quick-node
An ethers provider exposing Quick Node APIs, aimed to be used with your base provider.

## How does it work ?

Describe how the project works in a few sentences.
- Give a Layman description of the project.
- Show some schemas if required.
- How some of the project's dependencies help.

## Getting started !

### Installing from npm
```shell
yarn add <package_name>
```

### Quickstart

Below is a quick example of the package usage:
```typescript
// Creates the base provider.
const parentProvider = new ethers.providers.JsonRpcProvider(
    'https://fittest-sleek-log.discover.quiknode.pro/YOUR_TOKEN_HERE/'
);

// Creates a Quick Node provider using the parent provider's context.
const qnProvider = new QuickNodeProvider(parentProvider);

const res = await qnProvider.fetchNFTCollectionDetails([
   '0x2106C00Ac7dA0A3430aE667879139E832307AeAa',
]);

console.log(res)
```

Which should print:
````shell
[
  {
    name: 'Loopy Donuts',
    description: 'Loopy Donuts',
    address: '0x2106C00Ac7dA0A3430aE667879139E832307AeAa',
    genesisBlock: 13145256,
    genesisTransaction: '0xcd0d897674123c1859c406b4e14f8a1a6dd5b8b8161bc7563d012e950a26301d',
    erc721: true,
    erc1155: false
  }
]

````

## Author

Made with ❤️ by the folks at 🚀 [Quartz Technology](https://github.com/quartz-technology) 🚀
