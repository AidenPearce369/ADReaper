package ldapquery

//Need to add some additional attributes

var user_field = map[string]string{
	"SAM Account Name : ":    "sAMAccountName",
	"Common Name (CN) : ":    "cn",
	"User Principal Name : ": "userPrincipalName",
	"Object Category : ":     "objectCategory",
	//UAC Flag
	//ObjectSID
	//ObjectGUID

}

var user_logs_field = map[string]string{
	"SAM Account Name : ": "sAMAccountName",
	// "Bad Password Count : ": "badPwdCount",
	// "Bad Password Time : ":  "badPasswordTime",
	// "Password Last Set : ":  "pwdLastSet",
	// "Last LogOn : ":         "lastLogon",
}

var computer_field = map[string]string{
	"SAM Account Name : ": "sAMAccountName",
	"Common Name (CN) : ": "cn",
	//whenCreated
	//whenChanged
}

var group_field = map[string]string{
	"SAM Account Name : ": "sAMAccountName",
}

var dc_field = map[string]string{
	"SAM Account Name : ": "sAMAccountName",
	"Common Name (CN) : ": "cn",
	"Logon Count : ":      "logonCount",
}

var gpo_field = map[string]string{
	"GPO Display Name : ":  "displayName",
	"Common Name (CN) : ":  "cn",
	"GPO File Sys Path : ": "gPCFileSysPath",
}

var spn_field = map[string]string{
	"SAM Account Name : ":             "sAMAccountName",
	"Common Name (CN) : ":             "cn",
	"Service Principal Name (SPN) : ": "servicePrincipalName",
	"User Principal Name : ":          "userPrincipalName",
}
var never_loggedon_field = map[string]string{
	"SAM Account Name : ":    "sAMAccountName",
	"User Principal Name : ": "userPrincipalName",
}

var mssql_field = map[string]string{
	"SAM Account Name : ":             "sAMAccountName",
	"Common Name (CN) : ":             "cn",
	"Service Principal Name (SPN) : ": "servicePrincipalName",
	"User Principal Name : ":          "userPrincipalName",
}

var admin_priv_field = map[string]string{
	"SAM Account Name : ": "sAMAccountName",
	"Common Name (CN) : ": "cn",
}

var domain_trust_field = map[string]string{
	"Domain : ":           "sourcedomain",
	"Common Name (CN) : ": "cn",
}

var ou_field = map[string]string{
	"OU Name : ": "ou",
}
