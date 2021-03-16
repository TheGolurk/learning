#!/bin/bash

for i in $( ls ); do
	echo item: $i
done

# C/PERL Like
COUNTER=0
while [	$COUNTER -lt 10 ]; do
	echo The counter is $COUNTER
	let COUNTER=COUNTER+1
done
