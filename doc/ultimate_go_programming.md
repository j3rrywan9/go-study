# Ultimate Go Programming

## Introduction

I have taken four years of writing and thinking about Go and been able to put it together in this class.
So in this class, what we're gonna be covering is three things.
It's three things that I think are incredibly important to learning anything.
History, mechanics and then design.
So I'm gonna give you a little bit of history throughout this course about where these idioms, where these things are coming from.
Then I'm gonna be sharing you the mechanics of how things work inside and out, and finally we get to talk about design, and how to leverage all of this information to do things in a productive and correct way.

And finally of course you're gonna want that concurrency stuff, Goroutines, data races, channels, patterns, ways to think about writing multithreaded software in Go, and once all that's done and you have that strong foundation, we'll talk about testing and really cool things around how you can profile that code, but I always want us to be focusing on optimizing for correctness first.
This is gonna be a very big theme throughout this class.

## Lesson 1: Design Guidelines

### Prepare Your Mind

Now, a lot of programming languages you may have been using, your Java, your C#, your Ruby.
Their model is based on a virtual machine but Go is based on a real machine.
And this is a huge advantage to us because as we're gonna see we can look at code and know how it's gonna run on that machine.
Have some indication, some expectation of this, and it's not something that we ever wanna lose again.

I want you to be a champion of quality, efficiency, and simplicity.
I'm gonna show you how the language is a champion for these things.
How if we follow the idioms, if we follow the design guidelines, we can do this.

### Productivity versus Performance

Productivity has always mattered more.

### Correctness versus Performance

What we need to do is optimize for correctness.
Write code that is correct.

### Code Reviews

And our number one priority must always be integrity. Integrity must be first.
In fact, Go puts integrity above all else.
It is built deep into the language.
Now here's another buzzword, integrity.
I mean, what does integrity mean?
What integrity means is that we become serious about writing code that is reliable.

Now, for me, integrity comes in two pieces, a micro level and a macro level.
At the micro-level, integrity is about every read, every write, every memory allocation that we perform being accurate, consistent, and efficient.

This idea that you hide complexity without losing your ability to understand the cost of your decisions.

## Lesson 2: Language Syntax

### Variables

Because type provides the compiler, it provides us two pieces of critical information that's going to lend itself to integrity and readability.
That is, the amount of memory that we need to allocate and what that memory represents.
We need both pieces, without it the compiler cannot do its job and we cannot do our job.
Type is life.
```go
var a int
var b string
var c float64
var d bool
```
Now there's another big thing here, Go has this concept of zero value.
It might be one of the most important concepts that Go has.
And what zero value says is, every single variable or value we create, must be initialized.
And if we don't specify the initialization ourselves, then it gets initialized to its zero value.
I'm not saying default, I'm saying zero.
What it means is the entire allocation of memory, the four bytes, the eight bytes, whatever that is, we reset every bit to zero.
Zero value is a function of integrity, integrity first.
We will always take a cost if it means we get integrity.
So zero value is critical.

This is the short variable declaration operator, `:=`.
It is a productivity operator.
It allows us to declare a variable and initialize it at the same time.
At the exact same time.
Because the language is able to look at the type of the value on the right hand side, and use that to declare the variable on the left hand side.
When we are declaring something **not** set to its zero value, I want you to do this.
I want you to use that short variable declaration operator.
Declaring a variable set to its zero value, we're gonna start seeing `var`.
You set a variable not to its zero value, we're using the short variable declaration operator.

### Struct Types

Now if a struct like this is important, and we need to minimize the amount of padding possible, what we are told is always lay the fields out from highest to smallest.
This will push any padding down to the bottom. 

### Pointers

#### Pass by Value

Now, the one thing we have to understand about Go is everything is about pass by value.
And pass by value for me is WYSIWYG.
What you see is what you get, and you get integrity and readability through pass by value.
So we're gonna see that now by just looking at some pointer mechanics to understand things.
Now there's another thing you have to understand about pointers.
Pointers serve one purpose and one purpose only, and that is sharing.
If the word shared does not come out of your mouth then you do not need pointers.
Pointers are for sharing a value across a program boundary. 

Threads are our path of execution at the operating system level.
All of the code you're writing at some point gets into machine code, and the operating system's job is to choose a path of execution, a thread to execute those instructions one after the other.

The goroutine only has direct access to memory for the frame that it is operating on.
Understand that this frame is serving a really important purpose.
It's creating a sandbox, a layer of isolation.
It's giving us a sense of immutability that the goroutine can only mutate or cause problems here and nowhere else in our code.
This is very, very powerful constructs that we're gonna wanna leverage and it starts to allow us to talk about things like semantics.

You're gonna hear me use the words mechanics and semantics.
When I talk about mechanics, I'm talking about how things work.
When I talk about semantics, I'm talking about how things behave.
The semantics are gonna let us understand what the behavior is, the impact we're gonna have.
The mechanics are gonna allow us to visualize and see how all that works.
We need both.
The semantics, though, are very, very important and powerful.

#### Escape Analysis

Now, as I shared before, we don't have constructors in Go.
Yay, we don't want that, right?
It hides cost, but what we do have is what I call factory functions.
Factory function is a function that creates a value, initializes it for use, and returns it back to the caller.
This is great for readability, it doesn't hide cost, we can read it, and lends to simplicity in terms of construction.

