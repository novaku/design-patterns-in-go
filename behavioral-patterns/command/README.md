**Command** in Go
=================

![Command design pattern](https://refactoring.guru/images/patterns/content/command/command-en.png)

[Learn more about Command](https://refactoring.guru/design-patterns/command)

Conceptual Example
------------------

Let’s look at the Command pattern with the case of a TV. A TV can be turned ON by either:

*   ON Button on the remote;
*   ON Button on the actual TV.

We can start by implementing the ON command object with the TV as a receiver. When the execution method is called on this command, it, in turn, calls the `TV.on` function. The last part is defining an invoker. We’ll actually have two invokers: the remote and the TV itself. Both will embed the ON command object.

Notice how we have wrapped the same request into multiple invokers. The same way we can do with other commands. The benefit of creating a separate command object is that we decouple the UI logic from underlying business logic. There’s no need to develop different handlers for each of the invokers. The command object contains all the information it needs to execute. Hence it can also be used for delayed execution.