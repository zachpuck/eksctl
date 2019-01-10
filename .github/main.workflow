workflow "Build every commit" {
  on = "push"
  resolves = ["Run: make build"]
}

action "Run: make build" {
  uses = "./build/"
  runs = "make build"
}