The compiler is able to perform static code analysis, and in this case what the compiler will do is perform what's called escape analysis.
And this static code analysis called escape analysis will determine whether a value gets to be placed on the stack, which is what we want, or it escapes to the heap.
Notice that our first priority is that a value stays on the stack.
And this is because that memory is already there.
It's very very fast to leverage the stack.
Also stacks are self-cleaning, which means that the garbage collector doesn't even get involved until a value escapes a stack and ends up on the heap.
And in Go, that is what we call an allocation.
An allocation in Go is when an escape analysis determines that a value cannot be constructed on the stack, but has to be constructed on the heap because to keep it on the stack, like in this particular case would have an integrity issue.
And now the garbage collector has to get involved with anything that allocates to the heap.

The stack can give us a tremendous amount of performance because the memory is already allocated and it's self-cleaning.
We really want to try our best to leverage our value semantics and to keep values on the stack because again, one of the benefits of not only the isolation and the immutability and the reducing our side effects, but in many cases can also give us better performance because once something is allocated, the garbage collector has to get involved.
Again, we gotta learn how to balance our value and our pointer semantics.

It abstracts away the machine, so you don't necessarily have to care where data is, but when performance starts to matter and debugging starts to matter, it does begin to become important to understand the escape analysis and where things are.
So let's walk through this code with the understanding of what's really happening thanks to escape analysis.

Well the way escape analysis works is it doesn't care about construction.
Construction in Go tells you nothing.
What tells us everything is how a value is shared.
Sharing tells us everything, and because of line 46 the sharing of the value up the call stack, that's gonna clue us in that escape analysis is not gonna construct `u` user on the stack.
But it's going to go and construct this user value out on the heap.
Interesting, construction is happening immediately on the heap.
Now, what gets super interesting to me is this.
`u` represents a value of type `user`, but it represents a value of type `user` that's not on the stack frame, but on the heap.
And what we know is that if you want to access any sort of data that doesn't exist inside your frame, remember this is the active frame, our goroutine is operating here.
If you want to access anything that is not on this frame, you can only do that through a pointer.
From a syntax perspective, `u` is a value on the heap.
You get to keep your code simple by manipulating the `u` value through your value semantics; however, underneath the covers, the compiler knows well, I've to have that escape, so even though the syntax at the coding level is a value of type `user`, we'll convert that underneath to a pointer to be able to access to the heap value.
The syntax and the language is abstracting the details and that cognitive load from you, but it's important that you know that, but we're really working with `u` is truly a value on the heap that you get to work with as a value and not as a pointer, which historically you'd have to deal with yourself.

Yay, but there's a cost to this, right?
And that cost was an allocation.
This value now is out on the heap, and now the garbage collector has to manage it.
So, there's a guideline here that I want to provide because the ampersand is a very powerful readability operator, and you don't wanna walk away from it.
So, here is a general guideline.
I never want to use pointer semantics during construction.
I want to use value semantics during construction.
If you're going to assign that value that you're constructing to a variable so we can leverage the highest levels of readability in Go, we can show sharing down, we can show sharing up, you can do your own escape analysis.
The only time I want you to use pointer semantics on construction is if you're gonna do that on a `return`.
Okay, we're not assigning it to a variable, we're doing it on a `return`, or again, you're gonna do it inside a function call.
We're not assigning it to a variable.
We're doing it within the scope of a function call.
That's where that syntax makes sense.
But if your goal is to assign a value you're constructing to a variable, please, please use value semantic construction.
So, we want to make sure that we're using the right semantics and semantic consistency all of the time.
We don't want to construct values to variables using pointer semantics.
We want to leverage the ampersand in the right place.

#### Stack Growth

The operating system stack's about 1MB, and your Go stack is 2K.
What happens when you've got a Go routine that's making lots of function calls and eventually it runs out of stack space?
What's going to happen?
We can't just terminate the Go routine like we do with a thread, and we won't.
Remember, Go is about integrity first, it's about minimizing resources second.
We always want to try to reduce and minimize the amount of resources that we need.
We're all on cloud computing today, shared resources are important.
That 2K stack is a big part of that.
But when we run out of stack space, what we need now is to get a new stack.
This is something for me and it's very unique from a programming language, I'd never seen this before.
Basically, imagine that we had our stack, we had some value there, and imagine we were even sharing this value as we move down the call stack.
Eventually, we run out of stack space.
What Go does is it has what it called contiguous stacks.
What it's going to do is allocate a larger stack, 25% larger than the original one, and then, what it's got to do is copy all these frames back over, in this case, these pointers are relative, so they're very fast to fix.
But basically, that Go routine, during the function call, we're going to take a little bit of this latency hit on creating the larger stack, copying those frames over, and readjusting any of these pointers.
We do this in Go, we take this hit, one for integrity and two to minimize resources.
I told you nothing is free.
Integrity and minimizing resources come with a cost.
But this isn't something that's going to happen all of the time.
2K is usually more than enough for our stacks, because you usually don't go more than even like 10 function calls deep, you don't.
There's other optimizations the compiler can do to keep these frames very, very small.
This is a very good trade-off and balance.
But this can happen.
What this is telling us is, which is, again, really unique for me, is that values on your stack can potentially be moving around.
This is a whole new world.

