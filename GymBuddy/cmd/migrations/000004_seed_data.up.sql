-- =====================================================
-- WORKOUTS
-- =====================================================
INSERT INTO workouts (workout_name) VALUES
('chest'),
('back'),
('shoulders'),
('legs'),
('arms');

-- =====================================================
-- MUSCLES
-- =====================================================
-- Chest (workout_id = 1)
INSERT INTO muscles (workout_id, muscle_name) VALUES
(1, 'upper chest'),
(1, 'middle chest'),
(1, 'lower chest');

-- Back (workout_id = 2)
INSERT INTO muscles (workout_id, muscle_name) VALUES
(2, 'lats'),
(2, 'upper back'),
(2, 'lower back');

-- Shoulders (workout_id = 3)
INSERT INTO muscles (workout_id, muscle_name) VALUES
(3, 'front delts'),
(3, 'side delts'),
(3, 'rear delts');

-- Legs (workout_id = 4)
INSERT INTO muscles (workout_id, muscle_name) VALUES
(4, 'quadriceps'),
(4, 'hamstrings'),
(4, 'calves');

-- Arms (workout_id = 5)
INSERT INTO muscles (workout_id, muscle_name) VALUES
(5, 'biceps'),
(5, 'triceps'),
(5, 'forearms');

-- =====================================================
-- EXERCISES
-- =====================================================
-- Chest exercises
INSERT INTO exercises (muscle_id, exercise_name, max_weight, rep_count) VALUES
(1, 'incline bench press', 60, 10),
(2, 'flat bench press', 80, 8),
(3, 'decline bench press', 70, 10);

-- Back exercises
INSERT INTO exercises (muscle_id, exercise_name, max_weight, rep_count) VALUES
(4, 'lat pulldown', 55, 12),
(5, 'seated row', 65, 10),
(6, 'back extension', 30, 15);

-- Shoulders
INSERT INTO exercises (muscle_id, exercise_name, max_weight, rep_count) VALUES
(7, 'front raise', 12, 12),
(8, 'lateral raise', 10, 15),
(9, 'rear delt fly', 10, 15);

-- Legs
INSERT INTO exercises (muscle_id, exercise_name, max_weight, rep_count) VALUES
(10, 'leg press', 150, 10),
(11, 'romanian deadlift', 80, 8),
(12, 'standing calf raise', 50, 12);

-- Arms
INSERT INTO exercises (muscle_id, exercise_name, max_weight, rep_count) VALUES
(13, 'bicep curl', 20, 12),
(14, 'tricep pushdown', 35, 12),
(15, 'wrist curl', 15, 15);
