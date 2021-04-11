#!/bin/bash
set -ex

if [[ -z $EXPECTED_COV ]]; then EXPECTED_COV=0  ; fi;

run_tests () {
  make "$1" | tee test.out;
  cat test.out | grep "total:" | awk '{print substr($5, 1, length($5)-1)}' | {
      read result;
      if (( $(awk 'BEGIN {print ("'$result'" < "'$EXPECTED_COV'")}') ));
      then echo "FAILURE: Coverage is $result < $EXPECTED_COV%"; exit 1;
      else echo "SUCCESS: Coverage is $result expected was $EXPECTED_COV"; exit 0;
      fi
  }
}

EXPECTED_COV=55 run_tests docker-test
make end-test
