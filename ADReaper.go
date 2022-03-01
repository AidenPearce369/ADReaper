package main

import (
	"flag"
	"monish/ADReaper/ldapquery"
	"os"
	"strings"
)

// Main function

func main() {

	commandStr := "Command to run\n"
	commandStr += "\n\tusers           - to list all users"
	commandStr += "\n\tuser-logs       - to list user session activities"
	commandStr += "\n\tnever-loggedon  - to list users never logged on"
	commandStr += "\n\tgroups          - to list all groups with members"
	commandStr += "\n\tcomputers       - to list all computers"
	commandStr += "\n\tdc              - to list domain controllers"
	commandStr += "\n\tgpo             - to list group policy objects"
	commandStr += "\n\tspn             - to list service principal objects"
	commandStr += "\n\tadmin-priv      - to list AD objects with admin privilege"
	commandStr += "\n\tdomain-trust    - to list domain trust"
	commandStr += "\n\tou              - to list organizational units"
	commandStr += "\n\tms-sql          - to list MS-SQL servers"
	commandStr += "\n"

	var cmd = map[string]bool{
		"users":          true,
		"user-logs":      true,
		"never-loggedon": true,
		"groups":         true,
		"computers":      true,
		"dc":             true,
		"gpo":            true,
		"spn":            true,
		"admin-priv":     true,
		"domain-trust":   true,
		"ou":             true,
		"ms-sql":         true,
	}

	//getting args data
	ldapServer := flag.String("dc", "", "Enter the DC")
	ldapBind := flag.String("user", "", "Enter the Username")
	ldapPassword := flag.String("password", "", "Enter the Password")
	command := flag.String("command", "", commandStr)
	flag.Parse()
	if len(*ldapServer) == 0 || len(*ldapBind) == 0 || len(*ldapPassword) == 0 || len(*command) == 0 || !(cmd[*command]) {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	//formatting data
	s := strings.Split(*ldapServer, ".")
	baseDN := ""
	for x := 1; x < len(s); x++ {
		if x == len(s)-1 {
			baseDN += "DC=" + s[x]
		} else {
			baseDN += "DC=" + s[x] + ","
		}
	}
	*ldapServer += ":389"

	//Query LDAP
	// out, err := exec.Command("nltest", "/DSGETDC:", "/LDAPONLY").Output()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(out))

	//querying
	ldapquery.LDAPlistquery(*ldapServer, *ldapBind, *ldapPassword, baseDN, *command)
}
