package main

import "runtime"

func validFileFormat(path string, format string) {
	if path[:len(path)-len(format)] != format {
		panic(runtime.PanicNilError{})
	}
}
