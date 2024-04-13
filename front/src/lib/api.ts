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

export function put(url: string, data: any) {
  return fetch(url, {
    method: 'PUT',
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

export function getUsers() {
  return get(`${API_URL}/users`);
}

export function getUserByEmail(email: string) {
  return get(`${API_URL}/users/email/${email}`);

}

export function getUserByID(id: number) {
  return get(`${API_URL}/users/${id}`);
}
//GROUPS
export function getGroups() {
  return get(`${API_URL}/users/groups`);
}

export function getGroupName(id: number) {
  return get(`${API_URL}/groups/${id}`);
}

export function getGroupMembers(id: number) {
  return get(`${API_URL}/groups/${id}/users`);
}

export function createGroup(name: string) {
  return post(`${API_URL}/groups`, {name: name});
}

export function addUserToGroup(group_id: number, user_id: number) {
  return put(`${API_URL}/groups/${group_id}/users/${user_id}`, {});
}