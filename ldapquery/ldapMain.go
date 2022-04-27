package ldapquery

import (
	"fmt"
	"math"
	"strings"

	"gopkg.in/ldap.v2"
)

func LDAPconnect(
	ldapServer string,
	ldapBind string,
	ldapPassword string,
) (*ldap.Conn, error) {
	conn, err := ldap.Dial("tcp", ldapServer)
	if err != nil {
		return nil, fmt.Errorf("[+] Failed to connect with LDAP server. %s", err)
	}
	if err := conn.Bind(ldapBind, ldapPassword); err != nil {
		return nil, fmt.Errorf("[+] Failed to connect. %s", err)
	}
	return conn, nil
}

func LDAPfilter(
	needle string,
	filterDN string,
) string {
	res := strings.Replace(
		filterDN,
		"{username}",
		needle,
		-1,
	)
	return res
}

func LDAPauth(
	conn *ldap.Conn,
	baseDN string,
	filterDN string,
	loginUsername string,
	loginPassword string,
) error {
	result, err := conn.SearchWithPaging(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		LDAPfilter(loginUsername, filterDN),
		[]string{"dn"},
		nil,
	), math.MaxInt32)

	fmt.Print(result.Entries)
	fmt.Print("Entries : ")
	fmt.Print(len(result.Entries))

	if err != nil {
		return fmt.Errorf("[+] Failed to find user. %s", err)
	}

	if len(result.Entries) < 1 {
		return fmt.Errorf("[+] User does not exist")
	}

	if len(result.Entries) > 1 {
		return fmt.Errorf("[+] Too many entries returned")
	}

	if err := conn.Bind(result.Entries[0].DN, loginPassword); err != nil {
		fmt.Printf("[+] Failed to auth. %s", err)
	} else {
		fmt.Printf("[+] Authenticated successfuly!")
	}

	return nil
}

func LDAPlistquery(
	ldapServer string,
	ldapBind string,
	ldapPassword string,
	baseDN string,
	command string,
	filter string,
	name string,
) {
	conn, err := LDAPconnect(ldapServer, ldapBind, ldapPassword)
	if err != nil {
		fmt.Printf("Failed to connect the LDAP server. %s", err)
		return
	}
	if err := LDAP_QueryData(conn, baseDN, command, filter, name); err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer conn.Close()
}
