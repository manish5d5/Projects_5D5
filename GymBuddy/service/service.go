package service

import (
	"context"
	"GymBuddy/error"
	"GymBuddy/models"
	"log"
	repositories "GymBuddy/repositories"
)

type WorkoutService struct {
	repo *repositories.WorkoutRepository
}

func NewWorkoutService(r *repositories.WorkoutRepository) *WorkoutService {
	return &WorkoutService{repo: r}
}

func (s *WorkoutService) CreateWorkout(ctx context.Context, w models.Workouts) (int, error) {

	if w.Workout_name == "" {
		log.Println("(service.go)Invalid input: WorkoutFor is empty")
		return 0, apperror.ErrInvalidInput
	}
	if w.Muscle_name == "" {
		log.Println("(service.go)Invalid input: Muscle is empty")
		return 0, apperror.ErrInvalidInput
	}
	if w.Exercise_name == "" {
		log.Println("(service.go)Invalid input: Exercise is empty")
		return 0, apperror.ErrInvalidInput
	}
	if w.MaxWeight <= 0 {
		log.Println("(service.go)Invalid input for MaxWeight ")
		return 0, apperror.ErrInvalidInput
	}
	if w.Reps <= 0 {
		log.Println("(service.go)Invalid input for Reps ")
		return 0, apperror.ErrInvalidInput
	}

	return s.repo.Create(ctx, w)
}

func (s *WorkoutService) GetAllWorkouts(ctx context.Context) ([]models.Workouts, error) {
	return s.repo.GetAll(ctx)
}

func (s *WorkoutService) GetByWorkoutName(ctx context.Context, workoutName string) ([]models.Workouts, error) {
	if workoutName == "" {
		log.Println("(service.go)Invalid input: workoutName is empty")	
		return nil, apperror.ErrInvalidInput
	}
	return s.repo.GetByWorkoutName(ctx, workoutName)
}

func (s *WorkoutService) UpdateExercise(ctx context.Context, w models.Workouts) error {
	if w.Exercise_id <= 0 {
		log.Println("(service.go)Invalid input: ID must be greater than zero")
		return apperror.ErrInvalidInput
	}
	if w.Exercise_name == "" {
		log.Println("(service.go)Invalid input: WorkoutFor is empty")
		return apperror.ErrInvalidInput
	}
	if w.Muscle_id <= 0 {
		log.Println("(service.go)Invalid input: Muscle ID must be greater than zero")
		return apperror.ErrInvalidInput
	}
	if w.MaxWeight < 0 {
		log.Println("(service.go)Invalid input: MaxWeight cant be negative")
		return apperror.ErrInvalidInput
	}
	if w.Reps < 0 {
		log.Println("(service.go)Invalid input: Reps cant be negative")
		return apperror.ErrInvalidInput
	}
	return s.repo.UpdateExercise(ctx, w)
}

func (s *WorkoutService) PatchExercise(ctx context.Context, w models.Workouts) error {
	if w.Exercise_id <= 0 {
		log.Println("(service.go)Invalid input: ID must be greater than zero")
		return apperror.ErrInvalidInput
	}
	return s.repo.PatchExercise(ctx, w)
}

func (s *WorkoutService) DeleteExercise(ctx context.Context, id int) error {
	if id <= 0 {
		log.Println("(service.go)Invalid input: ID must be greater than zero")	
		return apperror.ErrInvalidInput
	}
	return s.repo.DeleteExercise(ctx, id)
}
