package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Column struct {
	TableName     string
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnDefault *string
}

// func main() {
// 	conn, err := pgx.Connect(context.Background(), "postgres://humanizati_27ab570d10874ad8b1a44a0425092a0f_p_postgres_svcs:69X68Pl1lL@95.217.155.57:30034/humanizati_27ab570d10874ad8b1a44a0425092a0f_p_postgres_svcs?sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close(context.Background())

// 	rows, err := conn.Query(context.Background(), `
//         SELECT table_name, column_name, data_type, is_nullable, column_default
//         FROM information_schema.columns
//         WHERE table_schema = 'public'
//     `)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// Store table columns
// 	tables := make(map[string][]Column)

// 	for rows.Next() {
// 		var col Column
// 		if err := rows.Scan(&col.TableName, &col.ColumnName, &col.DataType, &col.IsNullable, &col.ColumnDefault); err != nil {
// 			log.Fatal(err)
// 		}
// 		tables[col.TableName] = append(tables[col.TableName], col)
// 	}

// 	// Generate DBML
// 	var dbml strings.Builder
// 	for table, columns := range tables {
// 		dbml.WriteString(fmt.Sprintf("Table %s {\n", table))
// 		for _, col := range columns {
// 			dbml.WriteString(fmt.Sprintf("  %s %s", col.ColumnName, col.DataType))
// 			if col.IsNullable == "NO" {
// 				dbml.WriteString(" [not null]")
// 			}
// 			if col.ColumnDefault != nil {
// 				dbml.WriteString(fmt.Sprintf(" [default: %s]", *col.ColumnDefault))
// 			}
// 			dbml.WriteString("\n")
// 		}
// 		dbml.WriteString("}\n\n")
// 	}

// 	fmt.Println(dbml.String()) // Copy this to dbdiagram.io
// }

type Table struct {
	ID   string
	Slug string
}

type Field struct {
	TableID    string
	Slug       string
	Type       string
	IsNullable bool
	IsUnique   bool
}

type Relation struct {
	TableFrom string
	FieldFrom string
	TableTo   string
	FieldTo   string
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://moshina_0eb9943fc4024f53bd6c6e9bf777fa2e_p_postgres_svcs:nzHy1eGHgw@142.93.164.37:30032/moshina_0eb9943fc4024f53bd6c6e9bf777fa2e_p_postgres_svcs?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Fetch tables
	tableRows, err := conn.Query(context.Background(), `SELECT id, slug FROM "table" where is_system = false`)
	if err != nil {
		log.Fatal("poqeiweoiwioeqiwe", err)
		return
	}
	defer tableRows.Close()

	tables := []Table{}
	for tableRows.Next() {
		var t Table
		if err := tableRows.Scan(&t.ID, &t.Slug); err != nil {
			log.Fatal(err)
		}
		tables = append(tables, t)
	}

	// Fetch fields (columns)
	fieldRows, err := conn.Query(context.Background(), `
		SELECT table_id, slug, "type", is_system, "unique"
		FROM "field"
	`)
	if err != nil {
		log.Fatal("sdfasdfasdf", err)
		return
	}
	defer fieldRows.Close()

	fields := []Field{}
	for fieldRows.Next() {
		var f Field
		if err := fieldRows.Scan(&f.TableID, &f.Slug, &f.Type, &f.IsNullable, &f.IsUnique); err != nil {
			log.Fatal(err)
		}
		fields = append(fields, f)
	}

	// Fetch relations
	relationRows, err := conn.Query(context.Background(), `
		SELECT table_from, field_from, table_to, field_to 
		FROM relation
	`)
	if err != nil {
		log.Fatal("relation", err)
		return
	}
	defer relationRows.Close()

	relations := []Relation{}
	for relationRows.Next() {
		var r Relation
		if err := relationRows.Scan(&r.TableFrom, &r.FieldFrom, &r.TableTo, &r.FieldTo); err != nil {
			log.Fatal(err)
		}
		relations = append(relations, r)
	}

	// Generate DBML
	var dbml strings.Builder

	// Add tables
	for _, t := range tables {
		dbml.WriteString(fmt.Sprintf("Table %s {\n", t.Slug))
		for _, f := range fields {
			if f.TableID == t.ID {
				dbml.WriteString(fmt.Sprintf("  %s %s", f.Slug, FIELD_TYPES[f.Type]))
				if f.IsUnique {
					dbml.WriteString(" [unique]")
				}
				dbml.WriteString("\n")
			}
		}
		dbml.WriteString("}\n\n")
	}

	// Add relations
	for _, r := range relations {
		dbml.WriteString(fmt.Sprintf("Ref: %s.%s > %s.%s\n", r.TableFrom, r.FieldFrom, r.TableTo, r.FieldTo))
	}

	fmt.Println(dbml.String()) // Copy this into dbdiagram.io
}

var (
	FIELD_TYPES = map[string]string{
		"SINGLE_LINE":                 "VARCHAR",
		"MULTI_LINE":                  "VARCHAR",
		"PICK_LIST":                   "VARCHAR",
		"LOOKUP":                      "VARCHAR",
		"EMAIL":                       "VARCHAR",
		"PHOTO":                       "VARCHAR",
		"PHONE":                       "VARCHAR",
		"UUID":                        "VARCHAR",
		"INCREMENT_ID":                "VARCHAR",
		"RANDOM_NUMBERS":              "VARCHAR",
		"PASSWORD":                    "VARCHAR",
		"FILE":                        "VARCHAR",
		"CODABAR":                     "VARCHAR",
		"INTERNATION_PHONE":           "VARCHAR",
		"FORMULA_FRONTEND":            "VARCHAR",
		"TIME":                        "VARCHAR",
		"DATE_TIME":                   "TIMESTAMP",
		"DATE_TIME_WITHOUT_TIME_ZONE": "TIMESTAMP",
		"DATE":                        "DATE",
		"NUMBER":                      "FLOAT",
		"FLOAT":                       "FLOAT",
		"FLOAT_NOLIMIT":               "FLOAT",
		"FORMULA":                     "FLOAT",
		"CHECKBOX":                    "BOOL",
		"SWITCH":                      "BOOL",
		"MULTISELECT":                 "TEXT[]",
		"LOOKUPS":                     "TEXT[]",
		"DYNAMIC":                     "TEXT[]",
		"LANGUAGE_TYPE":               "TEXT[]",
		"MULTI_IMAGE":                 "TEXT[]",
		"MULTI_FILE":                  "TEXT[]",
		"MONEY":                       "TEXT[]",
		"INCREMENT_NUMBER":            "SERIAL",
		"COLOR":                       "VARCHAR",
		"JSON":                        "VARCHAR",
		"MAP":                         "VARCHAR",
	}
)
