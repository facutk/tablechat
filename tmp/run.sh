#!/usr/bin/env bash
set -e

# load .env if present and export everything it defines
if [ -f .env ]; then
  set -a
  . .env
  set +a
fi

# exec the built binary so signals are forwarded correctly
exec ./tmp/main "$@"