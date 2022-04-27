package ldapquery

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/ldap.v2"
)

//Get list of all users
func LDAP_ListResolver(t table.Writer, conn *ldap.Conn, baseDN string, command string) {
	result, err := conn.SearchWithPaging(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		LDAPfilter("*", ldap_queries[command]),
		[]string{},
		nil,
	), math.MaxInt32)
	if err != nil {
		fmt.Println(err)
	}
	if len(result.Entries) > 0 {
		for _, mem := range result.Entries {
			t.AppendRows([]table.Row{
				{" - " + mem.GetAttributeValue("sAMAccountName")},
			})

		}
		t.AppendSeparator()
		if command == "users" {
			t.AppendRows([]table.Row{
				{strconv.Itoa(len(result.Entries)) + " Domain Users Found"},
			})
		}
		if command == "computers" {
			t.AppendRows([]table.Row{
				{strconv.Itoa(len(result.Entries)) + " Domain Computers Found"},
			})
		}
		if command == "groups" {
			t.AppendRows([]table.Row{
				{strconv.Itoa(len(result.Entries)) + " Domain Groups Found"},
			})
		}
	}
}

//Get full properties of all users
func LDAP_FullDataResolver(t table.Writer, conn *ldap.Conn, baseDN string, command string) {
	result, err := conn.SearchWithPaging(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		LDAPfilter("*", ldap_queries[command]),
		[]string{},
		nil,
	), math.MaxInt32)
	if err != nil {
		fmt.Println(err)
	}
	if len(result.Entries) > 0 {
		for _, mem := range result.Entries {
			LDAP_Resolver(mem, t)
			t.AppendSeparator()
		}
	}
}

//get all properties of specific object
func LDAP_SpecificFullDataResolver(t table.Writer, conn *ldap.Conn, baseDN string, name string, command string) {
	result, err := conn.SearchWithPaging(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		LDAPfilter("*", strings.ReplaceAll(ldap_queries["specific-"+command], "{SAM}", name)),
		[]string{},
		nil,
	), math.MaxInt32)
	if err != nil {
		fmt.Println(err)
	}
	if len(result.Entries) > 0 {
		for _, mem := range result.Entries {
			LDAP_Resolver(mem, t)
			t.AppendSeparator()
		}
	}
}

//resolve memerOf for Users
func LDAP_UserMembershipResolver(t table.Writer, conn *ldap.Conn, baseDN string, name string, command string) {
	result, err := conn.SearchWithPaging(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		LDAPfilter("*", strings.ReplaceAll(ldap_queries["specific-"+command], "{SAM}", name)),
		[]string{},
		nil,
	), math.MaxInt32)
	if err != nil {
		fmt.Println(err)
	}
	if len(result.Entries) > 0 {
		for _, mem := range result.Entries {
			if len(mem.GetAttributeValue("memberOf")) > 0 {
				t.AppendRows([]table.Row{
					{"Member Of : \n\t" + strings.Join(mem.GetAttributeValues("memberOf"), "\n\t")},
				})
			}
		}
	}
}

