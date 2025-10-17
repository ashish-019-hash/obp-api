package repositories

import (
	"context"
	"obp-api-backend/internal/models"
	"obp-api-backend/pkg/db"
)

type userLockRepository struct{}

func NewUserLockRepository() UserLockRepository {
	return &userLockRepository{}
}

func (r *userLockRepository) Create(ctx context.Context, lock *models.UserLock) error {
	database := db.GetDB()
	query := `
		INSERT INTO user_locks (user_id, type_of_lock, last_lock_date)
		VALUES (?, ?, ?)
	`
	_, err := database.ExecContext(ctx, query,
		lock.UserID,
		lock.TypeOfLock,
		lock.LastLockDate,
	)
	return err
}

func (r *userLockRepository) GetByUserID(ctx context.Context, userID string) ([]*models.UserLock, error) {
	database := db.GetDB()
	query := `
		SELECT id, user_id, type_of_lock, last_lock_date
		FROM user_locks
		WHERE user_id = ?
	`

	rows, err := database.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locks []*models.UserLock
	for rows.Next() {
		lock := &models.UserLock{}
		err := rows.Scan(
			&lock.ID,
			&lock.UserID,
			&lock.TypeOfLock,
			&lock.LastLockDate,
		)
		if err != nil {
			return nil, err
		}
		locks = append(locks, lock)
	}

	return locks, rows.Err()
}

func (r *userLockRepository) IsLocked(ctx context.Context, userID string) (bool, error) {
	database := db.GetDB()
	query := `SELECT COUNT(*) FROM user_locks WHERE user_id = ?`

	var count int
	err := database.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userLockRepository) Delete(ctx context.Context, lockID int) error {
	database := db.GetDB()
	query := `DELETE FROM user_locks WHERE id = ?`
	_, err := database.ExecContext(ctx, query, lockID)
	return err
}
