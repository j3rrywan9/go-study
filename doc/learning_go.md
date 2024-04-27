# Learning Go

## Chapter 2. Primitive Types and Declarations

### Built-in types

Go has many of the same built-in types as other languages: booleans, integers, floats, and strings.
Using these types idiomatically is sometimes a challenge for developers who are transitioning from another language.
We are going to look at these types and see how they work best in Go.
Before we review the types, let's cover some of the concepts that apply to all types.

#### The zero value

Go, like most modern languages, assigns a default *zero value* to any variable that is declared but not assigned a value.
Having an explicit zero value makes code clearer and removes a source of bugs found in C and C++ programs.
As we talk about each type, we will also cover the zero value for the type.

#### Literals

In order to make it easier to read longer integer literals, Go allows you to put underscores in the middle of your literal.

*Rune literals* represent characters and are surrounded by single quotes.
Unlike many languages, single quotes and double quotes are not interchangeable.

#### Booleans

The `bool` type represents Boolean variables.
Variables of bool type can have one of two values: `true` or `false`.
The zero value for a `bool` is `false`.
```go
var flag bool // no value assigned, set to false
var isAwesome = true
```

#### A Taste of Strings and Runes

#### `var` vs `:=`

For a small language, Go has a lot of ways to declare variables.
There's a reason for this; each declaration style communicates something about how the variable is used.
Let's go through the ways you can declare a variable in Go, and see when each is appropriate.

#### Unused variables

#### Naming variables and constants

Even though underscore is a valid character in a variable name, it is rarely used, because idiomatic Go doesn't use snake case (names like `index_counter` or `number_tries`).
Instead, Go uses camel case (names like `indexCounter` or `numberTries`) when an identifier name consists of multiple words.

Within a function, favor short variable names.
The smaller the scope for a variable, the shorter the name that's used for it.
It is very common in Go to see single letter variable names.

## Chapter 3. Composite Types

### Arrays

### Slices

Most of the time, when you want a data structure that holds a sequence of values, a slice is what you should use.
What makes slices so useful is that the length is *not* part of the type for a slice.
This removes the limitations of arrays.

Working with slices looks quite a bit like working with arrays, but there are subtle differences.
The first thing to notice is that we don’t specify the size of the slice when we declare it:
```go
var x = []int{10, 20, 30}
```
This creates a slice of 3 `int`s using a *slice literal*.
Just like arrays, we can also specify only the indices with values in the slice literal:
```go
var x = []int{1, 5: 4, 6, 10: 100, 15}
```
You read and write slices using bracket syntax, and just like arrays, you can't read or write past the end or use a negative index:
```go
x[0] = 0
fmt.Println(x[2])
```
So far, slices have seemed identical to arrays.
We start to see the differences between arrays and slices when we look at declaring slices without using a literal.
```go
var x []int
```
This creates a slice of `int`s.
Since no value is assigned, `x` is assigned to the zero value for a slice, which is something we haven't seen before: `nil`.
In Go, `nil` is an identifier that represents the lack of a value for some types.
Like the untyped numeric constants we saw in the previous chapter, `nil` has no type, so it can be assigned or compared against values of different types.
A `nil` slice contains nothing.

#### `len`

#### `append`

The built-in `append` function is used to grow slices:
```go
var x []int
x = append(x, 10)
```
The append function takes at least two parameters, a slice of any type and a value of that type.
It returns a slice of the same type.
The returned slice is assigned back to the slice that's passed in.

#### Capacity

#### `make`

### Strings and Runes and Bytes

### Maps

Like most languages, Go provides a built-in data type for situations where you want to associate one value to another.
The map type is written as `map[keyType]valueType`.
Let's take a look at a few different ways to declare maps.
First, you can use a `var` declaration to create a map variable that's set to its zero value:
```go
var nilMap map[string]int
```
In this case `nilMap` is declared to be a map with `string` keys and `int` values.
The zero value for a map is `nil`.
A `nil` map has a length of 0.
Attempting to read a `nil` map always returns the zero value for the map's value type.
However, attempting to write to a `nil` map variable causes a panic.

We can use a `:=` declaration to create a map variable by assigning it a *map literal*:
```go
totalWins := map[string]int{}
```
In this case, we are using an empty map literal.

The key for a map can be any comparable type.
This means you cannot use a slice or a map as the key for a map.

#### Reading and Writing a Map

#### The comma ok idiom