#### Garbage Collection

So, the big thing about the garbage collector right now is that we call it a tri-color mark and sweep concurrent collector. 
ne of the big things is that it's not a compacting garbage collector, memory on our heap does not move around, which is getting interesting because memory on our stacks potentially are.
Once an allocation is made on the heap, it stays there until it gets swept away.
But the tri-color, is a mechanism.
We'll talk about how that is and that it runs concurrently is also very, very important.

### Constants

Constants have a parallel type system all to themselves.
And the realty also is that constants only exist in compile time.
That is really cool and much different.

Most of the time we think about constants as being read-only variables, and it's absolutely not the case in Go.

Now what's also interesting about constants is that they can be typed, and they can be untyped.
And when a constant is not typed, we consider it to be of a kind.
And what's special about constants of a kind is that they can be implicitly converted by the compiler.
```go
const ti int =  12345
const tf float64 = 3.141592
```
Now if you apply a type to a constant, like I'm doing on line 20 and 21 here, there's still constants.
They still only exist to compile time.
They still exist within that parallel type system, but we've handcuffed these constants.
We've now limited them to this particular level or precision, and they can no longer be complicity converted.
They're bound by all the laws of type.
So constants of a kind have a lot of flexibility, as opposed to constants of a type. 

### Functions

What's really special about Go is that your functions can return more than one value.

Now, this is really common in Go, to return both some sort of value and an error, and I want you to understand that Go doesn't have constructors.
So, this idea of a function getting executed during construction doesn't exist.
This is a good thing, remember we talked about in the very first lesson that readability is about not hiding the cost, and constructors are a big part of hiding the cost of things.
And so, in Go, what we have is what I call factory functions.
A function that is called that returns an initialized value, and without the function, that value could not have been initialized properly in any other way.
A lot of factory functions in Go do start with the name `New`, but I would consider `retrieveUser` here a factory function as well.
It is creating a value of type user, it's sharing it back up the call stack after it gets initialized for use.
So, a function like this is very, very common in Go.

You might see this happen a lot in Go.
You might see a combination if statement calling a function and then checking, let's say, an error, and moving on.
But this is something that's really unique in Go.
Something that I think is incredibly powerful.
Your `if` statement, your `for` statement, your `switch` statement.
These statements, they actually come with their own block of scope, something that we don't have in other languages.
In other words, the `u` variable that's being declared by this function call through the short variable declaration operator, this `u` variable is a unique variable to the `if` block.

Go has something that's called the blank identifier.
What this does is allows us to not declare a variable for a value when we're gonna be required to.

## Lesson 3: Data Structures

### Data-Oriented Design

Now, I've said a few times already in this training material that what data-oriented design is about is the understanding that every problem you solve is a data problem.
We are all data scientists at the end of the day.
Integrity comes from the data, our performance is going to be coming from the data, everything we do, our mental models, everything's going to be coming from the data.
If you don't understand the data you're working with you don't understand the problem you're trying to solve, because all problems are specific and unique to the data that you are working with, and data transformations are at the heart of everything we do, every function, every method, every encapsulation is really around the data transformations that we're working on.

I want to just keep reiterating this thing, that data-oriented design is about understanding the data, writing the code that we need, the algorithms that we need and eventually decoupling those algorithms that we have in the concrete to deal with the data changes.
Everything we must do must be focused around minimizing, simplifying, and reducing the amount of code we need to solve every problem.

### Arrays

#### Mechanical Sympathy

Also, small is fast, when we say small is fast, what we mean is, if the data you're working with is small enough to fit into the caches, and small enough to stay as close to the hardware thread as possible, you're also gonna see some better performance.

Now the link list sits somewhere in between.
We're probably getting cache line misses, because this data is not guaranteed to be on a predictable stride, but we're probably getting this data all on the same page.
This is why, when we talk about a two meg page, if you're dealing with a system, like a database, or something that's gonna be very, very large amounts of memory in data storage, those two meg pages can come in really, really handy, because the TLB means more data on a page, and your TLB cache will be better off for it.
So we're probably getting cache line misses on the link list, but we're not getting so many TLB misses, and we're falling somewhere in between, really interesting.

#### Semantics

But, the `for range` is a very special and powerful iterator in Go.
The reason why it's so powerful is because the `for range` comes with two different semantics.
Here we go again with semantics.
There are value semantics and pointer semantics associated with the `for range`.

### Slices

#### Declare and Length and Reference Types

