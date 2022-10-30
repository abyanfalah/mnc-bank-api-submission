# MNC Bank Payment API

This is my submission project of MNC Bank Backend Developer test.

This project uses Gin as web framework.


## Installation
You can directly download the build [here](https://github.com/abyanfalah/mnc-bank-api-submission/releases/tag/bin)
and run the app by executing it:
```bash
./mnc-bank-api
```

### Running the source code

If you downloaded the source code, you can go to the source code root directory, and run the app with:

```bash
go run .
```

### Building and running the app
or build the source code with:
```bash
go build
```
and run the app by executing it:
```bash
./mnc-bank-api
```
## Configuration
This app requires zero config.

This app also requires no RDBMS. It uses json files as database.

And the migration will create the required directory and files for the database once you run the app.

You can just execute or run the app and it will tell you which port this app is listening to.

The starting point is `localhost:8000`
and will iterate by itself if current port is used until unused one found.

Make sure to match the `base_url` in environment variables of your API client to the port where the API listens to.

## Usage
You can get a brief guide and simple usage explanation by importing [this file](https://github.com/abyanfalah/mnc-bank-api-submission/blob/main/Request%20collection.postman_collection.json) to Postman API client.

## Testing
To run the test of this app, go to test directory and execute the following command:
```bash
go test ./* -v
```