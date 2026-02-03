# Go Learning Repository 🚀

A comprehensive collection of Go (Golang) examples and exercises organized for progressive learning from basics to advanced concepts.

## 📚 About This Repository

This repository contains hands-on examples and 30 progressive exercises covering various Go programming concepts. The content is organized into logical categories to make learning systematic and easy to follow.

## 🗂️ Repository Structure

```
golang/
├── 01-basics/              # Variables, types, constants, operators
├── 02-functions/           # Function basics, anonymous, variadic, defer
├── 03-control-flow/        # If, switch, loops, range
├── 04-data-structures/     # Arrays, slices, maps, structs
├── 05-pointers/            # Pointer fundamentals
├── 06-interfaces/          # Interface concepts and implementation
├── 07-concurrency/         # Goroutines, channels, select
├── 08-packages/            # Package organization and modules
├── exercises/              # 30 progressive exercises (01-30)
└── projects/               # Practical projects
```

## 🎯 Learning Path

### 1. Basics (`01-basics/`)

**Variables & Data Types:**
- `variables.go` - Variable declarations and usage
- `datatypes.go` - Basic data types in Go
- `basicdatatypes.go` - More data type examples
- `constants.go` - Working with constants
- `iota.go` - Using iota for enumerations
- `conversion-variablescope.go` - Type conversion and variable scope
- `printingvarwithnovalues.go` - Zero values demonstration

**Numbers & Operations:**
- `binary-and-hexadecimal.go` - Binary and hexadecimal numbers
- `bitwiseshifting.go` - Bitwise operations
- `ASCII.go` - Working with ASCII values
- `randmath.go` - Random number generation
- `interfaces-mathrand.go` - Math random with interfaces

### 2. Functions (`02-functions/`)

**Function Basics:**
- `intro-to-function/` - Introduction to functions
- `functionwithparameter.go` - Functions with parameters
- `functionwithoutreturnvalue.go` - Void functions
- `functionswapthereturn.go` - Multiple return values
- `functionwithvariablevalues.go` - Functions with variable arguments
- `functionwithshortdeclarationoperator.go` - Short declaration in functions

**Advanced Functions:**
- `anonymous-function/` - Anonymous functions and closures
- `variadic-parameter-in-functions/` - Variadic functions
- `defer-function/` - Defer keyword usage
- `wrapper-function-to-log/` - Wrapper functions for logging
- `unfurling-slice-demo-function/` - Slice unfurling with variadic functions
- `method-reciver-function/` - Methods with receivers

### 3. Control Flow (`03-control-flow/`)

**Conditionals:**
- `if-statements/` - If statement basics
- `switch-statement/` - Switch statements
- `logical-statements/` - Logical operators
- `statements-statements-idiom/` - Go statement idioms

**Loops:**
- `loopsingo/` - Loop basics
- `nestedloop/` - Nested loops
- `rangeingo/` - Range keyword for iteration
- `modulus/` - Modulus operator in loops

### 4. Data Structures (`04-data-structures/`)

**Arrays & Slices:**
- `arrays/` - Array basics
- `slices/` - Slice fundamentals
- `appendslices/` - Appending to slices
- `slice-make/` - Creating slices with make
- `slicing-slice/` - Slicing operations
- `delete-slice/` - Deleting from slices
- `multi-dimensional-slice/` - Multi-dimensional slices
- `implications-of-slice/` - Slice internals part 1
- `implications-of-slices-02/` - Slice internals part 2

**Maps:**
- `maps/` - Map basics and operations

**Structs:**
- `structs/` - Struct basics
- `anonymous-struct/` - Anonymous structs
- `struct-embeddings/` - Struct embedding and composition

### 5. Pointers (`05-pointers/`)

- `pointer.go` - Pointer basics
- `secondpointer.go` - More pointer examples

### 6. Interfaces (`06-interfaces/`)

- `interfaceandconstans.go` - Interfaces and constants
- `interface-polymorphism/` - Polymorphism with interfaces
- `stringer-interface-go-internals/` - Stringer interface
- `writer-implicit-interface/` - Writer interface (implicit implementation)

### 7. Concurrency (`07-concurrency/`)

- `select-statements/` - Select statement for channel operations

### 8. Packages & Modules (`08-packages/`)

