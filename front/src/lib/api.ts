export const API_URL = 'http://api.localhost';

export function get(url: string) {
	return fetch(url).then((res) => res.json());
}

export function post(url: string, data: any) {
  return fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data),
  }).then((res) => res.json());
}

// MOVIES
export function getTopRatedMovies() {
	return get(`${API_URL}/tmdb/movies/toprated`);
}

export function getPopularMovies() {
	return get(`${API_URL}/tmdb/movies/popular`);
}
export function getMovieDetails(id: number) {
	return get(`${API_URL}/tmdb/movies/${id}/details`);
}

export function addToMyMovies(id: number, rating: number) {
	return post(`${API_URL}/users/movies/${id}`, 
  {
		rating: rating,
		viewed: true
	});
}

export function getMyMovies() {
  return get(`${API_URL}/users/movies`);

}

export function deleteMovie(id: number) {
  return fetch(`${API_URL}/users/movies/${id}`, {
    method: 'DELETE',
  }).then((res) => res.json());

}
// USERS
export function getMe() {
	return get(`${API_URL}/users/me`);
}
