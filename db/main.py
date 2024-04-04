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
    
    c=db.cursor()
    # max_price=5
    # c.execute("""SELECT spam, eggs, sausage FROM breakfast
    #         WHERE price < %s""", (max_price,))
    
    c.execute("""SELECT name, age FROM user""")
    
    return c.fetchone()

@app.post("/users/{id}", status_code=200)
def add_user_infos(id: int):
    
    return username+" got it's infos updated"

@app.delete("/users/{id}", status_code=200)
def remove_movie_from_watched(id: int):
    
    return username+" is now deleted"

@app.post("/users", status_code=200)
def create_user(id: int):
    
    return username+" is now created"


# GROUPS

@app.post("/groups", status_code=200)
def create_group(id: int):
    
    return groupname+" is now created"

@app.get("/groups/{id}", status_code=200)
def get_group_infos(id: int):
    
    return group_infos


# GROUPS
# USERS

@app.get("/groups/{id}/users", status_code=200)
def get_group_users(id: int):
    
    return list_of_group_users

@app.post("/groups/{id}/users", status_code=200)
def add_user_to_group(id: int):
    
    return username+" is now added to "+groupname

@app.delete("/groups/{id}/users", status_code=200)
def remove_user_from_group(id: int):
    
    return username+" is now deleted from "+groupname


# GROUPS
# TASTES

@app.get("/groups/{id}/tastes", status_code=200)
def get_group_tastes(id: int):
    
    return list_of_group_tastes