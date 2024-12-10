# Observer Pattern

- It is a behavioral design pattern
- Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically
- Useful when you need to notify multiple objects about the changes in the state of another object
- Useful when you need to decouple the objects, so they don't need to know about each other

### Elements:
- Subject: the object that is being observed, it maintains and notifies the observers
- Observer: the object that observes the subject, it receives the updates from the subject

##  When to use:
- When you want to keep structs loosely connected
- When you want to easily add or remove observers without changing the main subject.
- When you’re dealing with event systems that require various components to react without direct connections.

## When not to use:
- When the subject and observers are tightly coupled
- When observers are fixed and won't change
- When the relationships between objects are simple and don’t require notifications.
- When the order of notifications is crucial, as observers may be notified in an unpredictable sequence.

# Important
Go allow us to make this pattern with highly concurrency and without blocks (with channels and gorutines). That's why there are two examples of  **Observer**. observer.go has the classic pattern and the channels.go has the "go" version of it.

[Source](https://www.geeksforgeeks.org/observer-pattern-set-1-introduction/)