//resolve properties of LDAP object
func LDAP_Resolver(entry *ldap.Entry, t table.Writer) {
	if len(entry.DN) > 0 {
		t.AppendRows([]table.Row{
			{"DN : " + entry.DN},
		})
	}
	if len(entry.GetAttributeValue("sAMAccountName")) > 0 {
		t.AppendRows([]table.Row{
			{"SAM Account Name : " + entry.GetAttributeValue("sAMAccountName")},
		})
	}
	if len(entry.GetAttributeValue("sAMAccountType")) > 0 {
		t.AppendRows([]table.Row{
			{"SAM Account Type : " + entry.GetAttributeValue("sAMAccountType")},
		})
	}
	if len(entry.GetAttributeValue("mail")) > 0 {
		t.AppendRows([]table.Row{
			{"Mail ID : " + entry.GetAttributeValue("mail")},
		})
	}
	if len(entry.GetAttributeValue("uid")) > 0 {
		t.AppendRows([]table.Row{
			{"UID : " + entry.GetAttributeValue("uid")},
		})
	}
	if len(entry.GetAttributeValue("ou")) > 0 {
		t.AppendRows([]table.Row{
			{"OU : " + entry.GetAttributeValue("ou")},
		})
	}
	if len(entry.GetAttributeValue("cn")) > 0 {
		t.AppendRows([]table.Row{
			{"CN : " + entry.GetAttributeValue("cn")},
		})
	}
	if len(entry.GetAttributeValue("givenName")) > 0 {
		t.AppendRows([]table.Row{
			{"Given Name : " + entry.GetAttributeValue("givenName")},
		})
	}
	if len(entry.GetAttributeValue("sn")) > 0 {
		t.AppendRows([]table.Row{
			{"SN : " + entry.GetAttributeValue("sn")},
		})
	}
	if len(entry.GetAttributeValue("description")) > 0 {
		t.AppendRows([]table.Row{
			{"Description : " + entry.GetAttributeValue("description")},
		})
	}
	if len(entry.GetAttributeValue("title")) > 0 {
		t.AppendRows([]table.Row{
			{"Title : " + entry.GetAttributeValue("title")},
		})
	}
	if len(entry.GetAttributeValue("c")) > 0 {
		t.AppendRows([]table.Row{
			{"Country : " + entry.GetAttributeValue("c")},
		})
	}
	if len(entry.GetAttributeValue("co")) > 0 {
		t.AppendRows([]table.Row{
			{"Country Code : " + entry.GetAttributeValue("co")},
		})
	}
	if len(entry.GetAttributeValue("l")) > 0 {
		t.AppendRows([]table.Row{
			{"City : " + entry.GetAttributeValue("l")},
		})
	}
	if len(entry.GetAttributeValue("st")) > 0 {
		t.AppendRows([]table.Row{
			{"State : " + entry.GetAttributeValue("st")},
		})
	}
	if len(entry.GetAttributeValue("streetAddress")) > 0 {
		t.AppendRows([]table.Row{
			{"Street Address : " + entry.GetAttributeValue("streetAddress")},
		})
	}
	if len(entry.GetAttributeValue("postalCode")) > 0 {
		t.AppendRows([]table.Row{
			{"Postal Code : " + entry.GetAttributeValue("postalCode")},
		})
	}
	if len(entry.GetAttributeValue("postOfficeBox")) > 0 {
		t.AppendRows([]table.Row{
			{"Post Office Box : " + entry.GetAttributeValue("postOfficeBox")},
		})
	}
	if len(entry.GetAttributeValue("company")) > 0 {
		t.AppendRows([]table.Row{
			{"Company : " + entry.GetAttributeValue("company")},
		})
	}
	if len(entry.GetAttributeValue("instanceType")) > 0 {
		t.AppendRows([]table.Row{
			{"Instance Type : " + entry.GetAttributeValue("instanceType")},
		})
	}
	if len(entry.GetAttributeValue("objectClass")) > 0 {
		t.AppendRows([]table.Row{
			{"Object Class : " + strings.Join(entry.GetAttributeValues("objectClass"), ", ")},
		})
	}
	if len(entry.GetAttributeValue("objectCategory")) > 0 {
		t.AppendRows([]table.Row{
			{"Object Category : " + entry.GetAttributeValue("objectCategory")},
		})
	}
	if len(entry.GetAttributeValue("memberOf")) > 0 {
		t.AppendRows([]table.Row{
			{"Member Of : \n\t" + strings.Join(entry.GetAttributeValues("memberOf"), "\n\t")},
		})
	}
	if len(entry.GetAttributeValue("dNSHostName")) > 0 {
		t.AppendRows([]table.Row{
			{"DNS Host Name : " + entry.GetAttributeValue("dNSHostName")},
		})
	}
	if len(entry.GetAttributeValue("servicePrincipalName")) > 0 {
		t.AppendRows([]table.Row{
			{"Service Principal Name (SPN) : \n\t" + strings.Join(entry.GetAttributeValues("servicePrincipalName"), "\n\t")},
		})
	}
	if len(entry.GetAttributeValue("operatingSystem")) > 0 {
		t.AppendRows([]table.Row{
			{"Operating System : " + entry.GetAttributeValue("operatingSystem")},
		})
	}
	if len(entry.GetAttributeValue("operatingSystemVersion")) > 0 {
		t.AppendRows([]table.Row{
			{"Operating System Version : " + entry.GetAttributeValue("operatingSystemVersion")},
		})
	}
	if len(entry.GetAttributeValue("userAccountControl")) > 0 {
		uacFlag, err := strconv.ParseInt(entry.GetAttributeValue("userAccountControl"), 0, 64)
		if err != nil {
			//do nothing
		} else {
			t.AppendRows([]table.Row{
				{"UAC Flag : " + strings.Join(ParseLDAPFlag(uacFlag, uac_flags)[:], ",")},
			})
		}
	}
	if len(entry.GetAttributeValue("userAccountControl")) > 0 {
		t.AppendRows([]table.Row{
			{"User Account Control : " + entry.GetAttributeValue("userAccountControl")},
		})
	}
	if len(entry.GetAttributeValue("trustType")) > 0 {
		trustType, err := strconv.ParseInt(entry.GetAttributeValue("trustType"), 0, 64)
		if err != nil {
			//do nothing
		} else {
			t.AppendRows([]table.Row{
				{"Trust Type : " + strings.Join(ParseLDAPFlag(trustType, trust_type)[:], ",")},
			})
		}
	}
	if len(entry.GetAttributeValue("trustDirection")) > 0 {
		trustDirection, err := strconv.ParseInt(entry.GetAttributeValue("trustDirection"), 0, 64)
		if err != nil {
			//do nothing
		} else {
			t.AppendRows([]table.Row{
				{"Trust Direction : " + strings.Join(ParseLDAPFlag(trustDirection, trust_directions)[:], ",")},
			})
		}
	}
	if len(entry.GetAttributeValue("trustAttributes")) > 0 {
		trustAttributes, err := strconv.ParseInt(entry.GetAttributeValue("trustAttributes"), 0, 64)
		if err != nil {
			//do nothing
		} else {
			t.AppendRows([]table.Row{
				{"Trust Attribute : " + strings.Join(ParseLDAPFlag(trustAttributes, trust_flags)[:], ",")},
			})
		}
	}
	if len(entry.GetAttributeValue("objectGUID")) > 0 {
		guid, err := ParseGUID(entry.GetRawAttributeValue("objectGUID"))
		if err != nil {
			//do nothing
		} else {
			t.AppendRows([]table.Row{
				{"Object GUID : " + guid},
			})
		}
	}
	if len(entry.GetAttributeValue("objectSid")) > 0 {
		sidbyte := []byte(entry.GetAttributeValue("objectSid"))
		sid, err := ParseSID(sidbyte)
		if err != nil {
			//do nothing
		} else {
			t.AppendRows([]table.Row{
				{"Object SID : " + sid},
			})
		}
	}
	LDAP_TimeResolver(entry, t)
}