However, you sometimes do need to find out if a key is in a map.
Go provides the *comma ok idiom* to tell the difference between a key that's associated with a zero value and a key that's not in the map.
```go
m := map[string]int {
    "hello": 5,
    "world": 0,
}
v, ok := m["hello"]
fmt.Println(v, ok)

v, ok := m["world"]
fmt.Println(v, ok)

v, ok := m["goodbye"]
fmt.Println(v, ok)
```
Rather than assign the result of a map read to a single variable, with the comma ok idiom you assign the results of a map read to two variables.
The first gets the value associated with the key.
The second value returned is a bool.
It is usually named `ok`.
If `ok` is `true`, the key is present in the map.
If `ok` is `false`, the key is not present.

#### Deleting From Maps

Key-value pairs are removed from a map via the built-in `delete` function:
```go
m := map[string]int {
    "hello": 5,
    "world": 0,
}
delete(m, "hello")
```
The `delete` function takes a map and a key and then removes the key-value pair with the specified key.
If the key isn't present in the map or if the map is `nil`, nothing happens.
The `delete` function doesn't return a value.

### Structs

When you have related data that you want to group together, you should define a *struct*.
Most languages have a concept that's similar to a struct and the syntax that Go uses to read and write structs should look familiar.
```go
type person struct {
    name string
    age int
    pet string
}
```
A struct type is defined with the keyword `type`, the name of the struct type, the keyword `struct`, and a pair of braces (`{}`).
Within the braces, you list the fields in the struct.
Just like we put the variable name first and the variable type second in a `var` declaration, we put the struct field name first and the struct field type second.
Also note that unlike map literals, there are no commas separating the fields in a struct declaration.
You can define a struct type inside or outside of a function.
A struct type that's defined within a function can only be used within the function.

Once a struct type is declared, we can define variables of that type.
```go
var fred person
```
Here we are using a `var` declaration.
Since no value is assigned to `fred`, it gets the zero value for the `person` struct type.
A zero value struct has every field set to the field's zero value.

A *struct literal* can be assigned to a variable as well.
```go
bob := person{}
```
Unlike maps, there is no difference between assigning an empty struct literal and not assigning a value at all.
Both initialize all of the fields in the struct to the fields' zero values.
There are two different styles for a non-empty struct literal.
A struct literal can be specified as a comma-separated list of values for the fields inside of braces:

A field in a struct is accessed with dotted notation:
```go
bob.name = "Bob"
fmt.Println(bob.name)
```

#### Anonymous Structs

## Chapter 4. Blocks, Shadows, and Control Structures

### Blocks

Each place where a declaration occurs is called a *block*.
Variables, constants, types and functions declared outside of any functions are placed in the *package* block.
We've used import statements in our programs to gain access to printing and math functions.
They define names for other packages that are valid for the file that contains the `import` statement.
These names are in the *file* block.
All of the variables defined at the top level of a function (including the parameters to a function) are in a block.
Within a function, every set of braces (`{}`) defines another block, and in a bit we will see that the control structures in Go define blocks of their own.

You can access an identifier defined in any outer block from within any inner block.
This raises the question: what happens when you have a declaration with the same name as an identifier in a containing block?
If you do that, you *shadow* the identifier created in the outer block.

## Chapter 5. Functions

### Declaring and Calling Functions

We've already seen functions being declared and used.
Every program we've written has a `main` function that's the starting point for every Go program, and we've been calling the `fmt.Println` function to display output to the screen.

#### Simulating Named and Optional Parameters

Before we get to the unique function features that Go has, let's mention two that Go *doesn't* have: named and optional input parameters.
If you want to emulate named and optional parameters, define a struct that has fields that match the desired parameters, and pass the struct to your function.
Here's a snippet of code that demonstrates this pattern:
```go
type MyFuncOpts struct {
    FirstName string
    LastName string
    Age int
}

func MyFunc(opts MyFuncOpts) error {
    // do something here
}

func main() {
    MyFunc(MyFuncOpts {
        LastName: "Patel",
        Age: 50,
    })
    MyFunc(MyFuncOpts {
        FirstName: "Joe",
        LastName: "Smith",
    })
}
```

#### Variadic Input Parameters and Slices

Like many languages, Go supports *variadic parameters*.
The variadic parameter must be the last (or only) parameter in the input parameter list.
You indicate it with three dots (`...`) before the type.
The variable that's created within the function is a slice of the specified type.
You use it just like any other slice.

#### Multiple Return Values

The first difference that we'll see between Go and other languages is that Go allows for multiple return values.

There are a few changes to support multiple return values.
When a Go function returns multiple values, the types of the return values are listed in parentheses, separated by commas.
Also, if a function returns multiple values, you must return all of them, separated by commas.

