package ldapquery

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/ldap.v2"
)

func LDAP_QueryData(
	conn *ldap.Conn,
	baseDN string,
	command string,
	filter string,
	name string,
) error {
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
	if command == "users" && filter == "list" && name == "" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	if command == "users" && filter == "full-data" && name == "" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "users" && name != "" && filter != "membership" {
		LDAP_SpecificFullDataResolver(t, conn, baseDN, name, command)
	}
	if command == "users" && name != "" && filter == "membership" {
		LDAP_UserMembershipResolver(t, conn, baseDN, name, command)
	}
	if command == "computers" && filter == "list" && name == "" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	if command == "computers" && filter == "full-data" && name == "" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "computers" && name != "" {
		LDAP_SpecificFullDataResolver(t, conn, baseDN, name, command)
	}
	if command == "groups" && filter == "list" && name == "" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	if command == "groups" && filter == "full-data" && name == "" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "groups" && name != "" && filter != "membership" {
		LDAP_SpecificFullDataResolver(t, conn, baseDN, name, command)
	}
	if command == "groups" && filter == "membership" && name != "" {
		LDAP_GroupResolver(t, conn, baseDN, name)
	}
	if command == "dc" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "domain-trust" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "spn" && filter == "list" && name == "" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	if command == "spn" && filter == "full-data" && name == "" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "spn" && name != "" {
		LDAP_SpecificFullDataResolver(t, conn, baseDN, name, command)
	}
	if command == "never-loggedon" && filter == "list" && name == "" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	if command == "gpo" && name == "" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "ou" && name == "" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "ms-sql" && filter == "list" && name == "" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	if command == "ms-sql" && filter == "full-data" && name == "" {
		LDAP_FullDataResolver(t, conn, baseDN, command)
	}
	if command == "ms-sql" && name != "" {
		LDAP_SpecificFullDataResolver(t, conn, baseDN, name, command)
	}
	if command == "asreproast" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	if command == "unconstrained" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	if command == "admin-priv" {
		LDAP_ListResolver(t, conn, baseDN, command)
	}
	t.Render()
	return nil
}
