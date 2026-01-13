# 🚀 Gogen - Go CRUD Generator CLI

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/zaheershaikh936/gogen)](https://github.com/zaheershaikh936/gogen/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaheershaikh936/gogen)](https://goreportcard.com/report/github.com/zaheershaikh936/gogen)

> A powerful CLI tool to generate complete CRUD resources for Go Fiber framework in seconds ⚡

Stop writing repetitive boilerplate code! Gogen generates production-ready CRUD APIs with models, controllers, services, repositories, and routes—all in one command.

## ✨ Features

- 🎯 **One Command Generation** - Generate complete CRUD structure instantly
- 📁 **Clean Architecture** - Follows best practices with separated concerns
- 🔥 **Fiber Framework** - Built specifically for Go Fiber
- ⚡ **Lightning Fast** - Generate resources in milliseconds
- 🛠️ **Customizable** - Flexible output directory configuration
- 📦 **Zero Dependencies** - Simple installation, no complex setup

## 🎬 Demo
```bash
$ gogen resource User

🚀 Generating CRUD resource...
Model: User

Files to be generated:
  ✓ ./users/controllers/user_controller.go
  ✓ ./users/services/user_service.go
  ✓ ./users/repositories/user_repository.go
  ✓ ./users/routes/user_routes.go

✅ Done!
```

## � Installation

### Using Go Install (Recommended)
```bash
go install github.com/zaheershaikh936/gogen@latest
```

### Download Binary
Download the latest release for your OS from [Releases](https://github.com/zaheershaikh936/gogen/releases/latest)

**Linux:**
```bash
wget https://github.com/zaheershaikh936/gogen/releases/latest/download/gogen_Linux_x86_64.tar.gz
# Extract and move to bin
```

**macOS:**
```bash
curl -L https://github.com/zaheershaikh936/gogen/releases/latest/download/gogen_Darwin_arm64.tar.gz -o gogen.tar.gz
# Extract and move to bin
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
