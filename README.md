# unofficial ysap linter

Bash linter and formatter enforcing the [ysap.sh style guide](https://style.ysap.sh).

WARNING:  **WIP** experimental, needs testing, issues welcome

## Installation

```bash
go install github.com/emileFRT/uysaplint/cmd/ysaplint@latest
```

## Usage

```bash
# Format to stdout
ysaplint fmt script.sh

# Lint only
ysaplint script.sh

```

## Rules

Those are my extraction/understanding of the [ysap.sh style guide](https://style.ysap.sh).

| Rule | Auto-Fix | Severity | Description |
|------|----------|----------|-------------|
| `shebang` | ✅ | warning | Use \#!/usr/bin/env bash |
| `semicolon` | ✅ | warning | Avoid semicolons except in control statements |
| `function-keyword` | ✅ | error | Use `name()` not `function name` |
| `test-bracket` | ❌ | error | Use \[[ ]] instead of `[ ]` or `test` |
| `seq-usage` | ❌ | error | Use `{1..5}` or `((i=0;i<n;i++))` |
| `backticks` | ✅ | error | Use `$(...)` instead of backticks |
| `let-command` | ❌ | error | Use `((...))` instead of `let` |
| `parse-ls` | ❌ | error | Never parse `ls` output |
| `unquoted-var` | ❌ | warning | Quote variables: `"$var"` |
| `useless-cat` | ❌ | warning | Use `< file` not `cat file |` |
| `eval-usage` | ❌ | error | Never use `eval` |
| `set-errexit` | ❌ | warning | Avoid `set -e` |
| `block-statement` | ❌ | error | `then`/`do` on same line as `if`/`while`/`for` |
| `blank-lines` | ✅ | warning | Max 1 blank line in a row |
| `var-naming` | ❌ | warning | Use lowercase variable names |
| `declaration` | ✅ | warning | Don't use `readonly` or `declare -i` |

## Why so few rules support formatting

Short answer: either it might change the script behaviour (ex: removing `set -e`, substituing `[` to `[[`), or it is somewhat hard t implement and i can't be bothered for now (issues discussion welcome if needed/implementation proposal)

## Why in Go (and not Bash or <insert lang>...)

Leveraging `mvdan.cc/sh/v3` AST walk is straight forward, clean and almost perfectly complete. Unlike the others options i thought about.

## Caveats

Same as [shfmt's](https://github.com/mvdan/sh/tree/master?tab=readme-ov-file#caveats) .
Please report any other in an Issue, it will be appreciated.

## License

BSD v3, like mvdan.cc/sh/v3.
