# ADReaper

```ADReaper``` is a tool written in ```Golang``` which enumerate a Active Directory environment with LDAP queries within few seconds

## Installation 

You can download precompiled executable binaries for Windows/Linux from [latest releases](https://github.com/AidenPearce369/ADReaper/releases/tag/ADReaper)

### Install from source

To build from source, clone the repo and build it with GO

```c
$ git clone https://github.com/AidenPearce369/ADReaper
$ cd ADReaper/
$ go build
```

## Usage

ADReaper performs enumeration with various commands that performs LDAP queries with respective to it

```c
monish@chimera:/ADReaper$ ./ADReaper
  -command string
        Command to run

                users           - to list all users
                user-logs       - to list user session activities
                never-loggedon  - to list users never logged on
                groups          - to list all groups with members
                computers       - to list all computers
                dc              - to list domain controllers
                gpo             - to list group policy objects
                spn             - to list service principal objects
                admin-priv      - to list AD objects with admin privilege
                domain-trust    - to list domain trust
                ou              - to list organizational units
                ms-sql          - to list MS-SQL servers

  -dc string
        Enter the DC
  -password string
        Enter the Password
  -user string
        Enter the Username
        
```

## To-Do

Looking forward for contributors to build the next version

Planned features,
- Custom LDAP querying
- Filters with existing commands
- PrivEsc checker
- LAPS enumeration
- Kerberoasting
- Local admin access hunting
- Registry analysis

If interested, ping me :)
