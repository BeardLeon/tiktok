CREATE DATABASE IF NOT EXISTS tiktok DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- After create all table
ALTER TABLE videos ADD FOREIGN KEY (author_id) REFERENCES users(id);

ALTER TABLE comments ADD FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE comments ADD FOREIGN KEY (video_id) REFERENCES videos(id);

ALTER TABLE follows ADD FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE follows ADD FOREIGN KEY (follower_id) REFERENCES users(id);

ALTER TABLE likes ADD FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE likes ADD FOREIGN KEY (video_id) REFERENCES videos(id);