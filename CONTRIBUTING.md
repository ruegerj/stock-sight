# Contributing

## Local tools

In order to develop locally on stock-sight, make sure you have the following tools installed in your environment:

- [go 1.24](https://go.dev/)
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)
- [golangci-lint](https://golangci-lint.run/welcome/install/#local-installation)
- [task](https://taskfile.dev/installation/)
- [act (optional)](https://nektosact.com/installation/index.html)

## Project Management

All workitems (stories, tasks etc.) are tracked as [GitHub issues](https://github.com/ruegerj/stock-sight/issues).
A [KanBan board](https://github.com/users/ruegerj/projects/2) (GitHub project) is used to get an overview of the status of each task.
The current requirements and stories are tracked [here](./docs/requirements.md).

## Git Workflow

The workflow should be pragmatic and enable productivity, however the following guardrails apply:

- every meaningful change should be done in a dedicated feature branch
- every feature branch should be either _squash-merged_ or _rebased_ back onto `main`
- every meaningful change should be peer-review via pull-request
- CI should be passing before a change gets onto `main`
