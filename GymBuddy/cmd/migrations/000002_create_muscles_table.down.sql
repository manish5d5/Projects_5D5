ALTER TABLE muscles DROP CONSTRAINT IF EXISTS unique_muscle_per_workout;
DROP TABLE IF EXISTS muscles CASCADE;
