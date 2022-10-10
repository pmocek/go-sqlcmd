package config

func Initialize(handler func(err error), configFile string) {
	if handler == nil {
		panic("Please provide an error handler")
	}
	errorHandlerCallback = handler
	configureViper(configFile)
	load()
}
