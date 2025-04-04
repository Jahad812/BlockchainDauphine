package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/smartcontractkit/chainlink/core/scripts/keystone/src"
)

type command interface {
	Run([]string)
	Name() string
}

func main() {
	commands := []command{
		src.NewProvisionKeystoneCommand(),
		src.NewDeployAndInitializeCapabilitiesRegistryCommand(),
		src.NewToolkit(),
		src.NewGenerateLocalOCR3ConfigCommand(),
	}

	commandsList := func(commands []command) string {
		var scs []string
		for _, command := range commands {
			scs = append(scs, command.Name())
		}
		return strings.Join(scs, ", ")
	}(commands)

	if len(os.Args) >= 2 {
		requestedCommand := os.Args[1]

		for _, command := range commands {
			if command.Name() == requestedCommand {
				command.Run(os.Args[2:])
				return
			}
		}
		fmt.Println("Unknown command:", requestedCommand)
	} else {
		fmt.Println("No command specified")
	}

	fmt.Println("Supported commands:", commandsList)
	os.Exit(1)
}
