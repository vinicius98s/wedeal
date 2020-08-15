# Wedeal

### Pre requisites

You must have docker-compose installed in your machine.

### Installing dependencies:

If you are on Windows, you can try installing [make](https://pt.wikipedia.org/wiki/Make) with [chocolatey](https://chocolatey.org/install) and then running:

```sh
choco install make
make install
```

If you are on a Unix based OS it's probably already installed.

### Getting started

You can run `make dev` to get the project running.

```sh
make dev
```

Once you've started the server you should see a message like `Running on port 8080`, then if you hit the [http://localhost:8080/graphql](http://localhost:8080/graphql) url you should be able to see the Graphql Playground.
