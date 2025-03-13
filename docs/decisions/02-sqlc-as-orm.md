# 02 - SQLC as ORM for accessing data

**Status:** Proposed

## Context and Problem Statement

When interacting with a database writing everything from hand (connection, queries, models etc.) is rarely a good idea. Object-
Relational-Mappers (ORM's) help to abstract some of these parts in order to improve productivity, reliability, security and other
aspects.

## Considered Options

1. [GORM](https://gorm.io/index.html)
2. [SQLC](https://sqlc.dev/)

## Decision Outcome

Chosen option: "SQLC", since it ensures compatibility and type-safety between queries and db-schema at the same time. Additionally
it allows/demands one to write all queries from hand in order to optimize for the schema at hand and fine tune them if necessary.

### Consequences

- Good, because it aligns db-schema and queries in order to prevent a discrepancy between them causing errors during runtime.
- Good, because it forces one to write their own queries, allowing for customization and fine-tuning.
- Good, because it consists of minimal abstractions, making the tool easily understandable for newcomers.
- Bad, because it forces one to write their own queries, slowing down the development speed and increasing the initial building cost.
- Potentially bad, because it lacks some quality of life features compared to GORM. However these wont be needed for the current
  scope of the project.