//resolving time object of LDAP attributes
func LDAP_TimeResolver(entry *ldap.Entry, t table.Writer) {
	if len(entry.GetAttributeValue("badPwdCount")) > 0 {
		t.AppendRows([]table.Row{
			{"Bad Password Count : " + entry.GetAttributeValue("badPwdCount")},
		})
	}
	if len(entry.GetAttributeValue("badPasswordTime")) > 0 {
		t.AppendRows([]table.Row{
			{"Bad Password Time : " + ParseTimeStamp(entry.GetAttributeValue("badPasswordTime")).String()},
		})
	}
	if len(entry.GetAttributeValue("whenCreated")) > 0 {
		t.AppendRows([]table.Row{
			{"When Created : " + entry.GetAttributeValue("whenCreated")},
		})
	}
	if len(entry.GetAttributeValue("whenChanged")) > 0 {
		t.AppendRows([]table.Row{
			{"When Changed : " + entry.GetAttributeValue("whenChanged")},
		})
	}
	if len(entry.GetAttributeValue("pwdLastSet")) > 0 {
		t.AppendRows([]table.Row{
			{"Password Last Set : " + ParseTimeStamp(entry.GetAttributeValue("pwdLastSet")).String()},
		})
	}
	if len(entry.GetAttributeValue("lastLogon")) > 0 {
		t.AppendRows([]table.Row{
			{"Last Log On : " + ParseTimeStamp(entry.GetAttributeValue("lastLogon")).String()},
		})
	}
	if len(entry.GetAttributeValue("lastLogoff")) > 0 {
		t.AppendRows([]table.Row{
			{"Last Log Off : " + ParseTimeStamp(entry.GetAttributeValue("lastLogoff")).String()},
		})
	}
	if len(entry.GetAttributeValue("accountExpires")) > 0 {
		t.AppendRows([]table.Row{
			{"Account Expires On : " + ParseTimeStamp(entry.GetAttributeValue("accountExpires")).String()},
		})
	}
	if len(entry.GetAttributeValue("uSNChanged")) > 0 {
		t.AppendRows([]table.Row{
			{"uSNChanged : " + ParseTimeStamp(entry.GetAttributeValue("uSNChanged")).String()},
		})
	}
	if len(entry.GetAttributeValue("uSNCreated")) > 0 {
		t.AppendRows([]table.Row{
			{"uSNCreated : " + ParseTimeStamp(entry.GetAttributeValue("uSNCreated")).String()},
		})
	}
}

//get group DN of specific group
func LDAP_GroupDN(conn *ldap.Conn, baseDN string, name string) string {
	result, err := conn.SearchWithPaging(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		LDAPfilter("*", strings.ReplaceAll(ldap_queries["specific-groups"], "{SAM}", name)),
		[]string{},
		nil,
	), math.MaxInt32)
	if err != nil {
		fmt.Println(err)
	}
	if len(result.Entries) > 0 {
		for _, mem := range result.Entries {
			return mem.DN
		}
	}
	return ""
}

//get all users from a specific group
func LDAP_GroupResolver(t table.Writer, conn *ldap.Conn, baseDN string, name string) {
	result, err := conn.SearchWithPaging(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		math.MaxInt32,
		0,
		false,
		LDAPfilter("*", strings.ReplaceAll(ldap_queries["group-members"], "{DN}", LDAP_GroupDN(conn, baseDN, name))),
		[]string{},
		nil,
	), math.MaxInt32)
	if err != nil {
		fmt.Println(err)
	}
	if len(result.Entries) > 0 {
		t.AppendRows([]table.Row{
			{"Domain users in group : "},
		})
		for _, mem := range result.Entries {
			t.AppendRows([]table.Row{
				{" - " + mem.GetAttributeValue("sAMAccountName")},
			})

		}
	}
}
