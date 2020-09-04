package postgres

import "strings"

// CreateConnectionString is used to obtain the fitting connection String
func (d *Driver) CreateConnectionString(rawDataSource, username, password string) string {
	dataSource := removeUserPassword(rawDataSource)

	var strBuild strings.Builder
	strBuild.WriteString(dataSource)
	strBuild.WriteString(" ")
	strBuild.WriteString("user")
	strBuild.WriteString("=")
	strBuild.WriteString(username)
	strBuild.WriteString(" ")
	strBuild.WriteString("password")
	strBuild.WriteString("=")
	strBuild.WriteString(password)

	return strBuild.String()
}
