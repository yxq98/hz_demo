package mysql

import "database/sql"

func checkErr[T any](value *T, err error) (*T, error) {
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return value, nil
}
