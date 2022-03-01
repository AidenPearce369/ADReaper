package ldapquery

var ldap_queries = map[string]string{
	"users":                   "(objectClass=user)",
	"user-logs":               "(objectClass=user)",
	"groups":                  "(objectClass=group)",
	"computers":               "(objectClass=Computer)",
	"dc":                      "(&(objectCategory=Computer)(userAccountControl:1.2.840.113556.1.4.803:=8192))",
	"gpo":                     "(objectClass=groupPolicyContainer)",
	"spn":                     "(&(&(servicePrincipalName=*)(UserAccountControl:1.2.840.113556.1.4.803:=512))(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))",
	"unconstrained-users":     "(&(&(objectCategory=person)(objectClass=user))(userAccountControl:1.2.840.113556.1.4.803:=524288))",
	"unconstrained-computers": "(&(objectCategory=computer)(objectClass=computer)(userAccountControl:1.2.840.113556.1.4.803:=524288))",
	"ms-sql":                  "(&(objectCategory=computer)(servicePrincipalName=MSSQLSvc*))",
	"never-loggedon":          "(&(objectCategory=person)(objectClass=user)(|(lastLogonTimestamp=0)(!(lastLogonTimestamp=*))))",
	"group-members":           "(&(objectCategory=user)(memberOf={DN}))",
	"admin-priv":              "(adminCount=1)",
	"domain-trust":            "(objectClass=trustedDomain)",
	"ou":                      "(&(objectCategory=organizationalUnit)(ou=*))",
}

//Planned LDAP queries

// "locked-accounts":         "((objectCategory=person)(objectClass=user)(useraccountcontrol:1.2.840.113556.1.4.803:=16))",
// "password-never-expire":   "((objectCategory=person)(objectClass=user)(userAccountControl:1.2.840.113556.1.4.803:=65536))",
// "password-expired":        "((objectCategory=person)(objectClass=user)(pwdLastSet=0)(!useraccountcontrol:1.2.840.113556.1.4.803:=2))",
// 	nested group mem = "(&(objectClass=user)(memberof:1.2.840.113556.1.4.1941:={}))".format(groupDN)
// 	usersWithPwdLastSetThreeMonths    = "(&(sAMAccountType=805306368)(pwdLastSet<={pwdLastSetUnix}))"
// 	adUsersWithEmailID                = "(objectcategory=person)(mail=*)"
// 	adUsersWithNoEmailID              = "(objectcategory=person)(!mail=*)"
// 	usersHiddenFromExchgAddrBook      = "(&(sAMAccountType=805306368)(msExchHideFromAddressLists=TRUE))"
// 	adUsersCreatedCurrentYear         = "(&(&(&(objectClass=User)(whenCreated>={createdZTime}))))"
// 	adComputersRunningSpecificOS      = "(&(objectCategory=computer)(operatingSystem={Windows 10*}))"
// 	adComputersRunningSpecificOSBuild = "(&(&(objectCategory=computer)(operatingSystem=Windows 10*)(operatingSystemVersion={*buildNo*})))"
// 	exchgDistributionGroupsInAD       = "(&(objectCategory=group)(!groupType:1.2.840.113556.1.4.803:=2147483648))"
