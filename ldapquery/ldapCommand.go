package ldapquery

var ldap_commands = map[string]string{
	"users":          "Users",
	"user-logs":      "User Properties",
	"groups":         "Groups",
	"computers":      "Computers",
	"dc":             "Domain Controllers",
	"gpo":            "Group Policy Objects",
	"spn":            "Service Principal Names",
	"never-loggedon": "Users Never LoggedOn",
	"ms-sql":         "MS-SQL Servers",
	"admin-priv":     "Admin Priv",
	"domain-trust":   "Trusted Domain",
	"ou":             "Organizational Units",
	"asreproast":     "AS-REP Roastable Accounts",
	"unconstrained":  "Unconstrained Delegation",
}
