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

@app.get("/users/{id}/movies", status_code=200)
def get_seen_movies(id: int):
    
    return list_of_seen_movies

@app.post("/users/{id}/movies", status_code=200)
def add_movie_to_watched(id: int):
    
    return title+" is now watched"

@app.delete("/users/{id}/movies", status_code=200)
def remove_movie_from_watched(id: int):
    
    return title+" is now unwatched"


# USERS
# RECOMMENDATIONS

@app.get("/users/{id}/recommendations", status_code=200)
def get_recommendations(id: int):
    
    return list_of_recommendations

@app.post("/users/{id}/recommendations", status_code=200)
def add_recommendations(id: int):
    
    return "Recommendations added"


# USERS
# TASTES

@app.get("/users/{id}/tastes", status_code=200)
def get_tastes(id: int):
    
    return list_of_tastes


# USERS

@app.get("/users/{id}", status_code=200)
def get_user_infos(id: int):
    c.execute("""SELECT * FROM user WHERE user_id LIKE %s""" % (id))
    
    return c.fetchone()

@app.post("/users/{id}", status_code=200)
def add_user_infos(id: int, name: str, age: int):
    sql = "UPDATE user SET name = %s, age = %s WHERE user_id LIKE %s"    
    values = (name, age, id)
    c.execute(sql, values)
    
    db.commit()
    
    return "User : "+name+" got updated"

@app.delete("/users/{id}", status_code=200)
def remove_user(id: int):
    c.execute("""DELETE FROM user WHERE user_id LIKE %s""" % (id))
    
    db.commit()
    
    return "User got deleted"

@app.post("/users", status_code=200)
def create_user(name: str, age: int):
    sql = "INSERT INTO user (name, age) VALUES (%s, %s)"    
    values = (name, 24)
    
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

@app.get("/groups/{id}", status_code=200)
def get_group_infos(id: int):
    c.execute("""SELECT * FROM user_group WHERE group_id LIKE %s""" % (id))
    
    return c.fetchone()


# GROUPS
# USERS

@app.get("/groups/{id}/users", status_code=200)
def get_group_users(id: int):
    c.execute("""SELECT name FROM user, user_group_membership WHERE user_group_membership.group_id LIKE %s AND user.user_id = user_group_membership.user_id""" % (id))
    
    return c.fetchall()

@app.post("/groups/{id}/users", status_code=200)
def add_user_to_group(id: int, user_id: int):
    sql = "INSERT INTO user_group_membership (user_id, group_id) VALUES (%s, %s)"    
    values = (user_id, id)
    
    c.execute(sql, values)
    
    db.commit()
    
    return "User added to group"

@app.delete("/groups/{id}/users", status_code=200)
def remove_user_from_group(id: int):
    
    return username+" is now deleted from "+groupname


# GROUPS
# TASTES

@app.get("/groups/{id}/tastes", status_code=200)
def get_group_tastes(id: int):
    
    return list_of_group_tastes