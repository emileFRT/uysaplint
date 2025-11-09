# unofficial-ysap-fmt

Bash linter and formatter enforcing the [ysap.sh style guide](https://style.ysap.sh).

WARNING: experimental, needs testing, issues welcome

## Installation

```bash
go install github.com/emileFRT/unofficial-ysap-fmt/cmd/unofficial_ysap_fmt@latest
```

## Usage

```bash
# Format to stdout
unofficial-ysap-fmt script.sh

# Format in-place
unofficial-ysap-fmt -i script.sh

# Lint only
unofficial-ysap-fmt lint script.sh

# With shfmt integration (recommended)
unofficial-ysap-fmt -s script.sh

# Disable specific rules
unofficial-ysap-fmt -d eval-usage,useless-cat script.sh
```

## Options

| Flag | Description |
|------|-------------|
| `-i, --inplace` | Edit file in place |
| `-s, --shfmt` | Run shfmt formatter |
| `-d, --disable` | Disable rules (comma-separated) |
| `--binary-next-line` | shfmt: binary ops at start of next line |
| `--no-switch-case-indent` | shfmt: don't indent switch cases |
| `--no-space-redirects` | shfmt: don't space after redirects |
| `--keep-padding` | shfmt: keep padding |
| `--func-next-line` | shfmt: function brace on next line |

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

## Why Some Rules Aren't Auto-Fixed

Short answer: either it might change the script behaviour (ex: removing `set -e`, substituing `[` to `[[`), or it is "complex" and i can't be bothered for now (PR & issues/discussion welcome)

## Why in Go (and not Bash)

Leveraging `mvdan.cc/sh/v3` AST walk is straight forward and clean. Unlike the others option i thought about (and yeah, i use `shfmt)

## Caveats

Same as [shfmt's](https://github.com/mvdan/sh/tree/master?tab=readme-ov-file#caveats) .
Please report any other in an Issue, it will be appreciated.

## License

BSD v3, like mvdan.cc/sh/v3.