So we just finished talking about arrays and I showed you a lot of mechanical sympathies around them and we started seeing more of the value on the pointer semantics around that.
But slices are the most important data structure in Go.
This is something that you must learn, you must master, you can't cheat on because all of the data you'll be working with or at least the majority of it should be and probably will be stored in slices.
This is your go-to data structure.
Are there times where the slice is not reasonable/practical to use?
Absolutely, but right now until you know that it's not reasonable/practical, this is the direction.
And a slice of values over a slice of pointers again when it is reasonable and practical.
Let's start off with some slice code right here.
```go
fruits := make([]string, 5)
fruits[0] = "Apple"
fruits[1] = "Orange"
fruits[2] = "Banana"
fruits[3] = "Grape"
fruits[4] = "Plum"

fmt.Println(fruits)
```
Now notice that I'm using the `make` function.This is the built in function in Go and what it does is it allows us to create three of the reference types that are also built into the language.
Up until now we've been using the built in types and we've been using our user defined or struct types.
Go has another class of types called reference types.
That is gonna be our slice, our maps, our channels, our interface values and our functions.
They're all reference types.
They're reference types because they are data structures that have a pointer.
They're also reference types because when any of these types are set to their zero value, if we create a variable of a slice or a map and they're set to their zero value they're considered to be `nil` in this language.
It's like a pointer set to its zero type is `nil`.
A string is actually very close to being a reference type.
The problem is that when a string is set to its zero value, it's not nil, it's empty so I can't really put it in that class.
But we're using the built in function `make` to create this slice and we're gonna use `make` when we already know ahead of time how much memory to allocate towards its backing data structure which is an array.
Now a slice is a three word or 24 byte data structure on our AMD64 architectures and it's very similar to the string where you get a pointer and you have the length of bytes.
In this case not bytes though, but the length here of five which means we're gonna have five strings.
So what we're gonna end up here now is an array of five elements, zero, one, two, three, four, and our pointer pointing to that so kinda similar to the string.
But we end up also having a third value.
That's why it's three words.
This is capacity.
And when the length is only set on the `make` call, then the capacity matches the length.
So here's the question.
What is the difference between length and capacity?
Length represents the total number of elements you can access from this pointer position and I want you to know as well that the `make` call is going to set all of this here in the backing array, set to its zero value state.
So length is the total number of elements that you can access from that pointer position.
Capacity is the total number of elements that exist in the backing array from that pointer position.
Capacity can be larger than length, but not the other way around.
Your length sets the bounds just like an array has a bound that's set at compile time.

And this is also really important.
I want you to notice that we're passing a value of this slice into the print call.
In other words the slice value just like the string is designed to be using value semantics.
As we're gonna learn our built in, our reference types are designed around value semantics.
They're designed to be kept on the stack.
Notice again I've got this slice value on the frame here.
Below us now on the frame below, the one for print line, we've got this string value and thanks to the pointer they share the backing array.
There's efficiency in sharing and we also have consistency in copy cost.
This array could be millions of strings, but as we pass it around the world these program boundaries the cost is always consistent.
It's 24 bytes.
How beautiful is this?
The only thing that has to end up on the heap if anything and in this case right now it doesn't and the only thing that have to end up on the heap is the thing being shared, the backing array.
This is beautiful, this is back balance and minimizing allocations and only allocating those things that need to which are the things that are gonna be shared.

Now let's take a look at a slice where the capacity's larger than the length.
Here I am again using the make call, but instead of just saying five I've said five, eight.
What I've now said here is I want a slice whose length is five, but also has a capacity of eight.
And what that basically means is that I'm going to have three more elements right in here, three more elements, one, two, three, five, six and seven, where we have extra capacity.
A capacity is for growth.
I want you to remember that.
Length is what we have access for today and capacity is for growth so we have the ability to grow this slice beyond that length of five here efficiently up to eight elements.
Capacity is there for growth, but efficient growth.

#### Appending Slices

So, there's two things that we do with slices all the time in Go code.
We append values to them and we take slices of them, which is where the name comes from.
Let's start with the idea of appending.
```go
var data []string
data := []string{}
```
There's another special type in Go.
It's called the empty struct.
And the empty struct is declared like this.
What's interesting about the empty struct is that it's a zero allocation type.
I can create a million values of this empty struct, and there would be zero allocation.
And that's because there's eight-byte value tied inside the run time, like a global variable, that this empty struct is referencing.
It means all million empty struct values would have the same address.
So, basically, this pointer points to the empty struct.
That's what it does to represent empty, really cool stuff.

Notice that `append`, notice the API for append, this is important to me, the API for append is really using value semantics.
I love `append` because append is what we call a value semantic mutation API.
Notice that `append` does mutate, but we're not using pointers, we're not sharing.
`append` gets its own copy of the slice value, mutates it and returns it back.
This is critical.
`append` is able to do mutations and isolation without causing side effects because it's leveraging a value semantic mutation API.
Plus, I told you, slices should be using value semantics.
We'll talk more about that I promise you when we get to methods, but it's using value semantics so the APIs have to respect it.

#### Taking Slices of Slices

What you can do is take any existing slice value, like we have here, and create a new slice value sharing the backing array for efficiency.
The idea again, is the only thing that has to be on the heap is the backing array.
If anything has to be on the heap, if it has to escape, the only thing that should ever escape is the backing array, the thing being shared.
Our slice values get to stay on our stack, so what this syntax is saying is take your slice and what we're going to do is say `[a, b]`, and what that `[a, b]` means is index `a` through index `b`, `a` through `b`.

Try to avoid these, but sometimes you have to do it and that's why Go gives you the built in function `copy`, and the built-in function `copy` will let you take the source of one slice and make it the destination of the other.
You could see here that I'm making a slice of the same length as our `slice1`.

#### Slice and References

