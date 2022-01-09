**Visitor** in Go
=================

![Visitor Design&nbsp;Pattern](https://refactoring.guru/images/patterns/content/visitor/visitor.png)

**Visitor** is a behavioral design pattern that allows adding new behaviors to existing class hierarchy without altering any existing code.

[Learn more about Visitor](https://refactoring.guru/design-patterns/visitor)

Conceptual Example
------------------

The Visitor pattern lets you add behavior to a struct without actually modifying the struct. Let’s say you are the maintainer of a lib which has different shape structs such as:

*   Square
*   Circle
*   Triangle

Each of the above shape structs implements the common shape interface.

Once people in your company started to use your awesome lib, you got flooded with feature requests. Let’s review one of the simplest ones: a team requested you to add the `getArea` behavior to the shape structs.

There are many options to solve this problem.

The first option that comes to the mind is to add the `getArea` method directly into the shape interface and then implement it in each shape struct. This seems like a go-to solution, but it comes at a cost. As the maintainer of the library, you don’t want to risk breaking your precious code each time someone asks for another behavior. Still, you do want other teams to extend your library somehow.

The second option is that the team requesting the feature can implement the behavior themselves. However, this is not always possible, as this behavior may depend on the private code.

The third option is to solve the above problem using the Visitor pattern. We start by defining a visitor interface like this:

    type visitor interface {
        visitForSquare(square)
        visitForCircle(circle)
        visitForTriangle(triangle)
    }

The functions `visitForSquare(square)`, `visitForCircle(circle)`, `visitForTriangle(triangle)` will let us add functionality to squares, circles and triangles respectively.

Wondering why can’t we have a single method `visit(shape)` in the visitor interface? The reason is that the Go language doesn’t support method overloading, so you can’t have methods with the same names but different parameters.

Now, the second important part is adding the `accept` method to the shape interface.

    func accept(v visitor)

All of the shape struct needs to define this method, similarly to this:

    func (obj *square) accept(v visitor){
        v.visitForSquare(obj)
    }

Wait a second, didn’t I just mentioned that we don’t want to modify our existing shape structs? Unfortunately, yes, when using the Visitor pattern, we do have to alter our shape structs. But this modification will only be done once.

In case adding any other behaviors such as `getNumSides`, `getMiddleCoordinates`, we will use the same `accept(v visitor)` function without any further changes to the shape structs.

In the end, the shape structs just need to be modified once, and all future requests for different behaviors could be handled using the same accept function. If the team requests the `getArea` behavior, we can simply define the concrete implementation of the visitor interface and write the area calculation logic in that concrete implementation.