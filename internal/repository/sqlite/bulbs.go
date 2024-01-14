package sqlite

import (
	model "cmd/main.go/internal/model"
	"database/sql"
	"fmt"

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

func (r *BulbsRepo) Get(ip string) (model.Bulb, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE ip_addr = ?;", r.tableName)

	rows, err := r.db.Query(query, ip)
	if err != nil {
		return model.Bulb{}, fmt.Errorf("could not get bulb from bulbs table: %v", err)
	}

	defer rows.Close()

	var bulb model.Bulb

	for rows.Next() {
		err := rows.Scan(
			&bulb.ID, &bulb.IP, &bulb.BulbType, &bulb.Luminance, &bulb.Red, &bulb.Green, &bulb.Blue,
		)
		if err != nil {
			return model.Bulb{}, fmt.Errorf("could not scan bulb from bulbs table: %v", err)
		}
	}

	return bulb, nil
}
