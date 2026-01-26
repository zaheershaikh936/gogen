package text

import "strings"

func Pluralize(word string) string {
	if word == "" {
		return ""
	}
	word = strings.ToLower(word)

	irregulars := map[string]string{
		"person": "people",
		"child":  "children",
		"man":    "men",
		"woman":  "women",
		"tooth":  "teeth",
		"foot":   "feet",
		"mouse":  "mice",
		"goose":  "geese",
	}

	if plural, ok := irregulars[word]; ok {
		return plural
	}

	if strings.HasSuffix(word, "y") {
		return strings.TrimSuffix(word, "y") + "ies"
	}

	if strings.HasSuffix(word, "ss") || strings.HasSuffix(word, "sh") ||
		strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "x") ||
		strings.HasSuffix(word, "z") {
		return word + "es"
	}

	if strings.HasSuffix(word, "s") {
		return word
	}

	if strings.HasSuffix(word, "f") {
		return word[:len(word)-1] + "ves"
	}
	if strings.HasSuffix(word, "fe") {
		return word[:len(word)-2] + "ves"
	}

	return word + "s"
}

func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}

func ToPascalCase(s string) string {
	if s == "" {
		return ""
	}
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func Capitalize(s string) string {
	return ToPascalCase(s)
}

func TrimTrailingSlash(s string) string {
	return strings.TrimSuffix(s, "/")
}
