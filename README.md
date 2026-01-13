# Gogen

Gogen is a powerful CLI tool designed to rapid-prototype Go web applications using the [Fiber](https://gofiber.io/) framework. It automates the generation of standard CRUD boilerplate, following a clean architecture (Controllers, Services, Repositories, Routes).

## Features

- **CRUD Resource Generation**: Generates complete vertical slices (Route -> Controller -> Service -> Repository) in one command.
- **Smart Pluralization**: Automatically handles resource naming. Inputting `user` or `users` correctly generates a `users` resource. Handles irregular plurals (e.g., `category` -> `categories`, `box` -> `boxes`).
- **Dynamic Module Detection**: Automatically reads your `go.mod` file to generate correct internal import paths. No hardcoded module names!
- **Fiber-Ready**: Generates code pre-configured for use with the Gofiber v2 framework.

## 🚀 Installation

### Option 1: Go Install (Recommended)
If you have Go installed on your system, this is the quickest way to get started:
```bash
go install github.com/zaheershaikh936/gogen@latest
```
*Note: Ensure your `$GOPATH/bin` is in your system `PATH`.*

### Option 2: Pre-compiled Binaries
Download the latest binary for your OS from the [Releases](https://github.com/zaheershaikh936/gogen/releases) page.

**macOS/Linux:**
```bash
chmod +x gogen
sudo mv gogen /usr/local/bin/
```

**Windows:**
Download the `.zip` file, extract `gogen.exe`, and add it to your system Environment Variables.

---

## ⚡ Quick Start

1. **Initialize a new Fiber project:**
   ```bash
   mkdir my-api && cd my-api
   go mod init my-api
   go get github.com/gofiber/fiber/v2
   ```

2. **Generate your first resource:**
   ```bash
   gogen resource user --output ./internal/api
   ```

---

## 🛠 Usage

The `resource` command generates a complete vertical slice (Route -> Controller -> Service -> Repository).

```bash
gogen resource [model_name] --output [path]
```

### Examples

**Standard Resource:**
```bash
gogen resource product --output ./src/api
```

**Irregular Pluralization:**
Gogen handles complex naming for you. For example, `category` becomes `categories`:
```bash
gogen resource category --output ./src/api
```

### Generated Structure
```text
src/api/
└── products/
    ├── controllers/
    │   └── products_controller.go
    ├── repositories/
    │   └── products_repository.go
    ├── routes/
    │   └── products_routes.go
    └── services/
        └── products_service.go
```

## 🏗 Generated Code Layers

- **Routes**: Fiber router groups mapping HTTP methods to controller functions.
- **Controllers**: Request parsing and input validation.
- **Services**: Business logic and orchestration.
- **Repositories**: Data access layer (with Fiber context placeholders).

## 📄 License

[MIT](LICENSE)