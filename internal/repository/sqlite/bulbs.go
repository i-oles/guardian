package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"cmd/main.go/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

type BulbsRepo struct {
	db        *sql.DB
	tableName string
}

func NewBulbsRepo(db *sql.DB, tableName string) *BulbsRepo {
	return &BulbsRepo{
		db:        db,
		tableName: tableName,
	}
}

func (r *BulbsRepo) Get(id string) (model.Bulb, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE bulb_id = ? LIMIT 1", r.tableName)

	var bulb model.Bulb
	err := r.db.QueryRow(query, id).Scan(
		&bulb.ID, &bulb.Name, &bulb.BulbType, &bulb.Brightness, &bulb.Red, &bulb.Green, &bulb.Blue,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return model.Bulb{}, fmt.Errorf("bulb not found with id: %s", id)
	case err != nil:
		return model.Bulb{}, fmt.Errorf("failed to get bulb: %w", err)
	default:
		return bulb, nil
	}
}

func (r *BulbsRepo) GetOfflineBulbs(onlineIDs []string) ([]model.Bulb, error) {
	if len(onlineIDs) == 0 {
		return r.getAllBulbs()
	}

	placeholders := make([]string, len(onlineIDs))
	args := make([]interface{}, len(onlineIDs))
	for i, id := range onlineIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE bulb_id NOT IN (%s)",
		r.tableName,
		strings.Join(placeholders, ","),
	)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query offline bulbs: %w", err)
	}

	defer rows.Close()

	bulbs, err := getBulbs(rows)
	if err != nil {
		return nil, fmt.Errorf("could not get bulbs rows: %v", err)
	}

	return bulbs, nil
}

func getBulbs(rows *sql.Rows) ([]model.Bulb, error) {
	var bulbs []model.Bulb

	for rows.Next() {
		var bulb model.Bulb

		if err := rows.Scan(
			&bulb.ID, &bulb.Name, &bulb.BulbType, &bulb.Brightness, &bulb.Red, &bulb.Green, &bulb.Blue,
		); err != nil {
			return nil, fmt.Errorf("failed to scan bulb: %w", err)
		}

		bulbs = append(bulbs, bulb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return bulbs, nil
}

func (r *BulbsRepo) getAllBulbs() ([]model.Bulb, error) {
	query := fmt.Sprintf("SELECT * FROM %s", r.tableName)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all bulbs: %w", err)
	}
	defer rows.Close()

	bulbs, err := getBulbs(rows)
	if err != nil {
		return nil, fmt.Errorf("could not get bulbs rows: %v", err)
	}

	return bulbs, nil
}
