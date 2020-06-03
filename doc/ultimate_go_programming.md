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

## Design Guidelines

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

## Language Syntax

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

Constants are something that really have a much different feel and flavor to than your common variables.

Most of the time we think about constants as being read-only variables, and it's absolutely not the case in Go.

So constants are really, really powerful in Go.
Remember now that there's two types of constants, constants of a kind and constants of a type.
Your literal values in Go are constants of a kind, they're unnamed constants, constants of a kind can be implicitly converted by the compiler, which means that you can't really have enumerations in Go.
You're not gonna get those compiler protections.
Once a compiler is based on a type, then the full laws of type are gonna restrict its ability to be anything other than its particular precision. 

## Data Structures

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

But slices are the most important data structure in Go.
This is something that you must learn, you must master, you can't cheat on because all of the data you'll be working with or at least the majority of it should be and probably will be stored in slices.
