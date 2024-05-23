package impl

import "github.com/AndreeJait/go-utility/sqlw/postgres"

const (
	queryGetUserByID postgres.QueryString = `
		SELECT 
			user_id,
			email,
			username,
			password,
			role
		FROM 
		    users 
		WHERE
		    user_id = $1
	`

	queryGetUserByEmail postgres.QueryString = `
		SELECT 
			user_id,
			email,
			username,
			password,
			role
		FROM 
		    users 
		WHERE
		    email = $1
	`
)
