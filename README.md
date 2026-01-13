# Gogen

Gogen is a powerful CLI tool designed to rapid-prototype Go web applications using the [Fiber](https://gofiber.io/) framework. It automates the generation of standard CRUD boilerplate, following a clean architecture (Controllers, Services, Repositories, Routes).

## Features

- **CRUD Resource Generation**: Generates complete vertical slices (Route -> Controller -> Service -> Repository) in one command.
- **Smart Pluralization**: Automatically handles resource naming. Inputting `user` or `users` correctly generates a `users` resource. Handles irregular plurals (e.g., `category` -> `categories`, `box` -> `boxes`).
- **Dynamic Module Detection**: Automatically reads your `go.mod` file to generate correct internal import paths. No hardcoded module names!
- **Fiber-Ready**: Generates code pre-configured for use with the Gofiber v2 framework.

## Installation

### For Developers (Standard Go Install)

The easiest way to install `gogen` is using the `go install` command:

```bash
go install github.com/zaheershaikh936/gogen@latest
```

Make sure your `$GOPATH/bin` is in your system `PATH`.

### For Users (Pre-compiled Binaries)

You can download the latest pre-compiled binaries for your operating system from the [Releases](https://github.com/zaheershaikh936/gogen/releases) page.

### Local Build (Development)

To build and run the tool locally:

```bash
go build -o gogen
```

## Usage

Generate a new resource (model) with the `resource` command:

```bash
./gogen resource [model_name] --output [path]
```

### Example

```bash
./gogen resource user --output ./src/api
```

This command will create the following structure:
```text
src/api/
└── users/
    ├── controllers/
    │   └── users_controller.go
    ├── repositories/
    │   └── users_repository.go
    ├── routes/
    │   └── users_routes.go
    └── services/
        └── users_service.go
```

## Generated Code Structure

- **Routes**: Defines the Fiber router groups and maps HTTP methods to controller functions.
- **Controllers**: Handles request parsing and calls the appropriate service methods.
- **Services**: Contains business logic and orchestrates data movement between controllers and repositories.
- **Repositories**: Direct data access layer (currently initialized with Fiber context placeholders).

## License

[MIT](LICENSE) (or your preferred license)