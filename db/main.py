from typing import Union,List
from fastapi import FastAPI, Response, HTTPException
import os
import MySQLdb
import MySQLdb.cursors

from src.types import *


db_config = {
    'host': os.environ['DB_HOST'] if 'DB_HOST' in os.environ else '127.0.0.1',
    'port': int(os.environ['DB_PORT']) if 'DB_PORT' in os.environ else 3306,
    'user': os.environ['DB_USER'] if 'DB_USER' in os.environ else 'moviefinder',
    'passwd': os.environ['DB_PASS'] if 'DB_PASS' in os.environ else 'moviefinder',
    'db': os.environ['DB_NAME'] if 'DB_NAME' in os.environ else 'moviefinder',
    'cursorclass': MySQLdb.cursors.DictCursor,
}

db = MySQLdb.connect(**db_config)
c=db.cursor()

tags_metadata = [
    {
        "name": "Movies of Users",
        "description": "Manage the movies seen by the user.",
    },
    {
        "name": "Recommendations",
        "description": "Manage the recommendations for the users and the groups.",
    },
    {
        "name": "Tastes of Users",
        "description": "Manage the tastes in genre of a specific user.",
    },
    {
        "name": "Users",
        "description": "CRUD for users.",
    },
    {
        "name": "Groups",
        "description": "CR for groups.",
    },
    {
        "name": "Users from Groups",
        "description": "CRUD for users in a group.",
    },
    {
        "name": "Group tastes",
        "description": "Get the tastes for a group.",
    },
]

app = FastAPI(openapi_tags=tags_metadata)

# USERS
# MOVIES

@app.get("/users/{user_id}/movies", status_code=200, tags=["Movies of Users"])
def get_seen_movies(user_id: int) -> List[str]:
    c.execute("""SELECT movie_id FROM review WHERE user_id LIKE %s""" % (user_id))

    res = c.fetchall()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return res

@app.post("/users/{user_id}/movies/{movie_id}", status_code=200, tags=["Movies of Users"])
def add_movie_to_watched(user_id: int, movie_id: int, rating: int) -> ApiResponse:
    sql = "INSERT INTO review (rating, user_id, genre_id) VALUES (%s, %s, %s)"
    values = (rating, user_id, genre_id)

    c.execute(sql, values)

    db.commit()

    return {"message": "Movie added to user"}

@app.delete("/users/{user_id}/movies/{movie_id}", status_code=200, tags=["Movies of Users"])
def remove_movie_from_watched(user_id: int, movie_id: int) -> ApiResponse:
    c.execute("""DELETE FROM review WHERE user_id LIKE %s AND movie_id LIKE %s""" % (user_id, movie_id))

    db.commit()

    return {"message": "Movie deleted from user"}


# USERS
# RECOMMENDATIONS

@app.get("/users/{user_id}/recommendations", status_code=200, tags=["Recommendations"])
def get_recommendations(user_id: int) -> List[MovieReccomendationResponse]:
    c.execute("""SELECT (movie_id,accuracy) FROM movie_user_recommendation WHERE user_id LIKE %s""" % (user_id))

    res = c.fetchall()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return res

@app.post("/users/{user_id}/recommendations", status_code=200, tags=["Recommendations"])
def add_recommendations(user_id: int, movie_id: int, accuracy: float) -> ApiResponse:
    sql = "INSERT INTO movie_user_recommendation (accuracy, user_id, movie_id) VALUES (%s, %s, %s, %s)"
    values = (accuracy, user_id, movie_id)

    c.execute(sql, values)

    db.commit()

    return {"message": "Recommendation added"}


# USERS
# TASTES

@app.get("/users/{user_id}/tastes", status_code=200, tags=["Tastes of Users"])
def get_tastes(user_id: int) -> List[GenreResponse]:
    c.execute("""SELECT genre_name FROM genre, user_genre_preferences WHERE user_genre_preferences.user_id LIKE %s AND genre.genre_id = user_genre_preferences.genre_id""" % (user_id))

    res = c.fetchall()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return res

@app.post("/users/{user_id}/tastes/{genre_id}", status_code=200, tags=["Tastes of Users"])
def add_taste_for_user(user_id: int, genre_id: int) -> ApiResponse:
    sql = "INSERT INTO user_genre_preferences (user_id, genre_id) VALUES (%s, %s)"
    values = (user_id, genre_id)

    c.execute(sql, values)

    db.commit()

    return {"message": "Taste added to user"}

@app.delete("/users/{user_id}/tastes/{genre_id}", status_code=200, tags=["Tastes of Users"])
def remove_taste_from_user(user_id: int, genre_id: int) -> ApiResponse:
    c.execute("""DELETE FROM user_genre_preferences WHERE user_id LIKE %s AND genre_id LIKE %s""" % (user_id, genre_id))

    db.commit()

    return {"message": "Taste deleted from user"}

# USERS

@app.get("/users/{email}/bymail", status_code=200, tags=["Users"])
def get_user_infos_by_mail(email: str, response: Response) -> UserResponse:
    c.execute("""SELECT * FROM user WHERE mail LIKE '%s'""" % (email))

    res = c.fetchone()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    print(res)
    return res

