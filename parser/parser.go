package parser

import (
	"strings"
	"fmt"
	"strconv"
	"regexp"
)

// FIXME: refactor shared code between each Parse[..]Entitlement functions
// FIXME: create error objects

// Matches alphanum string containing alphanum substrings separated by single dashes
var isAlphanumOrDash = regexp.MustCompile(`^[a-zA-Z0-9]+(\-[a-zA-Z0-9]+)*$`).MatchString

func IsValidDomainName(domain string) bool {
	return isAlphanumOrDash(domain)
}

func IsValidDomainNameList(domain []string) bool {
	for _, domainField := range domain {
		if IsValidDomainName(domainField) == false {
			return false
		}
	}

	return true
}

func IsValidIdentifier(identifier string) bool {
	return isAlphanumOrDash(identifier)
}

// ParseVoidEntitlement() parses an entitlement with the following format: "domain-name.identifier"
func ParseVoidEntitlement(entitlementFormat string) (domain []string, id string, err error) {
	stringList := strings.Split(entitlementFormat, ".")
	if len(stringList) < 2 {
		return nil, "", fmt.Errorf("Parsing of entitlement %s failed: either domain or id missing")
	}

	id = stringList[len(stringList) - 1]
	domain = stringList[0:len(stringList) - 1]

	if IsValidDomainNameList(domain) == false {
		return nil, "", fmt.Errorf("Parsing of entitlement %s failed: domain must be alphanumeric and can contain '-'. '.' is a domain separator")
	}

	if IsValidIdentifier(id) == false {
		return nil, "", fmt.Errorf("Parsing of entitlement %s failed: identifier must be alphanumeric and can contain '-'")
	}

	return
}

// ParseIntEntitlement() parses an entitlement with the following format: "domain-name.identifier=int64-value"
func ParseIntEntitlement(entitlementFormat string) (domain []string, id string, value int, err error) {
	stringList := strings.Split(entitlementFormat, ".")
	if len(stringList) < 2 {
		return nil, "", 0, fmt.Errorf("Parsing of int entitlement %s failed: either domain or id missing")
	}

	idAndArgString := stringList[len(stringList) - 1]
	domain = stringList[0:len(stringList) - 2]

	if IsValidDomainNameList(domain) == false {
		return nil, "", 0, fmt.Errorf("Parsing of int entitlement %s failed: domain must be alphanumeric and can contain '-'. '.' is a domain separator")
	}

	idAndArgList := strings.Split(idAndArgString, "=")
	if len(idAndArgList) != 2 {
		return nil, "", 0, fmt.Errorf("Parsing of int entitlement %s failed: format required 'domain-name.identifier=int-value'")
	}

	id = idAndArgList[0]
	valueString := idAndArgList[1]

	if IsValidIdentifier(id) == false {
		return nil, "", 0, fmt.Errorf("Parsing of int entitlement %s failed: identifier must be alphanumeric and can contain '-'")
	}

	value, err = strconv.Atoi(valueString)
	if err != nil {
		return nil, "", 0, fmt.Errorf("Parsing of int entitlement %s failed: entitlement argument must be a 64bits integer")
	}

	return
}

// ParseStringEntitlement() parses an entitlement with the following format: "domain-name.identifier=string-value"
func ParseStringEntitlement(entitlementFormat string) (domain []string, id, value string, err error) {
	stringList := strings.Split(entitlementFormat, ".")
	if len(stringList) < 2 {
		return nil, "", "", fmt.Errorf("Parsing of string entitlement %s failed: either domain or id missing")
	}

	idAndArgString := stringList[len(stringList) - 1]
	domain = stringList[0:len(stringList) - 2]

	if IsValidDomainNameList(domain) == false {
		return nil, "", "", fmt.Errorf("Parsing of string entitlement %s failed: domain must be alphanumeric and can contain '-'. '.' is a domain separator")
	}

	idAndArgList := strings.Split(idAndArgString, "=")
	if len(idAndArgList) != 2 {
		return nil, "", "", fmt.Errorf("Parsing of string entitlement %s failed: format required 'domain-name.identifier=param'")
	}

	id = idAndArgList[0]
	value = idAndArgList[1]

	if IsValidIdentifier(id) == false {
		return nil, "", "", fmt.Errorf("Parsing of string entitlement %s failed: identifier must be alphanumeric and can contain '-'")
	}

	// FIXME: should we add constraints on the allowed characters in entitlement parameters and check integrity?

	return
}