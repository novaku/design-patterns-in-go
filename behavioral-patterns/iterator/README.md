**Iterator** in Go
==================

![Iterator design pattern](https://refactoring.guru/images/patterns/content/iterator/iterator-en.png)

**Iterator** is a behavioral design pattern that allows sequential traversal through a complex data structure without exposing its internal details.

Thanks to the Iterator, clients can go over elements of different collections in a similar fashion using a single iterator interface.

[Learn more about Iterator](https://refactoring.guru/design-patterns/iterator)

Conceptual Example
------------------

The main idea behind the Iterator pattern is to extract the iteration logic of a collection into a different object called iterator. This iterator provides a generic method of iterating over a collection independent of its type.