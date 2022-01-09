**Observer** in Go
==================

![Observer Design Pattern](https://refactoring.guru/images/patterns/content/observer/observer.png)

**Observer** is a behavioral design pattern that allows some objects to notify other objects about changes in their state.

The Observer pattern provides a way to subscribe and unsubscribe to and from these events for any object that implements a subscriber interface.

[Learn more about Observer](https://refactoring.guru/design-patterns/observer)

Conceptual Example
------------------

In the e-commerce website, items go out of stock from time to time. There can be customers who are interested in a particular item that went out of stock. There are three solutions to this problem:

1.  The customer keeps checking the availability of the item at some frequency.
2.  E-commerce bombards customers with all new items available, which are in stock.
3.  The customer subscribes only to the particular item he is interested in and gets notified if the item is available. Also, multiple customers can subscribe to the same product.

Option 3 is most viable, and this is what the Observer pattern is all about. The major components of the observer pattern are:

*   Subject, the instance which publishes an event when anything happens.
*   Observer, which subscribes to the subject events and gets notified when they happen.