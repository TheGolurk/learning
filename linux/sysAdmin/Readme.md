### Working with users and groups

## Admin powers: Best Practices
* Avoid using the root account
* Create unique accounts for each user
* Assign only necesary authority to each user
* Use admin power only via sudo

```bash
# Contains encrypted users passwords
$ less /etc/shadow

$ less /etc/passwd

$ less /etc/group

$ id {user}

$ who

$ w

$ last | less

# Create new user
$ sudo useradd -m jane

# skeleton files and directory
$ ls -a /etc/ske

# skeleton files and directory
$ ls -a /etc/skell
```
