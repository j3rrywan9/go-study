# Ultimate Go Programming, Second Edition

## Introduction

Welcome to the Ultimate Go Programming live lessons, second edition, yes it's the second time we're here and if you've seen the first edition, there's a lot more going on here.
My name is Bill Kennedy, I am a trainer in the Go space, I've been training companies for over four years, hundreds of companies and thousands of developers.
And this class is really about somebody who wants to learn this language inside and out.
I have taken four years of writing and thinking about Go and been able to put it together in this class.
I assume that you're already a software developer, I'm not gonna be teaching you basic syntax, I'm gonna be teaching what you need to know, the history, the mechanics, and the semantics of this language so you can learn how to read code and leverage Go's tooling like never before.
So in this video, we've got lots of lessons and I really, really need you to start from lesson one.
Please don't skip around, I build layers of foundation that are gonna be very important for you from beginning to end.
So we're gonna start with like design guidelines, I've gotta get your mind in place for what Go is, where Go is trying to take you, so you can appreciate everything I'm gonna show you moving forward.
And we are gonna talk about language syntax, and a lot of the language syntax we talk about will revolve around, again, the history of the language, the mechanics, how things work, the semantics, how things behave, this is so important.
So we're not just gonna be, we're not gonna really be talking about syntax like you think, I'm not gonna teach if statements and for statements, I'm gonna teach you the semantics about these things that are unique or very important in Go.
Our data structures, our arrays, our slices, and our maps, our core data structures in Go.
I'm not gonna just teach you how to use em, I'm gonna teach you why they're there, why Go only has these data structures, and the mechanical sympathies, and the data oriented design aspects behind it.
Now we always gotta learn decoupling, right?
We're gonna learn how to work in the concrete, our performance comes from the concrete, so our data structures, our language syntaxes, but decoupling is key.
Gotta be able to minimize those cascading changes throughout a code base.
So we're gonna learn the mechanics behind decoupling, the cost of decoupling, and a lot of the semantics there as well.
And we're gonna get into design, so we've done that, we've learned that, now we get into design, we learn how to leverage the beauty of composition in Go.
We're gonna start to design our software from the concrete up through the use of composition.
Error handling is everything.
We're gonna learn how important integrity is in Go, and that error handling piece is gonna teach us how to write API's that respect the people who are using our API's, we're gonna respect and give people enough context to make informed decisions.
Now I see a lot of companies lacking on the idea of package during a design.
How do we structure projects in Go?
How do we design packages?
We're gonna talk about that from a design standpoint.
I'm gonna try to give you as much guidelines as I can around packaging.
Then we're gonna move into the concurrency, which is where you're probably gonna wanna try to start.
I'm gonna try to hold you off again throughout the class.
Start from the beginning, get to the concurrency sections when they come up because we're gonna be using a lot of what we learned as we move in there.
So I'm gonna teach you Goroutines and how to manage Goroutines and how to deal with synchronization and orchestration, especially when we get into data races.
We'll talk a lot about synchronization in our data race lessons and then channels, we come back and start talking about orchestration.
So all of this concurrency stuff including special patterns that are gonna help you really summarize all of this stuff and show you the real power of Go.
And once we get through all the concurrency stuff, I'll show you how testing works in Go, unit tests, sub test, mocking, I'll show you all of that stuff.
I'll show you how cool benchmarking is in Go with some basic benchmarking there, validating what you're doing in terms of benchmarking, and then we're gonna have a ton of fun learning how to do profiling and tracing.
Every single lesson there I've just mentioned, sits on top of the other, so really take the time from beginning to end to walk through this material, there's a lot here, a lot here, I normally teach this material over three days and you get the opportunity to sit there in your home, wherever you are, and take your time to learn this stuff.
It's going to be very valuable for you if you're really looking to be a high-level performance oriented Go developer.
All right, that's everything, so let's get going.

## Lesson 1: Design Guidelines

Welcome to Lesson One, Design Guidelines.
Please take the time to listen to what I have to say here.
I'm gonna be setting up the stage for the history of where Go comes from.
We're gonna talk about a lot of quotes.
I'm gonna give you a lot of quotes from very important people, very interesting, from now to 40 years ago.
You're gonna understand why Go is the way it is, and what we're going to do in this class, and why I'm teaching it the way I teach it.
So the design guidelines here are very, very important for you to be able to set your mind, okay?
Set yourself up to have a very successful time going through this video.

### 1.1 Prepare Your Mind

We're gonna learn that lines of code really can affect how healthy a project is, how healthy a team is.
And I want to start moving away from being oppressed with projects with large amounts of code.

Now here's another thing that I want us to focus on and understand that Go is trying to solve.
We were taught to create these large abstractions in our code base to deal with change.
In fact, we've been taught to put abstractions on top of abstractions.
This is something that we also need to get rid of, because all this is doing is adding more code to our project, which is gonna be causing more pain.
I think one of the most beautiful things about Go is that we're gonna learn how we can create these very thin layers of decoupling and still achieve our ability to deal with change without all these extra generic layers of abstraction.
Now here's another one that a lot of people lose, and it's not fair.
The languages you may have been using could have been using a VM, like Java and C#, where there's a virtual machine there taking care of what we will call mechanical sympathy.
We're gonna learn mechanical sympathy in this class.
But what's important here is this, is that if performance matters, and it does, that's why we're here using Go, then it's the hardware that matters, because it is the hardware that executes instructions.
Your performance is gonna come from the hardware, nothing else.
So we either need to be sympathetic with the hardware or not.
The hardware is the platform at the end of the day, if performance matters, and we all know that it does.
And the last thing here, and again, I think it's lost because of the languages that we've been working with, is that engineering is about understanding cost.
Every decision you make comes with a cost.
There's a benefit for that decision and a cost, and if the benefit is not outweighing the cost, you shouldn't be taking the decision, shouldn't be making that decision.
So part of this class, a big part of this class is understanding that engineering is about understanding the costs you are taking.
And when you don't, you're not coding, you're not engineering.
You're hacking, and we can't have that today at all.
Now another thing that's very important, these are things that I believe Go is trying to solve.
These are problems that Go has seen, and we're trying to solve them with Go.
And this other one is that the idea is that you cannot throw more hardware at the problem.
Now not many of us have data centers anymore.
We're using the cloud.
And so I can apply that to the idea that you can't through more cloud computing at the problem anymore.
It's expensive stuff.
So one of the big things about Go is we're trying to write software that minimizes the use of resources, minimizes the use of that hardware, minimizing the use of cloud.
Why?
We'll keep those costs down.
We're on shared machines, making all these things that are gonna tie back in together into the language as we continue to learn it.
And we already know that you can't throw more developers at a problem.
At least historically it hasn't been the case.
How many times have you been on a team, and a manager, somebody in management comes and says, "well, you know, if I throw another developer on this project, could we get it done faster?"
And I'm always like, look, this isn't accounting.
It's not like I can throw an accountant on here and we're gonna move faster.
This is software development, and there are things that cause us to have pain as a team grows.
I think Go is solving a lot of these problems also around package-oriented design.
We're gonna talk about these things.And there are ways to structure, package projects together where we could put more projects, developers on a team and potentially get more work done.
But we have to know and think about this up front, and the language has kind of done that.
So throughout this class, throughout this class, throughout the 15 hours that we're gonna be together here, I need you to always be a champion for quality, efficiency, and simplicity.
I'm gonna try to teach you how to do that.
It's critically important that you always have a point of view and self-reflect about who you are as a developer.
But I don't want you to hold back on your point of view.
I want you to analyze and self-reflect
What is my philosophy here?
What is driving the engineering decisions that I'm making?
Am I improving as a developer?
It usually means that you're improving your point of view, that it's changing.
It's okay to change your mind on something.
That's growth, and we want to have that all the time.
So as we continue to focus on this class, I want you to open up your mind.
Look, technology is changing quickly.
We feel it and see it every day.
But our minds change slowly.
We resist to change.
And it's always easy to adopt new technology, but it's hard to adopt a new way of thinking.
And one of the things I've gotta do in this class is kind of break you down and build you back up.
There are things that I'm gonna be asking you not to do that you've been doing historically or been taught to do.
Go has different costs for things, and that means that you've got to make different decisions.
And it's not that what you were doing before was wrong.
It's just that it's not necessarily the best or the right thing to do when we are coding in Go.
And these are things that I'm gonna try to teach you over the time that we're gonna spend together in this class.
Now as we continue to talk about this, I want to focus on this idea of legacy software.
If you ever met me in a bar and I've been drinking a little whiskey, and you happen to tell me that you're a software developer, well, one of the very first things I'm gonna ask you is what is the legacy you're leaving behind?
This is a critical, critical question that you should be asking yourself.
Now I've got some really great historical quotes here.
I think history is critically important, mechanics after that, and semantic behavior.
We're gonna talk about all of these things throughout the class.
But I love this quote from Peter. "There are two kinds of software projects, those that fail and those that turn into legacy horrors."
This is great.
Every time I say it, I crack up a little bit.
But it's so true, right?
What is a project that fails?
A project that fails is a project that never makes it into production.
But a project that gets into production and is no longer able to be maintained, that's a legacy horror.
And if you really think about it, we've got 40 years of software in this country, around the world, that is really a legacy nightmare.
It is.
I can't tell you how many times I walk down the street and just wait for the lights to turn off.
I'm always saying that I don't want Hollywood anymore to write movies about earthquakes and asteroids.
I want a movie about legacy software systematically failing all over the planet and no one being able to do anything about it, because there aren't any unit tests to deal with it.
And I've got clients today, today that are writing legacy code out of the box, mainly because, one, they don't know how to do it any other way, and two, they create these incredibly short-term milestones for themselves which just really can't be met.
And at the end of the day, all they're doing is hacking.
Look, Chuck Moore said it best, "legacy software is an unappreciated but serious problem. It is gonna be the downfall of our civilization."
I really believe that.
But we today have opportunities to learn how to write software that isn't legacy out of the box.
Go is at the forefront of this.
We've just got to learn how to rethink how we do software development, restructure our software development, and keep this in the front of our minds at all time.
But I think one of my favorite quotes throughout this entire class in all the material comes from Sarah Mei.
If you don't know Sarah Mei, you need to look her up.
She is a powerhouse, a thought leader in our industry.
And though she spends most of her time coding in Ruby, I don't care.
What she talks about, and the things that she tells us applies to almost any engineering that we're doing in software development.
And she said this, and I love this quote.
"We think awful code is written by awful devs. But in reality, it's written by reasonable devs in awful circumstances."
And I want us to understand that, because all of us at some point are gonna be asked to look at a piece of code that we haven't written.
And when we look at that code, we might come up with a conclusion that I don't know what kind of drugs this developer was on, but they shouldn't have been sharing them, because this is out of control.
But I want us to step back and think, and say, whoa, whoa, whoa, whoa, whoa.
Let's not trash the developer who wrote this piece of code.
Let's come back and ask ourselves, what was the condition or situation they were in when they had to write this piece of software?
And now we're here to maintain it, and maybe we can even improve it.
And so I want us to all have empathy for everything that we're doing, because every environment is different.
Every situation is different.
And we're going to be asked also to maintain code that exists.
But as you continue to write new code, I want you to think about that person that has to maintain that code after you.
That's going to help us move away from these legacy issues.
Sarah Mei, brilliant quote.
Now if we really are gonna focus around not creating legacy code out of the box, for me the number one thing, the number one thing we have to keep in mind are mental models of the code base.
Mental models are everything.
And it's gonna be part of your daily refactoring of the code.
Mental models, number one.
Now Tom Love was the inventor of Objective-C, who said this.
And Tom Love works for the Department of Defense on very, very, very large projects.
But he said that, "let's imagine a project that's gonna end up with a million lines of code or more. The probability of that project being successful at least in the United States is about 50%."
This is a nasty thought that once projects get over a million lines of code, they're probably going to fail.
This is scary.
But why is Tom saying this?
He's saying it because he identified that a ream of copy paper is about 10,000 lines of code.
Now put that in your head.
I mean, we can all imagine what a ream of copy paper looks like.
I'm sure we've all touched a ream at one point of time in our lives.
That's 10,000 lines of code.
What Tom is saying is that a developer, the average developer really can't maintain a mental model of code beyond that 10,000.
When we talk about maintaining a mental model, I don't mean memorizing every line of code.
What we mean is that you know where everything in that code is.
If I ask you where this function is, if I ask you where certain things are, you can find it quickly.
You know how the flows work.
You know how variables are named.
You understand the code base intimately enough where if there is a problem, you can probably go to the code, identify where to start looking, and start reading that code very, very quickly.
That's the mental model.
This is what we're looking for.
And Go is going to allow us to do this, because we're gonna be able to write less code.
But this is really what happens.
If the average developer can only handle 10,000 lines of code, then once a project goes beyond 10,000 lines of code, you need a new developer on that project.
And when we start talking about code bases with a million lines of code, we're talking about 100 developers on that project if we really want to maintain our mental models and the health of the project.
And I don't know about you, but historically I've only ever worked on teams of about three or four people.
And I can tell you, it is difficult to keep three or four people rowing the boat in the same direction every day.
Imagine trying to do that with 100 people.
This is why these projects fail.
As the larger the code base gets, it's not even so much a technical issue anymore beyond the fact that you have to manage more people.
You have to maintain these mental models.
You probably don't have even one person that knows the full mental model of that code base.
It really becomes both a technical and a management issue.
So we really want to focus on lines of code and reduce it and to identify the health of our teams.
We really need to understand is there, that 10,000 to one developer balance.
Now cognitive load is really, really important.
Brian Kernighan said this best, "the hardest bugs are those where your mental model of the situation is just wrong, because you can't see the problem at all."
And everybody knows that debugging a program is twice as hard as writing it, so if you're clever when you write it, how are you gonna debug it?
People find this really interesting, but historically I've been writing software for 30 years.
Historically, I never allowed developers on my team to use a debugger without asking for permission.
This is primarily because when there is a problem in production, all you have is your mental model of the code base and the logs that you're writing.
But if these are not helping you find bugs, you're done, you're lost.
You're not gonna hook your debugger up into your production environment.
And I'm not saying that debuggers are necessarily bad.
They are a tool.
But they shouldn't be your first tool to finding the bugs.
They're great if you've got a brand-new code base and you want to understand how it flows, but I have a 20 minute rule, and that 20 minute rule would say leverage your mental model, leverage your logging to find that bug.
And if you can't after that, we'll turn it on.
But we've got a refactoring problem in front of us at this point, because we've lost it, right?
So I don't want us becoming attached to debuggers.
I want us to focus on our mental models and really also making sure that we create logs that have more signal to noise.
Because I love this quote, "debuggers don't remove bugs. They really don't. They just show them in slow motion."
At the end of the day, we've got to refactor that code cleanly so those bugs disappear.
But at the same time, if more bugs start to show, I don't need to waste time on the debugger.
Our mental models, our logging, it all helps.

