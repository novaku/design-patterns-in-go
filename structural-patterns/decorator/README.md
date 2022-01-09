**Decorator** in Go
===================

![Decorator design pattern](https://refactoring.guru/images/patterns/content/decorator/decorator.png)

**Decorator** is a structural pattern that allows adding new behaviors to objects dynamically by placing them inside special wrapper objects.

Using decorators you can wrap objects countless number of times since both target objects and decorators follow the same interface. The resulting object will get a stacking behavior of all wrappers.

[Learn more about Decorator](https://refactoring.guru/design-patterns/decorator)

Conceptual Example
------------------

Decorator design pattern is a structural design pattern. It lets you provide additional functionality or decorates an object without altering that object.

It is better understood with an example. Imagine you are opening a pizza chain. You started with two kinds of pizzas

*   Veggie Mania Pizza
*   Peppy Tofu pizza

Each of the above pizza had its price. So you would create a pizza interface as below

    package main
    
    type pizza interface {
    	getPrice() int
    }

You need to also create two pizza struct with a **getPrice** function which will return the price. These two pizza structs implement the pizza interface as they define the getPrice() method

Later on, you started to offer toppings along with the pizza with some additional price for each of the topping. So the original base pizza now needs to be decorated with a topping. Imaging you added below two toppings to the menu

*   Tomato Topping
*   Cheese Topping

Also, remember that pizza along with the topping is also a pizza. Customer can choose their pizza in different ways. For eg

*   Veggie mania with tomato topping
*   Veggie main with tomato and cheese topping
*   Peppy Paneer pizza without any topping
*   Peppy Paneer pizza with cheese topping
*   …

So how would you design now given that you now have the toppings as well. Decorator pattern will come into picture. It can help additional functionality without actually modifying any of the existing structs. Decorator pattern recommends in this case to create separate structs for each of the topping available. Each topping struct will implement the pizza interface above and also have an embed and instance of pizza.

We now have separate structs for different types of pizza and separate struct for the types of topping available. Each of the pizza and topping has its own price. And whenever you add any topping to a pizza then price of that topping is added to the price of base pizza and that is how you get a resultant price.

So the decorator pattern let’s you decorate the original base pizza object without altering that pizza object. The pizza object knows nothing about toppings. It just knows its price and nothing else.