@app.get("/users/{user_id}", status_code=200, tags=["Users"])
def get_user_infos(user_id: int) -> UserResponse:
    c.execute("""SELECT * FROM user WHERE user_id LIKE %s""" % (user_id))

    res = c.fetchone()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return res

@app.post("/users/{user_id}", status_code=200, tags=["Users"])
def add_user_infos(user_id: int, name: str) -> ApiResponse:
    sql = "UPDATE user SET name = %s WHERE user_id LIKE %s"
    values = (name, user_id)
    c.execute(sql, values)

    db.commit()

    return {"message": "User updated"}

@app.delete("/users/{user_id}", status_code=200, tags=["Users"])
def remove_user(user_id: int) -> ApiResponse:
    c.execute("""DELETE FROM user WHERE user_id LIKE %s""" % (user_id))

    db.commit()

    return {"message": "User deleted"}

@app.post("/users", status_code=200, tags=["Users"])
def create_user(mail: str, name: str) -> ApiResponse:
    sql = "INSERT INTO user (mail, name) VALUES (%s, %s)"
    values = (mail, name)

    c.execute(sql, values)

    db.commit()

    return {"message": "User created"}


# GROUPS

@app.post("/groups", status_code=200, tags=["Groups"])
def create_group(gp_name: str) -> ApiResponse:
    sql = "INSERT INTO user_group (group_name) VALUES (%s)"
    values = [gp_name]

    c.execute(sql, values)

    db.commit()

    return {"message": "Group created"}

@app.get("/groups/{group_id}", status_code=200, tags=["Groups"])
def get_group_infos(group_id: int) -> GroupResponse:
    c.execute("""SELECT * FROM user_group WHERE group_id LIKE %s""" % (group_id))

    res = c.fetchone()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return res


# GROUPS
# USERS

@app.get("/groups/{group_id}/users", status_code=200, tags=["Users from Groups"])
def get_group_users(group_id: int) -> List[int]:
    c.execute("""SELECT u.* FROM user u JOIN user_group_membership um ON u.user_id = um.user_id WHERE um.group_id LIKE %s""" % (group_id))

    res = c.fetchall()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return res

@app.post("/groups/{group_id}/users/{user_id}", status_code=200, tags=["Users from Groups"])
def add_user_to_group(group_id: int, user_id: int) -> ApiResponse:
    sql = "INSERT INTO user_group_membership (user_id, group_id) VALUES (%s, %s)"
    values = (user_id, group_id)

    c.execute(sql, values)

    db.commit()

    return {"message": "User added to group"}

@app.delete("/groups/{group_id}/users/{user_id}", status_code=200, tags=["Users from Groups"])
def remove_user_from_group(group_id: int, user_id: int) -> ApiResponse:
    c.execute("""DELETE FROM user_group_membership WHERE user_id LIKE %s AND group_id LIKE %s""" % (user_id, group_id))

    db.commit()

    return {"message": "User deleted from group"}


@app.get("/groups/{group_id}/recommendations", status_code=200, tags=["Recommendations"])
def get_group_recommendations(group_id: int) -> List[MovieReccomendationResponse]:
    c.execute("""SELECT (movie_id,accuracy) FROM movie_group_recommendation WHERE group_id LIKE %s""" % (group_id))

    res = c.fetchall()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return res

@app.post("/groups/{group_id}/recommendations", status_code=200, tags=["Recommendations"])
def add_group_recommendations(group_id: int, movie_id: int, accuracy: float) -> ApiResponse:
    sql = "INSERT INTO movie_group_recommendation (accuracy, group_id, movie_id) VALUES (%s, %s, %s, %s)"
    values = (accuracy, group_id, movie_id)

    c.execute(sql, values)

    db.commit()

    return {"message": "Recommendation added"}

@app.get("/groups/{group_id}/tastes", status_code=200, tags=["Group tastes"])
def get_group_tastes(group_id: int) -> List[GenreResponse]:
    c.execute("""SELECT genre_name FROM genre, group_genre_preferences WHERE group_genre_preferences.group_id LIKE %s AND genre.genre_id = group_genre_preferences.genre_id""" % (group_id))

    res = c.fetchall()
    if res is None:
        raise HTTPException(status_code=404, detail="Item not found")
    return res

@app.post("/groups/{group_id}/tastes/{genre_id}", status_code=200, tags=["Group tastes"])
def add_taste_for_group(group_id: int, genre_id: int) -> ApiResponse:
    sql = "INSERT INTO group_genre_preferences (group_id, genre_id) VALUES (%s, %s)"
    values = (group_id, genre_id)

    c.execute(sql, values)

    db.commit()

    return {"message": "Taste added to group"}

@app.delete("/groups/{group_id}/tastes/{genre_id}", status_code=200, tags=["Group tastes"])
def remove_taste_from_group(group_id: int, genre_id: int) -> ApiResponse:
    c.execute("""DELETE FROM group_genre_preferences WHERE group_id LIKE %s AND genre_id LIKE %s""" % (group_id, genre_id))

    db.commit()

    return {"message": "Taste deleted from group"}
