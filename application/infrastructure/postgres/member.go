package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mongmx/system-d/application/domain/member"
)

// NewMemberRepository creates new member repository
func NewMemberRepository(db *sql.DB) (member.Repository, error) {
	sqlDB := sqlx.NewDb(db, "postgres")
	r := memberRepository{sqlDB}
	return &r, nil
}

type memberRepository struct {
	db *sqlx.DB
}
