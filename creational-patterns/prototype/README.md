**Prototype** in Go
===================

![Prototype Design Pattern](https://refactoring.guru/images/patterns/content/prototype/prototype.png)

**Prototype** is a creational design pattern that allows cloning objects, even complex ones, without coupling to their specific classes.

All prototype classes should have a common interface that makes it possible to copy objects even if their concrete classes are unknown. Prototype objects can produce full copies since objects of the same class can access each other’s private fields.

[Learn more about Prototype](https://refactoring.guru/design-patterns/prototype)

Conceptual Example
------------------

Let’s try to figure out the Prototype pattern using an example based on the operating system’s file system. The OS file system is recursive: the folders contain files and folders, which may also include files and folders, and so on.

Each file and folder can be represented by an `inode` interface. `inode` interface also has the `clone` function.

Both `file` and `folder` structs implement the `print` and `clone` functions since they are of the `inode` type. Also, notice the `clone` function in both `file` and `folder`. The `clone` function in both of them returns a copy of the respective file or folder. During the cloning, we append the keyword “\_clone” for the name field.