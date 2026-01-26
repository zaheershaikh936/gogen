# 🧪 Testing and Verification Guide

Since `gogen` is a CLI tool that generates code, testing involves verifying both the **CLI behavior** and the **integrity of the generated code**.

## 1. Testing the CLI Tool

To test the `gogen` tool itself, you can run it directly from the source or build it.

### Basic CLI Verification
Run the following commands to ensure the CLI is responding correctly:

```bash
# Check version (if implemented) or help
go run main.go --help

# Attempt to generate a resource without arguments (should show error/help)
go run main.go resource
```

### Resource Generation Test
Generate a sample resource to verify the file creation logic:

```bash
# Generate a 'book' resource in a temporary directory
go run main.go resource book --output ./test_output
```

**What to check:**
- [ ] Does the command exit successfully with a "success" message?
- [ ] is the TUI (terminal interface) rendering correctly?
- [ ] Are the files created in the correct location (`./test_output/book/...`)?

---

## 2. Verifying Generated Code

Once the files are generated, you should verify their contents for correctness and idiomatic naming.

### File Structure Check
Verify that the `clean architecture` folders are created:
```
test_output/book/
├── controllers/
├── repositories/
├── routes/
└── services/
```

### Content Verification
Open the generated files and check for:
- **Package Names**: Should be consistent with the folder (e.g., `package controllers`).
- **Imports**: Check if the module name (from `go.mod`) is correctly prepended to internal imports.
- **Naming**: Ensure `book` was converted to `Book` (PascalCase) in function names like `GetBook` and `CreateBook`.

### Compilation Test
The best way to "check" the generated code is to try and compile a project that uses it.

1. Create a dummy Fiber app.
2. Generate a resource into it.
3. Run `go build ./...` to ensure there are no syntax errors or circular dependencies.

---

## 3. Automated Testing (Future)
To implement automated tests for this project, you would typically use Go's `testing` package to:
1. Run the generator function with specific inputs.
2. Compare the resulting file content against "golden files" (expected output templates).