#### Ignoring Returned Values

If a function returns multiple values, but you don't need to read one or more of the values, assign the unused values to the name `_`.
Surprisingly, Go does let you implicitly ignore *all* of the return values for a function.

#### Named Return Values

#### Blank Returns - Never Use These!

### Functions Are Values

Just as in many other languages, functions in Go are values.
The type of a function is built out of the keyword `func `and the types of the parameters and return values.
This combination is called the *signature* of the function.
Any function that has the exact same number and types of parameters and return values meets the type signature.

Since functions are values, you can declare a function variable:
```go
var myFuncVariable func(string) int
```

The default zero value for a function variable is `nil`.

#### Function Type Declarations

Just as you can use the `type` keyword to define a `struct`, you can use it to define a function type too (I'll go into more details on type declarations in Chapter 7):
```go
type opFuncType func(init, int) int
```

#### Anonymous Functions

You can not only assign functions to variables, but also define new functions within a function and assign them to variables.

Inner functions are *anonymous*; they don't have a name.
You declare an anonymous function with the keyword `func` immediately followed by the input parameters, the return values, and the opening brace.
It is a compile-time error to try to put a function name between `func` and the input parameters.

Just like any other function, an anonymous function is called by using parentheses.

You don't have to assign an anonymous function to a variable.
You can write them inline and call them immediately.
The previous program can be rewritten into this:
```go
func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}
```

Now, this is not something that you would normally do.
If you are declaring and executing an anonymous function immediately, you might as well get rid of the anonymous function and just call the code.
However, declaring anonymous functions without assigning them to variables is useful in two situations: `defer` statements and launching goroutines.

Before using a package-level anonymous function, be very sure you need this capability.
Package-level state should be immutable to make data flow easier to understand.
If a function's meaning changes while a program is running, it becomes difficult to understand not just how data flows, but how it is processed.

### Closure

Functions declared inside functions are special; they are *closures*.
This is a computer science word that means that functions declared inside functions are able to access and modify variables declared in the outer function.

#### Passing Functions as Parameters

Since functions are values and you can specify the type of a function using its parameter and return types, you can pass functions as parameters into functions.
If you aren't used to treating functions like data, you might need a moment to think about the implications of creating a closure that references local variables and then passing that closure to another function.
It's a very useful pattern that appears several times in the standard library.

#### Returning Functions from Functions

In addition to using a closure to pass some function state to another function, you can also return a closure from a function.

### `defer`

Programs often create temporary resources, like files or network connections, that need to be cleaned up.
This cleanup has to happen, no matter how many exit points a function has, or whether a function completed successfully or not.
In Go, the cleanup code is attached to the function with the `defer` keyword.

First, you make sure that a filename was specified on the command line by checking the length of `os.Args`, a slice in the `os` package.
The first value in `os.Args` is the name of the program.
The remaining values are the arguments passed to the program.

### Go Is Call by Value

You might hear people say that Go is a _call-by-value_ language and wonder what that means.
It means that when you supply a variable for a parameter to a function, Go _always_ makes a copy of the value of the variable.

## Chapter 6. Pointers

### A Quick Pointer Primer

A _pointer_ is simply a variable that holds the location in memory where a value is stored.

Every variable is stored in one or more contiguous memory locations, called *addresses*.
Different types of variables can take up different amounts of memory.

A pointer is simply a variable whose contents are the address where another variable is stored.

While different types of variables can take up different numbers of memory locations, every pointer, no matter what type it is pointing to, is always the same size: a number that holds the location in memory where the data is stored.

A pointer holds a number that indicates the location in memory where the data being pointed to is stored.
That number is called the _address_.

The zero value for a pointer is `nil`.

Go's pointer syntax is partially borrowed from C and C++.
Since Go has a garbage collector, most of the pain of memory management is removed.
Furthermore, some of the tricks that you can do with pointers in C and C++, including *pointer arithmetic* are not allowed in Go.

The `&` is the address operator.
It precedes a value type and returns the address where the value is stored:
```go
x := "hello"
pointerToX := &x
```

The `*` is the indirection operator.
It precedes a variable of pointer type and returns the pointed-to value.
This is called _dereferencing_:
```go
x := 10
pointerToX := &x
fmt.Println(pointerToX)  // prints a memory address
fmt.Println(*pointerToX) // prints 10
z := 5 + *pointerToX
fmt.Println(z)           // prints 15
```
Before dereferencing a pointer, you must make sure that the pointer is non-nil.
Your program will panic if you attempt to dereference a `nil` pointer:
```go
var x *int
fmt.Println(x == nil) // prints true
fmt.Println(*x)       // panics
```
A _pointer type_ is a type that represents a pointer.
It is written with a `*` before a type name.
A pointer type can be based on any type:
```go
x := 10
var pointerToX *int
pointerToX = &x
```
The built-in function `new` creates a pointer variable.
It returns a pointer to a zero-value instance of the provided type.
```go
var x = new(int)
fmt.Println(x == nil) // prints false
fmt.Println(*x)       // prints 0
```
The `new` function is rarely used.
For structs, use an `&` before a struct literal to create a pointer instance.
You can't use an & before a primitive literal (numbers, booleans, and strings) or a constant because they don't have memory addresses; they exist only at compile time.
When you need a pointer to a primitive type, declare a variable and point to it:
```go
x := &Foo{}
var y string
z := &y
```

### Don't Fear The Pointers

The first rule of pointers is to not be afraid of them.
If you are used to Java, JavaScript, Python, or Ruby, you might find pointers intimidating.
However, pointers are actually the familiar behavior for classes.
It's the nonpointer structs in Go that are unusual.

What you are seeing is that every instance of a class in these languages is implemented as a pointer.
When a class instance is passed to a function or method, the value being copied is the pointer to the instance.

The difference between Go and these languages is that Go gives you the _choice_ to use pointers or values for both primitives and structs.
Most of the time, you should use a value.
Values make it easier to understand how and when your data is modified.
A secondary benefit is that using values reduces the amount of work that the garbage collector has to do.

### Pointers Indicate Mutable Parameters

As you've already seen, Go constants provide names for literal expressions that can be calculated at compile time.
Go has no mechanism to declare that other kinds of values are immutable.

However, if a pointer is passed to a function, the function gets a copy of the pointer.
This still points to the original data, which means that the original data can be modified by the called function.

### Pointers Are a Last Resort

### Pointer Passing Performance

If a struct is large enough, using a pointer to the struct as either an input parameter or a return value improves performance.
The time to pass a pointer into a function is constant for all data sizes, roughly one nanosecond.
This makes sense, as the size of a pointer is the same for all data types.

### The Zero Value Versus No Value

Again, JSON conversions are the exception that proves the rule.
When converting data back and forth from JSON (yes, I'll talk more about the JSON support in Go's standard library in "encoding/json"), you often need a way to differentiate between the zero value and not having a value assigned at all.
Use a pointer value for fields in the struct that are nullable.

### The Difference Between Maps and Slices

### Slices as Buffers

### Reducing the Garbage Collector's Workload

### Tuning the Garbage Collector

## Chapter 7. Types, Methods, and Interfaces

As you saw in earlier chapters, Go is a statically typed language with both built-in types and user-defined types.
Like most modern languages, Go allows you to attach methods to types.
It also has type abstraction, allowing you to write code that invokes methods without explicitly specifying the implementation.

However, Go's approach to methods, interfaces, and types is very different from most other languages in common use today.
Go is designed to encourage the best practices that are advocated by software engineers, avoiding inheritance while encouraging composition.
In this chapter, we'll take a look at types, methods, and interfaces, and see how to use them to build testable and maintainable programs.

### Types in Go

Back in "Structs", you saw how to define a struct type:
```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}
```
This should be read as declaring a user-defined type with the name `Person` to have the _underlying type_ of the struct literal that follows.
In addition to struct literals, you can use any primitive type or compound type literal to define a concrete type.
Here are a few examples:
```go
type Score int
type Converter func(string)Score
type TeamScores map[string]Score
```

### Methods

Like most modern languages, Go supports methods on user-defined types.

The methods for a type are defined at the package block level:
```go
type Person struct {
    FirstName string
    LastName string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
```
Method declarations look just like function declarations, with one addition: the *receiver* specification.
The receiver appears between the keyword `func` and the name of the method.
Just like all other variable declarations, the receiver name appears before the type.
By convention, the receiver name is a short abbreviation of the type's name, usually its first letter.
It is non-idiomatic to use `this` or `self`.

There is one key difference between declaring methods and functions: methods can be defined _only_ at the package block level, while functions can be defined inside any block.

Just like functions, method names cannot be overloaded.
You can use the same method names for different types, but you can't use the same method name for two different methods on the same type.
While this philosophy feels limiting when coming from languages that have method overloading, not reusing names is part of Go's philosophy of making clear what your code is doing.

We'll talk more about packages in Chapter 10, but be aware that methods must be declared in the same package as their associated type; Go doesn't allow you to add methods to types you don't control.
While you can define a method in a different file within the same package as the type declaration, it is best to keep your type definition and its associated methods together so that it's easy to follow the implementation.

Method invocations should look familiar to those who have used methods in other languages:
```go
p := Person{
    FirstName: "Fred",
    LastName:  "Fredson",
    Age:       52,
}
output := p.String()
```

#### Pointer Receivers and Value Receivers

As I covered in Chapter 6, Go uses parameters of pointer type to indicate that a parameter might be modified by the function.
The same rules apply for method receivers too.
They can be _pointer receivers_ (the type is a pointer) or _value receivers_ (the type is a value type).
The following rules help you determine when to use each kind of receiver:

#### Code Your Methods for `nil` Instances

#### Methods Are Functions Too

#### Functions Versus Methods

#### Type Declarations Aren't Inheritance

#### Types Are Executable Documentation

### `iota` Is for Enumerations - Sometimes

Many programming languages have the concept of enumerations, which allow you to specify that a type can have only a limited set of values.
Go doesn't have an enumeration type.
Instead, it has `iota`, which lets you assign an increasing value to a set of constants.

### Use Embedding for Composition

While Go doesn't have inheritance, it encourages code reuse via built-in support for composition and promotion:
```go
type Employee struct {
    Name string
    ID string
}

func (e Employee) Description() string {
    return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
    Employee
    Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
    // do business logic
}
```
Note that `Manager` contains a field of type `Employee`, but no name is assigned to that field.
This makes `Employee` an *embedded field*.
Any fields or methods declared on an embedded field are *promoted* to the containing struct and can be invoked directly on it.
That makes the following code valid:
```go
m := Manager{
    Employee: Employee{
        Name: "Bob Bobson",
        ID:   "12345",
    },
    Reports: []Employee{},
}
fmt.Println(m.ID)            // prints 12345
fmt.Println(m.Description()) // prints Bob Bobson (12345)
```

### Embedding Is Not Inheritance

Built-in embedding support is rare in programming languages (I'm not aware of another popular language that supports it).
Many developers who are familiar with inheritance (which is available in many languages) try to understand embedding by treating it as inheritance.
That way lies tears.
You cannot assign a variable of type `Manager` to a variable of type `Employee`.
If you want to access the `Employee` field in `Manager`, you must do so explicitly.

### A Quick Lesson on Interfaces

While Go's concurrency model (which we cover in Chapter 12) gets all of the publicity, the real star of Go's design is its implicit interfaces, the only abstract type in Go.
Let's see what makes them so great.

Let's start by taking a quick look at how to declare interfaces.
At their core, interfaces are simple.
Like other user-defined types, you use the `type` keyword.

Here's the definition of the `Stringer` interface in the `fmt` package:
```go
type Stringer interface {
    String() string
}
```
In an interface declaration, an interface literal appears after the name of the interface type.
It lists the methods that must be implemented by a concrete type to meet the interface.
The methods defined by an interface are the method set of the interface.

Interfaces are usually named with "er" endings.

### Interfaces are Type-Safe Duck Typing

So far, nothing that's been said about the Go interface is much different from interfaces in other languages.
What makes Go's interfaces special is that they are implemented `implicitly`.
As you've seen with the `Counter` struct type and the `Incrementer` interface type that you've used in previous examples, a concrete type does not declare that it implements an interface.
If the method set for a concrete type contains all the methods in the method set for an interface, the concrete type implements the interface.
Therefore, that the concrete type can be assigned to a variable or field declared to be of the type of the interface.

To understand why, let's talk about why languages have interfaces.
Earlier I mentioned that _Design Patterns_ taught developers to favor composition over inheritance.
Another piece of advice from the book is "Program to an interface, not an implementation."
Doing so allows you to depend on behavior, not on implementation, allowing you to swap implementations as needed.
This allows your code to evolve over time, as requirements inevitably change.

## Chapter 16. Here Be Dragons: Reflect, Unsafe, and Cgo

All of those things are true, and for the vast majority of the Go code that you'll write, you can be assured that the Go runtime will protect you.
But there are escape hatches.
Sometimes your Go programs need to venture out into less defined areas.
In this chapter, you're going to look at how to handle situations that can't be solved with normal Go code.
For example, when the type of the data can't be determined at compile time, you can use the reflection support in the `reflect` package to interact with and even construct data.
When you need to take advantage of the memory layout of data types in Go, you can use the `unsafe` package.
And if there is functionality that can only be provided by libraries written in C, you can call into C code with `cgo`.
