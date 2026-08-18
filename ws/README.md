# Aspect Bazel Starter

This repository uses [Aspect Workflows](https://aspect.build) to provide an excellent Bazel developer experience.
It was generated from the Aspect Workflows template — create your own with `aspect init`
(see https://aspect.build/docs/cli) or from a starter at https://github.com/aspect-starters.

## Getting started

1. Install the Aspect CLI — see https://aspect.build/docs/cli/install. The `aspect`
   command pins this workspace's CLI version (via `.aspect/version.axl`) and wraps
   Bazel, so use it in place of `bazel`.
2. Set up a Bazel-based developer environment with [direnv](https://direnv.net): run `direnv allow`
   and follow the prompts to `bazel run //tools:bazel_env` (puts the project's tools — including
   `aspect` itself — on your PATH, so an `aspect` installed in step 1 isn't required once direnv is active).
3. Build and test everything:

   ```shell
   aspect build //...
   aspect test //...
   ```




## Installing dev tools

For developers to be able to run additional CLI tools without needing manual installation:

1. Add the tool to `tools/tools.lock.json`
2. Run <code>bazel run //tools:bazel_env</code> (following any instructions it prints)
3. When working within the workspace, tools will be available on the PATH

See https://aspect.build/blog/run-tools-installed-by-bazel for details.










