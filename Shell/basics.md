# Bash Scripting
* Shell is a command interpreter (i.e. REPL)
* Supports shell scripting
* Usually runs POSIX commands (i.e. the standard for command line and scripting interface)
* Example of shell: Bash, Zsh, fish

## Profile
* Edit prompt by changing `$PS1`

## I/O Redirection
* `process > out.txt` or `process >2 err.txt` for writting to fd 1 or 2 respectively for stdout or stderr
* Use `process >& outerr.txt` to output both stdout and stderr to common channel

## Commands
* Each command is associated to a program
* The program is usually retrieved from `$PATH`
* Use `which <command>` to check the abs path of command
* `~` Accesses home, `/` accesses root (not the same unless home is set as root)
* Command support args, flags and options (flag with value)
* Use `man <command` to access command manual or use [https://tldr.ostera.io ](https://tldr.ostera.io) 

### chmod, chown chgrp
* Each file/dir has permission of type read, write, exec (rwx) in permission group of user/owner, group, others (ugo)
* Each file/dir can only have 1 user and 1 group.

Command to change permission:
* `chmod <who><op><permission> <file/dir>` example `chmod u=rw <file/dir>`

File:
* `x` permits the running of a executable `./file_with_x.exe`
* Without `x` permission, you have to `sh file_with_x.ext` to run the interpreter over file

Directory:
* `r` permits the reading of the entries in a directory (e.g. `ls <dir>` works)
* `w` by itself **does not permits** the editing the directory entries to create/edit/remove content with it
* `x` permits the accessing of the directory and the executing of command on the content within
* Note:
  * `wx` permits create/edit/remove of directory's file/dir within (as long as you are aware of the path of the file/dir) 
  * `x` permits the editing of the **content** of file/dir within the dir (e.g. editing the file text) since this does not involve the changing of the directory's entries. But no renaming, deleting and creating of file/dir.
  * `t` to prevents the deletion of file/dir by non-owner

### sudo
* To run command as other user e.g. `sudo chmod 777 a.txt`
* Note: Not to be confused with `su` to switch user. Often used together `sudo su` to switch to root user by running the command as root first

### find, fd, locate
* Search by file name, type and other attributes

### cron
* To setup scheduled commands
* Or use `at`

### grep
* Search through content

### Process Metric
* `w` or `uptime` to check system avg usage e.g. average load per 1, 5 and 15 mins
* `top` stats of running processes
* `ps` for list of processes

### Jobs
* By default process run in foreground, to run in background on start `<command> &`
* To suspend foreground process, `CTRL-Z`, to terminate `CTRL-C`
* Run `jobs` to view jobs and pid and use `bg/fg <pid>` to bring process back/fore (without pid the default will be the most recent)

## Scripting
### Special Character
* [Special Character](https://tldp.org/LDP/abs/html/special-chars.html)

### Shebang
* start with a `#!/bin/bash`
* For resolve program where path is not known use env command `#!/usr/bin/env python`

### Boolean Expression
* Read more on `man test`
