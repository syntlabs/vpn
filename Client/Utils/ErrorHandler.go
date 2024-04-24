package main

import (
	"fmt"
	"os"
	"strings"
)

var netErr = networkErrors{
	Http404:              []string{"Http404", "Not found"},
	Http403:              []string{"Http403", "Bad db"},
	Http500:              []string{"Http500", "Server error"},
	Http200:              []string{"Http200", "Ok"},
	BrokenResponse:       []string{"BrokenResponseError", "Error occured when tried to get response"},
	BrokenRequest:        []string{"BrokenRequestError", "Validate request params"},
	NoInternet: []string{"NoInternetError", "Please make sure thar you are connected to the internet"},
	JsonConversion:  []string{"JsonConversionError", "Failed to convert response body to []byte"},
}

var valErr = valueErrors{
	Match:       []string{"MatchError", "Values do not match"},
	OutOfBounds: []string{"OutOfBoundsError", "Value is out of bounds in given range"},
	TimePointer: []string{"TimePointerError", "Subscription Time value type should be equal to predefined type"},
	WrongType:   []string{"WrongTypeError", "Value type is wrong"},
	WrongLanguage: []string{"WrongLanguageError", "Language does not match required"}
}

var logErr = logicErrors{
	Logic:     []string{"LogicError", "Some logic error happened here"},
	Salt:    []string{"SaltError", "Salts sizeof is bigger than max uin8(255) or salt is nil"},
	CantOpenFile: []string{"CantOpenFileError", "System can not open file because its either not exist or broken"},
	FileFormat: []string{"FileFormatError", "File format does not match required"},
}

var cyphErr = cypherErrors{
	Block:  []string{"BlockError", "Error occured when tries to creat block cypher"},
	Sizeof: []string{"SizeofError", "Make sure bytes sizeof does not exceed max unit8"},
	Decode: []string{"DecodeError", "Error occured when tries to decode cypher"},
	Method: []string{"MethodError", "Method cypher error"},
	Key: []string{"KeyError", "No public cert key was found"},
	Cert: []string{"CertError", "No cert was found"},
	PairCert: []string{"PairCertError", "Error occured when tried to creade cert-pubKey pair"},
}

var Err = errHandler{
	net: netErr, val: valErr, logic: logErr,
	cyph: cyphErr,
}

func annihilate() {
	os.Exit(012)
}

func raise(raisingError []string) {
	fmt.Print(strings.Join(raisingError, " "))
}