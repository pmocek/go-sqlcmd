package config

func GetContainerId() (containerId string) {
	currentContextName := config.CurrentContext

	if currentContextName == "" {
		panic("currentContextName must not be empty")
	}

	for _, c := range config.Contexts {
		if c.Name == currentContextName {
			for _, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					if e.ContainerDetails == nil {
						panic("Endpoint does not have a container")
					}
					containerId = e.ContainerDetails.Id

					if len(containerId) != 64 {
						panic("container id must be 64 characters")
					}

					return
				}
			}
		}
	}
	panic("Id not found")
}

func CurrentContextEndpointHasContainer() (exists bool) {
	currentContextName := config.CurrentContext

	if currentContextName == "" {
		panic("currentContextName must not be empty")
	}

	for _, c := range config.Contexts {
		if c.Name == currentContextName {
			for _, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					if e.AssetDetails != nil {
						if e.AssetDetails.ContainerDetails != nil {
							exists = true
						}
					}
					break
				}
			}
		}
	}
	return
}

func FindFreePortForTds() (portNumber int) {
	const startingPortNumber = 1433

	portNumber = startingPortNumber

	for {
		foundFreePortNumber := true
		for _, endpoint := range config.Endpoints {
			if endpoint.Port == portNumber {
				foundFreePortNumber = false
				break
			}
		}

		if foundFreePortNumber == true {
			// Check this port is actually available on the local machine
			if isLocalPortAvailableCallback(portNumber) {
				break
			} else {
				foundFreePortNumber = false
			}
		}

		portNumber++

		if portNumber == 5000 {
			panic("Did not find an available port")
		}
	}

	return
}
