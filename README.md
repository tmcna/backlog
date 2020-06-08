# BacklogCLI

BacklogCLI ia a CLI application for Backlog users.

# Installation
  
```
$ git clone https://github.com/tmcna/backlog
$ cd backlog 
$ cd cmd/backlogcli/
$ go build

```

# Setup

```
$ mkdir ~/.backlogcli
$ export BACKLOG_CLI=~~/.backlogcli
$ cd $BACKLOG_CLI
$ echo "API_KEY" > apikey.txt
$ echo "https://space-name.backlog.com" > space.txt
```


# Usage
```
NAME:
   backlogcli - A CLI application for Backlog users.

USAGE:
   backlogcli [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   user, u     List of users in your space.  
   act, a      Recent updates in your space.  
   notify, n   Updates space notification.  
   space, s    Information about space disk usage.  
   project, p  Operations project  
   issue, i    Operations issue  
   comment, c  Operations comment  
   help, h     Shows a list of commands or help for one command  

GLOBAL OPTIONS:
   --help, -h     show help (default: false)  
   --version, -v  print the version (default: false)  
```

# Author
Kazuhiro Kawachi

# License
MIT License
