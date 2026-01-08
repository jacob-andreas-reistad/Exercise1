Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> concurrency is when a program does several logical operations concurrently, meaning that the operations might overlap in time, but are not executed at the same time. Parallelism is when several logical operations are executed simultaneously, meaning that they are executed at the same time. Parallelism requires several CPUs while concurrency does not.

What is the difference between a *race condition* and a *data race*? 
> A race condition is when the result of a program depends on the order of operations in time. A data race occurs when two or more threads access the same memory address simultaneously, with at least one thread performing a write operation, and there is no synchronization, resulting in unpredictable behavior. In other words, a data race is a type of race condition
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler decides what thread to run next.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> *Your answer here*

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Your answer here*

What do you think is best - *shared variables* or *message passing*?
> *Your answer here*


