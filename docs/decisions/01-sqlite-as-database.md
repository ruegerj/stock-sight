# 01 - SQLite as Database for Persistence

**Status:** Proposed

## Context and Problem Statement

For stock-sight various data like stock-buys / -sells, tracked stocks etc. are entered by the user. This data should be persisted in
order to be accessible in the future (you want to know what stocks you bought and when). According to a [non-functional requirement
nr. 4](https://github.com/ruegerj/stock-sight/blob/main/docs/requirements.md#requirements) a _portable_ datastore is demanded.

## Considered Options

1. [MongoDB](https://www.mongodb.com/)
2. [PostgreSQL](https://www.postgresql.org/)
3. [SQLite](https://www.sqlite.org/)

## Decision Outcome

Chosen option: "SQLite", since it is the only database that is truly portable without any caveats (thus supporting the
non-functional requirement mentioned above). In addition a local installation of the other options would be possible, however it
would lead to both installation complexity and performance overhead on the users machine.

### Consequences

- Good, because it is just a single file at rest and thus truly portable.
- Good, because it doesn't add any performance overhead on the users machine.
- Potentially Bad, because the scalability could become a theoretical issue with a huge amount of data. However in practice this is
  neglectful, since projects like [PocketBase](https://pocketbase.io/) prove the immense load capabilities of SQLite, which aren't
  expected to be exceeded in the projects scope.
