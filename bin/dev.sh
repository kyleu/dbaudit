#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Starts the app, reloading on changes

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/.."

# $PF_SECTION_START(keys)$
export db_user=sa
export db_password=aStrongPassword123!
export db_database=dbaudit
# $PF_SECTION_END(keys)$

[[ -f "$HOME/bin/oauth" ]] && . "$HOME/bin/oauth"
export dbaudit_encryption_key=TEMP_SECRET_KEY

# include env file
if [ -f ".env" ]; then
  while IFS= read -r line || [ -n "$line" ]; do
    if [[ ! $line =~ ^#.* ]]; then
      export "$line"
    fi
  done < ".env"
fi

./bin/templates.sh
go mod tidy

ulimit -S -n 65536

air
