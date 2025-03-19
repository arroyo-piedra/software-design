# Prototype Pattern

- Creational design pattern
- Allows cloning existing objects instead of creating new ones from scratch

# Advantages:

- _Cloning Instead of Instantiation_: New objects are created by copying an existing one.
- _Flexibility_: Different prototypes can be used to create objects with slight variations.
- _Performance Optimization_: Reduces the overhead of creating objects dynamically.

# Important Bullet in golan:

- If your object contains pointers, ensure deep copying to avoid modifying the original.

# When to use?

- When object/struct creation is expensive or complex

# Important

There are 2 examples of prototypes, this is because there is an easy example to understand the
basics behaviors of the pattern and there is a complex one to understand the nested handling of
nested data.
