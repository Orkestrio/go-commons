package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, _ := q.FormattedQuery()

	now := time.Now()
	fmt.Println(string(query))
	fmt.Println(now.Sub(q.StartTime))

	return nil
}
