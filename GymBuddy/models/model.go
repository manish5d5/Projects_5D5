package models


type Workouts struct {
    Workout_id    int       `json:"wid"`
    Workout_name  string    `json:"workout_name"` 

    Muscle_id     int       `json:"mid"`
    Muscle_name   string    `json:"muscle_name"`
    
    Exercise_id   int       `json:"eid"`
    Exercise_name string    `json:"exercise_name"`
    MaxWeight     float64   `json:"max_weight"`        // weight used
    Reps          int       `json:"reps"`  
}
