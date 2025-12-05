CREATE TABLE exercises (
    exercise_id SERIAL PRIMARY KEY,
    muscle_id INT REFERENCES muscles(muscle_id) ON DELETE CASCADE,
    exercise_name VARCHAR(100) NOT NULL,
    max_weight FLOAT CHECK (max_weight >= 0),
    rep_count INT CHECK (rep_count > 0)
);

-- Prevent duplicate exercise name under the same muscle
ALTER TABLE exercises
ADD CONSTRAINT unique_exercise_per_muscle UNIQUE (muscle_id, exercise_name);
