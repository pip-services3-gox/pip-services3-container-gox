# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> IoC container for Golang

This module is a part of the [Pip.Services](http://pipservices.org) polyglot microservices toolkit. It provides an inversion-of-control (IoC) container to facilitate the development of services and applications composed of loosely coupled components.

The module containes a basic in-memory container that can be embedded inside a service or application, or can be run by itself.
The second container type can run as a system level process and can be configured via command line arguments.
Also it can be used to create docker containers.

The containers can read configuration from JSON or YAML files use it as a recipe for instantiating and configuring components.
Component factories are used to create components based on their locators (descriptor) defined in the container configuration.
The factories shall be registered in containers or dynamically in the container configuration file.

The module contains the following packages:

- [**Container**](https://godoc.org/github.com/pip-services3-gox/pip-services3-container-gox/container) - Component container and container as a system process
- [**Build**](https://godoc.org/github.com/pip-services3-gox/pip-services3-container-gox/build) - Container default factory
- [**Config**](https://godoc.org/github.com/pip-services3-gox/pip-services3-container-gox/config) - Container configuration
- [**Refer**](https://godoc.org/github.com/pip-services3-gox/pip-services3-container-gox/refer) - Container references

<a name="links"></a> Quick links:

* [Configuration](https://www.pipservices.org/recipies/configuration)
* [API Reference](https://godoc.org/github.com/pip-services3-gox/pip-services3-container-gox/)
* [Change Log](CHANGELOG.md)
* [Get Help](https://www.pipservices.org/community/help)
* [Contribute](https://www.pipservices.org/community/contribute)

## Use

Get the package from the Github repository:
```bash
go get -u github.com/pip-services3-gox/pip-services3-container-gox@latest
```

## Develop

For development you shall install the following prerequisites:
* Golang v1.12+
* Visual Studio Code or another IDE of your choice
* Docker
* Git

Run automated tests:
```bash
go test -v ./test/...
```

Generate API documentation:
```bash
./docgen.ps1
```

Before committing changes run dockerized test as:
```bash
./test.ps1
./clear.ps1
```

## Contacts

The Golang version of Pip.Services is created and maintained by:
- **Volodymyr Tkachenko**
- **Sergey Seroukhov**
- **Mark Zontak**

The documentation is written by:
- **Levichev Dmitry**