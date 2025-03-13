# 03 - Dependency Injection in combination with an IoC Container

**Status:** Proposed

## Context and Problem Statement

Most a bit more complex applications have components coupled to others, for instance: _a service needs an instance of a repository in
order to access data stored in the database_. Depending where the repository is instantiated the coupling is lighter or tighter
(outside vs. inside the service). Additionally it has a direct impact how testable a component is e.g. how easily some mocks/fakes
can be provided for isolating the scope of tests.

## Considered Options

- Dependency Injection as principle on its own
- Dependency Injection with an IoC (Inversion of Control) container
- none of them

## Decision Outcome

Chosen option: "Dependency Injection with an IoC container", since the benefits of dependency injection principles (testability,
interchangeability & lighter coupling) lie at hand. When combined with an IoC container the ergonomics of automatically provisioning
the correct dependencies for each component and the other features outweigh the initial setup efforts.

[fx](https://github.com/uber-go/fx) by Uber was chosen as an IoC container since it is "battle tested" on Uber's production services
and provides some nice features like lifecycle hooks for controlling the application startup/shutdown.

### Consequences

- Good, because it implicitly enforces testable code-design
- Good, because it provides interchangeability of components and lighter coupling between them
- Good, because it automatically provisions the necessary dependencies for each component instance
- Bad, because it requires some additional boiler plate and initial setup efforts, however the return of investment should be quite
  fast.
