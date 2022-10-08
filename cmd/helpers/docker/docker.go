package docker

func Initialize(handler errorHandlerService) {
	if handler == nil {
		panic("Please provide an error handler")
	}

	errorHandlerCallback = handler
}
