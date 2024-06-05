package postgres

import (
	"arch/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var ctx = context.Background()

func Conn(cfg config.Config) (*pgx.Conn, error) {

	dbUrl := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s`,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	fmt.Println(dbUrl)

	return pgx.Connect(ctx, dbUrl)
}
