# 🚀 Gogen - Go CRUD Generator CLI

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/zaheershaikh936/gogen)](https://github.com/zaheershaikh936/gogen/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaheershaikh936/gogen)](https://goreportcard.com/report/github.com/zaheershaikh936/gogen)

> A powerful CLI tool to generate complete CRUD resources for Go Fiber framework with a premium TUI experience 💎

Stop writing repetitive boilerplate code! Gogen generates production-ready CRUD APIs with models, controllers, services, repositories, and routes—all in one command.

# gogen

A command-line tool for generating CRUD resource scaffolding for the Go Fiber web framework. gogen automates the creation of controllers, services, repositories, and routes, following clean architecture principles.

## Features

- Generates complete CRUD resource structure for Go Fiber
- Produces controller, service, repository, and route files
- Supports custom output directory
- Minimal dependencies, fast execution
- Clean, idiomatic Go code structure

## ✨ Premium Experience

- 🪄 **Interactive Wizard** - Guided resource generation by just running `gogen resource`
- 💎 **Premium TUI** - Beautifully designed terminal interface with borders and boxed reports
- ⚡ **Real-time Feedback** - Animated spinners and smooth transitions
- 🎯 **One Command Generation** - Fast, scriptable resource creation

## Installation

### Using Go

```bash
go install github.com/zaheershaikh936/gogen@latest
```

Ensure that your `$GOPATH/bin` is in your `PATH`.

### Download Prebuilt Binary

1. Visit [Releases](https://github.com/zaheershaikh936/gogen/releases/latest)
2. Download the binary for your OS (Windows, Linux, macOS)
3. Add the binary location to your `PATH`

## CLI Usage

### Command Syntax

```bash
gogen resource <model> [flags]
```

### Arguments

- `<model>`: Required. The singular name of the resource to generate (e.g., `user`, `product`). The tool will pluralize as needed for folder and file names.

### Flags

| Flag         | Short | Default | Description                        |
|--------------|-------|---------|------------------------------------|
| `--output`   | `-o`  | `./`    | Output directory for generated files |
| `--help`     | `-h`  |         | Show help for the command          |

### Behavior Notes

- The tool creates a new folder for the resource under the specified output directory.
- Existing files with the same name will be overwritten.
- Only one resource can be generated per command invocation.

## Usage Examples

Generate a `user` resource in the current directory:

```bash
gogen resource user
```

Generate a `product` resource in a custom directory:

```bash
gogen resource product --output ./api
```

Use the short flag for output directory:

```bash
gogen resource order -o ./src
```

Show help for the resource command:

```bash
gogen resource --help
```

## End-to-End Example

Command:

```bash
gogen resource invoice --output ./internal/api
```

Generated folder structure:

```
internal/api/invoice/
├── controllers/
│   └── invoice_controller.go
├── services/
│   └── invoice_service.go
├── repositories/
│   └── invoice_repository.go
└── routes/
    └── invoice_routes.go
```

## Sample CLI Output

```text
  ╭─────────────────────────────────────────────╮  
  │                                             │
  │                  G O G E N                  │
  │     The professional Go/Fiber generator     │
  │                                             │
  ╰─────────────────────────────────────────────╯


 ⣻  Generating your resource...

  ╔══════════════════════════════════════════════╗
  ║                                              ║
  ║  ✓ Files generated successfully              ║
  ║                                              ║
  ║  ↳ invoice/routes/invoice_routes.go          ║
  ║  ↳ invoice/controllers/invoice_controller.go ║
  ║  ↳ invoice/services/invoice_service.go       ║
  ║  ↳ invoice/repositories/invoice_repository.go║
  ║                                              ║
  ╚══════════════════════════════════════════════╝

  ✨ Resource invoice is ready to use!
```

## Generated Architecture

gogen follows a clean, modular structure:

- `controllers/`: HTTP handlers for CRUD endpoints
- `services/`: Business logic layer
- `repositories/`: Data access layer
- `routes/`: Fiber route definitions

Each resource is self-contained under its own directory, supporting maintainable and testable code organization.

## Contributing

Contributions are welcome. To contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'feat: add new feature'`)
4. Push to your branch (`git push origin feature/your-feature`)
5. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

Zaheer Shaikh  
GitHub: [@zaheershaikh936](https://github.com/zaheershaikh936)
```

**Windows:**
Download `gogen_Windows_x86_64.zip` from releases and add to your PATH.

## 🚀 Quick Start

# Setup
### Check where Go installs binaries
```
# Add PATH 
export PATH=$PATH:$(go env GOPATH)/bin
```


### Generate a CRUD Resource
```bash
# Basic usage
gogen resource user

# Custom output directory
gogen resource product --output ./api

# Short flag
gogen resource order -o ./src
```

### What Gets Generated?
```
users/
├── controllers/
│   └── users_controller.go   # HTTP handlers (GetAll, Get, Create, Update, Delete)
├── services/
│   └── users_service.go      # Business logic layer
├── repositories/
│   └── users_repository.go   # Database operations
└── routes/
    └── users_routes.go       # Fiber route definitions
```

## 🏗️ Project Structure

Gogen follows clean architecture principles:
```
your-project/
├── controllers/   # HTTP handlers (presentation layer)
├── services/      # Business logic
├── repositories/  # Data access layer
└── routes/        # API route definitions
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🌟 Show Your Support

If you find this project useful, please give it a ⭐️ on GitHub!

## 📬 Contact

**Zaheer Shaikh** - [@zaheershaikh936](https://github.com/zaheershaikh936)

Project Link: [https://github.com/zaheershaikh936/gogen](https://github.com/zaheershaikh936/gogen)

---

**Made with ❤️ by Zaheer Shaikh**
