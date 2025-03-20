<img src="./docs/img/logo.png" width="128" />

# stock-sight

[![CI](https://github.com/ruegerj/stock-sight/actions/workflows/ci.yaml/badge.svg)](https://github.com/ruegerj/stock-sight/actions/workflows/ci.yaml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ruegerj_stock-sight&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ruegerj_stock-sight)

Terminal based stock watcher (SWAT project)

## Contents

- [Requirements, Scope & Stories](./docs/requirements.md)
- [Architecture Decisions](./docs/decisions/decisions.md)
- [Contributing](./CONTRIBUTING.md)
- [Development Environment](#development-environment)

## Links

- [KanBan Board](https://github.com/users/ruegerj/projects/2)
- [SonarQube](https://sonarcloud.io/project/overview?id=ruegerj_stock-sight)

## Development Environment

The commands listed below are the most frequentily used. In order to list all available commands with their description run the following command:

```bash
task --list
```

### Run

```bash
# example: task run cmd="hello --name anon"
task run cmd="<cmd + args>"
```

### Build

```bash
task build
```

### Test

```bash
task test
```

### Schema Generate

(Re-) Generates the query code based on the given [db- & query-definitions](./internal/embedded/db)

```bash
task query:generate
```
