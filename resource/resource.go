package resource

// A Resource is the base for all three action resources.
type Resource interface {
	// Source must return a pointer to a struct, where the "source" definition is stored
	Source() (source interface{})
}

// A CheckResource implements all required functions for the "check" action
type CheckResource interface {
	Resource
	// Version must return a pointer to a struct, where the "version" definition is stored
	Version() (version interface{})
	// Check for new resource versions and return them
	Check() (version interface{}, err error)
}

// A InResource implements all required functions for the "in" action
type InResource interface {
	Resource
	// Version must return a pointer to a struct, where the "version" definition is stored
	Version() (version interface{})
	// Params must return a pointer to a struct, where the "params" definition is stored
	Params() (params interface{})
	// Fetch the resource and return the version and metadata
	In(dir string) (version interface{}, metadata []interface{}, err error)
}

// A OutResource implements all required functions for the "out" action
type OutResource interface {
	Resource
	// Params must return a pointer to a struct, where the "params" definition is stored
	Params() (params interface{})
	// Put the resource and return the new version and metadata
	Out(dir string) (version interface{}, metadata []interface{}, err error)
}
