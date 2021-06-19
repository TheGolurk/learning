# Red Hat Enterprise Linux shell fundamentals

### What is a shell?
Kernel -> Direct access to hardware
	   -> Memory management
	   -> Scheduling
	   -> etc...

 Shell: Outer layer that acts as an interface to kernel level services


see father and child process
```bash
$ pstree
```

### Virtual Consoles
Linux give us a number of virtual consoles to we can connect to

In RHEL based systems we get 7 virtual consoles, and we can access with ALT + F1 to F7 
and CTRL + ALT + F1 to F7 in GUI

### Streams
Shell has 3 standard streams, 1 for error, 1 for input, 1 for output

Shell -> standard out (stdout) File descriptor: 1
	  -> standard in (stdin) File descriptor: 0
	  -> standard error (stderr) File descriptor: 2


#### STDOUT

```bash
$ echo 'Warning alert!!!!' > /tmp/alertfile
$ cat /tmp/alertfile
```

Eviting clobbering
```bash
$ set -o noclobber
```

```bash
$ echo 'test!!' >| /tmp/lsout
$ cat /tmp/lsout
```

Back to clobbering
```bash
$ set +o noclobber
```


OR APPEND TO A FILE

```bash
$ echo 'hola' >> /tmp/alertfile

```

### REDIRECTION
Everything in Unix/Linux is a file, for example sending out from stdout to /dev/pts/2 screen
```bash
$ tty
> /dev/pts/1
# Sending output to the next screen
$ ls -ls /etc/ > /dev/pts/2
$ echo "virus deployed" > /dev/pts/2
```

### STANDARD ERROR

|Stream |Description  | File Descriptor| Operator|
|:--- | ---: | :---:| :---:|
|stdin| Shell standard input stream|0| <|
|stdout| Shell standard output stream|1| >|
|stderr| Shell standard error stream|2| >|

```bash
#stdout
$ ls -l 1> /tmp/file
# > is shorthand for 1> and redirects stdout

#stderr
$ ls -l 2> /tmp/file
```

An example
```bash
$ ls -l /random/folder /home/golurk
> ls: cannot access /random/folder: No such file or directory
> /home/golurk/:
> total 0
> drwxr-xr-x. 2 golurk golurk 17 Jun 8 21:03 Desktop
> ....

# Redirecting the stderr to /tmp/stderr
$ ls -l /random/folder /home/golurk 2> /tmp/stderr
> /home/golurk/:
> total 0
> drwxr-xr-x. 2 golurk golurk 17 Jun 8 21:03 Desktop
.....
$ cat /tmp/stderr
> ls: cannot access /random/folder: No such file or directory

# Files of std allocated in
$ ls -l /dev/std*
```


### The shell environment

```bash
$ echo $BASH_VERSION
$ echo $HOSTNAME
$ echo $SHELL

# Define shell variables
$ defcon=5
$ echo $defcon

$ defcon=1
$ echo $defcon

```
