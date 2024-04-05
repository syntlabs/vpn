package Antares_vpn

type networkErrors struct {
	http404,
	http403,
	http500,
	http200,
	brokenResponse,
	noInternetConnection,
	jsonConversionError,
	brokenRequest []string
}

type valueErrors struct {
	match,
	outofbounds,
	wrongType,
	timepointer,
	wrongLanguage []string
}

type logicErrors struct {
	logicErr,
	saltError,
	cantOpenFile []string
}

type cypherErrors struct {
	cypherBlockError,
	cypherMethodError,
	cypherDecodeError,
	cypherSizeofError []string
}

type errHandler struct {
	network networkErrors
	value   valueErrors
	logic   logicErrors
	cypher  cypherErrors
}
