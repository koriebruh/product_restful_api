package helper

import "database/sql"

// <-- digunakan untuk mengecek jika ada error maka transaksi akan dirollback

func CommitOrRollback(tx *sql.Tx, err error) {
	func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			IfErrNotNul(errRollback)

		} else {
			errCommit := tx.Commit()
			IfErrNotNul(errCommit)
		}
	}()
}