### 1.2 Productivity versus Performance

This is a big one.
As an industry, if I asked people, for a long time, "Is performance your highest priority?"
They would say yes.
But the reality is that productivity has been our highest priority.
It has been in just the languages we've been designing, the languages we've been using, people wanna be productive and get things done, and we've kind of looked at performance, it's secondary, though.
If you ask anybody, they're too afraid to tell you that, because then you don't look like you're technical, or your chops are in place.
We put people on a pedestal.
We can get on stage and talk about how they can shave a few milliseconds or microseconds off of an algorithm, and I really kinda wanna move away from that.
Productivity really is, and should be, our highest priority, but we do need performance, it does matter.
And what I find interesting about our industry, is that for the last 30 years we have kind of been on the same path that Niklaus Wirth has said, back in 1987, he said that "The hope is that the progress in hardware is gonna cure our software ills."
In other words, he's saying that we as software developers have always said "We don't have to worry about performance, the hardware will come in and save us."
We've continued to kind of have this mantra for the last 30 years.
Back in 1987, after I graduated high school, I was on a plane to Italy and I sat next to a software developer and that software developer told me, this is 1987, all of us, right now, watching this video, would be Basic developers.
We'd all be coding in Basic today, because the hardware would be so fast we wouldn't need anything else.
The reality is that has not happened.
I mean, Henry, back in 2015, wrote, "The most amazing achievement of the computer software industry is its continuing cancellation of the steady and staggering gains made by the computer hardware industry."
We're still talking about this stuff, 30 years ago, remember I told you before, the hardware is the platform.
If performance matters, then we have to be sympathetic with the hardware, mechanical sympathy.
And look, you know, the hardware folks really haven't changed hardware in the last 15 plus years.
They really haven't.
No, we've had these multi-core processors, multiple hardware threads per core, and the reason why they're not doing it is because we haven't been able to take full advantage of what's there now.
We're starting to see now Intel talk about their 24, 36 core processors, and that's great and it sounds amazing, but if we don't engineer for this hardware, if we don't understand how it works, those processors are actually gonna slow us down not speed us up.
So we're gonna spend some time in this class talking about how caching works, and some of the mechanical sympathies the language drives through us and allows us, so we can get to these levels of performance that we want.
And hopefully take advantage of the hardware that's there, so that the hardware industry stops playing games, starts giving us more.
But that's on us, it's not on them.
But I think a core part of this class is gonna be this next quote from Brian Kernighan, and it starts, him talking about C.
He's asked a question about the C programming language, but I think you can replace C with Go, and it absolutely mirrors it exactly.
"Go is the best balance I've seen between power and expressiveness. You can do almost anything you want to do by programming fairly straightforwardly, and have a good mental model of what's going to happen on the machine, and you can predict reasonably well, and how quickly it's gonna run, and understand what's going on."
This is Go in a nutshell.
I'm going to teach you how to read code in this language, so that holds true not just for me, but for you.
And it's possible because Go's model is not a virtual machine, it is a real machine.
It is the real machine, and once you're coding against the real machine then you can do all of these things if you just know a little bit about the mechanics and the semantics of the language.
When I say mechanics, I'm talking about how things work, when I say semantics, I'm talking about how things behave.
And I'm gonna teach you throughout this class how to be able to do that, and that's going to drive you being able to make some better engineering choices.

### 1.3 Correctness versus Performance

So this is a big one for me, because we've also been trained as developers to kinda focus on performance, again, as a priority, and we tend to write code with the idea that this gonna be faster than that.
The reality is is that unless you've done a thousand times and you already know, you're guessing, and usually code that's focused on performance is clever code.
I wanna reverse the tables on this.
I don't wanna optimize for performance.
I wanna optimize for correctness when we're writing code, I wanna use our debuggers and our profiles, not really our debuggers, I apologize, our profiles, our ability to benchmark, our ability to trace, to identify where we may have a performance problem, and then fix the performance issues like that, still trying to maintain levels of correct or readable code.
Look, the performance of your software is gonna come from four places, or I should say the lack of performance is gonna come from four places.
First is latency, and it's gonna be latency on networking and IO, disk IO, those types of things.
That's going to be a real killer of performance, and those are things you probably want to focus on first.
Your next potential performance strain is going to be on memory allocations.
We're gonna talk about that in this class, and memory allocations can cause performance problems, and we'll see that when we start talking about the garbage collector, and the garbage collector's doing an amazing job, but we don't wanna walk away from it.
When we talk about performance in GO, it's not about is it the fastest, it's about is it fast enough, but garbage collection and memory allocations, that will be your number two.
Number three is going to be how you access data, how efficiently you access data.
We're gonna talk about that in this class as well.
And then, number four would be algorithm efficiencies, can we streamline an algorithm to take less steps than more, and we tend to, if we're optimizing performance, focus on that piece, and the reality is the hardware's gonna be so fast if we solve the first three problems, that we can have a little algorithm inefficiency, and still get more than we need in terms of performance, and now we're optimizing for correctness, and this is really where we wanna go.
And there's some great quotes by some amazing people in this area.
Wes Dyer, "Make it correct, make it clear, make it concise, make it fast."
I want us to listen to this history.
These are people who've already learned these lessons.
I don't wanna make the same mistakes.
Alright, JBD who works over at Google, and again, you gotta look her up, she's also a tremendous thought leader in this space.
Engineering's not about finding the perfect solutions, it's about understanding the trade-offs, we talked about this.
Engineering's about understanding the costs you're taking for the benefits that you need.
It also is about optimizing for correctness, knowing what we're doing in terms of our code.
And correctness of the implementation is the most important concern.
Listen what Al is saying here, "but there's no royal road", listen to Al again, what he's gonna say is that it takes several things, invariants, that's the engineering costs, testing and code reviews.
And testing and code reviews really require our ability to read and understand code.
Without our ability to read and understand code, we're really going to be lost.
There's a great quote that says that we are one of the few industries where we ask people how to write code before we teach them how to read, and this is really mind boggling when you start to think about it, because everything that we're gonna talk about, about correctness, is going to stem from our ability to read code first.
And this class is all about our ability to read code, to optimize for correctness, to be productive.
That's what we're trying to do, that's what Go is trying to do, and if we continue to look at quotes even from Bryan, right, the basic ideas here of good style, of correctness, they're fundamental to writing clear and simple code, and we've known this, he's known this for over 35 years.
I'm not telling you anything new, these are not my ideas, these are not my philosophies, I'm trying to bring to you industry leaders who are telling us over and over again correctness over performance when it comes to the code, and to do that, we've gotta be very good at reading code.
But I love what Jason Fried said here, maybe again, next to Sarah Mei, one of my favorite quotes here in the material, "Problems can usually be solved with mundane solutions. That means there's no glamorous work. You don't get to show off."
Okay, you build something, it works, you move on, okay?
And this is really important.
This is one of those quotes I want you to put in front of you at all times, because if you are a backend developer, like I am, and what you're looking to become if you're not already, then the idea is that your software should be working 24/7 without issues.
The reality is nobody should know your name if you're doing your job well.
If somebody knows your name, probably means your software's failing or causing problems.
You want people to know your name?
Become a frontend dev.
Frontend devs get the oohs and the aahs, but I promise you this, you're gonna spend two weeks working on a screen, killing yourself to get things in a certain location, certain words, and the first thing somebody's gonna do when they look at it is point something out that they don't like.
Aah, that used to drive me crazy.
I'm a back-end dev, okay?
I build systems to the point where nobody knows I even exist.
That's a good day in the office.

### 1.4 Code Reviews

