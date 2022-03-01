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
}

var ldap_variables = map[string]map[string]string{
	"users":          user_field,
	"user-logs":      user_logs_field,
	"groups":         group_field,
	"computers":      computer_field,
	"dc":             dc_field,
	"gpo":            gpo_field,
	"spn":            spn_field,
	"never-loggedon": never_loggedon_field,
	"ms-sql":         mssql_field,
	"admin-priv":     admin_priv_field,
	"domain-trust":   domain_trust_field,
	"ou":             ou_field,
}
