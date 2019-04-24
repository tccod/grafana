package login

import (
	"github.com/grafana/grafana/pkg/models"
	LDAP "github.com/grafana/grafana/pkg/services/ldap"
)

var newLDAP = LDAP.New
var readLDAPConfig = LDAP.ReadConfig

var loginUsingLdap = func(query *models.LoginUserQuery) (bool, error) {
	enabled, config := readLDAPConfig()

	if !enabled {
		return false, nil
	}

	if len(config.Servers) == 0 {
		return true, ErrNoLDAPServers
	}

	for _, server := range config.Servers {
		auth := newLDAP(server)
		err := auth.Login(query)
		if err == nil || err != LDAP.ErrInvalidCredentials {
			return true, err
		}
	}

	return true, LDAP.ErrInvalidCredentials
}
