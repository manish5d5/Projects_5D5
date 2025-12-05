package main

import (
	"context"
	"log"
	"net/http"

	handler "GymBuddy/http"
	"GymBuddy/repositories"
	"GymBuddy/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	// Connect to PostgreSQL
	db, err := repositories.Connect(ctx, repositories.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1710519",
		Dbname:   "GymBuddy",
	})
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	// Initialize layers
	workoutRepo := repositories.NewWorkoutRepository(db)
	workoutService := service.NewWorkoutService(workoutRepo)
	workoutHandler := handler.NewWorkoutHandler(workoutService)

	// Routes
	r := chi.NewRouter()

	r.Route("/api/v2/workouts", func(r chi.Router) {
    r.Post("/", workoutHandler.CreateWorkout)
    r.Get("/", workoutHandler.GetAllWorkouts)
    r.Get("/search_workout/{workout_Name}", workoutHandler.GetByWorkoutName)

    r.Put("/exercises/{id}", workoutHandler.UpdateExercise)
    r.Patch("/exercises/{id}", workoutHandler.PatchExercise)
    r.Delete("/exercises/{id}", workoutHandler.DeleteExercise)
})


	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}