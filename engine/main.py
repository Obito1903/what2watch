from typing import Union,List,Dict
from pydantic import BaseModel

from fastapi import FastAPI,BackgroundTasks
import requests
import os
import tmdbsimple as tmdb
from fastapi.middleware.cors import CORSMiddleware


DBApi = os.getenv("DB_API") or "http://db-api.localhost"
tmdb.API_KEY = os.getenv("TMDB_API_KEY")

app = FastAPI()

origins = [
    "http://what2watch.localhost",
    "http://localhost:5173",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class UserRequest(BaseModel):
    user_id: int

class GroupRequest(BaseModel):
    group_id: int

class Recommendation(BaseModel):
    movie_id: int
    score: float

@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/queue/add/user")
def queue_add_group(user_id: int, tasks: BackgroundTasks):
    tasks.add_task(create_user_recommendation, user_id)
    return {"group": user_id}

def create_user_recommendation(user_id: int):
    # Get list of movies liked by user
    print("Getting movies for user " + str(user_id))
    resp = requests.get(DBApi + "/users/" + str(user_id) + "/movies")
    if resp.status_code != 200:
        print("Error getting movies for user " + str(user_id))
        print(resp.status_code)
        print(resp.text)
        return []
    movies = resp.json()
    print(movies)
    recommendations: Dict[str,Recommendation] = {}
    for movie in movies:
        if movie['viewed'] == True and movie['rating'] > 3:
            # Get recommendations for movie
            recs = tmdb.Movies(movie['movie_id']).recommendations()
            rec = recs['results'][0]
            recommendations[rec['id']] = Recommendation(movie_id=rec['id'], score=(movie['rating']/10) * rec['popularity'])

    # Sort recommendations by score
    sorted_recommendations = sorted(recommendations.values(), key=lambda x: x.score, reverse=True)
    # Add recommendations to database
    for rec in sorted_recommendations:
        print(rec)
        requests.post(DBApi + "/users/" + str(user_id) + "/recommendations", json={"movie_id": rec.movie_id, "accuracy": rec.score})

@app.get("/queue/add/group")
def queue_add_group(group_id: int, tasks: BackgroundTasks):
    tasks.add_task(create_group_recommendation, group_id)
    return {"group": group_id}

def create_group_recommendation(group_id: int):
    # Get list of users in group
    print("Getting users for group " + str(group_id))
    resp = requests.get(DBApi + "/groups/" + str(group_id) + "/users")
    if resp.status_code != 200:
        print("Error getting users for group " + str(group_id))
        print(resp.status_code)
        print(resp.text)
        return []
    users = resp.json()
    print(users)
    recommendations: Dict[str,Recommendation] = {}
    for user in users:
        # Get recommendations for user
        recs = requests.get(DBApi + "/users/" + str(user) + "/recommendations")
        if recs.status_code != 200:
            print("Error getting recommendations for user " + str(user))
            print(recs.status_code)
            print(recs.text)
            continue
        for rec in recs.json():
            recommendations[rec['movie_id']] = Recommendation(movie_id=rec['movie_id'], score=rec['accuracy'])

    # Sort recommendations by score
    sorted_recommendations = sorted(recommendations.values(), key=lambda x: x.score, reverse=True)
    # Add recommendations to database
    for rec in sorted_recommendations:
        print(rec)
        requests.post(DBApi + "/groups/" + str(group_id) + "/recommendations", json={"movie_id": rec.movie_id, "accuracy": rec.score})
