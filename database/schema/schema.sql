CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  role VARCHAR(10) NOT NULL DEFAULT "user", 
  password VARCHAR(255) NOT NULL
);
DROP TABLE IF EXISTS categories;
CREATE TABLE IF NOT EXISTS categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(255) UNIQUE NOT NULL
);
CREATE TABLE IF NOT EXISTS posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  is_approved BOOLEAN DEFAULT FALSE,
  approved_by INTEGER,
  approval_date TIMESTAMP,
  visibility VARCHAR(20) DEFAULT "pending",
  user_id INTEGER NOT NULL,
  creat_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (approved_by) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS comments (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  content TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  date_creation TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  post_id INTEGER NOT NULL,
   is_approved BOOLEAN DEFAULT FALSE,
  approved_by INTEGER, 
  approval_date TIMESTAMP,
  visibility VARCHAR(20) DEFAULT "pending",
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (post_id) REFERENCES posts(id),
   FOREIGN KEY (approved_by) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS categorie_report (
  id INTEGER PRIMARY KEY AUTOINCREMENT, 
  name VARCHAR(255) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS report(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
 
  user_id INTEGER NOT NULL,
   post_id INTEGER NOT NULL,
   comment_id INTEGER,
   report_category_id INTEGER NOT NULL ,
   report_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  status VARCHAR(20) DEFAULT "pending", -- "pending", "reviewed", "rejected", "accepted"
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (post_id) REFERENCES posts(id),
  FOREIGN KEY (comment_id) REFERENCES comments(id),
  FOREIGN KEY (report_category_id) REFERENCES report_categories(id),

  CHECK ((post_id IS NULL AND comment_id IS NOT NULL) OR (post_id IS NOT NULL AND comment_id IS NULL))
);

CREATE TABLE IF NOT EXISTS reactPost(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  post_id INTEGER NOT NULL,
  react_type  VARCHAR(255) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (post_id) REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS reactComment(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  comment_id INTEGER NOT NULL,
  react_type VARCHAR(255) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (comment_id) REFERENCES comments(id)
);


CREATE TABLE IF NOT EXISTS sessionss (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  token VARCHAR(255) NOT NULL,
  expired_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id), 
  UNIQUE(user_id)
);

CREATE TABLE IF NOT EXISTS post_categorie (
  post_id INTEGER NOT NULL,
  categorie_id INTEGER NOT NULL,
  FOREIGN KEY (post_id) REFERENCES posts(id),
  FOREIGN KEY (categorie_id) REFERENCES categories(id),
  PRIMARY KEY (post_id, categorie_id)
);
CREATE TABLE IF NOT EXISTS moderator_requests (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  request_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  status VARCHAR(20) DEFAULT "pending", -- "pending", "approved", "rejected"
  review_date TIMESTAMP,
  reason TEXT, 
  FOREIGN KEY (user_id) REFERENCES users(id)
);


INSERT INTO categories (name) 
VALUES ('Sport'), ('Music'), ('Movies'), ('Science'), ('Politics'), ('Culture'), ('Technology')
ON CONFLICT (name) DO NOTHING;

INSERT INTO categorie_report (name) 
VALUES ('Irrelevant'), ('Obscene'), ('Illegal'), ('Insulting')
ON CONFLICT (name) DO NOTHING;

INSERT INTO users (username, email, role, password)
VALUES ('admin', 'admin@gmail.com', 'admin', 'admin12345')
ON CONFLICT (username) DO NOTHING;