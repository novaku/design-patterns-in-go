**Mediator** in Go
==================

![Mediator design pattern](https://refactoring.guru/images/patterns/content/mediator/mediator.png)

**Mediator** is a behavioral design pattern that reduces coupling between components of a program by making them communicate indirectly, through a special mediator object.

The Mediator makes it easy to modify, extend and reuse individual components because they’re no longer dependent on the dozens of other classes.

[Learn more about Mediator](https://refactoring.guru/design-patterns/mediator)

Conceptual Example
------------------

An excellent example of the Mediator pattern is a railway station traffic system. Two trains never communicate between themselves for the availability of the platform. The `stationManager` acts as a mediator and makes the platform available to only one of the arriving trains while keeping the rest in a queue. A departing train notifies the stations, which lets the next train in the queue to arrive.