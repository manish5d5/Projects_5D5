package repositories

import (
	"context"
	"GymBuddy/error"
	"GymBuddy/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"fmt"
)

type WorkoutRepository struct {
	db *pgxpool.Pool
}

func NewWorkoutRepository(db *pgxpool.Pool) *WorkoutRepository {
	return &WorkoutRepository{
		db: db,
	}
}

//
// ─── CREATE WORKOUT ─────────────────────────────────────────────────────────────
// Create inserts a full workout → muscle → exercise set
func (r *WorkoutRepository) Create(ctx context.Context, w models.Workouts) (int, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	// ─────────────── 1. CHECK IF EXERCISE ALREADY EXISTS ───────────────
	var existingExerciseID int
	err = tx.QueryRow(ctx, `
		SELECT exercise_id 
		FROM exercises 
		WHERE exercise_name = $1
	`, w.Exercise_name).Scan(&existingExerciseID)

	if err == nil {
		// Exercise already exists → return immediately
		_ = tx.Commit(ctx)
		log.Println(" Exercise already exists → return immediately")
		return existingExerciseID, nil
	}

	// ─────────────── 2. CHECK IF MUSCLE EXISTS ───────────────
	var muscleID int
	err = tx.QueryRow(ctx, `
		SELECT muscle_id 
		FROM muscles
		WHERE muscle_name = $1
	`, w.Muscle_name).Scan(&muscleID)

	if err != nil { // muscle not found
		// ─────────────── 3. CHECK IF WORKOUT EXISTS ───────────────
		var workoutID int
		err = tx.QueryRow(ctx, `
			SELECT workout_id 
			FROM workouts
			WHERE workout_name = $1
		`, w.Workout_name).Scan(&workoutID)

		if err != nil { // workout not found → INSERT workout
			err = tx.QueryRow(ctx, `
				INSERT INTO workouts (workout_name)
				VALUES ($1)
				RETURNING workout_id
			`, w.Workout_name).Scan(&workoutID)
			if err != nil {
				return 0, fmt.Errorf("failed to insert workout: %w", err)
			}
		}

		// ─────────────── INSERT new muscle ───────────────
		err = tx.QueryRow(ctx, `
			INSERT INTO muscles (workout_id, muscle_name)
			VALUES ($1, $2)
			RETURNING muscle_id
		`, workoutID, w.Muscle_name).Scan(&muscleID)
		if err != nil {
			return 0, fmt.Errorf("failed to insert muscle: %w", err)
		}
	}

	// ─────────────── INSERT new exercise (because not found earlier) ───────────────
	var exerciseID int
	err = tx.QueryRow(ctx, `
		INSERT INTO exercises (muscle_id, exercise_name, max_weight, rep_count)
		VALUES ($1, $2, $3, $4)
		RETURNING exercise_id
	`, muscleID, w.Exercise_name, w.MaxWeight, w.Reps).Scan(&exerciseID)

	if err != nil {
		return 0, fmt.Errorf("failed to insert exercise: %w", err)
	}

	// Commit transaction
	if err = tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("failed to commit: %w", err)
	}

	return exerciseID, nil
}


