# template-npm
A template for building, publishing and maintaining npm packages.

## ⚠️ Setup ⚠️

Please read carefully this section as it describes how to properly set up the repository once you've forked this
template.

There are a few things you **need** to change and know before adding your own code:
1. Change the `name`, `version`, `description`, maybe the `author` and the `license` fields in the `package.json`
file.
2. This template forces you to use the [conventional-commit](https://www.conventionalcommits.org/en/v1.0.0/) specification.
3. In order for the releaser to work properly, you must set the `NPM_TOKEN` and `GH_TOKEN` secrets in your fork.
   1. Both of these can be found/generated respectively [here](https://docs.npmjs.com/creating-and-viewing-access-tokens)
   and [here](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).
   2. Triggering releases is done by pushing code to the main branch. Versioning will be managed by the semantic-release
   GitHub action, which you can learn more about [here](https://semantic-release.gitbook.io/semantic-release/#how-does-it-work).
4. You must uncomment the last step in the `.github/workflows/release.yaml` file.

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

```

## Author

Made with ❤️ by the folks at 🚀 [Quartz Technology](https://github.com/quartz-technology) 🚀
