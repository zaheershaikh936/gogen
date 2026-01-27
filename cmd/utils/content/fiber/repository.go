package content

import (
	"fmt"

	"github.com/zaheershaikh936/gogen/cmd/utils/text"
)

const repository_content = `package repositories

import (
	"github.com/gofiber/fiber/v2"
	"%[1]s/%[2]s/models"
)
func GetAll%[3]s(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func Get%[3]s(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func Create%[3]s(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func Update%[3]s(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func Delete%[3]s(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
`

func RepositoryContent(model string, output string, moduleName string) string {
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
	content := fmt.Sprintf(repository_content, importPath, model, text.ToPascalCase(model))
	return text.PrependHeader(content)
}
