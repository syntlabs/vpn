package main

func validFileFormat(path string, format string) {
	if path[:len(path)-len(format)] != format {
		raise(Err.logic.FileFormat)
	}
}