Alright, since we're going to be focusing on reading code in this class, we've got to have a measuring stick, we've got to have priorities in reading code.
And I do these things in my head while I'm coding, we do these in code reviews, and this is also, from my perspective, Go's priority when we're reading and writing code.
Number one is integrity.
And if integrity isn't your number one, I'm counting right now, you can't work for me.
Integrity is everything.
Integrity is about being serious, about reliability, and it's really important.
In fact, I love asking this question.
This is my cell phone, you've got one, I have one, everybody you know has one, we can't live without these things.
So, here's an interesting question, if this cell phone stops working for the next hour, your phone, my phone, everybody's phone in the United States, it stops working for the next hour, how many people lose their lives?
I know it's a serious question, but the reality is that if this phone stops working for the next hour, the answer to that question is not zero, somebody's gonna lose their lives.
Our lives are so embedded in technology today that if things can go bad and people are gonna lose their lives, this is a very serious situation, this is why integrity has to be number one.
Do you realize that we are one of the very few industries where people's lives are at risk, and we're not licensed to do our job.
I mean, how long do you think it's going to last that we as software developers are not gonna need a license to do our jobs because at the end of the day everyone's life is at risk.
Now you might say "Bill, come on man, I'm in the gaming industry."
Or "I'm in some industry where people are having fun."
No, the reality is that, even with the mental illness crisis in this country, your software, your bug, your problem, may be the one that puts somebody over the edge, over the edge, and so I need you to constantly put into your mind that regardless of the industry, and regardless of the software right?
People's lives, their happiness, everything is being dependent on the technology that you're building, and we got to take integrity seriously.
Now integrity comes in two places.
There's a micro level and a macro level.
At the micro level, I want you to understand that every line of code you write, every line of code either does one of three things.
It allocates memory, reads that memory, and it writes to that memory.
What's really interesting is, all our code is doing is reading and writing to memory.
What's really interesting is all we are doing in our code is reading and writing to memory.
And yet all of those reads and writes cause everything around us to happen.
Lights are on, you're watching this video, phones are ringing.
It's all reads and writes.
What's even crazier is we're really just reading and writing integers at the end of the day to really mind-blowing stuff.
So at a micro level, integrity comes from the fact that those reads and writes are always accurate consistent and efficient.
When they're not, then we have an integrity issue.
We've got corruption, software now has to shut down, because if it doesn't, we're gonna cause more damage than good.
But there's also integrity at a macro level.
Want you to understand that really what we do every day is solve data transformation problems.
Data-oriented design is gonna be a big part of this class.
We work with data.
And data or InDesign says that if you don't understand the data, you do not understand the problem that you're working on, because the problem is the data.
And our functions and our methods, and these encapsulations that we write are all about the data transformations that we have to perform.
And so if your data transformations are not accurate, consistent and efficient, you also have integrity issues.
And we have to apply this idea of data-oriented design, Go is very focused on data-oriented design.
I'm gonna show you this in the language as opposed to object-oriented design.
Data-oriented design is gonna be very big here, and it's again all about our data transformations being accurate, consistent and efficient.
You know, one of the things I like to do is, I have a friend who's in the data science world, and I always say to him, "You know man, you guys are a trip over there you know, you guys came up with two funny words, a machine learning and AI, and suddenly you guys created this whole kind of software industry around, just data science."
I'm like "Dude we've been doing data science, machine learning and AI for 30 years, we're even building predictive algorithms."
Just because we didn't you know, formalize the models and all that stuff, doesn't mean we haven't been doing, everybody is.
What I want you to have in your head is that you are a data scientist.
If you are a software developer, you are a data scientist because you're working with data, you're transforming data, and it's what you do on a daily basis.
If you don't understand the data, you don't understand the problem.
If you don't understand the problem, how can you write a single line of code to solve it.
Now, integrity is going to really come in two places from a macro level, and things that we're gonna focus on in this class as well.
But I already told you before that lines of code is critical, the amount of lines of code in your project will determine the health of your team, and the health of your project.
I said before that the average developer cannot handle more than 10,000 lines of code, cannot maintain a mental model of more than 10,000 lines of code.
Once projects get larger than that, we really need more developers, and we keep focusing this idea that we need to write less code.
And there was a study done where they identified that for every thousand lines of code that's written, the average developer will produce 15 to 50 bugs.
Now I wanna put ourselves in the worst developer category which basically means that for every 20 lines of code you write and I write, we've added a bug to the software whether we like it or not.
Start putting that in your head.
So when you start thinking about let me add setters, and let me add getters, and let me add abstractions on top of abstractions, really what you're saying is let me throw more code at this project, and therefore let me throw more bugs at this project.
We need to take this idea in the reverse, less is always more.
And what's interesting is, when we have more code than we need right?
Then we've got more places for bugs to hide, and it also means you've got more incomplete tests because you've got to write that many more tests to deal with that many more lines of code, and eventually a codebase will collapse on itself, and instead of you maintaining your frameworks and your patterns, you start to just hack, and that's when legacy becomes a real threat to the project.

Now another big part of integrity outside of just writing less code is also dealing with error handling.
An error handling is a big big part of integrity.
Here's the reality.
Anybody could write a software that works when life is good.
A six-year-old can Google today and find code that works when life is good.
Engineering is not about when life is good, engineering is about when life is bad, and things are happening, and your software has to be able to maintain stability, and not add to the crisis, but try to recover from it.
Error handling is what we do every day, the majority of the code you write is focused on failure, not the good case.
And if that's not happening today, I need you to reevaluate what's going on.
I love the way Go deals with error handling, it was how I've been dealing with error handling for 30 years, even when I had exception handling in languages like C++ and C#.
I used it really just to deal with the unhandled stuff.
I would do very much what we're doing in Go today, and return errors out or codes out, and handle the errors right then and there, and I love the way error handling is done in Go.
We're gonna talk about these things when we get to the error handling section during design.
But there was a study done on some very critical software.
In fact Redis is on this list.
And how many lives today are dependent on Redis?
I mean there's a tremendous amount.
And they identify that out of these projects they were looking at, and the critical failures that caused this software to either go down or have corruption.
92% of those problems could have been avoided, 92% if the error handling was better, was more focused.
92%, see error handling is really really critical if integrity is important.
So we're gonna focus on that during design as well.
We're gonna be constantly focused on less code and error handling as we get into design.

The next step here on our code reviews is readability.
Readability is number two.
Doesn't trump integrity, nothing trumps integrity, you do in a code review and you find an integrity issue, that has to be fixed.
But readability is number two.
We talked about optimizing for correctness which is about readability.
Now readability means that we're going to structure our software in a way that's comprehensible.
And there's two aspects of readability that I wannna focus on.
The first one is the average developer.
The average developer on your team should be able to understand the full mental model of the codebase that we're working on, should be able to read every line of code, have a full knowledge of the mental models, and should be able to fix almost any bug, in fact every bug that comes across the codebase.
Now what I love doing is asking you to think about this.
I want you to think about you, and I want you to think about your team.
Have that in your head right now. Now ask yourself, am I the average developer on my team?
Do I comprehend the entire codebase we're working on? Do I have a strong mental model of what's going on?
Can I solve every single bug or problem that comes across this codebase?
The answer is no, then you're not the average developer.
Now I want you to think yourself for a second, I'm I less than average or I'm I more than average?
Now the more than average developers really have a burden on themselves because you have to have responsibility to not write code that is clever.
If you're less than average, you have the responsibility to come back up.
Now I want you to evaluate everybody on your team, who is average, who is less than average, who is more than average.
Are those less than average people working hard to come up to speed?
And are you helping them?
Or are you a hindrance?
As the average developer you have responsibility to bring them up.
Are you or who else on your team is more than average?
Is that person writing really clever code?
Have you taken the responsibility to tell them that their code is too clever?
And are they taking responsibility to not do that?
Look, you put me on a crypto team, I am less than average, I've got to come up.
You put me on a team building business applications, Web APIs, I'm gonna be above average, I've got to make sure my code comes down.
And if you're hiring somebody, and you don't know whether they're average less or more, you're really going to cause problems on your team dynamic.
If you're hiring people who are less than average, do you have the time to bring them up to speed?
The answer is probably no.
But if you do take the time, think about how an amazing developer they're gonna be for your team.
If you're hiring people who are more or above-average, do you have time to teach them how to be less clever?
Are they willing to be less clever?
That's going to cause more problems than probably the developers that are less than average.
I need you to do these evaluation for yourself, for your team, and when you're hiring, this is really important stuff.
But that's subjective right?
And it's based on the team, and it's based on the software you're building.
And you always wanna make sure that team dynamic is everybody knows who they are, what their responsibilities are, the codebase is clean and healthy with those mental models.
But that's subjective.
Let's talk about the piece that's not subjective.
What readability also means is not hiding the cost of the code you're writing, not hiding the cost.
Now, this one is hard for me to explain.
So, what I'm gonna use is a piece of code, it's written in C++ to show you what I mean about not hiding cost.
And even if you've never seen C++ before, I promise you, you'll be able to read this code.
So, let's take a look at this code that's written in C++.
What you see here on line seven is a type named `foo`.
Now, what's special about `foo` is that it has these features that you find in object-oriented programming languages like C++.
On line 8 there's a constructor, on line 9 there's the copy constructor, the move constructor, the destructor.
And if you're an object or a programming language developer, you know what these do right?
Constructors get called when a value is constructed, destructed.
If you're making a copy, if you're moving something, we call these features, because we're able to kind of tuck away code when these things happen in our codebase.
And you can also see on line 13, I've operator, overloaded the assignment, which means anytime we use the assignment operator against the value of type `foo`, this code is gonna get called, we call these features.
Now look on line 16.
On line 16, what we have is a function named `f1`, and `f1` creates an object of type `foo`, makes a copy of it on the return and brings that copy back out.
The caller gets a copy, it's own copy of the `foo` object being created on line 17.
Brilliant, I've shown you 18 lines of code I haven't hidden any of it from you, and this is what I wanna ask.
Look on line 23 and 27.
This is super interesting.
On line 23, we're calling the `f1` function which creates an object of type `foo` and returns a copy of it back out and it assigns it to a local variable `foo1`.
On line 27, we call `f1` again, the same function we call again, reassigning to our `foo1` variable.
So here's the question.
Who can look at this code right now, and I've shown you all of it, but let's just focus on the two lines, they're 23 and 27.
How many objects of type `foo` get created by the time this program runs and completes?
Does the copy constructor get called and when?
Does the move constructor get called and when?
And does the assignment operator ever get called and when?
Here's the reality is that you have no idea by looking on line 23 and 27 what the full scope of behavior is on that call.
And by the way, you don't ask me what version of C++ we're using, the behavior changes, the semantics change, also depending on what version of the language you're using.
This is a nightmare.
The fact that I cannot look on line 23 and get an understanding of how that code's gonna behave, how it's going to perform on the machine, and what my costs are is a nightmare.
This is why Go doesn't have constructors and destructors, and these features operate over features because they hide cost, they don't help us.
Now here's the reality.
When I run this code on the latest version of C++, you're gonna see that the first call to `foo` optimizes the copy away.
On line 23, the optimization by the compiler says, we don't need to make a copy, we're creating this object for the first time won't reassign it.
But on the next call, line 27, even though calling `f1` again and reassigning it, now constructors, assignment overloads, copy constructors, two objects were made.
Oh my God, the cost difference on line 23 to 27 is great.
We cannot write code like this, these features hide cost.
Go doesn't have it because Go doesn't want you to hide cost.
And when I see code that's hiding the cost or the impact, it's gonna be having on the machine.
Guess what, we have to fix it, we have to walk away from it.
So Go doesn't have those features.
And it's very very important that we remember that we are focusing on the real machine, that way we get to be able to write code that doesn't hide the cost, and we give that developer at every moment in time the ability to understand the cost they're taking and the impact it's going to have.
And I love this quote from Peter Bourgon, and he's another gentleman whose big part of the Go community, a thought leader, somebody should be looking up, somebody should be watching talks from.
"Making things easy to do is a false economy. Focus on making things easy to understand, and the rest will follow."

Now simplicity comes after readability.
Integrity, readability and simplicity.
Simplicity is interesting because what simplicity says is, that we wanna hide complexity.
Remember readability is about hiding, not hiding cost.
Simplicity's about hiding complexity, and our encapsulations right?
Our functions and our methods which are trying to read, reduce the boilerplate code, is trying to simplify things if those encapsulations cause our ability no longer understand the impact our code is having, we can't have it.
Simplicity is not something you do day one, it is something that you refactor into.
And one of the biggest things that I can tell you about simplicity is what Dijkstra said.
You know, it's a great virtue, it requires hard work to achieve, education to appreciate it.
But complexity sells.
Well you got to get away from that idea that complexity sells, and we've got to refactor into simplicity as we go.
Go is an amazing language as it comes to simplicity engineering.
I'm gonna try to point those things out to you as we go.
But there's two more big parts of simplicity here quotes that I wanna say again.
Again we're trying to hide complexity, but we're gonna focus on this in the class really well.
And Dijkstra said it again, the purpose of abstraction is not to be vague, but to create a new semantic level where one can be absolutely precise.
We're gonna talk about at some point, precise APIs, and how powerful they are in a language like Go.
Where we have a compiler, and have the ability to find and identify bugs at compile time as a opposed to production.
And I'm gonna be asking you to kind of throw away some of the API design that you've been taught around, object-oriented programming here, and really focus on the cost that that has over the benefits of this idea of precise API design.
And Joe Beda whose company called Heptio, great company, said this, and I love this also around simplicity, computing is all about abstractions. Those below are just details, and those above are limiting complicated crazy town.
So anytime I look at a function or a method, and I'm seeing that this is an abstraction, right?
We're seeing encapsulation, I wanna make sure that these encapsulations are not just testable, I wanna make sure that they're providing a new semantic precise level of understanding that is not hiding cost okay?
But hiding complexity.

