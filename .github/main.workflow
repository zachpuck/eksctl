workflow "Build" {
  on = "push"
  resolves = ["pull build image"]
}

action "pull build image" {
  uses = "actions/docker/cli@76ff57a6c3d817840574a98950b0c7bc4e8a13a8"
  runs = "pull weaveworks/eksctl:build"
}
