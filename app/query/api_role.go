package query

import (
	"github.com/Ketarin/pg-casbin/app/infra"
	jpgx "github.com/jackc/pgx"
)

type ApiRoleInterface interface {
	GetRbacGlobalDefinitions() ([]string, error)
	GetRbacUserDefinitions(user string) ([]string, error)
}

type ApiRole struct {
	infra.ConnectionInterface
}

func NewApiRole(conn infra.ConnectionInterface) ApiRoleInterface {
	return &ApiRole{ConnectionInterface: conn}
}

//GetRbacGlobalDefinitions returns all roles with their children
func (a *ApiRole) GetRbacGlobalDefinitions() ([]string, error) {

	sql := `SELECT
              CONCAT_WS(', ', 'p', arl.name, art.url, art.method)
            FROM api_route art
              INNER JOIN api_route_role arr ON art.api_route_id = arr.api_route_id
              INNER JOIN api_role arl ON arr.api_role_id = arl.api_role_id
            UNION
            SELECT
              CONCAT_WS(', ', 'g', arl2.name, arl1.name)
            FROM api_role_role arl
              INNER JOIN api_role arl1 ON arl1.api_role_id = arl.parent_api_role_id
              INNER JOIN api_role arl2 ON arl2.api_role_id = arl.child_api_role_id`

	rows, err := a.Query(sql)
	if err != nil {
		return nil, err
	}

	return a.generateRbacDefinitions(rows)
}

//GetRbacUserDefinitions returns all roles with their children
func (a *ApiRole) GetRbacUserDefinitions(user string) ([]string, error) {

	sql := `SELECT 
              CONCAT_WS(', ', 'g', u.name, arl.name)
            FROM api_role_user aru
              INNER JOIN api_role arl ON arl.api_role_id = aru.api_role_id
              INNER JOIN users u ON u.user_id = aru.user_id
            WHERE u.name = $1`

	rows, err := a.Query(sql, user)
	if err != nil {
		return nil, err
	}

	return a.generateRbacDefinitions(rows)
}

func (a *ApiRole) generateRbacDefinitions(rows *jpgx.Rows) (definitions []string, err error) {

	for rows.Next() {
		var definition string

		errScan := rows.Scan(&definition)
		if errScan != nil {
			return definitions, err
		}

		definitions = append(definitions, definition)
	}

	return definitions, nil
}
