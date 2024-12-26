CREATE TABLE comments (
    id PRIMARY KEY AUTOINCREMENT,
    content TEXT NOT NULL,
    date_creation CURRENT_TIMESTAMP
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
    );