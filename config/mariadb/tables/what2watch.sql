
CREATE TABLE genre (
    genre_id INT AUTO_INCREMENT PRIMARY KEY,
    genre_name VARCHAR(255) UNIQUE
);

CREATE TABLE user_group (
    group_id INT AUTO_INCREMENT PRIMARY KEY,
    group_name VARCHAR(255) UNIQUE
);

CREATE TABLE group_genres (
    group_genres_id INT AUTO_INCREMENT PRIMARY KEY,
    group_id INT,
    genre_id INT,
    FOREIGN KEY (group_id) REFERENCES user_group(group_id),
    FOREIGN KEY (genre_id) REFERENCES genre(genre_id)
);

CREATE TABLE user (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    mail VARCHAR(255) UNIQUE,
    name VARCHAR(255)
);

CREATE TABLE movie_group_recommendation (
    group_id INT,
    accuracy FLOAT,
    movie_id INT,
    PRIMARY KEY (movie_id, group_id),
    FOREIGN KEY (group_id) REFERENCES user_group(group_id)
);

CREATE TABLE movie_user_recommendation (
    user_id INT,
    accuracy FLOAT,
    movie_id INT,
    PRIMARY KEY (movie_id, user_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);


CREATE TABLE review (
    review_id INT AUTO_INCREMENT PRIMARY KEY,
    rating INT,
    viewed BOOLEAN,
    user_id INT,
    movie_id INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE user_group_membership (
    user_group_membership_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    group_id INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (group_id) REFERENCES user_group(group_id)
);

CREATE TABLE user_genre_preferences (
    user_genre_preferences_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    genre_id INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (genre_id) REFERENCES genre(genre_id)
);


INSERT INTO genre (genre_id, genre_name) VALUES
(28, 'Action'),
(12, 'Adventure'),
(16, 'Animation'),
(35, 'Comedy'),
(80, 'Crime'),
(99, 'Documentary'),
(18, 'Drama'),
(10751, 'Family'),
(14, 'Fantasy'),
(36, 'History'),
(27, 'Horror'),
(10402, 'Music'),
(9648, 'Mystery'),
(10749, 'Romance'),
(878, 'Science Fiction'),
(10770, 'TV Movie'),
(53, 'Thriller'),
(10752, 'War'),
(37, 'Western');