## Lesson 2: Language Syntax

This is where we begin to start learning how to read code in Go with the understanding that every decision you make comes with a cost.
Nothing is free, what benefit are you taking for the cost that you're taking?
And this is where we start to begin to look at the cost of things.
We're gonna learn how pointers work and memory works, which is so critical if you really want to understand the cost of things.
We're gonna learn how important it is to really understand how things work in terms of memory.
We're gonna get through all of that and language syntax that's important for our ability to be able to predict reasonably well how code is gonna run on the machine.

### 2.1 Variables

Variables, all right, so in this section, we're gonna talk about variables, but trust me, this material assumes that you are already a software developer, you know what variables are.
I'm not teaching you variables from the very beginning.
They are just the things here that I need to talk about to help us develop a base knowledge of material as we continue to grow in complexity in our ability to read and write code in Go, so I really hope you stick out this section.
I'm not teaching you what a variable is, but there are some very important things here.
So we're gonna scroll down now into this area where we have the code review.
We'll be using the Go Playground for a lot of this.

Now, one of the things that I want to talk to you about is type.
Type is everything.
Type is life, and I promise you one day, if I get a little too drunk, I'm gonna go to a tattoo parlor and I'm gonna write type life right here in my knuckles, that's how important it is.
Actually, if that ever happens, you can just laugh at me for the rest of your life, but type is everything.
Type is life.
Without type, we can't really have integrity and we can't really start to understand the cost of the decisions we are making in our code.
I mean, start with this right here.
Our basic unit of memory that we're gonna be working with in our programming model is a byte, and a byte is 8 bit, so now I've just drawn a byte right now here on the board with a particular bit pattern.
Now, what I would like to ask you is, what is the value of this byte?
What does this byte, what is the value I've just stored inside this box?
The reality is you cannot answer this question unless I give you the type information, unless you know what this bit pattern represents.
Now, if I tell you that this represents an integer, now you can tell me, well, Bill, that represents the number 10, brilliant.
But if I said nah, it doesn't represent an integer, it represents a color, now you gotta say, Bill, what is number 10 on your color chart?
And I might say to you, oh, yeah, yeah, no, no, that's red.
So you could see that this bit pattern can represent an infinite number of things, it's all based on what the type information is.
Now, type provides us two pieces of information, size, how much memory, how many bytes of memory we need to read or write at any given time, and its representation, what it represents, and Go is not being novel in many many cases that you really already know the syntax of this language, I don't have to really teach too much syntax, and just like other programming languages, Go has the built-in types, your numerics, your string, and your bool.
What you are seeing on lines 15 through 18 in the code is just that, those built-in types being used, and if you look at the name of the types, start at 17, I love looking at 17, the name of the type `float64` tells us two things, look, I want you to look at some of this early material differently than you ever have before, but I'm not necessarily teaching you anything new, I wanna give you a different perspective on what's happening here.
So again, we can have this foundational material and understanding as things get more complex.
When you look at the name `float64`, it's really interesting because it tells us both parts of the type information, `float64`, "64" tells us that it's an eight byte, 64 bit value that's giving us the cost in terms of memory footprint, and "float" tells us it's an IEEE 754 binary decimal.
You see how the name in this case is giving us both pieces of type information.
We know a bool is one byte, we only need one bit for that byte, on or off, true or false, and then when you get to `int`, it's really interesting.
There is precision based integers, `int8` and `int16` and `int32`, and `int64`.
But if I saw you using a precision based integer and it wasn't obvious to me why, that would raise a flag to me during a code review.
Now, when I say raising a flag, it doesn't mean it's wrong, it is not obvious to me why you're making a particular choice and I would wanna talk about it, and there are times when you need a precision based integer if we are using the atomic instructions or if you're doing something that has to be the same size across multiple platforms, but most of the time, we're just gonna use `int`, which can seem confusing because all this type name is giving us is half of the type information, all it's telling us is that whatever number of bytes we're going to allocate, they only represent an integer value.
So when the question comes, if we don't understand size, we don't understand cost, so how much memory is allocated for an integer?
Well, this is all gonna be based on the architecture that we are building for.
Now, most of you, if you type Go space version into your command line, you're gonna see that you're using AMD64 as your architecture.
Some of you might be on an ARM processor, but for the most part, you're probably all using AMD64, which tells us that we are on a 64-bit architecture.
Now, we stick with AMD64 for a second.
What that's gonna mean from a Go perspective is that since we're gonna work on a 64-bit architecture, that means that our pointer size or the address size is gonna be 64 bits or eight bytes.
And the way Go kinda works is this.
If we look at what the address size is gonna be for the architecture that we're on, AMD64, 64 bit, eight bytes.
Then we say, okay, our generic word size, when you hear me use the term word throughout the class, we're talking about a value that can change size depending on the architecture, so our word size is gonna match the architecture, that's 64 bits or eight bytes, and then Go says, if your address and your word size is 64 bits or eight bytes, let's just make the integer follow.
That will be 64 bits or eight bytes.
All three are always going to be the same.
Now, this gets interesting because you're seeing code on the Playground, the Playground is a single threaded environment, when we get to the concurrency stuff, we won't really be able to use it.
But what's also interesting is the Playground is an amd64p32 architecture, p32, it's a native client architecture.
What that means is that addresses are 32 bits on this architecture.
For all intents and purposes, it's a 32-bit architecture because our address scheme is four bytes or 32 bits, which now means what, addresses are four byte, 32 bits, word size is four byte, 32 bits, therefore our integer's four bytes or 32 bits.
So we see that our integer can change size depending on architecture, and there is mechanical sympathies for this.
The hardware or the architecture is basically saying that four byte integers is gonna be more efficient than eight byte integers on our 64 bit, we're saying eight byte integers are more efficient, right, so we're trying to gain mechanical sympathies for the platform that we are building on when it comes to the address, the word size and the integer, and that's why we want to use int and not a precision based integer unless it's something very specific like we have to and it's very obvious we wanna do that.

Now, there is another concept here, I would call it one of Go's wars.
There is a war in Go and it's that there's too many to declare variables and to create values in Go, and we'll talk about this early on in the class, but what you're noticing here is I'm using the keyword `var` to declare these variables.
Now, there is a concept called zero values.
Zero value's very very important, and it's an integrity play in Go, and I want us to understand that, again, the cost of integrity is performance, so this little performance hit will never show up in a profile, I mean machine's very very fast, but we want it.
And what zero value says is all memory that we allocate gets initialized at least to its zero value state.
In other words, any time we allocate a byte of memory, if we don't pre-initialize it with some value, and we're gonna run electrons through the machine and make sure that we set it all to zero.
This is going to help with a huge number of bugs that we found in the past where a program starts up, it allocates a variable, it doesn't initialize it, the bit pattern is reasonable enough and the program starts running, and it's running, but it's also in a corrupted state, so zero value's an integrity play and it's very very important.
One of the things we're gonna do in this code base, one of the things I follow, is that if I'm going to declare a variable, create some value like this, declare a variable and it's gonna be set to its zero value state, I'm gonna use `var`.
`var` is that kind of readability marker that we are declaring and initializing to zero value, and `var` gives us zero value 100% of the time in this language, and I'll show you some examples later on where if you use, if you're not using `var`, you're not getting zero value, and there's one more type here and it's called `string`.
And string's a really kind of made up datatype for any language, and there's different ways to implement strings.
Go has what I would call a unique way, I had never seen it before up until Go.
Now, when we look at on line 16 and we say this `var b string`, what we're seeing is a string being declared and created setting it to its zero value which is what we call an empty string, and so an empty string in Go is going to look something like this.
`b` is going to be a two word data structure, strings are two words.
Notice I'm using the term word because the size of a string will change depending on the architecture, on the Playground the word is gonna be two, four byte words, so that's eight bytes, on our 64 bit, it'll be 16 two eight byte words, there it is, where the first word is a pointer, we're gonna talk about pointers and the next word is the number of bytes.
Since this is an empty string, we don't have any pointer to a backing array yet, and there are no bytes.
So a word is a two word data value, eight to 16 bytes of depending upon architecture.
Again, I want to talk about cost, engineering cost, so we can predict reasonably well how things are going to run, how things are gonna happen in memory, so type is important, the size of things can be important, we really want to understand the cost of the things we're taking.
Now, you're probably not gonna always want to create, construct, and initialize to zero value.
There are gonna be times where you wanna pre-initialize something, and this is where the short variable declaration operator comes from, that is the `:=` that you see here on line 27.
A lot of people think that this syntax from the language comes from C, but actually, the Pascal programming language had a large influence also in the syntax of this language.
This is one of those places where Pascal comes from.
Now, this operator is really a productivity operator, it's allowing us to declare and initialize at the same time, I don't want to use the concept that that Go has duck typing, it doesn't.
It's really struct-based typing, but in this particular case, I want us to always look at the short variable declaration operator as a productivity operator, so `aa` is 10, `bb` is string, `cc` is a float, and `dd` is bool, and it's based on the fact that the compiler knows what type of value is sitting on the other side of that short variable declaration.
And if we go back to "hello" for a second, what "hello" would look like from a string now, is an array of five bytes, right, H-E-L-L-O.
Our pointer would point to that, and we would say we have five bytes right there.
We're gonna go deeper into this stuff as the class goes on, but I really believe in visuals, being able to visualize code and early on in my career, I used to keep a piece of paper at my desk with a pen and I would draw these things out.
So what you're seeing me draw on the board is really my mental models, my visual mental models of code and how things go and you're gonna see a lot of that throughout this class and I really think it helps and it's something that it might help you as well.
So we've got the short variable declaration operator when it's something that's gonna not be zero value.
Now, you might see code like this, and I could see this.
I'm using the short variable declaration operator to assign or declare an integer and set it a zero value.
There's nothing necessarily wrong with this.
A big part of what we are gonna learn about writing code and mental models is consistency.
If you wanna do this, then be consistent.
I wouldn't.
I would use `var` for zero value, because I think the readability marker of `var` is just way too strong to walk away, and if I was doing a code review, I would ask us to switch that to `var`, but it's not for any other reason for readability and consistency and readability.
So I don't want to do that, you won't see that in this code base, but that's up to you.
As long as everybody on your team is doing it the same way, then it's going to be okay.

