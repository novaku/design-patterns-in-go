**Facade** in Go
================

![Facade design&nbsp;pattern](https://refactoring.guru/images/patterns/content/facade/facade.png)

**Facade** is a structural design pattern that provides a simplified (but limited) interface to a complex system of classes, library or framework.

While Facade decreases the overall complexity of the application, it also helps to move unwanted dependencies to one place.

Conceptual Example
------------------

It’s easy to underestimate the complexities that happen behind the scenes when you order a pizza using your credit card. There are dozens of subsystems that are acting in this process. Here’s just a shortlist of them:

*   Check account
*   Check security PIN
*   Credit/debit balance
*   Make ledger entry
*   Send notification

In a complex system like this, it’s easy to get lost and easy to break stuff if you’re doing something wrong. That’s why there’s a concept of the Facade pattern: a thing that lets the client work with dozens of components using a simple interface. The client only needs to enter the card details, the security pin, the amount to pay, and the operation type. The Facade directs further communications with various components without exposing the client to internal complexities.