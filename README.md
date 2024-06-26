<h1 align="center">
  Golang Rest API Starter
</h1>

![CI](https://github.com/nealajpatel/golang-rest-api-starter/actions/workflows/code-ql.yaml/badge.svg?branch=main&event=push)
![CI](https://github.com/nealajpatel/golang-rest-api-starter/actions/workflows/pipeline.yaml/badge.svg?branch=main&event=push)


Boilerplate API built using Golang: (Gin Framework for routing, Testify for testing, Swagger for Documentation, Viper for configuration, and Go Mod for dependency management). The intention is for this to be used as a starter application to build out a scalable and organized API service. With minimal configuration, this application can be built and deployed in a container using docker. We have done some research and have tried to stick to Golang best practices as much as possible. The routes and testing can be expanded to meet your needs.

* [Gin Documentation](https://github.com/gin-gonic/gin)
* [Testify Documentation](https://github.com/stretchr/testify)
* [Swagger Documentation](https://swagger.io/docs/)
* [Viper Documentation](https://github.com/spf13/viper)
* [Go Mod Documentation](https://github.com/golang/go/wiki/Modules)

## 🚀 Quick start

1. **Copy directory into your group/namespace**

    This can be done by forking this repository into your namespace. From there you can remove the fork relationship and update the name. These can be done in the advanced settings

1. **Update naming.**

    Update the golang-rest-api-starter name throughout the project to your application name. This can be done with a simple find/replace

1. **Start coding!**

    You should now be able to start coding. You can start expanding, building, and deploying. You can use make to build/ run your app:

        make build
        make run

## 🧐 What's inside?

A quick look at the top-level files and directories you'll see in this project.

    .
    ├── build
    ├── config
    ├── controllers
    ├── docs
    ├── middlewares
    ├── server
    ├── tests
    ├── utils
    ├── Dockerfile
    ├── Makefile
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── .gitignore
    └── README.md

1. **`/build`**: This directory contains the binary after a build

1. **`/config`**: This directory contains the viper configration and the property files.

1. **`/controllers`**: This directory contains the controllers that the router/handlers are configured to call.

1. **`/docs`**: Directory for swagger files.

1. **`/middlewares`**: This directory is used to house the middlewares used by the controllers.

1. **`/server`**: This Directory houses the router and service files which configure the route and start the server

1. **`/tests`**: This directory contains all the tests for your application and get kicked off by:

        make test

1. **`/utils`**: This directory contains all the util classes used throughout the app.

1. **`Dockerfile`**: This file contains the buildsteps to build your image using a multi stage build. It is currently using a scratch image to keep it lightweight.

1. **`Makefile`**: This file allows the use of make to run tests, build locally, and is used to build in the pipeline. This can be expanded as needed

1. **`go.mod/go.sum`**: These files are generated by Go Mod dependency management and can be learned about in the documentation link provided above.

1. **`main.go`**: This file is the starting point for the application, which starts the server.

1. **`.gitignore`**: This file tells git which files it should not track / not maintain a version history for.

1. **`README.md`**: A text file containing useful reference information about your project.

## 🎓 Further Changes

There may be further configuration needed as you expand your application.

**Configmap template:** You can update the dev.yml file to include your properties.

### Contributors:

* Neal Patel
* Chris Elias
