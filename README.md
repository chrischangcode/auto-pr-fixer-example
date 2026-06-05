# auto-pr-fixer-example

Demo repo for [auto-pr-fixer](https://github.com/chrischangcode/auto-pr-fixer).

## Scenario

A dependency bump PR updates `go-chi/chi` from v5.0.12 to v5.2.0. The new
version renamed `middleware.Logger` to `middleware.DefaultLogger`, breaking
the build. auto-pr-fixer detects the CI failure, calls the GitHub Models API,
and commits the fix directly to the PR branch.
