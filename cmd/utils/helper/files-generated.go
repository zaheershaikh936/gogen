package helper

import (
	"os"

	fiberContent "github.com/zaheershaikh936/gogen/cmd/utils/content/fiber"
	ginContent "github.com/zaheershaikh936/gogen/cmd/utils/content/gin"
)

func GetFilePath(model string, output string, file_name string) string {
    filepath := "./" + output + "/" + model + file_name
    return filepath
}

func CreateFile(model string, output string, dir string, file_name string, content string) []string {
    dirPath := "./" + output + "/" + model + "/" + dir
    err := os.MkdirAll(dirPath, 0755); if err != nil { return []string{} }
    
    file_path := dirPath + "/" + model + file_name
    err = os.WriteFile(file_path, []byte(content), 0644); if err != nil { return []string{} }
    return []string{file_path}
}

func CreateModel(model string, output string) []string {
    err := os.MkdirAll("./" + output + "/" + model, 0755); if err != nil { return []string{} }
    return []string{"./" + output + "/" + model}
}

func RoutesGenerated(model string, output string, framework string) []string {
    moduleName := GetModuleName()
    var routes_content string
    if framework == "gin" {
        routes_content = ginContent.RoutesContent(model, output, moduleName)
    } else {
        routes_content = fiberContent.RoutesContent(model, output, moduleName)
    }
    return CreateFile(model, output, "routes", "_routes.go", routes_content)
} 

func ControllerGenerated(model string, output string, framework string) []string {
    moduleName := GetModuleName()
    var controller_content string
    if framework == "gin" {
        controller_content = ginContent.ControllerContent(model, output, moduleName)
    } else {
        controller_content = fiberContent.ControllerContent(model, output, moduleName)
    }
    return CreateFile(model, output, "controllers", "_controller.go", controller_content)
}

func ServiceGenerated(model string, output string, framework string) []string {
    moduleName := GetModuleName()
    var service_content string
    if framework == "gin" {
        service_content = ginContent.ServiceContent(model, output, moduleName)
    } else {
        service_content = fiberContent.ServiceContent(model, output, moduleName)
    }
    return CreateFile(model, output, "services", "_service.go", service_content)
}

func RepositoryGenerated(model string, output string, framework string) []string {
    moduleName := GetModuleName()
    var repository_content string
    if framework == "gin" {
        repository_content = ginContent.RepositoryContent(model, output, moduleName)
    } else {
        repository_content = fiberContent.RepositoryContent(model, output, moduleName)
    }
    return CreateFile(model, output, "repositories", "_repository.go", repository_content)
}