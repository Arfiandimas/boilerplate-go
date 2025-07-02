package library

type libraryExample struct {
}

// NewLibraryExample initializer for LibraryExample
func NewLibraryExample() LibraryExample {
	return &libraryExample{}
}

// LibExample implements LibraryExample
func (*libraryExample) LibExample() uint64 {
	panic("unimplemented")
}
