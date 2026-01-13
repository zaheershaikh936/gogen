package helper

import (
	"github.com/zaheershaikh936/gogen/cmd/utils/text"
)

func Pluralize(word string) string {
    return text.Pluralize(word)
}

func ToSnakeCase(s string) string {
    return text.ToSnakeCase(s)
}

func Capitalize(s string) string {
    return text.Capitalize(s)
}