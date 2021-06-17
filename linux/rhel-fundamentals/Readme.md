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
$ ls -ls /etc/ > /dev/pts/2
```

