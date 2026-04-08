CREATE DATABASE IF NOT EXISTS achievement_leveling;
USE achievement_leveling;

CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,
    google_id VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    picture VARCHAR(512),
    xp INT NOT NULL DEFAULT 0,
    level INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS achievement_categories (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(100),
    color VARCHAR(7),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS achievements (
    id CHAR(36) PRIMARY KEY,
    category_id CHAR(36),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    xp_reward INT NOT NULL DEFAULT 10,
    icon VARCHAR(100),
    difficulty ENUM('easy', 'medium', 'hard', 'epic') DEFAULT 'medium',
    repeatable BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES achievement_categories(id)
);

CREATE TABLE IF NOT EXISTS user_achievements (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    achievement_id CHAR(36) NOT NULL,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    times_completed INT DEFAULT 1,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (achievement_id) REFERENCES achievements(id),
    UNIQUE KEY unique_user_achievement (user_id, achievement_id)
);

CREATE TABLE IF NOT EXISTS level_thresholds (
    level INT PRIMARY KEY,
    xp_required INT NOT NULL,
    title VARCHAR(100)
);

-- Seed level thresholds
INSERT INTO level_thresholds (level, xp_required, title) VALUES
(1, 0, 'Novice'),
(2, 100, 'Beginner'),
(3, 300, 'Apprentice'),
(4, 600, 'Journeyman'),
(5, 1000, 'Adept'),
(6, 1500, 'Expert'),
(7, 2100, 'Veteran'),
(8, 2800, 'Master'),
(9, 3600, 'Grandmaster'),
(10, 4500, 'Legend');

-- Seed categories
INSERT INTO achievement_categories (id, name, icon, color) VALUES
('c1000000-0000-0000-0000-000000000001', 'Fitness', 'dumbbell', '#EF4444'),
('c1000000-0000-0000-0000-000000000002', 'Learning', 'book-open', '#3B82F6'),
('c1000000-0000-0000-0000-000000000003', 'Career', 'briefcase', '#10B981'),
('c1000000-0000-0000-0000-000000000004', 'Social', 'users', '#F59E0B'),
('c1000000-0000-0000-0000-000000000005', 'Creative', 'palette', '#8B5CF6');

-- Seed achievements
INSERT INTO achievements (id, category_id, title, description, xp_reward, difficulty, repeatable) VALUES
('a1000000-0000-0000-0000-000000000001', 'c1000000-0000-0000-0000-000000000001', 'First Workout', 'Complete your first workout session', 50, 'easy', FALSE),
('a1000000-0000-0000-0000-000000000002', 'c1000000-0000-0000-0000-000000000001', 'Week Warrior', 'Work out every day for a week', 200, 'medium', FALSE),
('a1000000-0000-0000-0000-000000000003', 'c1000000-0000-0000-0000-000000000002', 'Bookworm', 'Read a complete book', 100, 'easy', TRUE),
('a1000000-0000-0000-0000-000000000004', 'c1000000-0000-0000-0000-000000000002', 'Course Complete', 'Finish an online course', 300, 'hard', TRUE),
('a1000000-0000-0000-0000-000000000005', 'c1000000-0000-0000-0000-000000000003', 'Promotion', 'Get a job promotion', 500, 'epic', FALSE),
('a1000000-0000-0000-0000-000000000006', 'c1000000-0000-0000-0000-000000000003', 'Side Project', 'Launch a side project', 250, 'hard', TRUE),
('a1000000-0000-0000-0000-000000000007', 'c1000000-0000-0000-0000-000000000004', 'Team Player', 'Help a colleague with a task', 50, 'easy', TRUE),
('a1000000-0000-0000-0000-000000000008', 'c1000000-0000-0000-0000-000000000005', 'First Creation', 'Create something artistic', 75, 'easy', FALSE),
('a1000000-0000-0000-0000-000000000009', 'c1000000-0000-0000-0000-000000000005', 'Portfolio Piece', 'Complete a major creative project', 400, 'epic', TRUE),
('a1000000-0000-0000-0000-000000000010', 'c1000000-0000-0000-0000-000000000001', 'Marathon Runner', 'Run a full marathon', 500, 'epic', FALSE);
