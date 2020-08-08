# Wedeal

### Pre requisites

You must have at least Go 1.11 installed

### Installing dependencies:

If you are on Windows, you can try installing [make](https://pt.wikipedia.org/wiki/Make) with [chocolatey](https://chocolatey.org/install) and then running:

```sh
choco install make
make install
```

If you are on a Unix based os, then you can just:

```sh
make install
```

### Getting started

You can run `make dev` to get the project running.

```sh
make dev
```

Once you've started the server you should see a message like `Running on port 8080`, then if you hit the [http://localhost:8080/graphql](http://localhost:8080/graphql) url, you should be able to see the Graphql Playground
