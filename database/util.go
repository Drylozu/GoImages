package database

import (
	"math/rand"
	"reflect"
	"strings"
)

var invisibleCharacters = []rune{
	'\u2060',
	'\u2061',
	'\u2062',
	'\u2063',
	'\u2064',
	'\u2066',
	'\u2067',
	'\u2068',
	'\u2069',
	'\u206A',
	'\u206B',
	'\u206C',
	'\u206D',
	'\u206E',
	'\u206F',
	'\u200B',
	'\u200C',
	'\u200D',
	'\u200E',
	'\u200F',
	'\u061C',
	'\uFEFF',
	'\u180E',
	'\u00AD',
}

var characters = len(invisibleCharacters)

func GenerateName(l ...interface{}) string {
	length := 12
	if len(l) > 0 && reflect.TypeOf(l[0]).String() == "int" {
		length = l[0].(int)
	}

	var name strings.Builder

	for i := 0; i < length; i++ {
		name.WriteRune(invisibleCharacters[rand.Intn(characters)])
	}
	return name.String()
}
