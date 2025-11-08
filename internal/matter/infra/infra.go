package infra

import (
	"database/sql"
	"focussprint/internal/matter/domain"
)

type matterInfra struct {
	db *sql.DB
}

func NewMatterInfra(db *sql.DB) *matterInfra {
	return &matterInfra{db: db}
}

func (m *matterInfra) GetMatter(id int) (*domain.Matter, error) {
	query := `
		SELECT id, name, description, created_at, updated_at
		FROM matters
		WHERE id = ?
	`

	var matter domain.Matter
	err := m.db.QueryRow(query, id).Scan(
		&matter.ID,
		&matter.Name,
		&matter.Description,
		&matter.CreatedAt,
		&matter.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	// загружаем подпункты
	subQuery := `
		SELECT id, name, description, created_at, updated_at
		FROM sub_matters
		WHERE matter_id = ?
		ORDER BY id
	`

	rows, err := m.db.Query(subQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subMatters []*domain.SubMatter
	for rows.Next() {
		var sub domain.SubMatter
		err := rows.Scan(
			&sub.ID,
			&sub.Name,
			&sub.Description,
			&sub.CreatedAt,
			&sub.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		subMatters = append(subMatters, &sub)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	matter.SubMatter = subMatters
	return &matter, nil
}

func (m *matterInfra) CreateMatter(newMatter *domain.Matter) (id int, err error) {
	query := `
		INSERT INTO matters (name, description)
		VALUES (?, ?)
	`

	result, err := m.db.Exec(query, newMatter.Name, newMatter.Description)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func (m *matterInfra) UpdateMatter(idMatter int, newMatter *domain.Matter) (id int, err error) {
	query := `
		UPDATE matters
		SET name = ?, description = ?
		WHERE id = ?
	`

	_, err = m.db.Exec(query, newMatter.Name, newMatter.Description, idMatter)
	if err != nil {
		return 0, err
	}

	return idMatter, nil
}

func (m *matterInfra) DeleteMatter(id int) error {
	query := `DELETE FROM matters WHERE id = ?`

	_, err := m.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
