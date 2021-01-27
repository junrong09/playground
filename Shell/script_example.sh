#!/bin/bash

# Variable Assignment:
foo=bar # Assignment $foo with value "bar". Cannot have space e.g. "foo = bar" or it will be interpreted as foo with argument "=" and "bar"
echo "$foo"


# Script arguments:
# $0 - Name of the script
# $1 to $9 - Arguments to the script. $1 is the first argument and so on.
# $@ - All the arguments
# $# - Number of arguments
# $? - Return code of the previous command
# $$ - Process identification number (PID) for the current script
# !! - Entire last command, including arguments. A common pattern is to execute a command only for it to fail due to missing permissions; you can quickly re-execute the command with sudo by doing sudo !!
# $_ - Last argument from the last command. If you are in an interactive shell, you can also quickly get this value by typing Esc followed by .
echo "Running program $0 with $# arguments with pid $$"


# Command substitution:
echo "Starting program at $(date)" # $(command) will be resolve and output substituted
echo "${foo}_more_string" # ${variable} parameter expansion or can use $foo but cannot concate string behind
#echo "$foo_more_string" wrong as bash will treat the whole foo_more_string as variable name
diff <(ls .) <(ls .) # <(command) will be resolve, output written to temp file and temp filename being substituted


# Globbing:
touch {a,b,c}.{text,in} # Set permutation
rm ?.text # single character wildcard
rm *.in # multiple character wildcard


# I/O Redirection:
ls unknownFolder . > stdout.txt 2> stderr.txt # Output std out and error to different channel
ls unknownFolder . &> stdouterr.txt # Aggregate both std out and err into common channel
# User >> to append instead of replacing file
echo errmsg >&2 # Channel stdout to stderr (i.e. use it to write to stderr)
ls unknownFolder 2> /dev/null # To discard output
rm std{out,err,outerr}.txt # Clean up


# Functions:
# Declare function
function echo_sample_function {
    for i in $@; do
        echo $i
    done

    # echo $@
    local exitCode=2 # local variable
    foo=bar2 # Global scope, affects all $foo
    return $exitCode
    exit 2 # same as return
}
# Call function
echo_sample_function "input1" $foo; echo $?


# Boolean comparison:
# Return exit code, 0 is true, 1 is false
# For numeric; use lt (less), gt (greater), eq (equal), ne (not equal), le (less than or equal), ge (greater than or equal)
[ 5 -gt 2 ]; echo $? # Note: there must be spaces between square bracket and the number
[ 5 -lt 2 ]; echo $?
#[ 5 < 2 ] && [ 5 > 2 ] && echo "All true" # Return true because it will be treated as redirection
# For string; use >, <, =, != (there isnt >= or <=)
[ $foo == bar2 ]; echo $? # String matched, return true
[ ! $foo == foo ]; echo $? # Negation sign
[ -z "" ]; echo $? #True when string is empty
[ "a" -eq "b" ] # Error when used num comparator on strings
# For files; Read more under "CONDITIONAL EXPRESSIONS" in "man bash"
[ -e ./a.txt ]; echo $? #True when file exist
# Read More on new version of test [[]]: http://mywiki.wooledge.org/BashFAQ/031


# Conditional Statements:
val=200
if [ $val -le 100 ] && [ $val -gt 50 ]; then
    echo "Value less than or equal 100"
elif [ $val -le 200 ]; then
    echo "Value less than or equal 200"
else
    echo "Value greater than 200"
fi


# Loops:
# For loop
echo "cat" >> list.txt
echo "dog" >> list.txt
echo "pig" >> list.txt
for i in $(cat list.txt); do
    echo $i
done
rm list.txt
# While loop
count=0
while [ $count -lt 5 ]; do
    echo "counting ${count}..."
    let count+=1
done

# Arithmetic
# Using arith expansion
a=1
a=$(( $a + 1 ))
echo $a; echo $(( 100 * 234 ))
# Using expr; For evaluation
expr 2 + 2
a=$(expr 2 + 2)
# Using let; For evaluate and store in variable
let a=4 + 2
let a++
# Using bc; For decimal
echo 'scale=2; 8.3/2.4' | bc #division with result to 2 d.p.
echo 'scale=2; sqrt(50)' | bc