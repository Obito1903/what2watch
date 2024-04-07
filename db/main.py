from typing import Union
from fastapi import FastAPI
import MySQLdb

db_config = {
    'host': '127.0.0.1',
    'port': 3306,
    'user': 'moviefinder',
    'passwd': 'moviefinder',
    'db': 'moviefinder',
}

db = MySQLdb.connect(**db_config)
c=db.cursor()

app = FastAPI()

# USERS
# MOVIES

@app.get("/users/{user_id}/movies", status_code=200)
def get_seen_movies(user_id: int):
    
    return list_of_seen_movies

@app.post("/users/{user_id}/movies", status_code=200)
def add_movie_to_watched(user_id: int):
    
    return title+" is now watched"

@app.delete("/users/{user_id}/movies", status_code=200)
def remove_movie_from_watched(user_id: int):
    
    return title+" is now unwatched"


# USERS
# RECOMMENDATIONS

@app.get("/users/{user_id}/recommendations", status_code=200)
def get_recommendations(user_id: int):
    
    return list_of_recommendations

@app.post("/users/{user_id}/recommendations", status_code=200)
def add_recommendations(user_id: int):
    
    return "Recommendations added"


# USERS
# TASTES

@app.get("/users/{user_id}/tastes", status_code=200)
def get_tastes(user_id: int):
    
    return list_of_tastes


# USERS

@app.get("/users/{user_id}", status_code=200)
def get_user_infos(user_id: int):
    c.execute("""SELECT * FROM user WHERE user_id LIKE %s""" % (user_id))
    
    return c.fetchone()

@app.post("/users/{user_id}", status_code=200)
def add_user_infos(user_id: int, name: str, age: int):
    sql = "UPDATE user SET name = %s, age = %s WHERE user_id LIKE %s"    
    values = (name, age, user_id)
    c.execute(sql, values)
    
    db.commit()
    
    return "User : "+name+" got updated"

@app.delete("/users/{user_id}", status_code=200)
def remove_user(user_id: int):
    c.execute("""DELETE FROM user WHERE user_id LIKE %s""" % (user_id))
    
    db.commit()
    
    return "User got deleted"

@app.post("/users", status_code=200)
def create_user(name: str, age: int):
    sql = "INSERT INTO user (name, age) VALUES (%s, %s)"    
    values = (name, age)
    
    c.execute(sql, values)
    
    db.commit()
    
    return "User : "+name+" is now created"


# GROUPS

@app.post("/groups", status_code=200)
def create_group(gp_name: str):
    sql = "INSERT INTO user_group (group_name) VALUES (%s)"    
    values = [gp_name]
    
    c.execute(sql, values)
    
    db.commit()
    
    return "Group : "+gp_name+" is now created"

@app.get("/groups/{group_id}", status_code=200)
def get_group_infos(group_id: int):
    c.execute("""SELECT * FROM user_group WHERE group_id LIKE %s""" % (group_id))
    
    return c.fetchone()


# GROUPS
# USERS

@app.get("/groups/{group_id}/users", status_code=200)
def get_group_users(group_id: int):
    c.execute("""SELECT name FROM user, user_group_membership WHERE user_group_membership.group_id LIKE %s AND user.user_id = user_group_membership.user_id""" % (group_id))
    
    return c.fetchall()

@app.post("/groups/{group_id}/users/{user_id}", status_code=200)
def add_user_to_group(group_id: int, user_id: int):
    sql = "INSERT INTO user_group_membership (user_id, group_id) VALUES (%s, %s)"    
    values = (user_id, group_id)
    
    c.execute(sql, values)
    
    db.commit()
    
    return "User added to group"

@app.delete("/groups/{group_id}/users/{user_id}", status_code=200)
def remove_user_from_group(group_id: int, user_id: int):
    c.execute("""DELETE FROM user_group_membership WHERE user_id LIKE %s AND group_id LIKE %s""" % (user_id, group_id))
    
    db.commit()
    
    return "User got deleted from its group"


# GROUPS
# TASTES

@app.get("/groups/{id}/tastes", status_code=200)
def get_group_tastes(id: int):
    
    return list_of_group_tastes