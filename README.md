# Go Learning Repository 🚀

A comprehensive collection of Go (Golang) examples and exercises for learning the language from basics to advanced concepts.

## 📚 About This Repository

This repository contains hands-on examples and exercises covering various Go programming concepts. Each file/folder demonstrates a specific topic with practical code examples.

## 🎯 Learning Path

### 1. Basics & Fundamentals

**Variables & Data Types:**
- [`variables.go`](./variables.go) - Variable declarations and usage
- [`datatypes.go`](./datatypes.go) - Basic data types in Go
- [`basicdatatypes.go`](./basicdatatypes.go) - More data type examples
- [`constants.go`](./constants.go) - Working with constants
- [`iota.go`](./iota.go) - Using iota for enumerations
- [`conversion-variablescope.go`](./conversion-variablescope.go) - Type conversion and variable scope
- [`printingvarwithnovalues.go`](./printingvarwithnovalues.go) - Zero values demonstration

**Numbers & Operations:**
- [`binary-and-hexadecimal.go`](./binary-and-hexadecimal.go) - Binary and hexadecimal numbers
- [`bitwiseshifting.go`](./bitwiseshifting.go) - Bitwise operations
- [`modulus/`](./modulus/) - Modulus operator examples
- [`ASCII.go`](./ASCII.go) - Working with ASCII values
- [`randmath.go`](./randmath.go) - Random number generation
- [`interfaces-mathrand.go`](./interfaces-mathrand.go) - Math random with interfaces

### 2. Functions

**Function Basics:**
- [`intro-to-function/`](./intro-to-function/) - Introduction to functions
- [`functionwithparameter.go`](./functionwithparameter.go) - Functions with parameters
- [`functionwithoutreturnvalue.go`](./functionwithoutreturnvalue.go) - Void functions
- [`functionswapthereturn.go`](./functionswapthereturn.go) - Multiple return values
- [`functionwithvariablevalues.go`](./functionwithvariablevalues.go) - Functions with variable arguments
- [`functionwithshortdeclarationoperator.go`](./functionwithshortdeclarationoperator.go) - Short declaration in functions

**Advanced Functions:**
- [`anonymous-function/`](./anonymous-function/) - Anonymous functions and closures
- [`variadic-parameter-in-functions/`](./variadic-parameter-in-functions/) - Variadic functions
- [`defer-function/`](./defer-function/) - Defer keyword usage
- [`wrapper-function-to-log/`](./wrapper-function-to-log/) - Wrapper functions for logging
- [`method-reciver-function/`](./method-reciver-function/) - Methods with receivers

### 3. Data Structures

**Arrays & Slices:**
- [`arrays/`](./arrays/) - Array basics
- [`slices/`](./slices/) - Slice fundamentals
- [`appendslices/`](./appendslices/) - Appending to slices
- [`slice-make/`](./slice-make/) - Creating slices with make
- [`slicing-slice/`](./slicing-slice/) - Slicing operations
- [`delete-slice/`](./delete-slice/) - Deleting from slices
- [`multi-dimensional-slice/`](./multi-dimensional-slice/) - Multi-dimensional slices
- [`implications-of-slice/`](./implications-of-slice/) - Slice internals part 1
- [`implications-of-slices-02/`](./implications-of-slices-02/) - Slice internals part 2
- [`unfurling-slice-demo-function/`](./unfurling-slice-demo-function/) - Slice unfurling with variadic functions

**Maps:**
- [`maps/`](./maps/) - Map basics and operations

**Structs:**
- [`structs/`](./structs/) - Struct basics
- [`anonymous-struct/`](./anonymous-struct/) - Anonymous structs
- [`struct-embeddings/`](./struct-embeddings/) - Struct embedding and composition

### 4. Control Flow

**Conditionals:**
- [`if-statements/`](./if-statements/) - If statement basics
- [`switch-statement/`](./switch-statement/) - Switch statements
- [`logical-statements/`](./logical-statements/) - Logical operators
- [`statements-statements-idiom/`](./statements-statements-idiom/) - Go statement idioms

**Loops:**
- [`loopsingo/`](./loopsingo/) - Loop basics
- [`nestedloop/`](./nestedloop/) - Nested loops
- [`rangeingo/`](./rangeingo/) - Range keyword for iteration

### 5. Concurrency

- [`select-statements/`](./select-statements/) - Select statement for channel operations

### 6. Interfaces

- [`interfaceandconstans.go`](./interfaceandconstans.go) - Interfaces and constants
- [`interface-polymorphism/`](./interface-polymorphism/) - Polymorphism with interfaces
- [`stringer-interface-go-internals/`](./stringer-interface-go-internals/) - Stringer interface
- [`writer-implicit-interface/`](./writer-implicit-interface/) - Writer interface (implicit implementation)

### 7. Pointers

- [`pointer.go`](./pointer.go) - Pointer basics
- [`secondpointer.go`](./secondpointer.go) - More pointer examples

### 8. Packages & Modules

- [`new/`](./new/) - Package organization examples
- [`lion/`](./lion/) - Custom package example

### 9. Practice Exercises

All exercises are numbered and progressively build on concepts:

**Beginner Exercises (9-19):**
- [`excersie-#9/`](./excersie-#9/)
- [`exercise#10/`](./exercise#10/)
- [`excersie-#11/`](./excersie-#11/)
- [`excercise-12/`](./excercise-12/)
- [`excercise-19/`](./excercise-19/)

**Intermediate Exercises (22-35):**
- [`excercise-22/`](./excercise-22/) through [`excercise-35/`](./excercise-35/)

**Advanced Exercises (40-51):**
- [`excerxise-40/`](./excerxise-40/)
- [`excercise-41/`](./excercise-41/) through [`excercise-51/`](./excercise-51/)

### 10. Projects

- [`conting-frequeny-of-words-in-book/`](./conting-frequeny-of-words-in-book/) - Word frequency counter project

## 🚀 Getting Started

### Prerequisites

- Go 1.x or higher installed ([Download Go](https://golang.org/dl/))
- Basic understanding of programming concepts

### Running Examples

1. Clone this repository:
```bash
git clone https://github.com/JobinAj/golang-practise-code.git
cd golang-practise-code
```

2. Run any standalone file:
```bash
go run variables.go
```

3. Run examples in folders:
```bash
cd anonymous-function
go run main.go
```

4. Run the main program (if applicable):
```bash
go run main.go
```

## 📖 How to Use This Repository

1. **Follow the Learning Path**: Start with basics and progress sequentially
2. **Read the Code**: Each file contains comments explaining the concepts
3. **Experiment**: Modify the code and see what happens
4. **Complete Exercises**: Work through numbered exercises to practice
5. **Build Projects**: Try the practical projects to apply multiple concepts

## 🤝 Contributing

Feel free to:
- Report issues or bugs
- Suggest new examples or improvements
- Submit pull requests with additional exercises
- Share your learning experience

## 📝 Notes

- Some folder names have intentional variations (excercise vs exercise) - they're preserved from the learning journey
- Each example is self-contained and can be run independently
- Code comments include learning notes and explanations

## 📜 License

This repository is for educational purposes. Feel free to use it for learning Go!

## 🔗 Useful Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Tour](https://tour.golang.org/)

---

**Happy Learning! 🎉**

If you find this repository helpful, please consider giving it a ⭐!
