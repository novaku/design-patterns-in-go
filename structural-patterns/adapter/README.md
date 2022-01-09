**Adapter** in Go
=================

![Adapter design pattern](https://refactoring.guru/images/patterns/content/adapter/adapter-en.png)

**Adapter** is a structural design pattern, which allows incompatible objects to collaborate.

The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.

[Learn more about Adapter](https://refactoring.guru/design-patterns/adapter)

Conceptual Example
------------------

We have a client code that expects some features of an object (Lightning port), but we have another object called _adaptee_ (Windows laptop) which offers the same functionality but through a different interface (USB port)

This is where the Adapter pattern comes into the picture. We create a struct type known as _adapter_ that will:

*   Adhere to the same interface which the client expects (Lightning port).

*   Translate the request from the client to the adaptee in the form that the adaptee expects. The adapter accepts a Lightning connector and then translates its signals into a USB format and passes them to the USB port in windows laptop.