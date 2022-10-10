package mssql

func Initialize(handler func(err error)) {
	if handler == nil {
		panic("Please provide an error handler")
	}

	errorHandlerCallback = handler
}
