package content

import (
	"fmt"

	"github.com/zaheershaikh936/gogen/cmd/utils/text"
)

const routes_content = `package routes
import (
	"github.com/gofiber/fiber/v2"
	"%[1]s/%[2]s/controllers"
)
func %[3]sRoutes(router fiber.Router) {
	%[4]sRouter := router.Group("/%[4]s")	
	%[4]sRouter.Get("/", controllers.GetAll%[3]s)
	%[4]sRouter.Get("/:id", controllers.Get%[3]s)
	%[4]sRouter.Post("/", controllers.Create%[3]s)
	%[4]sRouter.Put("/:id", controllers.Update%[3]s)
	%[4]sRouter.Delete("/:id", controllers.Delete%[3]s)
}
`

func RoutesContent(model string, output string, moduleName string) string {
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
	return fmt.Sprintf(routes_content,
		importPath,
		model,
		text.Capitalize(model),
		text.ToSnakeCase(model),
	)
}
