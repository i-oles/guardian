package sqlite

import (
	model "cmd/main.go/internal/model"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type BulbsRepo struct {
	db        *sql.DB
	tableName string
}

func NewBulbsRepo(db *sql.DB, collName string) *BulbsRepo {
	return &BulbsRepo{
		db:        db,
		tableName: collName,
	}
}

func (r *BulbsRepo) Get(id string) (model.Bulb, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE bulb_id = ?;", r.tableName)

	rows, err := r.db.Query(query, id)
	if err != nil {
		return model.Bulb{}, fmt.Errorf("could not get bulb from bulbs table: %v", err)
	}

	defer rows.Close()

	var bulb model.Bulb

	for rows.Next() {
		err := rows.Scan(
			&bulb.ID, &bulb.Name, &bulb.BulbType, &bulb.Luminance, &bulb.Red, &bulb.Green, &bulb.Blue,
		)
		if err != nil {
			return model.Bulb{}, fmt.Errorf("could not scan bulb from bulbs table: %v", err)
		}
	}

	return bulb, nil
}

func (r *BulbsRepo) GetOfflineBulbs(onlineIDs []string) ([]model.Bulb, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE bulb_id NOT IN (%s);",
		r.tableName, strings.Join(quoteSlice(onlineIDs), ","))

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not get bulbs from bulbs table: %v", err)
	}

	defer rows.Close()

	var bulbs []model.Bulb

	for rows.Next() {
		var bulb model.Bulb

		err := rows.Scan(
			&bulb.ID, &bulb.Name, &bulb.BulbType, &bulb.Luminance, &bulb.Red, &bulb.Green, &bulb.Blue,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan bulb from bulbs table: %v", err)
		}

		bulbs = append(bulbs, bulb)
	}

	return bulbs, nil
}

func quoteSlice(slice []string) []string {
	quotedSlice := make([]string, len(slice))
	for i, s := range slice {
		quotedSlice[i] = fmt.Sprintf("'%s'", s)
	}

	return quotedSlice
}
