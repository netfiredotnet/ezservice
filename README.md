# ezservice Library

## Overview

The `ezservice` package simplifies the creation and management of system services in Go. It builds on top of the `github.com/kardianos/service` package to abstract away the boilerplate code needed to install, start, stop, and manage services.

## Features

- Simplified service lifecycle management.
- Easy installation and uninstallation of services.
- Supports cross-platform service management (Windows, Linux, macOS).
- Built-in logging capabilities.

## Installation

To use `ezservice` in your project, add it to your Go module:

```sh
go get github.com/netfiredotnet/ezservice
```

Ensure you also have the `kardianos/service` package:

```sh
go get github.com/kardianos/service
```

## Basic Usage

Here’s a quick example of how to use `ezservice` to create and manage a service:

```go
package main

import (
	"fmt"
	"time"

	"github.com/kardianos/service"
	"github.com/netfiredotnet/ezservice"
)

type Service struct{}

func main() {
	var err error
	// Service handling boilerplate is abstracted away in the ezservice package
	srv := &Service{}
	if err = ezservice.New("mysvc", "My Service", "This is a demo of the ezservice package", "v1.0", srv); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}

// This needs to return as quickly as possible, else service manager will block during this time. So the rest of the code is run within a goroutine.
func (r *Service) Start(s service.Service) (err error) {
	var logger service.Logger

	if logger, err = s.Logger(nil); err != nil {
		return
	}
	go longRunningTask(logger)
	return
}

// This needs to return as quickly as possible, else service manager will block during this time. So the rest of the code is run within a goroutine.
func (r *Service) Stop(s service.Service) (err error) {
	var logger service.Logger

	if logger, err = s.Logger(nil); err != nil {
		return
	}
	logger.Info("Shutting down service")
	return
}

func longRunningTask(logger service.Logger) {
	for {
		logger.Info(fmt.Sprintf("Service is running at %v", time.Now()))
		time.Sleep(5 * time.Second)
	}
}

```

### Service Initialization

When you call `ezservice.New`, it sets up the service with the provided name, display name, description, version, and a service instance. This handles registration with the service manager for the current operating system.

### Service Lifecycle Methods

- **Start**: This method is called when the service starts. It should return quickly, so any long-running tasks should be run in a separate goroutine.
- **Stop**: This method is called when the service stops. Use this to clean up any resources or gracefully shut down tasks.

## Supported Commands

When you run your binary, you can control the service using the following commands:

- **install-service**: Installs the service on the system.
- **uninstall-service**: Uninstalls the service from the system.
- **run**: Runs the service directly without installing it.
- **start**: Starts the installed service.
- **stop**: Stops the installed service.
- **status**: Displays the current status of the service.

### Example Usage

To install the service:

```sh
./your_binary install-service
```

To start the installed service:

```sh
./your_binary start
```

To stop the running service:

```sh
./your_binary stop
```

To uninstall the service:

```sh
./your_binary uninstall-service
```

To check the service status:

```sh
./your_binary status
```

To run the service directly without installing it:

```sh
./your_binary run
```

## Logging

The `ezservice` package integrates with the `kardianos/service` logging system. You can obtain a logger from the service instance and use it to log messages throughout your service lifecycle.

## Contributing

We welcome contributions to `ezservice`. Feel free to open issues or submit pull requests on the [GitHub repository](https://github.com/netfiredotnet/ezservice).

## License

`ezservice` is licensed under the MIT License. See the [LICENSE](https://github.com/netfiredotnet/ezservice/blob/main/LICENSE) file for details.

## Contact

For questions or support, please reach out to the maintainers via [GitHub issues](https://github.com/netfiredotnet/ezservice/issues).

Happy coding with `ezservice`!
