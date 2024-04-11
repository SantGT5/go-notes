package database

func CloseDB() error {
	if db == nil {
		return nil
	}

	dbSQL, err := db.DB()

	if err != nil {
		return err
	}

	err = dbSQL.Close()

	if err != nil {
		return err
	}

	return nil
}
