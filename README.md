# Zendesk Search tool

## Setup requirements

Go 1.16 was the development setup for this application. You can install golang by following the steps [here](https://golang.org/doc/install)
If you are on mac, you can run `brew install go` to install latest version of golang

## Running the application locally

Run the following command to start the application in your local terminal.

```bash
make run
```

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

