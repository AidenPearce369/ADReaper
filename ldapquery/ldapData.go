package ldapquery

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/ldap.v2"
)

// Lists all Users in current domain
func LDAP_QueryData(
	conn *ldap.Conn,
	baseDN string,
	command string,
) error {
	result, err := conn.Search(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		LDAPfilter("*", ldap_queries[command]),
		[]string{},
		nil,
	))
	if err != nil {
		return fmt.Errorf("[+] Failed to search query. %s", err)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{ldap_commands[command]})
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:     ldap_commands[command],
			WidthMin: 20,
			WidthMax: 100,
		},
	})
	for _, entry := range result.Entries {
		//DEBUG LINE
		// entry.Print()
		// fmt.Println(strings.Repeat("*", 60))
		t.AppendRows([]table.Row{
			{"Distinguished Name (DN) : " + entry.DN},
		})
		for x, y := range ldap_variables[command] {
			t.AppendRows([]table.Row{
				{x + entry.GetAttributeValue(y)},
			})
		}
		if LDAP_CommandChecker(command) {
			LDAP_Resolver(entry, t)
		}
		if command == "user-logs" {
			LDAP_TimeResolver(entry, t)
		}
		if command == "groups" {
			LDAP_GroupResolver(t, conn, baseDN, entry.DN)
		}
		t.AppendSeparator()
	}
	t.Render()
	return nil
}
