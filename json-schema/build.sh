#!/usr/bin/env sh
set -eu

cd schema
bundle exec prmd combine --meta meta.json schemata/ > schema.json \
&& echo 'Success generating Schema and Docs'
