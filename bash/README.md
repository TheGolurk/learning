# Learning bash

[Bash-HOW-TO](https://tldp.org/HOWTO/Bash-Prog-Intro-HOWTO-1.html)

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


