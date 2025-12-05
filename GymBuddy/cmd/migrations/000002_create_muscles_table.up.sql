CREATE TABLE muscles (
    muscle_id SERIAL PRIMARY KEY,
    workout_id INT REFERENCES workouts(workout_id) ON DELETE CASCADE,
    muscle_name VARCHAR(100) NOT NULL
);

-- Prevent duplicate muscle name under the same workout
ALTER TABLE muscles
ADD CONSTRAINT unique_muscle_per_workout UNIQUE (workout_id, muscle_name);
