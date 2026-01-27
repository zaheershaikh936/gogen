# Contributing Guide

Thank you for your interest in contributing to **Gogen**!  
This project is built with care, and contributions are most valuable when they come from a place of understanding and intention.

This guide exists to keep collaboration smooth, respectful, and enjoyable for everyone involved.

---

## 🌱 How We Think About Contributions

Contributions here are less about "just fixing something" and more about:

- Understanding how the system works
- Making changes that fit naturally into the existing codebase
- Sharing your thought process so others can learn from it too

Even small improvements are very welcome — especially when they're thoughtful and well explained.

---

## 🌿 Workflow & Branching

To keep things organized, we follow a simple and consistent workflow.

### 1. Fork the Repository

- Fork this repository to your own GitHub account
- This keeps experimentation safe and avoids conflicts

```bash
git clone https://github.com/<your-username>/gogen.git
cd gogen
```

### 2. Create a Branch in Your Fork

- Keep your fork's `main` branch clean and up to date
- Create a new branch for **each** feature, fix, or issue

Example workflow:

```bash
git checkout main
git pull upstream main
git checkout -b feature/short-description
```

Branch naming conventions:

- `feat/` – new features
- `fix/` – bug fixes
- `docs/` – documentation updates
- `chore/` – maintenance tasks
- `refactor/` – code refactoring

### 3. Make Your Changes

While working on your changes:

- Take time to explore the relevant parts of the codebase
- Follow existing patterns, naming, and structure
- Keep changes focused and avoid unrelated refactors

---

## 🤔 Understanding Over Speed

This project values understanding over quick fixes.

Before opening a pull request, it's encouraged to:

- Look beyond the single file you're editing
- Understand how your change fits into the broader system
- Think about maintainability and future impact

If something isn't clear, feel free to open a discussion or ask questions.

---

## 🏗️ Project Structure

Gogen follows a clean, modular architecture:

```
gogen/
├── cmd/                    # CLI commands and utilities
│   ├── resources.go        # Resource command registration
│   ├── root.go            # Root command setup
│   ├── version.go         # Version command
│   └── utils/             # Core generation logic
│       ├── resource.go    # Main resource generation
│       ├── content/       # Template content for generated files
│       ├── helper/        # Helper functions
│       ├── text/          # Text processing utilities
│       └── ui/            # TUI components
├── main.go                # Entry point
├── go.mod                 # Go module definition
└── README.md              # Project documentation
```

### Key Components

- **`cmd/utils/resource.go`**: Core resource generation logic
- **`cmd/utils/content/`**: Templates for controllers, services, repositories, and routes
- **`cmd/utils/helper/`**: Utility functions for file operations and naming conventions
- **`cmd/utils/ui/`**: Terminal UI components using Bubble Tea and Lipgloss

---

## 🧹 Coding Standards

This project follows standard Go best practices and conventions:

### General Guidelines

- Format code using `gofmt` or `goimports`
- Use meaningful variable and function names
- Keep functions small and focused
- Follow Go project layout conventions
- Avoid unnecessary dependencies

### Naming Conventions

Gogen uses idiomatic Go naming:

- **Exported functions/types**: `PascalCase` (e.g., `GenerateResource`, `UserController`)
- **Internal functions/variables**: `camelCase` (e.g., `formatFileName`, `resourceName`)
- **Constants**: `PascalCase` or `SCREAMING_SNAKE_CASE` depending on context

### Before Submitting

Run these commands to ensure code quality:

```bash
# Format code
gofmt -w .

# Run tests
go test ./...

# Build to verify compilation
go build
```

---

## 🛠️ About AI Tools

AI tools can be helpful when used responsibly.

**Acceptable uses include:**
- Syntax reminders
- Documentation lookups
- Language or framework references
- Boilerplate generation for repetitive patterns

**What's discouraged:**
- Submitting AI-generated code without understanding it
- Patchwork solutions that don't match the project's style or logic
- Copy-pasting without reviewing or testing

**You should always be able to explain why your solution works.**

---

## 📝 Pull Request Expectations

When opening a pull request, please include a clear explanation covering:

- **What approach did you take?**
- **How did you verify the changes?**

Even for small changes, a brief explanation is appreciated.

### Linking Issues

When your PR addresses an issue, use GitHub keywords in your PR description:

```markdown
Fixes #42
Closes #12
Resolves #45
```

This automatically links and closes the issue when the PR is merged.

### Example PR Description

```markdown
## Summary
Adds support for custom template directories

## Approach
I traced the template loading logic in `cmd/utils/content/` and added
a new flag `--template-dir` that allows users to specify custom templates.
The implementation follows the existing pattern used for the `--output` flag.

## Testing
- Tested with custom template directory
- Verified default behavior still works
- Ran existing tests: `go test ./...`

Fixes #42
```

---

## 🧪 Testing & Verification

Before submitting:

- **Test your changes locally** by running the CLI with various inputs
- **Ensure existing tests pass**: `go test ./...`
- **Build the binary**: `go build` to verify compilation
- If no tests exist for your change, describe how you verified it works

### Manual Testing Examples

```bash
# Test basic resource generation
gogen resource user

# Test with custom output
gogen resource product --output ./test-output

# Test edge cases
gogen resource user-profile
gogen resource order-item -o ./api
```

---

## 🐛 Reporting Bugs

If you encounter a bug, please help us improve Gogen by reporting it.

**Before opening a new issue:**

- Search existing issues to avoid duplicates
- Ensure you are using the latest version

**When reporting a bug, include:**

- A clear and descriptive title
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Your environment:
  - OS (Linux/macOS/Windows)
  - Go version (`go version`)
  - Gogen version (`gogen version`)

**Example:**

```text
Steps to reproduce:
1. Run `gogen resource user-order --output ./api`
2. Check generated files
3. Notice incorrect naming in routes file

Expected behavior:
Routes should use `UserOrder` naming consistently

Actual behavior:
Some files use `userOrder` while others use `user_order`

Environment:
- OS: Windows 11
- Go: 1.23.4
- Gogen: v0.2.0
```

---

## 💡 Suggesting Features

We love new ideas!

**To suggest a feature:**

- Open a GitHub issue with the `enhancement` label
- Clearly explain the problem your feature solves
- Describe how you expect it to work
- Include examples or use cases if possible

**Example:**

```markdown
## Feature Request: Support for Echo Framework

### Problem
Currently, Gogen only supports Fiber and Gin. Many Go developers use Echo.

### Proposed Solution
Add Echo as a third framework option in the interactive wizard.

### Use Case
Users could run `gogen resource user` and select Echo from the framework list.
```

---

## ✨ Final Note

This project values **curiosity, clarity, and care**.

- If you're unsure about something — **ask**.
- If you want feedback — **start a discussion**.
- If you want to contribute thoughtfully — **you're very welcome here**.

Thanks for taking the time to read and contribute! 🚀

---

## 🙌 Thank You

Your contributions make **Gogen** better for everyone.  
We appreciate your time and effort — happy coding!

**Made with ❤️ by the Gogen community**
