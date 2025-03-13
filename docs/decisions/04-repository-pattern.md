# 04 - Repository Pattern as data-access abstraction

**Status:** Proposed

## Context and Problem Statement

Data-access logic (e.g. queries) can easily pollute the domain code of an application as it grows and gets complexer. This leads to
deteriorated readability of the _actual_ business logic, making the code harder to maintain. In addition the coupling to the used
persistence technology is omnipresent, causing for big refactoring costs, should it be exchanged in the future.

## Considered Options

- [Repository Pattern](https://martinfowler.com/eaaCatalog/repository.html)
- no abstraction for data-access

## Decision Outcome

Chosen option: "Repository Pattern", because it provides a neat and lightweight abstraction which favors decoupling of the concrete
data-access- from the domain-logic. Thus enabling light coupling and easy interchangeability and testability of the
persistence-layer beneath.

### Consequences

- Good, because it enables light coupling, testable code and interchangeability of the persistence-layer.
- Good, because it reduces the complexity of business logic and thus enhances the readability and maintainability.
- Bad, because it adds some additional "boiler plate" code, however for the maintenance in the long run this shouldn't be a notable
  factor.
