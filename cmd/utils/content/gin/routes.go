package content

import (
	"fmt"

	"github.com/zaheershaikh936/gogen/cmd/utils/text"
)

const routes_content = `package routes

import (
	"github.com/gin-gonic/gin"
	"%[1]s/%[2]s/controllers"
)

func %[3]sRoutes(router *gin.RouterGroup) {
	%[4]sRouter := router.Group("/%[5]s")
	{
		%[4]sRouter.GET("/", controllers.GetAll%[3]s)
		%[4]sRouter.GET("/:id", controllers.Get%[3]s)
		%[4]sRouter.POST("/", controllers.Create%[3]s)
		%[4]sRouter.PUT("/:id", controllers.Update%[3]s)
		%[4]sRouter.DELETE("/:id", controllers.Delete%[3]s)
	}
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
		text.ToPascalCase(model),
		text.ToCamelCase(model),
		text.ToSnakeCase(model), // For URL path
	)
}
