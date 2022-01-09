**Memento** in Go
=================

![Memento design pattern](https://refactoring.guru/images/patterns/content/memento/memento-en.png)

**Memento** is a behavioral design pattern that allows making snapshots of an object’s state and restoring it in future.

The Memento doesn’t compromise the internal structure of the object it works with, as well as data kept inside the snapshots.

[Learn more about Memento](https://refactoring.guru/design-patterns/memento)

Conceptual Example
------------------

The Memento pattern lets us save snapshots for an object’s state. You can use these snapshots to revert the object to the previous state. It’s handy when you need to implement undo-redo operations on an object.