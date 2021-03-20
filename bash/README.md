# Learning bash

[Bash-HOW-TO](https://tldp.org/HOWTO/Bash-Prog-Intro-HOWTO-1.html)

## stdout ##
```
# stdout 2 file
ls -l > ls-l.txt

# stderr 2 file
grep da * 2> grep-errors.txt

# stdout 2 stderr
grep da * 1>&2

# stderr 2 stdout
grep * 2>&1

# stderr and stdout 2 file
# This will place every output of a program to a file. This is suitable sometimes for cron entries, if you want a command to pass in absolute silence.
rm -f $(find / -name core) &> /dev/null
```


## pipes ##
```
#  Here, the following happens: first the command ls -l is executed, and it's output, instead of being printed, is sent (piped) to the sed program, which in turn, prints what it has to. 
ls -l | sed -e "s/[aeio]/u/g"

ls -l | grep "\.txt$"
```

# Useful Commands #

## sed ##
Sed is a non-interactive editor. Instead of altering a file by moving the cursor on the screen, you use a script of editing instructions to sed, 
plus the name of the file to edit. You can also describe sed as a filter. Let's have a look at some examples:

```
        $sed 's/to_be_replaced/replaced/g' /tmp/dummy
```

 Sed replaces the string 'to_be_replaced' with the string 'replaced' and reads from the /tmp/dummy file. The result will be sent to stdout (normally the console)
 but you can also add '> capture' to the end of the line above so that sed sends the output to the file 'capture'.
