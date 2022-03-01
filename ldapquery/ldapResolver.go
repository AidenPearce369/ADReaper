package ldapquery

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/ldap.v2"
)

func LDAP_CommandChecker(command string) bool {
	switch command {
	case
		"user-logs",
		"user-log":
		return false
	}
	return true
}

func LDAP_Resolver(entry *ldap.Entry, t table.Writer) {
	if len(entry.GetAttributeValue("description")) > 0 {
		t.AppendRows([]table.Row{
			{"Description : " + entry.GetAttributeValue("description")},
		})
	}
	if len(entry.GetAttributeValue("dNSHostName")) > 0 {
		t.AppendRows([]table.Row{
			{"DNS Host Name : " + entry.GetAttributeValue("dNSHostName")},
		})
	}
	if len(entry.GetAttributeValue("servicePrincipalName")) > 0 {
		t.AppendRows([]table.Row{
			{"Service Principal Name (SPN) : " + entry.GetAttributeValue("servicePrincipalName")},
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
				{"UAC Flag : " + strings.Join(ParseLDAPFlag(trustAttributes, trust_flags)[:], ",")},
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
}

func LDAP_TimeResolver(entry *ldap.Entry, t table.Writer) {
	if len(entry.GetAttributeValue("badPwdCount")) > 0 {
		t.AppendRows([]table.Row{
			{"Bad Password Count : " + entry.GetAttributeValue("badPwdCount")},
			{"Bad Password Time : " + ParseTimeStamp(entry.GetAttributeValue("badPasswordTime")).String()},
			{"Password Last Set : " + ParseTimeStamp(entry.GetAttributeValue("pwdLastSet")).String()},
			{"Last Log On : " + ParseTimeStamp(entry.GetAttributeValue("lastLogon")).String()},
		})

	}
}

func LDAP_GroupResolver(t table.Writer, conn *ldap.Conn, baseDN string, DN string) {
	result, err := conn.Search(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		LDAPfilter("*", strings.ReplaceAll(ldap_queries["group-members"], "{DN}", DN)),
		[]string{},
		nil,
	))
	if err != nil {
		fmt.Println(err)
	}
	if len(result.Entries) > 0 {
		t.AppendRows([]table.Row{
			{"Group Members : "},
		})
		for _, mem := range result.Entries {
			t.AppendRows([]table.Row{
				{" - " + mem.GetAttributeValue("sAMAccountName")},
			})

		}
	}
}