I just wanna continue to iterate that when you're using pointer semantics, and we're mutating memory, we have to be aware of the side effects.
This is what makes programming difficult.
Functional programming says everything is value semantics, every piece of code, every function we call, operates on its own piece of data, but we don't get all the efficiencies we need to have the fastest programs that we want, and this is why GO is brilliant in giving you the responsibility of both value and pointer semantics, but we have to be aware of when we're using what, and when we're using the pointer semantics, we have to be aware of the mutations and the problems they can cause.

This is a side effect, these are the nasty bugs that are so hard to find, and so anytime we're working with pointer semantics, that's great, it's gonna give us levels of efficiency, right, we have to be careful there, but we also have to make sure that we're very clean with data updates, like with slices, and that our mutations are not going to cause problems, or the mutations are happening in the wrong place.
This is just a program to show you what a side effect really can look like in production-level code, and how careful we have to be when we're sharing these slices beyond program boundaries, especially if the call to append is going to happen.
And one of the things I do a lot is any time I see a call to append in code, I immediately put the breaks on the code review, because append can create some of the nastiest side effects you've ever seen.
So unless the code is decoding or unmarshalling, usually an append call is somewhat of a smell, and it's not necessarily wrong, but we're gonna have to double and triple check that we're not causing some sort of side effect with a backing array that's gonna get replaced.

#### Strings and Slices

Now, strings in Go are UTF-8 based. In fact, the entire language is UTF-8 based.
If you are not saving your source code as UTF-8 you're gonna have a real problem with those literal strings and raw strings that you're storing inside your code files.
Everything has to be UTF-8 encoding, all the way through.
Now, what's interesting about UTF-8 is that it's a three layer character set.
You have bytes at the very bottom, and really, we would always consider strings just to be bytes at the end of the day.
You've got bytes at the bottom, in the middle you have what are call code points.
And a code point is a 32-bit or 4-byte value.
And then, after code points, you have characters.
And the idea is that a code point is anywhere from one to four bytes.
And then a character is anywhere from one to multiple code points.
You have this, kind of like, n-tiered type of character set.

#### Range Mechanics

Again, what's beautiful about value semantics is that we're always operating on our own copy of the data.
That tends to keep us isolated, it tends to keep us safe, it tends to allow us to do mutations that can't affect other parts of our program.
Pointer semantics, however, means that we have shared access, which means we start mutating in the middle of an operation like this, we're going to cause ourselves pain.
I cannot stress this enough.
Now we still haven't learned about when to use values and when to use pointers, I've just been trying to show you that the mechanics of value pointer semantics are very, very real and they're really going to allow us to understand the behavior and the cost that our software has.

### Maps

Remember we had arrays, we just went through slices, and now we're gonna talk about maps.
And maps are just key value pairs.
Remember the built-in function `make` allows us to make slices.
Allows us to make maps.
We're gonna see it again later when we're talking about channels.
Now there's other interesting things about the maps and you have the ability to do literal construction.
```go
users := map[string]user{
        "Roy": {"Rob", "Roy"},
        "Ford": {"Henry", "Ford"},
        "Mouse": {"Mickey", "Mouse"},
        "Jackson": {"Michael", "Jackson"},
}
```
And instead of using `make` and then the assignment operators, what I'm doing now is using that literal construction.
Remember we saw literal construction with structs.
We saw it with our slices, we saw it with arrays.
Here it is again now with maps.
So this is gonna construct that map, allow it for use, and we're inside the literal form there, doing the key value pairs again with that colon there and the closing comma.
But notice, also, there's a built-in function called `delete`.
And `delete` will allow you to remove a key from the map.
Now there are restrictions on what a key can be.
Not a value; you can store anything you want in the map, it uses the empty interface.
But when it comes to the key, it has to be something that they can execute a hash against.
So always think about the key as being something that you can use inside of an `if` statement.
If you can compare the value, then you can use it as a key.

## Lesson 4: Decoupling

### Methods - Declare & Receiver Behavior

