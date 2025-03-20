# Adapter Pattern

- Structural design pattern
- Allows incompatible interfaces working together. Acts as a bridge between 2 interfaces

# Advantages

- Useful for legacy code or third party libraries because decouples the code from it
- Allows interface compatibility without modifying existing code
- Improves maintainability by keeping changes isolated in the adapter

# Example keys

In the example we have a logger of an application and a logger from a third party entity. We are
going to adapt the third party logger to our logger.
