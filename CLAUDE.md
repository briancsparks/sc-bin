# CLAUDE.md - Guidelines for sc-bin

## Commands
- No formal build system - these are primarily bash scripts
- No formal tests - manually test scripts as needed
- For linting bash scripts: `shellcheck scriptname.sh`

## Code Style Guidelines
- Scripts start with `#!/bin/bash -e`
- Use standard functions: `errcho()`, `die()`, `ensurecmd()`, `announce()`
- Error handling: use `die "error message"` to exit with error
- Tab indentation (use `.editorconfig` settings)
- Variable naming: lowercase_with_underscores
- Always quote variables: `"$var"` not `$var`
- Use `[[ ]]` for conditionals, not `[ ]`
- Path handling: get script directory with `script_dir="$( cd "$(dirname "$( readlink -f "${BASH_SOURCE[0]}" )" )" && pwd )"`
- Template comments format: `###A` for one-line description, `###B` for longer descriptions
- Include usage instructions in functions when appropriate
- Command validation: use `ensurecmd` for required commands

## Sourcing Common Functions
- Common functions can be placed in the `inc/` directory
- Use the new `inc/utils` file for standard functions (errcho, die, ensurecmd, announce)
- Source common functions at the top of your script:
  ```bash
  script_dir="$( cd "$(dirname "$( readlink -f "${BASH_SOURCE[0]}" )" )" && pwd )"
  source "$script_dir/../inc/utils"  # Adjust path as needed
  ```
- For colors and formatting, use: `source "$script_dir/../inc/colors"`
- Check if script is being run or sourced with: `if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then`