But here's the problem, when the data's changing and the problem is changing, we're changing our code, what we wanna try to do is minimize, minimize the cascading changes that are gonna rip through the entire code base, and this is where the decoupling comes in.
And so, we've been focusing down here on the concrete data.
What we now have to start dealing with is the decoupling, and how do you decouple?
Decoupling is gonna be a big part of what we do.
And decoupling is all done through behavior, and we're gonna start wanting to focus on behavior now.
Behavior is where we do design, architecture, and decoupling, and our real work, the real work we're doing is all down here in the concrete and the data.
And one of the those we're gonna wanna learn is we wanna start from the concrete and the data and move up this way.
Alright, so what we're gonna do is taking a look at some code here, and the code we're gonna look at is giving us the ability to allow data to have behavior.
I want you to think about this for a second.
To allow data to have behavior.
Now, Go has functions, we've been looking at functions since we started, but Go also has the concept of a method, and a method is a function, a function that has what we call a *receiver*.
So, let's take a look at what a method looks like in Go.
```go
import "fmt"

type user struct {
        name  string
        email string
}

func (u user) notify() {
        fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
        u.email = email
}
```
I always want those methods to come after the type.
Now, this is a method because what we have is a special parameter being defined.
It's a parameter, look at it as a parameter.
We call this parameter a *receiver*.
So, between the keyword `func` and the function name, we have a parameter, we call it the *receiver*, and now this is considered a method.
And it's these methods that allow data to have behavior.
Now, one of the things we have to learn as we continue to move forward, especially after the mechanics, is, when should a piece of data have behavior?
I want this to be the exception, not the rule.
And this is gonna be hard for you if you're coming from an object-oriented programming language because OOP wants you to always think about data as having the state, the data, and always having behavior.
Object-oriented programming wants you to try to say that everything in your program should resemble some sort of object, like in the real world, that has state and behavior.
And I don't think it's really the best way to go in Go.
This is not what we wanna do.
In fact, I'm gonna show you how Go is really separating state from behavior most of the time.
And if we do this, you're going to be able to write less code and things are going to be simple.
Functions should be your first choice until they're not reasonable or practical to do so.
This is where we're gonna get into method.
So, I still gotta teach you when you should take the exception and use a method.
Data drives the semantic, so once the semantic is chosen, either everything is value or everything is pointer, and the method and the code you write has to respect the semantic for the data.
Now, if you're used to programming languages like C++ and C# and Java, you actually deal with receivers.
You've heard of the `this` pointer before, you've heard of the `self` pointer.
Yeah, these are receivers, they're *pointer receivers* and they're named for you.
One of the things that Go does, which is brilliant, is they give you the choice of value and pointer semantics, value receivers, pointer receivers, but Go also allows you to name the receiver, and we never wanna use `this` and `self` when naming the receiver.
You want that receiver name to be short and sweet.
It should never even be more than three letters long.
This is at the highest level of context in your Go program.
So, our value semantic methods, like `notify`, means that the method is operating on its own copy when we make the call.
Our pointer semantic methods there, right?
Right there, `changeEmail` means that we are having shared access when we make the call.
There it is, we've got these two choices.
Now, while we continue to learn mechanics, you're gonna see code that has got a mix of value and pointer semantics.
Once we move to design, you will not see this anymore.
When it comes to calling a method, Go only cares that the data, remember, data, value, pointer, that the data is of type `user`.
It doesn't care that it's a value or a pointer of that data, all it cares is that we are working with `user` values in some form, which is brilliant because now when we make that call to `changeEmail`, what's gonna happen is Go's gonna be able to adjust to make the call.

### Methods - Value and Pointer Semantics

Engineering is all about choices, it's about cost, it's about knowing when to take those exceptions.
So these general guidelines are very, very good and they're gonna work for you most of the time and they're gonna keep you out of a lot of trouble.
Remember, semantic consistency is everything.

There are really three classes of types that we deal with, or three classes of data that we deal with that we've gotta make these decisions on.
If you are working with the built-in types, strings, numerics, and bool, you are to be using value semantics.
Fields in a struct, I don't wanna see pointers.
Everything should be value based.
Reference types, we're gonna follow the same rule.
Value semantics for reference types, I don't wanna see you take the address of a reference type.
There's one exception to this, however.
A slice and a map, you may take the address of a slice or a map only if you're sharing it down the call stack and to a function that's either named `decode` or `unmarshall`.
But not all data can leverage the value semantics, and when it comes to our struct types this is where things get interesting.
You've got to choose, if you're defining your own struct type, what semantic is going to be in play.
And if you're not sure what to use then we're gonna use those pointer semantics, okay?
And then if you're absolutely sure that we can use value semantics, guess what?
We want to use those value semantics.
They're really our first choice.
Pointer semantics are really the exception, when we need it.

You will see this throughout the entire standard library.
There is nothing random about code in the standard library.
It follows these semantic rules to a `T`, and there are very, very few exceptions, there, and you can go and search for this stuff yourself.
This is the key to having a code base that can be maintained by more than one person over time and can grow and maintain a consistency which is gonna, again, help reduce bugs and give everybody the ability to predict reasonably well how this code is gonna behave on the machine it's running on.

So to sum up. Built-in types, value semantics.
Reference types, value semantics, except for decoding and un-marshalling.
And then our struct types we have to make a choice, if you're not sure, you use the pointer semantics first.
I want you to remember that you can choose a semantic in the beginning, realize that you're wrong, and then just refactor that the other way.
We are working on mostly close-based source code systems, right?
Private source-code systems.
Unless you're doing open source, we have the right to break APIs.
We're gonna talk more about this during design.
I want to break APIs when we do things that are wrong.
If we're dealing with open source that's been versioned, then we can't break APIs, but we can always add new ones and deprecate old ones to make improvements.
We've got to constantly be reviewing the code we're writing, adjusting the code, refactoring that code and making it better.
So I don't wanna get hung up where we're like, stuck, because I don't know what semantic to choose.
Choose a semantic, be consistent, and if it's the wrong one, no problem.
We'll refactor later.

### Methods - Function/Method Variables

Now, one of the things I wanna make sure you understand is that methods are really made up.
They really don't exist.
All we have is state and all we have is functions.
Methods give us this syntactic sugar that a piece of data has behavior.
That's what it's there for, to give us syntactic sugar, this belief system that data has behavior, and there are times when it's very important for data to have behavior.
Again, we've got to learn when should data have behavior and when it shouldn't, and I want data having behavior to be the exception, not the rule.

All of the behavior is only still associated with data. Why is this?
One, because these methods are not declared inside the type, like you would see in a class-based system, they're declared outside the type.
This is Go, again, separating state from behavior.
We're not really wanting to combine it, we're keeping it separate.
Your syntax is also keeping it separated.

