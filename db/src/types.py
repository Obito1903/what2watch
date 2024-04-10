from pydantic import BaseModel

class Group(BaseModel):
    group_id: int
    group_name: str

class ApiResponse(BaseModel):
    message: str

class MovieReccomendationResponse(BaseModel):
    movie_id: str
    accuracy: float

class GenreResponse(BaseModel):
    genre_id: int
    name: str

class UserResponse(BaseModel):
    user_id: int
    name: str
    mail: str

class GroupResponse(BaseModel):
    group_id: int
    name: str
