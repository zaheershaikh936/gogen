# Contributing to Gogen 🎉

Thank you for your interest in contributing to **Gogen**!
We welcome contributions of all kinds—bug reports, feature suggestions, documentation improvements, and code enhancements.

Please take a moment to review this guide before getting started.

---

## 📋 Table of Contents

* How to Report Bugs
* How to Suggest Features
* How to Submit a Pull Request
* Coding Standards
* Running Tests Locally
* Code of Conduct

---

## 🐛 How to Report Bugs

If you encounter a bug, please help us improve Gogen by reporting it.

Before opening a new issue:

* Search existing issues to avoid duplicates
* Ensure you are using the latest version

When reporting a bug, include:

* A clear and descriptive title
* Steps to reproduce the issue
* Expected behavior
* Actual behavior
* Your environment:

  * OS (Linux/macOS/Windows)
  * Go version (`go version`)
  * Gogen version

Example:

```text
Steps to reproduce:
1. Run gogen resource user
2. Use custom output directory
3. Command fails with error

Expected behavior:
CRUD files should be generated successfully.
```

---

## 💡 How to Suggest Features

We love new ideas!

To suggest a feature:

* Open a GitHub issue
* Clearly explain the problem your feature solves
* Describe how you expect it to work
* Include examples or use cases if possible

---

## 🔀 How to Submit a Pull Request

### Step 1: Fork the repository

Click the **Fork** button on GitHub and clone your fork locally.

```bash
git clone https://github.com/<your-username>/gogen.git
cd gogen
```

---

### Step 2: Create a new branch

Create a branch for your change:

```bash
git checkout -b feature/your-feature-name
```

Branch naming conventions:

* `feat/` – new features
* `fix/` – bug fixes
* `docs/` – documentation updates
* `chore/` – maintenance tasks

---

### Step 3: Make your changes

Keep your changes focused and minimal.
Ensure your code is readable and well-structured.

---

### Step 4: Commit your changes

Use clear and meaningful commit messages:

```bash
git commit -m "docs: add contributing guide"
git commit -m "feat: add custom output flag"
git commit -m "fix: handle empty resource name"
```

---

### Step 5: Push and open a Pull Request

```bash
git push origin feature/your-feature-name
```

* Open a Pull Request against the `main` branch
* Reference the related issue (e.g. `Closes #4`)
* Clearly describe what your PR does

---

## 🧹 Coding Standards

This project follows standard Go best practices:

* Format code using `gofmt`
* Use meaningful variable and function names
* Keep functions small and focused
* Follow Go project conventions
* Avoid unnecessary dependencies

Before submitting, run:

```bash
gofmt -w .
```

---

## 🧪 Running Tests Locally

Before opening a Pull Request, ensure all tests pass:

```bash
go test ./...
```

If you add new features, please include relevant tests where possible.

---

## 🤝 Code of Conduct

Be respectful and inclusive in all interactions.

By contributing to this project, you agree to follow GitHub’s community guidelines and maintain a positive environment for everyone.

---

## 🙌 Thank You

Your contributions make **Gogen** better for everyone.
We appreciate your time and effort—happy coding! 🚀



