DROP DATABASE if exists radar_team;
create database radar_team;
USE radar_team;
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    nickname VARCHAR(50) NOT NULL DEFAULT 'guest',
    password_hash VARCHAR(255) NOT NULL,
    reset_token VARCHAR(255),
    reset_token_expiry DATETIME,
    email VARCHAR(255),
    account_status VARCHAR(20) DEFAULT 'active'
    );
DROP TABLE IF EXISTS radar_wiki;
CREATE TABLE radar_wiki (
       author_id INT AUTO_INCREMENT PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       oss_path VARCHAR(255) NOT NULL,
       author_nickname VARCHAR(50) NOT NULL,
       tag VARCHAR(50) NOT NULL DEFAULT 'None',
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       FOREIGN KEY (author_id) REFERENCES users(user_id)
);


INSERT INTO users (username, nickname, password_hash, email) VALUES
                                                                 ('userA', 'User A', 'hashed_password_for_userA', 'userA@example.com'),
                                                                 ('userB', 'User B', 'hashed_password_for_userB', 'userB@example.com'),
                                                                 ('userC', 'User C', 'hashed_password_for_userC', 'userC@example.com');


INSERT INTO radar_wiki (title, oss_path, author_nickname) VALUES
                                                              ('Article 1', 'oss://your-bucket/articles/article1.txt', 'UserA'),
                                                              ('Article 2', 'oss://your-bucket/articles/article2.pdf', 'UserB'),
                                                              ('Article 3', 'oss://your-bucket/articles/article3.docx', 'UserC');