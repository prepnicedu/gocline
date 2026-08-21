# Aspect Bazel Starter for Go

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

   A small `hello/` sample is included for each selected language as a starting point.




## Formatting code

`format` is put on your PATH by `bazel_env` (via `.envrc`).

- Run `format` to re-format all files locally.
- Run `format path/to/file` to re-format a single file.
- On CI, run the `format` task to verify formatting; see https://aspect.build/docs/cli/tasks-ci

## Linting code

This project uses [rules_lint](https://github.com/aspect-build/rules_lint) to run linters as Bazel
aspects. The linters are configured in `.aspect/config.axl` and run via the Aspect CLI's `lint`
command (not the upstream Bazel CLI), which collects the cached report files, applies fixes
interactively, and sets a matching exit code.

- Run `aspect lint //...` to check for lint violations.



## Installing dev tools

For developers to be able to run additional CLI tools without needing manual installation:

1. Add the tool to `tools/tools.lock.json`
2. Run <code>bazel run //tools:bazel_env</code> (following any instructions it prints)
3. When working within the workspace, tools will be available on the PATH

See https://aspect.build/blog/run-tools-installed-by-bazel for details.







## Working with Go modules

After adding a new `import` statement in Go code, run `aspect gazelle` to update the BUILD file.

If the package is not already a dependency of the project, you have to do some additional steps:

```shell
# Update go.mod and go.sum, using same Go SDK as Bazel (it comes from direnv)
% go mod tidy -v
# Update MODULE.bazel to include the package in `use_repo`
% bazel mod tidy
# Repeat
% aspect gazelle
```






## Delivering container images

`oci_push` targets (e.g. the `:image_push` from the `hello/` sample's `go_image`) are built and
pushed by `aspect delivery`. The delivery query and release flags are configured in
`.aspect/config.axl`; point each `oci_push`'s `repository` at your registry (the sample uses the
anonymous, ephemeral `ttl.sh`).

- Run `aspect delivery` to build + push deliverables.


