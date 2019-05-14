action "Lint" {
    uses = "actions-contrib/golangci-lint@master"
    args = "run --new-from-rev master"
}

workflow "Lint" {
    on = "push"
    resolves = ["Lint"]
}
