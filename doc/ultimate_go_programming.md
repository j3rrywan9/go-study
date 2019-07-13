# Ultimate Go Programming

## Design Guidelines

## Language Syntax

### Pointers

#### Pass by Value

Let's start with the idea that everything in Go is pass by value.
When I say pass by value, what I mean is wizzywig.
What you see is what you get.
I wanna give you an initial example of what we mean by pass by value and this idea that the wizzywig allows us to truly understand the impact that our program is having.

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
