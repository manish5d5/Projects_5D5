CREATE TABLE workouts (
    workout_id SERIAL PRIMARY KEY,
    workout_name VARCHAR(100) UNIQUE NOT NULL
);
