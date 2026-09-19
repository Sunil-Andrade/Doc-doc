package db

import (
	"context"
	"docdoc/models"

	"github.com/jackc/pgx/v5"
)

func GetDocumentByCode(ctx context.Context, conn *pgx.Conn, code string) (*models.Document, error) {
	var document models.Document

	err := conn.QueryRow(ctx,
		`SELECT id, code, file_path, created_at
		 FROM documents
		 WHERE code = $1`, code).Scan(&document.ID,
		&document.Code,
		&document.FilePath,
		&document.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &document, nil
}
