package impl

import "github.com/AndreeJait/go-utility/sqlw/postgres"

const (
	queryGetUserByID postgres.QueryString = `
		SELECT 
			user_id,
			email,
			username,
			role
		FROM 
		    users 
		WHERE
		    user_id = $1
	`
)
