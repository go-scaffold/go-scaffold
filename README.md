# go-scaffold

[![Go Report Card](https://goreportcard.com/badge/github.com/pasdam/go-scaffold)](https://goreportcard.com/report/github.com/pasdam/go-scaffold)
[![CI Status](https://github.com/pasdam/go-scaffold/workflows/Continuous%20integration/badge.svg)](https://github.com/pasdam/go-scaffold/actions)

Command line application that generates files/projects from a template.

The app is heavily inspired by [Helm](https://helm.sh/), but with the intent to
be general purpose.

You can look at the [examples](./examples) folder to checkout some examples.

**Note**: if you are looking for the legacy version of the application, please
check the `legacy` branch.

## Build and install

To build the executable:

```sh
make go-build
```

To build the docker image:

```sh
make docker-build
```

To build both:

```sh
make build
```

To build and install it:

```sh
make install
```

## Usage

### Template generation

In order to generate the output files from a template:

```sh
go-scaffold generate [<flags>] <template_oath> <output_dir>
```

i.e.:

```sh
go-scaffold generate -f ./examples/hello-world-markdown/values-project1.yaml ./examples/hello-world-markdown build/
```
