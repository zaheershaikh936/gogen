package content

import (
	"fmt"

	"github.com/zaheershaikh936/gogen/cmd/utils/text"
)

const controller_content = `package controllers

import (
	"github.com/gofiber/fiber/v2"
	"%[1]s/%[2]s/services"
)

func GetAll%[3]s(c *fiber.Ctx) error {
	return services.GetAll%[3]s(c)
}

func Get%[3]s(c *fiber.Ctx) error {
	return services.Get%[3]s(c)
}

func Create%[3]s(c *fiber.Ctx) error {
	return services.Create%[3]s(c)
}

func Update%[3]s(c *fiber.Ctx) error {
	return services.Update%[3]s(c)
}

func Delete%[3]s(c *fiber.Ctx) error {
	return services.Delete%[3]s(c)
}
`

func ControllerContent(model string, output string, moduleName string) string {
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
	content := fmt.Sprintf(controller_content, importPath, model, text.ToPascalCase(model))
	return text.PrependHeader(content)
}