- `new/` - Package organization examples
- `lion/` - Custom package example with dependencies

## 📝 Progressive Exercises (`exercises/`)

**30 exercises organized from basic to advanced:**

### **Basics (01-04)**
- **01** - Variables (var, const, local)
- **02** - Time formatting and Printf
- **03** - Importing packages
- **04** - Package dependencies

### **Control Flow - Conditionals (05-08)**
- **05** - Boolean operators (&&, ||, !)
- **06** - Switch statements
- **07** - If with comma-ok idiom (map lookup)
- **08** - If with comma-ok (float keys)

### **Control Flow - Loops (09-15)**
- **09** - Basic for loop
- **10** - For loop with switch
- **11** - For loop with condition
- **12** - For loop with break
- **13** - For loop with continue
- **14** - Nested loops
- **15** - Range over slice

### **Arrays & Slices - Basics (16-23)**
- **16** - Arrays fundamentals
- **17** - Slice basics and range
- **18** - Slice iteration with range
- **19** - Array/slice practice
- **20** - Appending to slices
- **21** - Slice slicing operations
- **22** - Append multiple values
- **23** - Deleting from slices

### **Slices - Advanced (24-25)**
- **24** - Creating slices with make
- **25** - Multi-dimensional slices

### **Maps (26-29)**
- **26** - Maps with range
- **27** - Map with slice values
- **28** - Map - adding entries
- **29** - Map - iteration patterns

### **Concurrency (30)**
- **30** - Goroutines and channels

## 🚀 Getting Started

### Prerequisites

- Go 1.x or higher installed ([Download Go](https://golang.org/dl/))
- Basic understanding of programming concepts

### Installation

1. Clone this repository:
```bash
git clone https://github.com/JobinAj/golang-practise-code.git
cd golang-practise-code
```

2. Verify Go installation:
```bash
go version
```

### Running Examples

**Run files from topic folders:**
```bash
# Run a specific example
go run 01-basics/variables.go

# Run examples in folders
cd 02-functions/anonymous-function
go run main.go
```

**Run exercises in sequence:**
```bash
# Start with exercise 01
cd exercises/exercise-01
go run main.go

# Progress through exercises
cd ../exercise-02
go run main.go
```

**Run projects:**
```bash
cd projects/conting-frequeny-of-words-in-book
go run main.go
```

## 📖 How to Use This Repository

### **For Beginners:**
1. Start with `01-basics/` to understand fundamentals
2. Move through folders sequentially (01 → 08)
3. Complete exercises `01-15` to reinforce basics
4. Practice with exercises `16-30` for data structures

### **For Intermediate Learners:**
1. Review `04-data-structures/` for arrays, slices, maps
2. Study `06-interfaces/` for polymorphism
3. Practice exercises `16-30`
4. Explore the word frequency project

### **For Advanced Topics:**
1. Deep dive into `07-concurrency/`
2. Study `08-packages/` for module organization
3. Complete exercise `30` (goroutines)
4. Build your own projects using learned concepts

## 💡 Learning Tips

- **Read the Code**: Each file contains comments explaining concepts
- **Experiment**: Modify code and observe the results
- **Run Everything**: Don't just read—execute and see output
- **Follow Order**: Exercises build on previous knowledge
- **Practice Daily**: Consistency is key to mastering Go

## 🤝 Contributing

Contributions are welcome! Please:
- Report issues or suggest improvements
- Add new examples with clear documentation
- Follow existing code style and organization
- Update README when adding new content

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## 📜 License

This repository is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Useful Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Tour](https://tour.golang.org/)
- [Go Playground](https://play.golang.org/)

## 🌟 Progress Tracker

Track your learning progress:
- [ ] Complete 01-basics
- [ ] Complete 02-functions
- [ ] Complete 03-control-flow
- [ ] Complete 04-data-structures
- [ ] Complete 05-pointers
- [ ] Complete 06-interfaces
- [ ] Complete 07-concurrency
- [ ] Complete 08-packages
- [ ] Complete exercises 01-15 (Basics)
- [ ] Complete exercises 16-30 (Advanced)
- [ ] Build a custom project

---

**Happy Learning! 🎉**

*If you find this repository helpful, please consider giving it a ⭐!*

**Repository maintained by:** [JobinAj](https://github.com/JobinAj)
