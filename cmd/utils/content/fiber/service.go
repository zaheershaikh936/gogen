package content

import (
	"fmt"

	"github.com/zaheershaikh936/gogen/cmd/utils/text"
)

const service_content = `package services

import (
	"github.com/gofiber/fiber/v2"
	"%[1]s/%[2]s/repositories"
)
func GetAll%[3]s(c *fiber.Ctx) error {
	return repositories.GetAll%[3]s(c)
}

func Get%[3]s(c *fiber.Ctx) error {
	return repositories.Get%[3]s(c)
}

func Create%[3]s(c *fiber.Ctx) error {
	return repositories.Create%[3]s(c)
}

func Update%[3]s(c *fiber.Ctx) error {
	return repositories.Update%[3]s(c)
}

func Delete%[3]s(c *fiber.Ctx) error {
	return repositories.Delete%[3]s(c)
}
`

func ServiceContent(model string, output string, moduleName string) string {
	importPath := output
	// Trim trailing slashes to avoid double slashes in imports
	importPath = text.TrimTrailingSlash(importPath)
	if importPath == "./" || importPath == "." || importPath == "" {
		importPath = moduleName
	} else {
		// Removing leading ./ if present for imports
		if len(importPath) >= 2 && importPath[0:2] == "./" {
			importPath = moduleName + "/" + importPath[2:]
		} else {
			importPath = moduleName + "/" + importPath
		}
	}
	content := fmt.Sprintf(service_content, importPath, model, text.ToPascalCase(model))
	return text.PrependHeader(content)
}
