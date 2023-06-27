# GitHub Actions

This project was configured in a way that you're able to run GitHub actions locally using https://github.com/nektos/act.

The suggested way of installing it is using Brew:

```bash
brew install act
```

## Running actions

To run the libraries test workflows execute:

```bash
act workflow_dispatch -j unit_test_service
```

In the root of the project.