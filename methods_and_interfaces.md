# Methods and Interfaces

## Methods

Go does not have classes.
However, you can define methods on types.

A method is a function with a special *receiver* argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

### Methods are functions

### Methods continued

You can declare a method on non-struct types, too.

You can only declare a method with a receiver whose type is defined in the same package as the method.
You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

### Pointer receivers

You can declare methods with pointer receivers.

This means the receiver type has the literal syntax `*T` for some type `T`.

Methods with pointer receivers can modify the value to which the receiver points.
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

### Pointers and functions

### Methods and pointer indirection

Methods with pointer receivers take either a value or a pointer as the receiver when they are called.

That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

## Interfaces

An *interface type* is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

### Interfaces are implemented implicitly

A type implements an interface by implementing its methods.
There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

### Interface values

### The empty interface

The interface type that specifies zero methods is known as the *empty interface*:
```go
interface{}
```
An empty interface may hold values of any type.

### Type assertions

### Type switches

A *type switch* is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

The declaration in a type switch has the same syntax as a type assertion `i.(T)` but the specific type `T` is replaced with the key word `type`.
