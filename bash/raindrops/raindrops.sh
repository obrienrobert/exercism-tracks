#!/usr/bin/env bash

main() {
  rainDrop=''

  if [[ $(($1 % 3)) == 0 ]]; then
    rainDrop+="Pling"
  fi

  if [[ $(($1 % 5)) == 0 ]]; then
    rainDrop+="Plang"
  fi

  if [[ $(($1 % 7)) == 0 ]]; then
    rainDrop+="Plong"
  fi

  if [[ -z "$rainDrop" ]]; then
    rainDrop="$1"
  fi

  echo "$rainDrop"
}

main "$@"
