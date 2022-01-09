**Bridge** in Go
================

![Bridge design pattern](https://refactoring.guru/images/patterns/content/bridge/bridge.png)

**Bridge** is a structural design pattern that divides business logic or huge class into separate class hierarchies that can be developed independently.

One of these hierarchies (often called the Abstraction) will get a reference to an object of the second hierarchy (Implementation). The abstraction will be able to delegate some (sometimes, most) of its calls to the implementations object. Since all implementations will have a common interface, they’d be interchangeable inside the abstraction.

[Learn more about Bridge](https://refactoring.guru/design-patterns/bridge)

Conceptual Example
------------------

Say, you have two types of computers: Mac and Windows. Also, two types of printers: Epson and HP. Both computers and printers need to work with each other in any combination. The client doesn’t want to worry about the details of connecting printers to computers.

If we introduce new printers, we don’t want our code to grow exponentially. Instead of creating four structs for the 2\*2 combination, we create two hierarchies:

*   Abstraction hierarchy: this will be our computers
*   Implementation hierarchy: this will be our printers

These two hierarchies communicate with each other via a Bridge, where the Abstraction (computer) contains a reference to the Implementation (printer). Both the abstraction and implementation can be developed independently without affecting each other.