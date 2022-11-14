package commander


//Per golang design doc "an unfortunately necessary kludge":
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#pointer-method-example
// https://www.reddit.com/r/golang/comments/uqwh5d/generics_new_value_from_pointer_type_with/
type PtrIsInterface[T any] interface {
	Commander
	*T
}

func NewCommand[T PtrIsInterface[Value], Value any]() T {
	var commander T = new(Value)
	commander.DefineCommand()
	return commander
}