Functions in Go are values, they are literally typed values, and that means that we can pass a function by its name anywhere we want in our program as long as, again, type information is the same, the function's signature.
But here, I'm creating a variable.
I'm creating a level of indirection, if you think about it, to be able to access and call `displayName`.
So, there's a cost to decoupling in Go, and it's really important, we're gonna continue to say this.
When you decouple a piece of concrete data in Go, you're going to have the cost of indirection and allocation.
There's no way around it right now, no way around it.
And so, we have to make sure that if we're decoupling, that the cost of indirection and allocation is worth it.
We do not wanna blindly decouple code in Go.
We wanna do it because it's the right thing to do, it's the right engineering choice.
It's reasonable and practical.
Because this cost, we know, is great.
Remember, this is number two on our lack of performance list in our programs.
Allocation is number two, so we don't wanna take these allocations lightly, but if the decoupling is adding value, then I'll take this cost all day long 'cause if I can make sure that I can minimize cascading changes throughout a code base, if I can make sure that I can work with different pieces of concrete data without blowing up a code base, well, I'll take the cost of allocation all day long.
Then it's worth it.
Alright, so our methods are giving us the ability to give a piece of data behavior.

### Interfaces - Polymorphism

*Polymorphism* means that you write a certain program and it behaves differently depending upon the data that it operates on.
I wanna change that up a little bit and just say, *polymorphism* means that a piece of code changes its behavior depending on the concrete data that it is operating on.
This is critical, I keep talking to you about, concrete data drives everything.
Concrete data and semantics drive everything.
And here it is, we're talking about decoupling, which is focused on behavior, because what's driving it?
The data.
And so when should a piece of data have behavior?
Well, one good reason, one good practical reason is when we need to implement polymorphism and that polymorphism's gonna give us that levels of decoupling.
So that's when a piece of data should have behavior, when we need that decoupling and we wanna be able to process all the different types of concrete data with a single piece of code.
That's good.
Now, there may be other times when you wanna have data have behavior, if it's an API that has to be, what I would say, stateful.
Right, it's gotta maintain sort of state around the API.
But we gotta be very careful about API design when we do that.
We'll talk about those things as we go forward.

Okay, so what I wanna do is give you an example of polymorphism in Go and continue to drive the idea around this idea that concrete data is driving the polymorphism because the concrete data has behavior.
```go
type reader interface {
        read(b []byte) (int, error)
}

type file struct {
        name string
}

func (file) read(b []byet) (int, error) {
        s := "<rss><channel><title>Going Go Programming</title></channel></rss>"
        copy(b, s)
        return len(s), nil
}
```
But interfaces are not real.
They only define a method set of behavior.
They define a contract of behavior.
There's nothing real about an interface, nothing at all.
There's an implementation detail behind `r`, but from our programming model, `r` does not exist.
It is not real, there's nothing concrete about it.
This is gonna be an important concept for us to follow through with as we continue to look at this code.
And one of things you know we're looking at mechanics and we'll talk more about this.
It's very important that your interfaces define a behavior.
Now, Go is about convention over configuration.
The concrete type `file` now implements the `reader` interface using value semantics.
You see Go is about convention over configuration.
We do not configure an interface to the concrete type like you might see in other languages.
In other languages, you probably have to do something like this.
This is not Go, this is configuration.
This is going to limit us at the end of the day.
It actually does limit us and cause our software to have more lines of code.
Go is about less lines of code. Go is about being more productive.
And because of the static code analysis, it's gonna give us a lot of benefits as I'm gonna show you as we move away from mechanics and into design.
So we don't have to do this in Go.
We just have to declare the method like we're going on line 20.
And the compiler compile time can identify interface compliance, satisfaction.

### Interfaces - Storage by Value

### Embedding

```go
type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // Embedded type
	level string
}
```
We are embedding a `user` value inside of the `admin` value, but, for me, a little bit better way of thinking about this is to create this inner type and outer type relationship.
I want you to look at `user` as the inner type, and `admin` as the outer type, and when you do what we have is this concept of *inner type promotion*.
Everything related to the inner type can be promoted up to the outer type.
In other words, promotion allows access to those things from the inner type through the outer type.
So look, let's do our construction like we did before.
Construction's very tedious in Go.
Types have to be very explicit, but you can see here that the `user` type looks and feels like a field during construction.
It's really not a field.
We're just trying to access the `user` value that we've embedded directly, but we do it through the type's name.
And so, now what we're doing is initializing the inner value directly, but here's where inner type promotion really comes in.
We can still access the notified behavior through the inner type value directly, but because of inner type promotion, I can also access notify directly through the outer type value.
Look, it's the same exact output for both calls.
Inner type promotion is giving us a sense of type reuse, but we're gonna use it for much more than that.
I now have two implementations of the `notifier` interface, of the method `notify`.
One from the inner type, one from the outer type.
But when the outer type implements the same things as the inner type, then there is no promotion.
The outer type will override the inner type's implementation, now if you look at this.
Now when we call our polymorphic function with the address to the outer type, it is the outer type's implementation that gets called.
Same thing when we call notify against the outer type's value, but in this particular case, we can always still access the inner type's behavior through the inner type value directly like we just did there. It's all there, so when I said that inner type promotion can, I meant can, because only a compile time does the compiler really check, and only a compile time do those things that are being accessed promote up.

