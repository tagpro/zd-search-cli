# Zendesk Search tool

[![Build](https://github.com/tagpro/zd-search-cli/actions/workflows/build.yml/badge.svg)](https://github.com/tagpro/zd-search-cli/actions/workflows/build.yml)
[![Lint](https://github.com/tagpro/zd-search-cli/actions/workflows/lint.yml/badge.svg)](https://github.com/tagpro/zd-search-cli/actions/workflows/lint.yml)

## Setup requirements

Go 1.16 was the development setup for this application. 

The easiest way to install go is using the official docs. Install golang by following the steps [here](https://golang.org/doc/install)

If you are on mac, you can run `brew install go` to install the latest version of golang.

If you are on Ubuntu, you can follow the steps here - https://github.com/golang/go/wiki/Ubuntu

## Running the application locally

Run the following command to start the application in your local terminal.

```bash
make run
```

## Running linter locally

This repo uses [golangci-lint](https://github.com/golangci/golangci-lint) to do linting. 

To install the tool, run `make install-linter` inside the terminal from the root of the repo. 
It requires `$GOPATH/bin` to be available inside the `$PATH` (official Go installer does this by default).

If you are on linux, setting the following in `~/.zshrc` or `~/.bashrc` or any other shell of your choosing and restarting the terminal should work 

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Alternatively, you can follow installations instructions [here](https://golangci-lint.run/usage/install/#local-installation) 

## Repo layout

- The entrypoint of the application is `cmd/search` which has the main golang file. 
- The application starts from `pkg/cmd/search` which is called by `cmd/search` to start the app. This is the first layer
  of the application. It is responsible to start and run the application. The logic to take user input is also defined
  in this layer
- The layer on top of `pkg/cmd/search` is called `pkg/serializer`. Serializer could be a misleading term to what it does. 
  It has the connection to the data store. It takes the input from the `pkg/cmd/search` and uses that to fetch the data 
  from the store. It also handles the logic to print the data on the screen.
- The layer on top of `pkg/serializer` is `pkg/store`, which is initialized by the app and sent to serializer for future
  use. `pkg/store` caches all the data from input files into memory to be fetched by the serializer quickly.

## Assumptions

- `_id` values are unique per record for all the users, tickets and, organisations.

## How to use the app

### To search the data, select the options displayed on the screen `Search Zendesk` and press enter.

![primary options](./docs/images/select-primary-option.png)

### After selecting the option, choose the type of data you want to search

![secondary options](./docs/images/select-secondary-option.png)

### After selecting the type, enter the details in the field

![secondary options](./docs/images/insert-details.png)

