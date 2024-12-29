-- Insert Users
INSERT INTO users (username, email, password) VALUES
('user1', 'user1@example.com', 'password1'),
('user2', 'user2@example.com', 'password2'),
('user3', 'user3@example.com', 'password3'),
('user4', 'user4@example.com', 'password4'),
('user5', 'user5@example.com', 'password5'),
('user6', 'user6@example.com', 'password6'),
('user7', 'user7@example.com', 'password7'),
('user8', 'user8@example.com', 'password8'),
('user9', 'user9@example.com', 'password9'),
('user10', 'user10@example.com', 'password10');

-- Insert Categories
INSERT INTO categories (name) VALUES
('Technology'),
('Health'),
('Finance'),
('Education'),
('Travel'),
('Food'),
('Lifestyle'),
('Entertainment'),
('Sports'),
('Science');

-- Insert Posts
INSERT INTO posts (title, content, user_id) VALUES
('The Rise of AI Technology', 'Artificial Intelligence is transforming industries.', 1),
('Healthy Living Tips', 'Here are some tips for a healthier lifestyle.', 2),
('Investing 101', 'Learn the basics of investing your money.', 3),
('Top 10 Travel Destinations for 2023', 'Explore these amazing places this year.', 4),
('Understanding Nutrition', 'A guide to understanding what you eat.', 5),
('The Future of Work', 'How technology is changing the workplace.', 6),
('Best Practices for Online Learning', 'Tips for effective online education.', 7),
('Delicious Recipes for Home Cooking', 'Try these easy recipes at home.', 8),
('Sports Events to Watch This Year', 'Don’t miss these exciting events!', 9),
('Scientific Discoveries of the Year', 'A look at the most important discoveries.', 10);

-- Insert Post Categories
INSERT INTO post_categories (post_id, category_id) VALUES
(1, 1), -- The Rise of AI Technology -> Technology
(2, 2), -- Healthy Living Tips -> Health
(3, 3), -- Investing 101 -> Finance
(4, 4), -- Top 10 Travel Destinations for 2023 -> Travel
(5, 2), -- Understanding Nutrition -> Health
(6, 1), -- The Future of Work -> Technology
(7, 4), -- Best Practices for Online Learning -> Education
(8, 6), -- Delicious Recipes for Home Cooking -> Food
(9, 9), -- Sports Events to Watch This Year -> Sports
(10, 10); -- Scientific Discoveries of the Year -> Science

-- Insert Comments
INSERT INTO comments (content, user_id, post_id) VALUES
('Great article on AI!', 2, 1),
('Very informative. Thanks!', 3, 2),
('I learned a lot from this post.', 4, 3),
('Can’t wait to travel again!', 5, 4),
('Nutrition is so important!', 6, 5),
('Interesting perspective on work.', 7, 6),
('Online learning has its challenges.', 8, 7),
('These recipes look delicious!', 9, 8),
('Excited for the upcoming games!', 10, 9),
('Science is fascinating!', 1, 10);

-- Insert Reactions to Posts
INSERT INTO reactPost (user_id, post_id, react_type) VALUES
(1, 1, 'like'),
(2, 1, 'like'),
(3, 2, 'dislike'),
(4, 3, 'like'),
(5, 4, 'like'),
(6, 5, 'like'),
(7, 6, 'dislike'),
(8, 7, 'like'),
(9, 8, 'like'),
(10, 9, 'dislike');

-- Insert Reactions to Comments
INSERT INTO reactComment (user_id, comment_id, react_type) VALUES
(1, 1, 'like'),
(2, 2, 'like'),
(3, 3, 'dislike'),
(4, 4, 'like'),
(5, 5, 'like');

