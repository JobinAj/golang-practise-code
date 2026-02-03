# Contributing to Go Learning Repository

Thank you for your interest in contributing to this Go learning repository! 🎉

## How to Contribute

### Reporting Issues

If you find any errors or have suggestions:
1. Check if the issue already exists
2. Create a new issue with a clear description
3. Include code examples if applicable

### Adding New Examples

We welcome new examples! Please ensure:

1. **Code Quality:**
   - Code is well-formatted (`go fmt`)
   - Code runs without errors
   - Include helpful comments explaining concepts

2. **Structure:**
   - Create a new directory for complex examples
   - Use standalone `.go` files for simple concepts
   - Include a descriptive name

3. **Documentation:**
   - Add clear comments in the code
   - Update the README.md to include your example in the appropriate section
   - Explain what concept the example demonstrates

### Improving Existing Examples

To improve existing code:
1. Fork the repository
2. Create a feature branch (`git checkout -b improve-example-name`)
3. Make your changes
4. Ensure code still runs correctly
5. Submit a pull request with a description of improvements

### Code Style Guidelines

- Follow standard Go conventions ([Effective Go](https://golang.org/doc/effective_go))
- Use meaningful variable names
- Add comments for complex logic
- Keep examples simple and focused on one concept
- Format code with `go fmt` before committing

### Submitting Pull Requests

1. Fork the repository
2. Create your feature branch
3. Make your changes
4. Test your code
5. Commit with clear, descriptive messages
6. Push to your fork
7. Open a pull request

### Example Contribution

```go
// Good example - clear and educational
package main

import "fmt"

// Demonstrates the defer keyword
// defer schedules a function call to run after the function completes
func main() {
    defer fmt.Println("This runs last")
    fmt.Println("This runs first")
    fmt.Println("This runs second")
}
```

## Questions?

Feel free to open an issue for any questions about contributing!

## Code of Conduct

Be respectful, constructive, and helpful. This is a learning environment for everyone.

---

Happy coding! 🚀