Now, there's one more concept here that I want to talk about, and it's called conversion.
Go doesn't have casting, it has conversion. And what conversion means really is that we may be taking a cost, a memory cost, as we convert values from one type to the other.
If you've never heard of casting before, what casting has done traditionally and it's really a part of helping with performance, is saying this following thing.
Let's say that we allocated a one byte integer.
There it is, there is my one byte integer.
Let's say for some silly crazy reason, I decide that I really want `a`, not to represent a one byte integer, but a four byte integer.
Casting will allow me to do something like this, where I could tell the compiler, look, you know and I know that `a` is an `int8` or one byte integer, but casting let's us pretend that what that memory really is a 4-byte integer, and the compiler trusting us will just say, okay, and suddenly now, if I'm casting `a` from one bytes to four bytes, I have the ability to read and write memory across those four bytes from this particular location.
I could be potentially corrupting a lot of memory here.
Now, this is a silly example because really, where casting traditionally comes in is kind of two places.
One, if you're dealing with data coming over the wire and you've got large number of bytes, you probably would like blindly copy those bytes somewhere in memory and then you would cast or overlay a structure on top of it, right?
And that's gonna be very very very fast.
You're just gonna say, hey, those bytes over there, those 20 bytes starting at that address location, they really represent this structure, and then that's gonna let you, because without type, you can't read and write to memory, it's gonna let us do that, but if you're off by one byte, then we've got real problems, so casting comes with the idea that if you're off, or seeing that one byte is you overlay that struct, now you are reading and writing the bytes you didn't.
And that's a real real problem with casting, so Go says, look, integrity is number one.
We care about integrity as our highest priority, so you can use the `unsafe` package to do some casting, what we'd rather have is conversion, what Go would say here is we don't wanna even set up this scenario, so if you really want `a` to be four bytes, then we're gonna convert `a` into a new value that will be four bytes and maybe, and in this particular case, even have to give it a new variable name, in this case in the sample, we used `aaa`.
But the idea of conversion over casting is again, is an integrated play.
There could be a cost of new memory but we always rather be safe than sorry.
So this is what I wanna share with you in the variable section, you already know what a variable is, but to sum up here, okay, type is everything.
Type is life.
It gives us two pieces of information, size of the memory that we're gonna be allocating and reading and writing and what that memory represents, and without the type information, we will have chaos.
It gives us the levels of integrity we need, right?
We have the idea of zero value in Go, and I wanna use the keyword `var` for that.
All memory's initialized at least to a zero value state.
We can use the short variable declaration operator when we are initializing something to it other than a zero value state.
There is exceptions to everything, and part of engineering is knowing when to take that exception, but these are the guidelines we're gonna be following, and we have conversion over casting, again, it's an integrity play to keep our software, again, and our data and our memory safe.

## Lesson 4: Decoupling

You are not allowed to start here, I really need you to have looked at those first three lessons before we get here, because everything happens in the concrete, and that's what decoupling is about, decoupling the concrete.
So to appreciate decoupling mechanics and semantics, we've gotta know and learn our concrete.
So in this section, we're gonna learn how we can create thin layers of decoupling that are precise.
We're also gonna learn the cost and the impact that decoupling has on your Go programs, because what I want is to minimize pollution and be able to maximize the advantages that decoupling gives us in our software, both in terms of design and architecture, and include the cost of what decoupling brings.
We wanna minimize these things, and we're gonna learn those things in this section.

### 4.1 Methods - Part 1 (Declare & Receiver Behavior)

We're about to shift gears now and we're gonna start talking about methods, but we're really gonna start talking about is decoupling.
Up until now, throughout the video, we've been really talking about the concrete, concrete data, and I've been trying to stress to you that everything we do is really built in and around the concrete data.
The problems you solve are based on the concrete data.
If you do not understand the data, you do not understand the problem.
If you don't understand the problem, you can't write code, right?
And if the problem is changing, that means the data is changing.
The problem is changing, your code's gonna have to change.
Our work, our daily lives are built in and around data.
Remember, we're all data scientists.
And what we also been trying to show is that your performance comes from the data.
Your performance comes from everything we're doing with data, and again, latency on network disk I/O and latencies on memory.
Latencies on how we access data on the machine, all of it's data, data, data, data, data.
But here's the problem, when the data's changing and the problem is changing, we're changing our code, what we wanna try to do is minimize, minimize the cascading changes that are gonna rip through the entire code base, and this is where the decoupling comes in.
And so, we've been focusing down here on the concrete data.
What we now have to start dealing with is the decoupling, and how do you decouple?
Decoupling is gonna be a big part of what we do.
And decoupling is all done through behavior, and we're gonna start wanting to focus on behavior now.
Behavior is where we do design, architecture, and decoupling, and our real work, the real work, the real work we're doing is all down here in the concrete and the data.
And one of the those we're gonna wanna learn is we wanna start from the concrete and the data and move up this way.
I don't ever, I really, really, really don't want you to work in this direction.
Too many developers are doing this.
They're starting with the behavior, they're starting with the decoupling, they're thinking about these, and you're guessing.
If you start up here, you are guessing.
I wanna take the guesswork out.
Solve the problem first in the concrete, then work towards decoupling.
I'm gonna show you this more anyway as we start getting into the design sections of this class, but right now I still have to teach you mechanics.
Gonna teach you the mechanics of what's happening above this line so we can get into design and architecture.

Alright, so what we're gonna do is take a look at some code here, and the code we're gonna look at is giving us the ability to allow data to have behavior.
I want you to think about this for a second.
To allow data to have behavior.
Now, Go has functions, we've been looking at functions since we started, but Go also has the concept of a method, and a method is a function, a function that has what we call a receiver.
So, let's take a look at what a method looks like in Go.
```go
import "fmt"

type user struct {
	name string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.changeEmail("bill@hotmail.com")
	bill.notify()
	
	joan := &user{"Joan", "joan@email.com"}
	joan.changeEmail("joan@hotmail.com")
	joan.notify()
}
```
Here's our user-defined type, it's `user`.
It's got a `name` and an `email`, and if I scroll up just a little bit above the type, you see two methods.
I always want those methods to come after the type.
Now, this is a method because what we have is a special parameter being defined.
It's a parameter, look at it as a parameter.
We call this parameter a receiver.
So, between the keyword `func` and the function name, we have a parameter, we call it the receiver, and now this is considered a method.
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
Really, the question is, when should a piece of data have behavior?
But let's start with some of the mechanics here.
Now, Go is interesting.
We've been talking about our value and our pointer semantics, and normally, value and pointer semantics don't hit Go developers in the face until we get to methods because now as we get into methods, you have to make a choice.
Should I be using a value receiver?
If you notice this, we have two receiver types that we can choose.
Should I be using a value receiver or should I be using a pointer receiver?
This is very interesting, and a lot of Go developers get stuck and they don't know what to do.
And too many Go developers, especially early on, start doing this.
Okay, I know.
If the method I am writing has to mutate the data, I'll use a pointer receiver.
I'll share the data in.
But if the method doesn't have to mutate, I'll use a value receiver.
I'll let it just have a copy.
This is very, very, very bad, we cannot do this.
I'm not gonna be able to stress enough that we must be consistent with our semantics.
Data drives the semantic, so once the semantic is chosen, either everything is value or everything is pointer, and the method and the code you write has to respect the semantic for the data.
Now, if you're used to programming languages like C++ and C# and Java, you actually deal with receivers.
You've heard of the `this` pointer before, you've heard of the `self` pointer.
Yeah, these are receivers, they're pointer receivers and they're named for you.
One of the things that Go does, which is brilliant, is they give you the choice of value and pointer semantics, value receivers, pointer receivers, but Go also allows you to name the receiver, and we never wanna use `this` and `self` when naming the receiver.
You want that receiver name to be short and sweet.
It should never even be more than three letters long.
This is at the highest level of context in your Go program.
So, you can see, I'm using `u` for `user`, `u` for the `*user`, any time we're using a `user` here, regardless of the semantic.
Okay, great.
So, our value semantic methods, like `notify()`, means that the method is operating on its own copy when we make the call.
Our pointer semantic methods there, right?
Right there, `changeEmail()` means that we are having shared access when we make the call.
There it is, we've got these two choices.
Now, while we continue to learn mechanics, you're gonna see code that has got a mix of value and pointer semantics.
Once we move to design, you will not see this anymore.
No more, no, no, no, no.
But I'm teaching mechanics, so I need both sides of this so I can show you the behavior.
Remember, mechanics, the semantics teach us the behavior, now we can understand how this code's going to perform.
Alright, great, so if I move into `main()`, I wanna show you a few things.
On line 32, I'm constructing a value of type Bill.
A value of type Bill, there it is right there, named Bill.
Let me clear this up a little bit so we can take a look at and visualize what's going to be happening in this code right here.
I'm going to be constructing a value of type `user` on line 32.
We know `bill` is a value of type `user`.
Now, notice the call on line 33 to `changeEmail()`.
`changeEmail()` was using our pointer semantics, changeEmail said, hey, hey, hey, you've gotta share a value with me.
Now, you might think, well, why would that compile?
Because `bill` is not a pointer to a `user`, it is a value to a `user`.
This is interesting.
When it comes to calling a method, Go only cares that the data, remember, data, value, pointer, that the data is of type `user`.
It doesn't care that it's a value or a pointer of that data, all it cares is that we are working with `user` values in some form, which is brilliant because now when we make that call to `changeEmail()`, what's gonna happen is Go's gonna be able to adjust to make the call.
It's going to take the address of `bill`, satisfy the receiver type, and share `bill` with `changeEmail()` because `changeEmail()`'s saying, I wanna use pointer semantics.
When we call line 34, `notify()`, which is using our value semantics, no big deal, `bill` is already a value and `notify()` will operate on its own copy.
Value semantics means we're operating on our own copy, pointer semantics, shared access.
Now, on line 38, I'm doing something I told you I don't want to do.
Any time I am constructing a value to a variable, I do not wanna use this syntax.
I don't wanna start my life out as a pointer.
Don't use those pointer semantics on construction if we're going to assign to a variable.
That being said, I wanted to show you something with the pointer.
Notice again, `changeEmail()`'s already using pointer semantics, no big deal, `joan` is a pointer, yada yada yada, but on `notify()`, `notify()` says, I want my own copy of the value.
Go, again, is gonna come in and say, no problem, I will take the value that the pointer points to for you and I will satisfy the call.
Now, what you're seeing me do on the board might be interesting because Go is adjusting to make the call, but now what we're doing is switching semantics in the code.
And I've tried to tell you over and over again, you start switching semantics, life is going to be very, very bad.
And in fact, this call on line 40 is the worst of them all.
You're really not allowed to take a copy of a value that a pointer points to.
Not all data can be copied, and if it's shared with you, you should pretty much assume that it's not copy.
So, if I ever saw code like on line 40, that would really scare me.
We wanna really maintain our value semantics.
We'll continue to talk about that, but I wanna show you here how Go's able to adjust to make the call to match the receiver type with the data that's there.
Okay, so what I wanna do next, now, is really try to give you some guidelines around when we should be using value semantics and when we should be using pointer semantics so we can come back into this code in a bit and get a better understanding of what we're doing in this code block and why.

### 4.1 Methods - Part 2 (Value & Pointer Semantics)