### Exporting

The last mechanic we need to learn before we can get into design here is exporting.
Exporting is kind of like the encapsulation piece of Go.
But it's not like your object-oriented programming languages around private and public.
It is a little bit different.
The basic unit of compilation in Go is a package.
A package in Go represents a folder.
I really wanna go deep into this later.
But encapsulation begins and ends really at the folder level.
Every folder represents a static library.
That is our smallest unit of compilation.
This folder also should represent the namespace for everything that's inside this folder or this package.
I also wanna see one file named after the package name; that's gonna help.
We'll talk more about that later.
Based on this basic unit of compilation, the idea is that symbols are either exported or available outside of the code in this folder, or un-exported and only available to the code in this folder or package.
Up until now, I've been very careful to try to keep everything lowercase because Go is all about convention over configuration.
This is one of the more brilliant conventions that Go has.
Notice that the folder name is `counters`.
Notice that the package directive name is `counters`.
You really wanna match the folder name with the package name because we import packages by their folder name, and if these two names don't match, you're gonna really confuse somebody because this is the namespace and that's the folder, and we expect these things to match.
All code inside this folder has to use the same package directive or the compiler's going to get upset.
What Go says is you just look at the first letter of anything that's been named, any identifier, and if it's starting with a capital letter, guess what?
It's exported;
it's available outside of this folder or this package.
If it's lowercase, it's unavailable or un-exported.
But I want you to understand it's really not private and public.
Private and public is about the data.
This is about the identifier.

## Lesson 5: Composition

### Grouping Types

We're going to start talking about some design related material now, and this is going to be around composition.
There's a lot of things to talk about around composition, and composition is something that you're really going to be focused on quite a bit as a Go developer, even much more than all of the concurrency multithreaded stuff.
What we want to start off with here is the concept of grouping, because grouping is going to be one of these things where I've really got to kind of break you down and build you back up.
Things are going to be a little different here in Go.
We can group everything by what we do.
Now, I also, at this point, I also want to be able to start reading code and identify smells.
This is very important.
When we see smells in code, it becomes also critical for our ability to be able to talk about, verbalize, explain, have the right language to help people understand what the smells are and how to fix it.
There's a lot of smells in this code as well when I go through it.
Look, just alone, when I see that `Animal` type right there, when I see that `Animal` type, I notice that this is really a generic type just being placed in the code for reusable state.
I'm really always concerned about that.
The bigger smell, too, is that I've defined a type called `Animal`, and yet we're not really using it.
I do not need a value of type `Animal` in this program in order for the program to be able to work.
I could get rid of `Animal`, embed those fields in `Cat` and `Dog`, and I wouldn't have to break a single method or function.
That, again, is a smell.
I only want types where we need values of that type.
I want to be very, very clear that to define a type, literally for reusable state, is a smell, there are exceptions to every rule.
But because we don't need an `Animal`, guess what?
We don't need the type.
From my perspective, it's type pollution.
Believe it or not, a little copying and pasting can go a very long way.
We've been taught these concepts of DRY, Do not Repeat Yourself.
But I think the cost of Do not Repeat Yourself in Go is worse than maybe some of these other languages you've worked on.
Coupling is always going to be a bigger cost than a little duplication.
We'll continue to see that as we move on.
Also, when we look at the `Speak` implementation for `Animal`, it is completely useless.
There's no value in it.
Think about what we've done.
I've added a type, I've added a method called `Speak`, these things need tests, which can't really be accurate.
They're not used.
This is code that we have to, now, increase our lines of code, increase our bug potential, increase our test cases, and it adds zero value to the software.
We really, really, really don't want to be polluting software like that.
I understand where this developer's coming from with the idea of grouping things by animal and creating this reasonable state.
But this is not how, at least, I want you to design and architect and write code in the future.
Then, how do we correct this program?
What are we going to tell this developer?
Well, simply what I'm going to tell the developer is, stop thinking about who you are.
Stop thinking about what cats and dogs are and start focusing on what cats and dogs do.
This is what's important, what cats and dogs do.
It's convention over configuration.
Let's do this instead.
Let's remove the `Animal` type, it is pure pollution as far as I'm concerned.
Let's just bring in an interface, the `Speaker` interface, that behavior.
Remember, interfaces should define behavior, not persons, places, and things.
I want the `Speaker` interface.
If I saw an `Animal` interface here, I'd be very, very sad.
That's not it.
Let's keep the concrete types as those persons, places, and things, and let's have the interface describe the behavior.
That's the convention we're looking for.
OK, so we've got the `Speaker` interface, one active behavior, `Speak`.
Then when what I'm going to do is a little copy and paste.
Yes, I'm going to copy those common fields into `Dog` and `Cat`.

### Decoupling - Part 1

A lot of you have been taught, to start with the interface, start with the behavior, try to figure out what those contracts are.
I don't want you to do that.
That is guessing.
I don't want you to do that.
Remember, the problem is solved in the concrete, not in the abstract.
So we need a concrete influenced, implementation solution first, in order to know how to decouple.
This is going to simplify your code, and allow you to focus on what's important, which is the problem, and remember, the problem is the data.
