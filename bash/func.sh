#!/bin/bash

function quit {
	exit
}

function say {
	echo $1!
}

say hello
say world
quit
echo foo
