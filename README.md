# 🚀 Gogen - The Premium Go CRUD Generator

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zaheershaikh936/gogen)](https://goreportcard.com/report/github.com/zaheershaikh936/gogen)

> Stop writing boilerplate. Start building. ⚡

**Gogen** is an interactive, high-end CLI tool designed to generate complete CRUD resources for the Go Fiber framework in seconds. It follows clean architecture principles and delivers a premium developer experience with a "shadcn-style" terminal UI.

---

## ✨ Features

- 💎 **Premium TUI** - Beautifully designed terminal interface with borders, boxed reports, and vibrant colors.
- 🪄 **Interactive Wizard** - Don't remember the syntax? Just run `gogen resource` and let the wizard guide you.
- ⚡ **Real-time Feedback** - Animated spinners and smooth transitions for a high-end feel.
- 📁 **Clean Architecture** - Automatically generates Controllers, Services, Repositories, and Routes.
- 🎯 **One Command Generation** - Non-interactive mode for fast, scriptable resource creation.

---

## 🛠️ Installation

```bash
go install github.com/zaheershaikh936/gogen@latest
```

Ensure your `$GOPATH/bin` is in your `PATH`.

---

## 🚀 Usage

### 🪄 Interactive Mode (Recommended)
Simply run the command and follow the prompts:
```bash
gogen resource
```

### 🎯 Quick Mode
Specify the model name directly for instant generation:
```bash
gogen resource user
```

### 📁 Custom Output
Generate files in a specific directory:
```bash
gogen resource product --output ./internal/api
```

---

## 🏗️ Generated Architecture

Every resource is scaffolded with a robust, production-ready structure:

```text
users/
├── controllers/    # Fiber HTTP handlers
├── services/       # Business logic layer
├── repositories/   # Data access layer
└── routes/         # API route definitions
```

---

## 🤝 Contributing

Contributions make the open-source community an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📝 License

Distributed under the MIT License. See `LICENSE` for more information.

Developed with ❤️ by **Zaheer Shaikh**
