package ezservice

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/kardianos/service"
)

var (
	parser *flags.Parser
)

func getServiceStatus(svc service.Service) (err error) {
	var stat service.Status
	var statStr string
	var supportedSystems []service.System
	// statStr = fmt.Sprintf("%v is ", SvcDisplayName)
	statStr = fmt.Sprintf("%v is ", "service")
	if stat, err = svc.Status(); err != nil {
		return
	}
	switch s := stat; s {
	case service.StatusUnknown:
		statStr += "UNKNOWN. (Is it running as a service?)\n"
	case service.StatusRunning:
		statStr += "RUNNING.\n"
	case service.StatusStopped:
		statStr += "STOPPED.\n"
	}
	fmt.Printf("%v\nSYSTEM SERVICE INFORMATION:\n", statStr)

	// Obtain available systems
	supportedSystems = service.AvailableSystems()

	for idx, s := range supportedSystems {
		fmt.Printf("%v:\n\tService system: %v\n", idx, s.String())
		fmt.Printf("\tAvailable: %v\n", s.Detect())
		fmt.Printf("\tInteractive: %v\n", s.Interactive())
	}
	return
}

func New(name string, displayName string, description string, version string, svcInterface service.Interface) (err error) {
	var a args
	var svc service.Service
	var logger service.Logger

	conf := &service.Config{
		Name:        name,
		DisplayName: displayName,
		Description: description,
	}

	if svc, err = service.New(svcInterface, conf); err != nil {
		return
	}

	if logger, err = svc.Logger(nil); err != nil {
		return
	}

	if !service.Interactive() {
		// We are being called by the service manager - start the service

		if err = svc.Run(); err != nil {
			logger.Error(err)
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	} else {
		// Parse args
		parser = flags.NewParser(&a, flags.Default)

		// Handle missing flags or other errors
		if _, err := parser.Parse(); err != nil {
			switch flagsErr := err.(type) {
			case flags.ErrorType:
				if flagsErr == flags.ErrHelp {
					os.Exit(0)
				}
				os.Exit(1)
			default:
				os.Exit(1)
			}
		}

		switch parser.Active.Name {
		case "run":
			// Run the service in the foreground
			svcInterface.Start(svc)
		case "start":
			if err = svc.Start(); err != nil {
				logger.Error(err)
			}
			logger.Infof("service started\n")
		case "install-service":
			if err = svc.Install(); err != nil {
				logger.Error(err)
			}
			logger.Infof("service installed\n")
		case "uninstall-service":
			if err = svc.Uninstall(); err != nil {
				logger.Error(err)
			}
			logger.Infof("service uninstalled\n")
		case "stop":
			if err = svc.Stop(); err != nil {
				logger.Error(err)
			}
			logger.Infof("service stopped\n")
		case "status":
			if err = getServiceStatus(svc); err != nil {
				logger.Error(err)
			}
		default:
			fmt.Println("Command not supported")
		}
	}

	return
}
