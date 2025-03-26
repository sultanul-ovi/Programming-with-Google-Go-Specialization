# Polymorphism

- ability for an object to have different "forms" depending on the context
- example: `Area()` function
  - rectangle, `area = base * height`
  - triangle, `area = 0.5 * base * height`
- identical at a high level of abstraction
- different at a low level of abstraction

## Inheritance

- subclass inherits the methods/data of the superclass
- example: `Speaker` superclass
  - `Speak()` method prints "< noise >"
- subclasses `Cat` and `Dog`
  - also have the `Speak()` method
- Cat and Dog are different forms of speaker
- remember: go does not have inheritence

## Overriding

- subclass redefines a method inherited from the superclass
- example: Speaker, Cat, Dog
  - Speaker `Speak()` prints "< noise >"
  - Cat `Speak()` prints "meow"
  - Dog `Speak()` prints "woof"
- `Speak()` is polymorphic
  - different implementations for each class
  - same signature (name, params, return)
