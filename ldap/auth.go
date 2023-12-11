package ldap

import (
	models "go-chi-api/model"

	"github.com/go-ldap/ldap/v3"
)

func ConnectToLDAP(serverURL, bindDN, bindPassword string) (*ldap.Conn, error) {
	conn, err := ldap.Dial("tcp", serverURL)
	if err != nil {
		return nil, err
	}

	err = conn.Bind(bindDN, bindPassword)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func AuthenticateUser(conn *ldap.Conn, username, password string) (*models.User, error) {
	// Implement logic to authenticate user against LDAP server
	return nil, nil
}
