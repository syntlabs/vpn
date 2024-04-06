package main

type networkErrors struct {
	Http404,
	Http403,
	Http500,
	Http200,
	BrokenResponse,
	BrokenRequest,
	NoInternet,
	JsonConversion []string
}

type valueErrors struct {
	Match,
	OutOfBounds,
	TimePointer,
	WrongType,
	WrongLanguage []string
}

type logicErrors struct {
	Logic,
	Salt,
	CantOpenFile,
	FileFormat []string
}

type cypherErrors struct {
	Block,
	Sizeof,
	Decode,
	Method,
	Key,
	Cert,
	PairCert []string
}

type errHandler struct {
	net networkErrors
	val   valueErrors
	logic   logicErrors
	cyph  cypherErrors
}
