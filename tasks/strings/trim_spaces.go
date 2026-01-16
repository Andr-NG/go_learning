// strings.Join(), strings.Fields()


package main

import "strings"

func TrimSpaces(s string) (any) {

	tempVar := strings.Fields(s)

	return strings.Join(tempVar, " ")

}
