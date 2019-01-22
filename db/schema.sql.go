package db

import (
	"fmt"
)

// GetDropSchema returns the sql requried to remove all tables from the database
func GetDropSchema() string {
	return `
    DROP TABLE IF EXISTS "main";
    `
}

// GetSchema returns the database schema for the stickerbot database
func GetSchema() string {
	sql := `
        CREATE TABLE IF NOT EXISTS "public"."main" (
            "id" VARCHAR(255) NOT NULL,
            "bot_state" VARCHAR(255),
            "templates" VARCHAR(255),
            "created_at" timestamptz NOT NULL DEFAULT now(),
            "updated_at" timestamptz NOT NULL DEFAULT now(),
            CONSTRAINT "main_pkey" PRIMARY KEY ("id")
        ) WITH (oids = false);
    `

	tableNames := []string{
		"main",
	}

	for _, table := range tableNames {
		sql += fmt.Sprintf("\nDROP TRIGGER IF EXISTS set_%s_updated_at ON %s;\n", table, table)
		sql += fmt.Sprintf("\nCREATE TRIGGER set_%s_updated_at\nBEFORE UPDATE ON %s\nFOR EACH ROW EXECUTE PROCEDURE update_modified_column();\n", table, table)
	}

	return sql
}
