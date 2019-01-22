package luxdb

// Migrate will drop all existing tables and create all tables
// from schema.sql. THIS WILL ERASE ALL DATA!!
func (d *DBSession) Migrate() error {
	_, err := d.Exec(GetSchema())
	if err != nil {
		return err
	}

	return nil
}

// DropTables will drop all tables from the database
func (d *DBSession) DropTables() error {
	_, err := d.Exec(GetDropSchema())
	if err != nil {
		return err
	}

	return nil
}
