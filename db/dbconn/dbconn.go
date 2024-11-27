package dbconn

import (
	"database/sql"
	"fmt"
)

func InitDB(db *sql.DB) error {
	// put db in WAL mode
	_, err := db.Exec("PRAGMA journal_mode=WAL")
	if err != nil {
		return fmt.Errorf("failed to enable WAL mode on sqlite database: %w", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return fmt.Errorf("failed to enable foreign key on sqlite database: %w", err)
	}

	_, err = db.Exec("PRAGMA busy_timeout = 5000")
	if err != nil {
		return fmt.Errorf("failed to set busy timeout on sqlite database: %w", err)
	}

	_, err = db.Exec("PRAGMA synchronous = NORMAL")
	if err != nil {
		return fmt.Errorf("failed to set syncronous mode on sqlite database: %w", err)
	}

	_, err = db.Exec("PRAGMA cache_size = 1000000000")
	if err != nil {
		return fmt.Errorf("failed to se cache size on sqlite database: %w", err)
	}

	_, err = db.Exec("PRAGMA temp_store = memory;")
	if err != nil {
		return fmt.Errorf("failed to set temp store on sqlite database: %w", err)
	}

	return nil
}

func DeactivateWAL(db *sql.DB) error {
	_, err := db.Exec("PRAGMA journal_mode=DELETE;")
	if err != nil {
		return fmt.Errorf("failed to deactivate WAL on sqlite database: %w", err)
	}
	return nil
}
