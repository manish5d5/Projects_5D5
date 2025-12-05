package http

import (
	"GymBuddy/error"
	"GymBuddy/models"
	"GymBuddy/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	service *service.WorkoutService
}

func NewWorkoutHandler(s *service.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{service: s}
}

//
// ─── HELPER FUNCTION TO WRITE JSON RESPONSE ─────────────────────────────────────
//
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

//
// ─── HELPER FUNCTION TO HANDLE OUR CUSTOM APP ERRORS ───────────────────────────
func handleError(w http.ResponseWriter, err error) {
    appErr, ok := err.(*apperror.AppError)
    if !ok {
        // Unknown error - return 500
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    // Known error
    w.WriteHeader(appErr.Code)
    json.NewEncoder(w).Encode(map[string]string{
        "error": appErr.Message,
    })
}


//
// ─── CREATE WORKOUT: POST /api/v1/workouts ─────────────────────────────────────
//
// ─── CREATE WORKOUT: POST /api/v1/workouts ─────────────────────────────────────
func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var wkt models.Workouts

	if err := json.NewDecoder(r.Body).Decode(&wkt); err != nil {
		log.Println("(handler.go)Error decoding workout JSON:", err)
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateWorkout(r.Context(), wkt)
	if err != nil {
		log.Println("(handler.go)Error creating workout:", err)	
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusCreated, map[string]any{
		"message": "workout created",
		"id":      id,
	})
}

//
// ─── GET ALL WORKOUTS: GET /api/v1/workouts ─────────────────────────────────────
//
// ─── GET ALL WORKOUTS: GET /api/v1/workouts ─────────────────────────────────────
func (h *WorkoutHandler) GetAllWorkouts(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAllWorkouts(r.Context())
	if err != nil {
		log.Println("(handler.go)Error getting all workouts:", err)	
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, data)
}

//
// ─── SEARCH BY workout_for: GET /api/v1/workouts/search/{workout_for} ───────────
//
func (h *WorkoutHandler) GetByWorkoutName(w http.ResponseWriter, r *http.Request) {
	workoutName := chi.URLParam(r, "workout_Name")

	data, err := h.service.GetByWorkoutName(r.Context(), workoutName)
	if err != nil {
		log.Println("(handler.go)Error getting workouts by workout_for:", err)	
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, data)
}

//
// ─── UPDATE WORKOUT: PUT /api/v1/workouts/{id} ──────────────────────────────────
//
func (h *WorkoutHandler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
    // 1. Extract ID from URL
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        fmt.Println("ErrInvalidInput1")
        handleError(w, apperror.ErrInvalidInput)
        return
    }

    // 2. Decode JSON body
    var wkt models.Workouts

    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() // ❗ blocks unknown JSON fields

    if err := decoder.Decode(&wkt); err != nil {
        // unknown field OR bad JSON
        fmt.Println("ErrInvalidInput2")
        handleError(w, apperror.ErrInvalidInput)
        return
    }

    // 3. PATCH must contain ONLY reps or max_weight
    if wkt.Reps <=0 && wkt.MaxWeight < 0 {
        // no valid patch fields
        fmt.Println("ErrInvalidInput3")
        handleError(w, apperror.ErrInvalidInput)
        return
    }

    // If JSON contains empty string for fields like "muscle": "",
    // they decode into default zero values. So reject them too:
    if wkt.Exercise_id != 0 || wkt.Muscle_id<0 {
        fmt.Println("ErrInvalidInput4")
        handleError(w, apperror.ErrInvalidInput)
        return
    }

    // 3. Set ID from URL (JSON should NOT contain ID)
    wkt.Exercise_id = id

    // 4. Call service layer
    if err := h.service.UpdateExercise(r.Context(), wkt); err != nil {
        log.Println("(handler.go) Error updating workout:", err)
        handleError(w, err)
        return
    }

    // 5. Success response
    respondJSON(w, http.StatusOK, map[string]string{
        "message": "workout updated",
    })
}



// ─── PATCH WORKOUT: PATCH /api/v1/workouts/{id} ───────────────────────────────
//

func (h *WorkoutHandler) PatchExercise(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil || id <= 0 {
        handleError(w, apperror.ErrInvalidInput)
        return
    }

    // 2. Decode JSON body
    var wkt models.Workouts

    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() // ❗ blocks unknown JSON fields

    if err := decoder.Decode(&wkt); err != nil {
        // unknown field OR bad JSON
        handleError(w, apperror.ErrInvalidInput)
        return
    }

    // 3. PATCH must contain ONLY reps or max_weight
    if wkt.Reps <0 && wkt.MaxWeight < 0 {
        // no valid patch fields
        handleError(w, apperror.ErrInvalidInput)
        return
    }

    // If JSON contains empty string for fields like "muscle": "",
    // they decode into default zero values. So reject them too:
    if wkt.Exercise_id != 0 ||wkt.Muscle_name != "" || wkt.Exercise_name != "" || wkt.Workout_name != "" || wkt.Muscle_id != 0 {
        handleError(w, apperror.ErrInvalidInput)
        return
    }
	wkt.Exercise_id = id
	if err := h.service.PatchExercise(r.Context(), wkt); err != nil {
		log.Println("(handler.go)Error patching workout:", err)	
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "workout updated"})
}



//
// ─── DELETE WORKOUT: DELETE /api/v1/workouts/{id} ───────────────────────────────
//
func (h *WorkoutHandler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := h.service.DeleteExercise(r.Context(), id); err != nil {
		log.Println("(handler.go)Error deleting exercise:", err)		
		handleError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "exercise deleted"})
}
