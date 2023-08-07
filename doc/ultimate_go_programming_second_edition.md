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
