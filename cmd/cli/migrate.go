package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()
	// run migration command

	switch arg2 {
	case "up":
		err := sco.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := sco.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := sco.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		err := sco.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		err = sco.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		showHelp()
	}
	return nil
}