So now I wanna give you some general guidelines on when to be using value semantics and when to be using pointer semantics.
There are exceptions to everything, that's important, but if you don't know when to take the exception, don't.
Engineering is all about choices, it's about cost, it's about knowing when to take those exceptions.
So these general guidelines are very, very good and they're gonna work for you most of the time and they're gonna keep you out of a lot of trouble.
Remember, semantic consistency is everything.
So let's go and start taking a look at some code, and we're gonna pull some code from the standard library, but before we start with this code, let's look at what we're dealing with here.
There are really three classes of types that we're gonna be working with in a Go program.
You're gonna have the built-in types, your numerics, your string and your bool, you're gonna have the reference types, slices, maps, channels, interface values and functions, right?
These are the reference types, and then you're gonna have the user defined types.
These are really gonna be your struct types.
There are really three classes of types that we deal with, or three classes of data that we deal with that we've gotta make these decisions on.
So I'm gonna make this as simple as I can for you.
If you are working with the built-in types, strings, numerics, and bool, you are to be using value semantics.
In other words, I don't wanna see any code that takes a pointer to an int, any numeric, a string, or a bool, I don't wanna see that.
Everybody can get their own copy of the built-in types, in fact, strings are designed, they're mutable, they're designed to be copied.
Right, we saw that before.
So when we're dealing with these built-in types, including fields in a struct, I don't wanna see pointers.
Everything should be value based.
Now, there are some exceptions.
I'll give you one exception.
If you're dealing with a SQL database, and you have a struct that's going to be used to be un-marshalling or marshalling data, the concept of row doesn't exist unless there's a pointer.
So if it's not obvious to me, but that struct's being used as some sort of relational database scheme and I see a pointer, you know, we've gotta have a conversation.
So when I say value semantics, I mean not only value's being passed in and out of our function, I also mean the fields when we're defining a struct.
Okay, so that's it, right?
Built-in types, value semantics.
Everybody gets the concept.
Reference types, we're gonna follow the same rule.
Value semantics for reference types, I don't wanna see you take the address of a reference type.
There's one exception to this, however.
A slice and a map, you may take the address of a slice or a map only if you're sharing it down the call stack and to a function that's either named `decode` or `unmarshall`.
That's it.
I see that, we're not gonna have a conversation, that's reasonable, practical, decoding and un-marshaling require pointer semantics, and those data structures, slices and maps, are collections of data.
Other than that, everybody should get a copy of the slice value, the map.
Since maps and channels, anyway, are pointers, there's no reason to take the address of the address, and in our slice you'll see our interface values, or data structures.
Still, it doesn't matter, they're designed to be copied to stay on stacks, that means fields should also respect the value semantic form, in other words, I don't wanna see pointers to slices, pointers to interfaces.
We're going to do that.
Value semantics all the way for those reference types, except for the exception I gave you.
Great, so that means 2/3 of the types are using value semantics, right?
This is fantastic, I've told you, value semantics are a really important part of your ability to keep the heap clean which is going to give us a tremendous amount of performance.
But not all data can leverage the value semantics, and when it comes to our struct types this is where things get interesting.
You've got to choose, if you're defining your own struct type, what semantic is going to be in play.
You're gonna hear me say that a lot, now.
What's in play, what's in play?
And if you're not sure what to use then we're gonna use those pointer semantics, okay?
And then if you're absolutely sure that we can use value semantics, guess what?
We want to use those value semantics.
They're really our first choice.
Pointer semantics are really the exception, when we need it.
But I wanna take a look at some standard library code, to kind of bring this stuff home.
And I want you to look at these two types that I've pulled from the `net` package.
Line 20 and line 21. Now, again, what we would say is, "Hey, we're declaring a new type named IP, which is based on a slice of bytes."
Line 21, new type named `IPMask` which is based on a slice of bytes.
And what we really have here now is two brand new types that are truly based on a slice.
That means that they are reference types.
IP is a reference type, IPMask is a reference type, and what do I tell you?
Reference types follow value semantics.
Now based on that, we shouldn't be very surprised about line 26.
Look at how this method is being declared.
Notice that it's using value semantics.
But notice that `Mask()` is a mutation API.
You might think, "Well, hmm, Mask() is gonna have to mutate, I shouldn be using pointer semantics."
No, you would be wrong.
You remember the code, the methods, the functions, the code has to respect the semantic, its type, semantic and code.
So notice that `Mask()` is a value semantics mutation API, it gets its own copy of `IP`.
It mutates that copy, and then returns a new copy out.
Look at what it's doing.
And this is a very safe mutation, isn't it?
Because the mutation happens in isolation in our sandbox.
It can't create side effects.
This is a beautiful thing.
But the choice of this API is driven off the fact that `IP` is a reference type based on its declaration of being based on a slice of bytes.
Notice you're also passing a value of `IPMask`.
Why?
It's a reference type.
And this goes beyond just the method receivers.
Here's a function, a value of type `IP`, returns a value of type `string`, why? Strings are built-in type value semantics, `IP`'s a reference type value semantic.
You will see this throughout the entire standard library.
There is nothing random about code in the standard library.
It follows these semantic rules to a `T`, and there are very, very few exceptions, there, and you can go and search for this stuff yourself.
This is the key to having a code base that can be maintained by more than one person over time and can grow and maintain a consistency which is gonna, again, help reduce bugs and give everybody the ability to predict reasonably well how this code is gonna behave on the machine it's running on.

Now, let's talk about those struct types again.
When there's a struct type you have to decide what semantic you're going to use.
Here's the struct type for `Time`. I pulled this out of the `Time` package.
So let me ask you a question.
If I asked you to write the Time API, and I've given you this struct, what semantic would you choose?
Would you choose value semantics or pointer semantics?
Remember, you have to choose a semantic before you write a line of code.
Let me give you a couple of questions you can ask yourself, which in many cases can help.
It doesn't help all the time, but sometimes this helps.
So let's start with Time.
If I gave you a value of type `Time`, here it is, here's my value, remember, data is concrete, here's my concrete `Time` value.
Now, if I add 5 seconds to this value, is it the same value, just slightly mutated, or is it a new value?
Is two different points in time discrete data, different data, or the same data mutated?
I would argue that if I add 5 seconds to this `Time` value, I have a brand new piece of data.
And when we have that scenario it's usually a good indication that value semantics could be at play here.
That we can make copies of this data and mutations make new copies of data.
We do this with strings, right?
You mutate a string, you get a new string.
So that's interesting.
Now what if I was dealing with a `User` type?
Should we be making copies of users?
And this isn't a technical question, I'm not asking you if technically it's safe, I'm asking should we, from a correctness standpoint, right?
It's not optimizing for performance, it's optimizing for correctness.
Now, I could make an argument here.
If I met you, and I just started calling you Jack for the rest of the day, like I decide I don't know your name but I don't care, I'm gonna change your name, you are now Jack for the rest of the day.
Does that make you somebody different?
If I take a user value and change their name is it the same user, just slightly modified?
Or is it a completely new user?
It's not a new person, right?
It isn't, it's just somebody whose name has changed.
That's a good indication that we probably should be using pointer semantics.
But I also want to apply the idea that since we can't make copies of people in the real world, why would we do something like that in our software?
Our brains would not reasonably predict that sort of behavior, because we don't do that in the real world.
We need to bring all of this stuff into a code base, when we're coding we don't want to confuse people.
So probably users should be using pointer semantics, because mutation doesn't create a new piece of data, it just modifies the existing data.
This is a great question to ask, it doesn't always help you answer the question, but it's a good starting point and when you're not sure you should always be using pointer semantics because it's always safer to share than make copies because not all data can be copied.

So let's see what the language team decided to use for `Time`.
Now, what I've done is I've pulled the factory function out of the standard library.
Your factory function should always come before the type.
And they are the most important function in the API, because the factory function tells you the semantic that is at play, the semantic that has been chosen by the developer.
Where is it coming from?
The return, look.
The `Now()` function is returning a value of type `Time`.
You are getting your own copy of the `Time` value.
What is the factory function saying to you?
It's saying that value semantics are at play.
I am creating a `Time` value, I'm giving you your own copy, you can give other people copy, store it as a value, leverage the value semantics, keep this data on the stack, to the extent that we can, make your fields also of the `Time` value.
Brilliant.
Factory functions are everything.
This is where we start, this is where we choose our semantic.
If you're working with a struct type that you're not familiar with, maybe part of a standard library or somebody else's library, look for that factory function, and I pray it's right after the type.
Shouldn't be hard to find.
And leverage this information that I've highlighted, to show you and tell you how to work with the data correctly.
Now, since value semantics are at play, look at `Add()`.
`Add()` is a mutation API.
But it's not using pointer semantics, it's using value semantics, right?
A value semantic mutation.
This is not an accident, this is not random.
The code, the functions, the method has to respect the semantic, not the other way around.
So `Add()` gets its own `Time` value, mutates its own `Time` value in isolation within its sandbox, right?
Levels of immutability.
And then returns a new one out.
This is gonna allow us not to have any side effects while all this stuff is going.
This is beautiful stuff.
`Time` can be using value semantics, the API is using value semantics.
Okay, here's another function called `div()`.
Takes a value of type `Time`, value of type `Duration`, value of type `int`, value of type `Duration`.
`Duration` too, leveraging value semantics.
If you saw the declaration of `Duration`, it's really a type that's based on an `int64`.
`int64` is a built-in type.
Aw, look what's happening.
Look, the reality now is that you can look at almost any type in your language, identify what it is, or its base, and start to understand already what semantic you should be using.
And nothing should be random.
Code that has semantic consistency is going to always be better over the inconsistencies.

Okay, here's another factory.
Oh, that's right, I forgot about this.
Look, I told you there are exceptions to everything, and here's one of those exceptions.
When you are implementing or using functions that have `unmarshall` or `decode` in it they require naturally to have pointer semantics.
You need to share because they're gonna be mutated.
So here are four APIs in the `Time` package that are switching semantics from value to pointer, but guess what?
Unmarshall, decode, unmarshall, unmarshall, they're implementing contracts or interfaces for this behavior.
Absolutely fine, that's reasonable, practical.
But if I saw the pointer semantics for anything else, it would raise a really big flag with me.

Okay, look, here is a factory function for the file, for OS.
Notice that this factory function returns not a value of type `File` but a pointer of type `File`.
You see that?
What does this mean?
It means that pointer semantics are at play.
It also means that you do not have a right, remember this, this is incredibly important.
You do not have a right to make a copy of a value that a pointer points to when it's been shared with you like that.
Assume that it is dangerous to make copies if something has been shared.
Put that in your head, that is a major violation of semantic law.
We do not make copies that pointers point to.
Assume that that is very, very bad.
Now look at the `Chdir()` method.
It doesn't mutate the value used to make the call, but it's still using pointer semantics because that type requires pointer semantics.
Again, with this function here, pointer semantics, because we have to use, the API, the code, has to respect the semantic and not the other way around.
It doesn't get to choose the semantic.
We define the data, we define the type, we get the semantic, and we write code.
This is critical.

So to sum up.
* Built-in types, value semantics.
* Reference types, value semantics, except for decoding and unmarshalling.
* And then our struct types we have to make a choice, if you're not sure, you use the pointer semantics first.
 
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

### 4.1 Methods - Part 3 (Function / Method Variables)

### 4.2 Interfaces - Part 1 (Polymorphism)

