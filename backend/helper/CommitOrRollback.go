package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			panic(err)
		}
	} else {
		err := tx.Commit()
		if err != nil {
			panic(err)
		}
	}

}
