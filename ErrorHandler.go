package Antares_vpn

import (
	"fmt"
	"os"
	"strings"
)

var netErr = networkErrors{
	http404:              []string{"404", "not found"},
	http403:              []string{"403", "bad gateway"},
	http500:              []string{"500", "server error"},
	http200:              []string{"200", "ok"},
	brokenResponse:       []string{"brokenResponseError", "Error occured when tried to get response"},
	brokenRequest:        []string{"brokenRequestError", "Validate request params"},
	noInternetConnection: []string{"NoInternetError", "Please make sure thar you are connected to the internet"},
	jsonConversionError:  []string{"JsonConversionError", "Failed to convert response body to []byte"},
}

var valErr = valueErrors{
	match:       []string{"MatchError", "Values do not match"},
	outofbounds: []string{"OutOfBoundsError", "Value is out of bounds in given range"},
	timepointer: []string{"TimePointerError", "Subscription Time value type should be equal to predefined type"},
	wrongType:   []string{"WrongTypeError", "Value type is wrong"},
}

var logErr = logicErrors{
	logicErr:     []string{"LogicError", "Some logic error happened here"},
	saltError:    []string{"SaltError", "Salts sizeof is bigger than max uin8(255) or salt is nil"},
	cantOpenFile: []string{"CantOpenFileError", "System can not open file because its either not exist or broken"},
}

var cyphErr = cypherErrors{
	cypherBlockError:  []string{"CypherBlockError", "Error occured when tries to creat block cypher"},
	cypherSizeofError: []string{"CypherSizeofError", "Make sure bytes sizeof does not exceed max unit8"},
	cypherDecodeError: []string{"CypherDecodeError", "Error occured when tries to decode cypher"},
	cypherMethodError: []string{"CypherMethodError", "Bad cypher error"},
}

var ErrorsGlobal = errHandler{
	network: netErr, value: valErr, logic: logErr,
	cypher: cyphErr,
}

func annihilate() {
	os.Exit(012)
}

func raise(raisingError []string) {

	fmt.Print(strings.Join(raisingError, " "))
}