//
// ─── GET ALL WORKOUTS ───────────────────────────────────────────────────────────
//
func (r *WorkoutRepository) GetAll(ctx context.Context) ([]models.Workouts, error) {
	query := `
		SELECT 
			w.workout_id,
			w.workout_name,
			m.muscle_id,
			m.muscle_name,
			e.exercise_id,
			e.exercise_name,
			e.max_weight,
			e.rep_count
		FROM workouts w
		JOIN muscles m ON m.workout_id = w.workout_id
		JOIN exercises e ON e.muscle_id = m.muscle_id
		ORDER BY w.workout_id DESC, m.muscle_id, e.exercise_id;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Println("(repo.go) Error fetching workouts:", err)
		return nil, apperror.ErrNotFound
	}
	defer rows.Close()

	var list []models.Workouts

	for rows.Next() {
		var w models.Workouts

		if err := rows.Scan(
			&w.Workout_id,
			&w.Workout_name,
			&w.Muscle_id,
			&w.Muscle_name,
			&w.Exercise_id,
			&w.Exercise_name,
			&w.MaxWeight,
			&w.Reps,
		); err != nil {
			log.Println("(repo.go) Error scanning workout:", err)
			return nil, err
		}

		list = append(list, w)
	}

	return list, nil
}

//
// ─── GET WORKOUTS BY workout_for  (like "chest") ───────────────────────────────
//
func (r *WorkoutRepository) GetByWorkoutName(ctx context.Context, workoutName string) ([]models.Workouts, error) {
	query := `
		SELECT 
			w.workout_id,
			w.workout_name,
			m.muscle_id,
			m.muscle_name,
			e.exercise_id,
			e.exercise_name,
			e.max_weight,
			e.rep_count
		FROM workouts w
		JOIN muscles m ON m.workout_id = w.workout_id
		JOIN exercises e ON e.muscle_id = m.muscle_id
		WHERE w.workout_name = $1
		ORDER BY m.muscle_id, e.exercise_id;
	`

	rows, err := r.db.Query(ctx, query, workoutName)
	if err != nil {
		log.Println("(repo.go) Error fetching workouts by workout name:", err)
		return nil, apperror.ErrNotFound
	}
	defer rows.Close()

	var list []models.Workouts

	for rows.Next() {
		var w models.Workouts

		if err := rows.Scan(
			&w.Workout_id,
			&w.Workout_name,
			&w.Muscle_id,
			&w.Muscle_name,
			&w.Exercise_id,
			&w.Exercise_name,
			&w.MaxWeight,
			&w.Reps,
		); err != nil {
			log.Println("(repo.go) Error scanning workout:", err)
			return nil, err
		}

		list = append(list, w)
	}

	if len(list) == 0 {
		return nil, apperror.ErrNotFound
	}

	return list, nil
}

//
// ─── UPDATE WORKOUT ─────────────────────────────────────────────────────────────
//
func (r *WorkoutRepository) UpdateExercise(ctx context.Context, w models.Workouts) error {

    // ─── 1. VALIDATE muscle_id exists ────────────────────────────────
    var exists bool
    err := r.db.QueryRow(ctx, `
        SELECT EXISTS (
            SELECT 1 FROM muscles WHERE muscle_id = $1
        )
    `, w.Muscle_id).Scan(&exists)

    if err != nil {
        log.Println("(repo.go) Error checking muscle_id:", err)
        return apperror.ErrUpdateFailed
    }

    if !exists {
        log.Println("(repo.go) muscle_id not found:", w.Muscle_id)
        return fmt.Errorf("muscle_id %d does not exist", w.Muscle_id)
    }

    // ─── 2. UPDATE exercise ────────────────────────────────────────────
    query := `
        UPDATE exercises
        SET exercise_name = $1,
            muscle_id     = $2,
            max_weight    = $3,
            rep_count     = $4
        WHERE exercise_id = $5
    `

    cmd, err := r.db.Exec(ctx, query,
        w.Exercise_name,
        w.Muscle_id,
        w.MaxWeight,
        w.Reps,
        w.Exercise_id,
    )

    if err != nil {
        log.Println("(repo.go) Error updating exercise:", err)
        return apperror.ErrUpdateFailed
    }

    if cmd.RowsAffected() == 0 {
        log.Println("(repo.go) No exercise found with ID:", w.Exercise_id)
        return apperror.ErrNotFound
    }

    return nil
}

//
// ─── Patc WORKOUT ─────────────────────────────────────────────────────────────
//
func (r *WorkoutRepository) PatchExercise(ctx context.Context, w models.Workouts) error {
	query := `
		UPDATE exercises
		SET max_weight = $1,
		    rep_count = $2
		WHERE exercise_id = $3
	`

	cmd, err := r.db.Exec(ctx, query,
		w.MaxWeight,
		w.Reps,
		w.Exercise_id,
	)
	if err != nil {
		log.Println("(repo.go) Error patching exercise:", err)
		return apperror.ErrUpdateFailed
	}

	if cmd.RowsAffected() == 0 {
		log.Println("(repo.go) No exercise found to patch with ID:", w.Exercise_id)
		return apperror.ErrNotFound
	}

	return nil
}
//
// ─── DELETE WORKOUT ─────────────────────────────────────────────────────────────
//
func (r *WorkoutRepository) DeleteExercise(ctx context.Context, id int) error {
	cmd, err := r.db.Exec(ctx, `
		DELETE FROM exercises 
		WHERE exercise_id = $1
	`, id)

	if err != nil {
		log.Println("(repo.go) Error deleting exercise:", err)
		return apperror.ErrDeleteFailed
	}

	if cmd.RowsAffected() == 0 {
		log.Println("(repo.go) No exercise found to delete with ID:", id)
		return apperror.ErrNotFound
	}

	return nil
}

