**Factory Method** in Go
========================

![Factory Method pattern](https://refactoring.guru/images/patterns/content/factory-method/factory-method-en.png)

**Factory method** is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.

Factory Method defines a method, which should be used for creating objects instead of direct constructor call (`new` operator). Subclasses can override this method to change the class of objects that will be created.

[Learn more about Factory Method](https://refactoring.guru/design-patterns/factory-method)

Conceptual Example
------------------

It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and inheritance. However, we can still implement the basic version of the pattern, the Simple Factory.

In this example, we’re going to build various types of weapons using a factory struct.

First, we create the `iGun` interface, which defines all methods a gun should have. There is a `gun` struct type that implements the iGun interface. Two concrete guns—`ak47` and `musket`—both embed gun struct and indirectly implement all `iGun` methods.

The `gunFactory` struct serves as a factory, which creates guns of the desired type based on an incoming argument. The _main.go_ acts as a client. Instead of directly interacting with `ak47` or `musket`, it relies on `gunFactory` to create instances of various guns, only using string parameters to control the production.