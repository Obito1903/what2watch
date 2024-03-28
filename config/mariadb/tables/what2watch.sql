-- TABLE
CREATE TABLE genre (
    genre_id INT PRIMARY KEY,
    genre_name VARCHAR(255) UNIQUE
);

CREATE TABLE user_group (
    group_id INT PRIMARY KEY,
    group_name VARCHAR(255) UNIQUE
);

CREATE TABLE user (
    user_id INT PRIMARY KEY,
    name VARCHAR(255),
    age INT,
    group_id INT,
    FOREIGN KEY (group_id) REFERENCES user_group(group_id)
);


CREATE TABLE user_group_membership (
    user_id INT,
    group_id INT,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (group_id) REFERENCES user_group(group_id)
);

CREATE TABLE user_genre_preferences (
    user_id INT,
    genre_id INT,
    PRIMARY KEY (user_id, genre_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (genre_id) REFERENCES genre(genre_id)
);

CREATE TABLE review (
    review_id INT PRIMARY KEY,
    rating INT,
    viewed BOOLEAN,
    user_id INT,
    movie_id INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (movie_id) REFERENCES movie(movie_id)
);


CREATE TABLE log {
    log_id INT PRIMARY KEY,
    log_level INT,
    log_category VARCHAR(255),
    log_time TIMESTAMP,
    log_message TEXT,
}