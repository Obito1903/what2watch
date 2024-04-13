from pydantic import BaseModel

class Group(BaseModel):
    group_id: int
    group_name: str

class ApiResponse(BaseModel):
    message: str

class MovieReviewsResponseEntry(BaseModel):
    review_id: int
    movie_id: int
    rating: int
    viewed: bool


class MovieReccomendationResponse(BaseModel):
    movie_id: int
    accuracy: float

class GenreResponse(BaseModel):
    genre_id: int
    name: str

# Users

class UserResponse(BaseModel):
    user_id: int
    name: str
    mail: str

class UserPostMovieRequest(BaseModel):
    rating: int
    viewed: bool

class UserPostRecommendationRequest(BaseModel):
    movie_id: int
    accuracy: float


class UserPostUpdateRequest(BaseModel):
    name: str

class UserPostCreateRequest(BaseModel):
    name: str
    mail: str

# Groups

class GroupRequest(BaseModel):
    name: str

class GroupResponse(BaseModel):
    group_id: int
    name: str

class GroupPostRecommendationRequest(BaseModel):
    movie_id: int
    accuracy: float
