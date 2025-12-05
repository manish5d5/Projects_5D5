ALTER TABLE exercises DROP CONSTRAINT IF EXISTS unique_exercise_per_muscle;
DROP TABLE IF EXISTS exercises CASCADE;