So we just finished the topic on methods, which teaches us how to or the mechanics behind how you add behavior to data.
And I told you that I want this to be the exception, not the rule.
But I want to start showing you now the mechanics, and we'll look at some of the semantics, of interfaces, okay.
And interfaces give us the ability to have polymorphism.
And there's a saying or a quote from the inventor of BASIC, Tom Kurtz.
And he describes what polymorphism is perfectly.
Polymorphism means that you write a certain program and it behaves differently depending upon the data that it operates on.
I wanna change that up a little bit and just say, polymorphism means that a piece of code changes its behavior depending on the concrete data that it is operating on.
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
Now, in this program, we'll start off right out of the box on line 10.
On line 10, what we have is the `reader` interface.
Notice that we're using the keyword `type`
This type `reader` is not based on a struct.
It's based on the interface.
Right, it's got one active behavior `read`, which takes a slice of bytes, int, error.
But let's go back and look at `reader` for a second because this interface type always boggles my mind.
It's a type, right?
That means that I can do this.
I can say `var r reader`.
I can do that, I can declare a variable of this `reader` type.
But what blows my mind is the fact that an interface type is not real.
We've gotta make this very, very clear throughout.
Interface types are not real.
`r` is not real, there's nothing concrete about an interface type.
Struct types, yes, concrete.
That's real data we can manipulate, we can have fun with it.
But interfaces are not real.
They only define a method set of behavior.
They define a contract of behavior.
There's nothing real about an interface, nothing at all.
There's an implementation detail behind `r`, but from our programming model, `r` does not exist.
It is not real, there's nothing concrete about it.
This is gonna be an important concept for us to follow through with as we continue to look at this code.
Okay, so I've defined this interface type as one active behavior `read()`.
And one of things you know we're looking at mechanics and we'll talk more about this.
It's very important that your interfaces define a behavior.
Verbs, right?
From my perspective, we're reading, we're writing, we're running, we're printing, right.
Interfaces are about behavior.
I don't wanna see interfaces that describe things.
I don't wanna see an animal interface or a house or a user.
These are not interfaces, those are not behaviors.
Those are things, that's your concrete data.
The more you get away from that and the more your interfaces describe behavior, here we've got `reader`, has one active behavior, `read()`, this is behavior, we're gonna be much better off with the decoupling because again, we're focused on what?
Behavior.
Behavior, that's it.
And we haven't been taught that kind of stuff in our object-oriented programming world because we wanna decouple everything.
I'm trying to go back to, let's use our concrete data 'cause that's the problem and let's decouple what we need to and that's the behavior.
And I wanna keep this philosophy.
I wanna keep these mechanics and semantics throughout our entire program.
Now, I've got the `reader` interface.
It's got one active behavior.
`read()` takes a slice of bytes, returns an `int` and an `error`.
Okay, but let me ask you another question before we get going here.
You know, I could have defined `read()` to look like this.
I could've said, you know what, let's do this instead.
Let's pass in the number of bytes I wanna read and return that slice of bytes out.
Many of you might argue that this is a simpler API.
But why is this API horrific?
Why is line 13 so bad of an API design choice in Go?
That's because if we design our code or write our code like line 13, every time we make a call to `read()`, we're gonna be allocating on the heap.
Remember, you have a responsibility in API design, not only to write precise APIs where you avoid misuse and fraud, but you also have a responsibility on the impact your APIs have on the user machine.
There's an impact here.
And this one is allocation.
You might say, Bill, where is the allocation coming from?
Look, you can't implement methods like this inside the type.
I'm just doing this, so I can just write a little code to show you that if we were implementing it, where these allocations would come from.
Look, your first allocation would come from the fact that you are gonna have to make this slice of bytes here based on `n`.
That's going to be an allocation because `n`, we don't know what compile time, what the backing array should be.
Immediate allocation, there's one potential allocation for this API design.
But you might say, you know what, Bill?
I'm not going to then allow you to pass an `n`.
I'm gonna make it even simpler.
I'm just gonna say `read()` and I'll hard code that.
Yup, okay, compiler knows what the size of the backing array is now.
We've gotten rid of that allocation, but guess what?
Eventually, you've gotta return the slice back up.
That's now gonna cause an allocation on the backing array because you can't have a pointer down the call stack.
Do you see that this API design is not the right choice?
Or we're gonna have an API that's gonna be in a tight loop.
The other API, the one that I've chosen here, doesn't allocate 'cause we're asking the caller to allocate the memory and share it down.
See, in Go, you do have a lot of control over your memory management because you understand escape analysis and you can drive data down, which means that the caller can define the slice, even hard coded size, share it down the call stack, and there's no allocation.
Okay, great, just wanna keep bringing up where we can where allocations are gonna be the escape analysis stuff.
This was one of those places.
Now, I've got my `reader` interface type, there it is.
I can declare variables.
Remember `r` is not real, interfaces are not real.
Put this into your head, the only real is concrete.
And on line 15, it's exactly what we've done.
I got a concrete type named `file`.
And this file represents a concrete piece of data, some sort of file that you might have on a file system, and I've got just a name in there.
But more importantly, look on line 20. On line 20, we have declared a method and this method has the same signature as the interface.
Look at that, `read()`, slice of bytes, `int`, `error`.
Now, Go is about convention over configuration.
Listen to every word that I say right now.
Listen to every single word 'cause not a single word here isn't important.
Because of the method declaration on line 20, I can now say the following.
The concrete type `file` now implements the `reader` interface using value semantics.
Think about everything I just said.
Every word's important.
Because of line 20, I can now say the concrete type `file` now implements the `reader` interface using value semantics.
You see Go is about convention over configuration.
We do not configure an interface to the concrete type like you might see in other languages.
In other languages, you probably have to do something like this.
This is not Go, this is configuration.
This is going to limit us at the end of the day.
It actually does limit us and cause our software to have more lines of code.
Go is about less lines of code.
Go is about being more productive.
And because of the static code analysis, it's gonna give us a lot of benefits as I'm gonna show you as we move away from mechanics and into design.
So we don't have to do this in Go.
We just have to declare the method like we're going on line 20.
And the compiler compile time can identify interface compliance, satisfaction.
Great, so now, the concrete piece of data associated with the concrete type `file`.
Notice I'm using the word "concrete".
That struct's concrete.
Now, implements the `reader` interface, you know, using the value semantics.
Line 20 again.
The concrete type `file` now implements the `reader` interface using value semantics, there we go.
Forget about the implementation, implementation's silly.
We're just pretending to read a file.
Now, look, I have a second concrete type named `pipe`, supposed to represent some sort of networking pipe, right.
But look on line 32.
I also have a method named `read()` with the same signature, right, that's part of the `reader` interface.
And therefore, since there's only one method in the `reader` interface, we are satisfying the entire method set of the interface.
Now, because of this, because on line 32, I can now say the following, the concrete type `pipe` now implements the `reader` interface using value semantics.
There we are again.
I have two distinct pieces of data, two concrete types, both with their unique implementation of the `reader` interface.
Brilliant, this is setting us up for polymorphism.
Remember what polymorphism says.
A piece of code changes its behavior depending upon the data, the concrete data, it is operating on.
Oh god, I love that saying, I love it.
It gives me chills every time I say it because I really wanna stress this data-oriented design and focusing on the concrete data and letting it drive everything we do, even through the decoupling.
Okay, let's look at what I would call our polymorphic function.
There it is on line 50, it's called `retrieve()`.
```go
func retrieve(r reader) error {
	data := make([]byte, 100)
	
	len, err := r.read(data)
	if err != nil {
		return err
	}
	
	fmt.Println(string(data[:len]))
	return nil
}
```
This is polymorphic and there's nothing about the concrete.
In fact, you might even be confused for a second.
Look at the parameter.
The parameter in this function seems to say the following.
It says, pass me a value of type `reader`.
Can you do that please?
I would love a value of type `reader`.
Guess what?
We already know this is impossible.
There are no such things as values of type `reader`.
Values of type `reader` do not exist because `reader` is an interface type and interface types are valueless.
Whoa, what did I just say?
Think about this.
Interface types are valueless.
So it is impossible for this polymorphic function to be asking for a value of type `reader` because it doesn't exist.
Well, if it doesn't exist and that's what it's saying, what is then really saying?
Well, actually, this is what `retrieve()` is saying.
It is saying, please pass me any piece of concrete data, any value or any pointer, that satisfies, that contains the behavior of `reader`, the full method set of `reader`.
Think about this again.
Remember we can only pass concrete data through those program boundaries, through our app.
That's what's real, that's what we manipulate, that's solving the problem.
So what is this function saying again?
Pass me any piece of concrete data, any value, any pointer, that satisfies the `reader` interface and implements this contract that has the data, the full suite of behavior, method sets, that satisfies the `reader`.
And don't we have two concrete pieces of data in this program already, `file` and `pipe`, that have the full method set of `reader`?
We do.
But this is my polymorphic function because it knows nothing about the concrete.
It can operate against any piece of concrete data that implements the `read()` method.
And you can see here that we're gonna be executing that behavior through the interface.
Okay, you know what, let's look at the mechanics on the board.
Let's draw all this stuff out, so you can see this polymorphism in action, both the semantics side of it and the mechanics.
Okay, so look what we do right here, right there on line... I gotta turn this around again, line 41.
Look at line 41, okay, brilliant.
What are we doing on line 41?
We are creating a value of type `file`.
There's `f`, there's our `file` value.
I'm also creating a value of type `pipe`, there it is, `p`.
I've got these two values.
Remember we're constructing to a variable.
We're gonna use our value semantic construction.
We've done that, great. But now, look on line 45. On line 45, we're making a call to `retrieve()`.
Now, I'm always gonna be asking you the same question.
What semantic is at play when we make this function call?
In other words, is the function receiving its own copy of the data, value semantics?
Or is it being shared the data, pointer semantics?
We look at this and we realize that this polymorphic function is being given a copy, its own copy of the file.
So we're passing a copy of `f` across this program boundary.
The compiler during compilation knows that this piece of data implements the `reader` interface using value semantics.
You can see that during static code analysis.
Remember retrieve is saying, pass me any concrete piece of data, any value or any pointer, that satisfies, implements the reader interface, has the full suite of behavior.
We know that this does, right.
We know that this piece of data has a `read()` method.
We know it's there.
Brilliant, now, on the other side though, even though we're giving it this copy, what we have is `r`.
And `r` is an interface value.
Now, there's only an implementation detail because from our programming model, `r` is not real.
It is valueless, there's nothing concrete about it.
But what is `r` then from an implementation detail?
`r` is really an interface value, which makes it a reference type.
And when `r` is set to its zero value, there it is, two pointers, `nil`, `nil`.
Now, when we pass this concrete piece of data over the program boundary, the whole idea here is there's a relationship between interface values and our concrete data.
And that is one of storage.
We store concrete data inside of interface values and when we store the data inside of interface value, that finally makes the interface value concrete.
Everything's always gotta get back to the concrete.
So the interfaces are really valueless.
Through the mechanism of storing, when we pass this concrete data across the program boundary, now, we have something that's more concrete.
And it's this second word that we're using for storage.
So look, we made a copy of `f`, we're now storing it inside of the interface.
This is giving us our levels of decoupling.
And because we made a copy of `f`, what does that mean?
It means what?
Allocation, there it is.
Allocation, every time I draw that little red, I'm showing hey, allocation.
So we got an allocation now.
Brilliant, what about the first word?
Okay, we're gonna get a little technical here.
The first word points to a very special internal table that we call the "iTable".
Now, if you've ever heard of vtables before in an object-oriented programming language, then an iTable's very much like the vtable, but it's an interface, so we replace "v" with "i".
And vtables are just a matrix of function pointers that allow us to have base class pointers go into derived class objects behavior.
And the iTable's kind of doing the same thing.
But the first word of the iTable will always describe the type of value that we're storing inside the interface.
In this case, we've got a value of type `file`.
The rest of the iTable will be that function pointer.
So this is going to basically point to the concrete implementation of `read()` for `file`.
Right, so what's going to happen now on line 53?
Remember interfaces are giving us a level of decoupling.
So on line 53, when we call read, when you call `read()` against the interface, we do an iTable lookup.
Guess what we have here again?
Indirection.
We do an iTable lookup to find out where the implementation of the actual `read()` is and then we call that implementation against, in our case, our copy.
And there we have indirection.
Indirection and allocation.
That is what decoupling costs us.
But guess what?
I've got a polymorphic function that can handle any kind of concrete data.
And remember, `retrieve()` is changing its behavior.
The call on line 53 is changing its behavior depending on the concrete data it's operating on.
Right now, `retrieve()` is accessing a file system.
How cool is that?
But let's go look at the next thing.
On line 46, now, we call `retrieve()` again.
Again, what semantic is at play?
Value semantics are at play.
`retrieve()` is gonna receive its own copy, not of `f` this time, but it's gonna retrieve its own copy of `p`.
And we know `p` also knows how to `read()`.
That means, this time, the iTable doesn't say `file` here.
It says, no, no, no, I've got a `pipe` value.
Brilliant, now this time, when I call `read()` against `r`, we're gonna call `read()` against our copy of `p`.
Allocation, but the behavior has changed.
Polymorphism means that a piece of code like `retrieve()` changes its behavior depending upon the concrete data it is operating on.
It's still all driven from real concrete data.
This has now made the interface value concrete.
But at the end of the day, interface values really are valueless.
They're not real.
Only the concrete data is real.
Now moving forward, I don't particularly care about iTables.
It's a little too technical.
I want you to know they're there.
I want you to know that this is a pointer.
But what I'm gonna be doing moving forward is putting inside of this word, the type of value that we have stored inside of the interface.
Okay, so I would be saying here that we have a value of type `pipe`.
That means we're using value semantics with the interface.
Here's the copy.
And now, when we're calling `read()` against our copy, then the iTable's giving us all that.
So there's no magic going on here.
Now, this is our basic polymorphism in Go.
It's really showing you when a piece of data should have behavior.
This is one of those cases when we need to implement polymorphism.
I cannot stress enough though there is a cost to decoupling here.
And again, it's the indirection and the allocation.

### 4.2 Interfaces - Part 2 (Method Sets and Address of Value)

