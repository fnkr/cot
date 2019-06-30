workflow "Test, Build, Release" {
  on = "push"
  resolves = ["Test", "Build", "Release"]
}

action "Test" {
  uses = "./.ci"
  args = "/test"
}

action "Build" {
  uses = "./.ci"
  args = "/build"
}

action "Release" {
  uses = "fnkr/github-action-ghr@v1"
  needs = ["Test", "Build"]
  secrets = ["GITHUB_TOKEN"]
  env = {
    GHR_PATH = "_build/"
    GHR_COMPRESS = "gz"
  }
}
