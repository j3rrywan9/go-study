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

#### Stack Growth