Alright, let's look at another polymorphic example here, and things are going to get a little more complicated, but that's okay, we're going to build up to this.
Now I've got another interface here called `Notifier`, one active behavior, `notify()`.
Again, I want that relationship of behavior.
Interfaces should describe behavior, not persons, places, and things.
And here it is, in `Notifier`, `notify()`.
We've got this one active behavior, there we are.
Now here's my concrete data, it's named `user`, `name` and `email`, and here we are, line 22, and remember what I can say, the concrete data `user` now implements the `Notifier` interface using pointer semantics.
We were using value semantics before, now we're using pointer semantics.
All I had to do was implement the `notify()` method based on the fact that that is the full method set now for `Notifier`.
I can say it again, the concrete type `user` now implements the `Notifier` interface using pointer semantics.
Now I go ahead and I look down here, and we have, again, our polymorphic function.
What is this function saying again?
It's saying, I will accept any piece of concrete data, any value, any pointer, that implements the `Notifier` interface.
It can't be asking for Notifier values, cause they don't exist, that's an interface type.
The interface types are valueless.
Pass me any concrete piece of data, any value, any pointer, that implements, that has the full method set of behavior for `Notifier`.
And then we will call into that behavior from a decoupled state through the interface value.
Brilliant, so let's go through this program now.
So I go ahead, on line 31, and I construct my user value, there it is, user right there, and we put "Bill" in it.
There it is, there's my concrete data.
Now just like we did before, on line 36, I call my polymorphic function using value semantics, just like we did before, which means we want to pass a copy of "Bill" across this program boundary.
But something very odd happens.
When I try to run this program, the compiler very quickly says, whoa, whoa, whoa, whoa, whoa!
Look, I'm sorry, but this piece of data you're trying to pass across this program boundary, it doesn't have the behavior that it needs, it doesn't have `notify()` behavior.
And we're like, what?
What you talking about, compiler?
I did everything you asked me to do.
I went ahead, and I defined a method, named `notify()`, and `notify()` matches the `notify()` method inside the interface.
I have compliance, there it is!
I've got a method right there named `notify()` for the concrete type `user` that implements the `Notifier` interface using pointer semantics.
Compiler, it's all there, why are you not letting me pass this value through?
Why does this piece of data not implement the `Notifier` interface?
Well, the compiler's absolutely right.
And what the compiler's trying to do here, I call it love.
You see, the compiler loves you, and the compiler doesn't want you to have an integrity issue in your software.
And if the compiler allowed you to use value semantics here, we would be in a tremendous amount of trouble.
Now I want to explain to you and show you where this integrity issue is, there's really two parts to it.
But before I can, we've got to talk about method sets.
There's a set of rules in the specification around method sets.
And these method sets are there to protect us and our software
Now I'm going to write the rules for method sets as it relates to the specification, and then I'm going to switch it up just a little bit to make it a little easier for you to read and understand and then understand where the integrity issues are with line 36.
But this is all love, I don't want you to look at it any other way.
This is the compiler showing you love.
Okay, here we go.
This is what the compiler says.
It says, if you're working with a value of type `T`, right, if we're working with our value semantics here, working with a value of type `T`, then what it's saying here is only those methods using value receivers, right, only those methods using value semantics, belong to the method set for this value, that's it.
And if you're working with pointers, right, we've got pointer semantics here, then what it's saying, your pointer receiver methods, and your value receiver methods, what it's saying is all the methods that you declare for that concrete type, they exist for pointers.
You get the full method declarations for your pointer data, but only your value methods for your value data.
In other words, these pointer receiver methods are left out of the method set for values.

||||
|---|---|---|
| T | | Value Receiver |
| *T | Pointer Receiver | Value Receiver |

And if we just quickly go back and look at what's happening here, notice that I implemented the `Notifier` interface using pointer semantics.
Pointer semantics, and now I'm trying to use value semantics.
I'm mixing semantics, am I not, and the compiler's going, whoa, you're trying to create a copy of `T`, but guess what, `T` doesn't have any methods.
Compiler was right, because the pointer receiver method didn't get included in the method set.
The real question again now is why?
Why can't we include pointer receiver methods for this value of `T`, but we can for pointers, and why do value receivers get to apply to both?
Okay, so the two integrity pieces here, one I would call minor, one I would call major.
Let's look at the minor integrity issue.
Now what if I said to you that not every value you work with has an address?
Think about this for a second.
What if I told you that you could not take the address of `T`?
Well, think about it, if you can't take the address of `T`, you can't call a pointer receiver method, you cannot be using pointer semantics if something can't be shared.
Remember, integrity is about a hundred percent.
If you can't have something a hundred percent of the time, you cannot have it, you're setting yourself up for integrity issue.
Look at this piece of code right here.
I define a type named `duration` based on an integer.
`duration` is based on an `int`.
And then I implement the `Notifier` interface using pointer semantics, I have a method using pointer semantics named `notify()`.
But look at what I do on line 18.
On line 18, I go ahead and I take the literal value 42, which we know is a constant of kind int, convert it to a value of type `duration`, and try to call `notify()`.
But the compiler says, dude, I'm sorry, this value has no address.
And the compiler's right, because still at the end of the day, this is a constant.
It's a constant of type `duration`, and we know constants only exist to compile time.
They never find themselves on a stack or on the heap.
There is no address, even though it does satisfy the interface, there is no address.
So this would be very bad to store this duration inside of our interface because there's no address, and therefore the pointer receivers can't be applied.
Look, we're getting at least compiler messages here.
This is a minor one for me.
This isn't the big one, this isn't the one that's like hugs and love here.
But I want you to understand that if you can't have something a hundred percent of the time, you can't have it at all.
So we can't assume that we can always get the address of `T`.
But if you already have an address to `T`, then we already know it's in memory, so we can always use the address, or we can always get a copy of the value that the pointer points to.
You can always do this, but you can't assume the address.
But what's the bigger play here?
Why are these rules really in place?
And remember, this stuff has been known for a decade.
Well, let's look at this a little differently.
When we talk about decoupling, I really want to focus on the behavior side, not necessarily on the data side.
And when we do that, something very interesting happens.
If we read this chart from right to left, read it in this direction, focus on the behavior, something really, really interesting happens.
And what this chart is saying is the following: If you've chosen pointer semantics, right, if you've chosen your pointer semantics, and want you to remember something, we define the type, we choose the semantics, and then we implement.
So if the pointer semantics were the choice, then what this chart is saying is, you can only share.
If you've chosen pointer semantics, then the only thing you're allowed to do is share, because that is the only safe thing to do.
If you choose value semantics, then we want you to make copies.
But I told you one thing.
I said to you that sometimes it is absolutely safe to share a value even if you're working in value semantics.
I don't really want to change semantics if I don't have to, but even if you're using value semantics, there are times where it's safe to share.
But what this chart is also saying is, it is never safe, it's never safe to make a copy of a value that a pointer points to.
It's never safe.
So if you're in pointer semantic mode, you're only allowed to share, you're never allowed to make a copy of the value that the pointer points to.
How amazing is this chart?
When we look at it from the behavior point of view, it all begins to make sense.
The compiler's saying, if you chose pointer semantics, you're only allowed to share.
If you change your semantic back to value, that is a major, major violation of law.
If you're using value semantics, please, let's make copies of things, let's give the interface its own copy.
But there are times where you might need to share something, like decoupling, like unmarshalling.
But we never, ever, never, ever want to go from pointer semantics to value semantics.
Never, ever, ever, it's a major violation, and what the compiler has done here is said, when we're storing data inside of an interface, we're never going to allow ourselves to break this major law of pointer semantics to value semantics, no, no, no, no, we're never allowed to make a copy of the pointer the pointer points to, you cannot assume that that is safe.
So when we come back and look at this code now, what we realize is we implemented the interface using pointer semantics, which now means that we can only share data with the interface.
We're no longer allowed to make copies of the data, which is what I was trying to do.
So if I turn around and now say, okay, let's maintain the right law here, since we're using pointer semantics to implement the interface, I can only share the data with the interface.
Boom, there it works.
How cool is this?
So this chart again is here to protect us, to maintain the right level of integrity when we're dealing with decoupling.
So if you are using pointer semantics, which was chosen at the time you defined the type, those methods were declared based on the type, then you can only share data.
If you're using value semantics, let's make copies, but you also have the ability to share when it is important.
This is never allowed, never allowed, using pointer semantics, you may not make copies of the value that the pointer points to.

### 4.2 Interfaces - Part 3 (Storage by Value)

So I wanna show you one more time how the semantics really drive our ability to read and understand the impact code's gonna have.
And that when we are decoupling in Go, it's gonna come with the cost of indirection and allocation.
And this is not unique to Go, by the way.
I mean, decoupling is just inherently an allocation type of operation.
But let's take a look at this code here.
We can really see this and again, begin to understand that when we understand our mechanics, and then we understand our semantics, we understand how code's gonna behave, and the impact, we can make some better engineering choices.
Here is the `printer` interface when active behavior `print()`.
And here's our concrete type `user`.
It has a `name` and it implements the `printer` interface using value semantics.
And what do we now know when we're implementing interfaces with value semantics?
We can both use value semantics and pointer semantics, but we prefer to use the value semantics and keep our semantics clean.
Okay, here we go, value semantics is at play with this concrete data user.
Now I go ahead and I create a `user` value, there it is, there is `u`, I put "Bill" in it, beautiful, there's my concrete data, that's our work.
And then on line 32, I do something really interesting.
I create a slice, a collection, of `printer` interface values.
I now have a collection of data that is not bound to just one concrete type, but any concrete data, any value or any pointer that implements the `printer` interface, the behavior of `print()`.
And I'm using my slice literal here to do construction.
And on line 35, I use my value semantics to store the first piece of concrete data.
So this will say `user` and this will have its own copy of `u`.
Value semantics are at play, which now means what?
It means that since we're using value semantics, what do we have here?
Allocation, there it is, we have a copy.
There's our allocation.
And then on line 38, what have we done?
We're now not using value semantics, we're using pointer semantics on line 38.
That means the interface is gonna look like this, a `*user`, it means that it's referencing the original `u`, pointer semantics, value semantics, again, because anytime we store a value inside of the interface, we're gonna have allocation.
Guess what?
Allocation, right there.
There's our cost, indirection and allocation, boom, there it is.
But when I go ahead on line 42, and I change out our data, only index one, not index zero will see the change, because index zero is using value semantics, index one is using pointer semantics, we can see and understand the impact that we're having.
Now on line 46, we're using the value semantic form of the `for range`.
This is not an accident, this is not random.
Think about it.
We are gonna range over a collection of interface values.
Interfaces are reference types.
Reference types use value semantics.
Guess what form of the `for range` we use.
Value semantics.
Do you see how everything is driven off of the data, and the data semantic?
What is the collection?
Interfaces.
Interfaces, value semantics, value semantics, there we go, right there, `for _, e`.
Now `e` then is an interface value, as well, local to the `for` loop.
On the first iteration, what do we get?
A copy of index zero, we're using those value semantics, which now means, this is `e`, which now means I have two interface values sharing, pointers are for sharing, efficiency, I'm hoping by now you see the efficiencies that we're getting through our pointers.
But we made a copy, value semantics, I've got two interface values, now it's sharing the same user value.
And then we call `print()`, right?
We call `print()` against our copy, which is gonna call `print()` against there, right there.
We call `print()` against `e`, which calls `print()` against our copy.
Now on the next iteration, what happens?
We get a copy of the next interface value.
So we're not gonna be pointing here anymore.
This is now gonna say `*user`, we're gonna come here, and now when we call `print()` against the same `e`, polymorphism, the behavior changes depending upon the concrete data that we're operating on, we now call `print()` against this, right here, against the original.
And we see the change.
Polymorphism, again, is all about decoupling through behavior.
But everything's still being driven through the concrete, and the concrete data having the behavior.
Polymorphism means that a piece of code changes its behavior depending upon the concrete data it's operating on.
And right here, we were operating on two different `user` values, one in value semantic mode, one in pointer semantic mode, and we see different behavior.
This is really important.
And if we maintain these semantics and these semantic consistencies throughout our code base, life is just going to be better for us and everybody on the team and everybody that we haven't even met yet that eventually has to come in and maintain this code.
