**State** in Go
===============

![State Design Pattern](https://refactoring.guru/images/patterns/content/state/state-en.png)

**State** is a behavioral design pattern that allows an object to change the behavior when its internal state changes.

The pattern extracts state-related behaviors into separate state classes and forces the original object to delegate the work to an instance of these classes, instead of acting on its own.

[Learn more about State](https://refactoring.guru/design-patterns/state)

Conceptual Example
------------------

Let’s apply the State Design Pattern in the context of vending machines. For simplicity, let’s assume that vending machine only has one type of item or product. Also for simplicity lets assume that a vending machine can be in 4 different states:

*   hasItem
*   noItem
*   itemRequested
*   hasMoney

A vending machine will also have different actions. Again for simplicity lets assume that there are only four actions:

*   Select the item
*   Add the item
*   Insert money
*   Dispense item

The State design pattern should be used when the object can be in many different states and depending upon incoming request the object needs to change its current state.

In our example, a vending machine can be in many different states and these states will constantly switch from one to another. Let’s say vending machine is in `itemRequested`. Once the action “Insert Money” occurs, the machine moves to `hasMoney` state.

Depending on its current state, the machine can behave differently to the same requests. For example, if user wants to purchase an item then the machine will proceed if it’s in `hasItemState` or it will reject if it’s in `noItemState`.

The code of the vending machine is not polluted with this logics, all the state-dependent code lives in respective state implementations.