#!/bin/env sh

go test -v  > test.out

test_output="$(cat test.out)"

test_pass_count=$(echo -e "$test_output" | grep "ok  	github.com/ipatser/vcs_ipatser" | wc -l)
test_fail_count=$(echo -e "$test_output" | grep "FAIL	github.com/ipatser/vcs_ipatser" | wc -l)

echo "===================="

# set status based on result
if [ "$test_pass_count" = "1" ]; then
    echo "TEST PASSED"
elif [ "$test_fail_count" = "1" ]; then
    echo "TEST FAILED"
else
    echo "TEST RESULT UNKNOWN"
fi

echo "===================="
echo -e "$test_output"
echo "===================="

# upload to artifactory
