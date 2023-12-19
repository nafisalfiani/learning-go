package repository

import (
	"database/sql"
	"fmt"
	"game-store/entity"
)

type branch struct {
	db *sql.DB
}

type Interface interface {
	List() ([]entity.Branch, error)
	Get(id string) (entity.Branch, error)
	Create(branch entity.Branch) (entity.Branch, error)
	Update(id string, newData entity.Branch) (entity.Branch, error)
	Delete(id string) error
}

func InitBranch(db *sql.DB) Interface {
	return &branch{
		db: db,
	}
}

// List fetch a list of branches
func (b *branch) List() ([]entity.Branch, error) {
	var branches []entity.Branch
	rows, err := b.db.Query("SELECT id, name, location FROM branch")
	if err != nil {
		return branches, err
	}
	defer rows.Close()

	for rows.Next() {
		var Branch entity.Branch
		err := rows.Scan(&Branch.Id, &Branch.Name, &Branch.Location)
		if err != nil {
			return branches, err
		}

		branches = append(branches, Branch)
	}

	return branches, nil
}

// Get fetch specific Branch by given id
func (b *branch) Get(id string) (entity.Branch, error) {
	var branch entity.Branch
	err := b.db.QueryRow("SELECT id, name, Location FROM branch WHERE id = ?", id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Location,
	)
	if err != nil {
		return branch, err
	}

	return branch, nil
}

// Create creates new Branch based on given Branch struct
func (b *branch) Create(Branch entity.Branch) (entity.Branch, error) {
	result, err := b.db.Exec("INSERT INTO branch (name, Location) VALUES (?, ?)",
		Branch.Name,
		Branch.Location,
	)
	if err != nil {
		return Branch, err
	}

	BranchID, err := result.LastInsertId()
	if err != nil {
		return Branch, err
	}

	return b.Get(fmt.Sprintf("%v", BranchID))
}

// Update updates existing Branch based on given Branch struct
func (b *branch) Update(id string, newData entity.Branch) (entity.Branch, error) {
	Branch, err := b.Get(id)
	if err != nil {
		return Branch, err
	}

	var name, location any
	if newData.Name == "" {
		name = nil
	} else {
		name = newData.Name
		Branch.Name = newData.Name
	}

	if newData.Location == "" {
		location = nil
	} else {
		location = newData.Location
		Branch.Location = newData.Location
	}

	if _, err := b.db.Exec("UPDATE branch SET name = COALESCE(?, name), location = COALESCE(?, location) WHERE id = ?",
		name,
		location,
		id,
	); err != nil {
		return Branch, err
	}

	return Branch, nil
}

// Delete deletes existing Branch
func (b *branch) Delete(id string) error {
	res, err := b.db.Exec("DELETE FROM branch WHERE id = ?", id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count < 1 {
		return sql.ErrNoRows
	}

	return nil
}
