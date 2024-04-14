import { kc } from '$lib/stores';
import { get as getStore } from 'svelte/store';
export const API_URL = 'http://api.localhost';


export function get(url: string) {
  return fetch(url, {
    headers: {
      'Authorization': `Bearer ${getStore(kc)?.token as string}`,
      'Content-Type': 'application/json',
    },
    credentials: 'same-origin',
  }).then((res) => res.json()).catch((err) => console.error(err));
}

export function post(url: string, data: any) {
  return fetch(url, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${getStore(kc)?.token as string}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
    credentials: 'same-origin',

  }).then((res) => res.json());
}

export function put(url: string, data: any) {
  return fetch(url, {
    method: 'PUT',
    headers: {
      'Authorization': `Bearer ${getStore(kc)?.token as string}`,
      'Content-Type': 'application/json'
    },
    credentials: 'same-origin',
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
export function refreshGroupRecommendations(group_id: number) {
  return get(`${API_URL}/groups/${group_id}/refresh_recommendations`);
}
export function getMyMovies() {
  return get(`${API_URL}/users/movies`);

}

export function deleteMovie(id: number) {
  return fetch(`${API_URL}/users/movies/${id}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${getStore(kc)?.token as string}`,
    },
    credentials: 'same-origin',
  }).then((res) => res.json());

}

export function getMovieGenres() {
  return get(`${API_URL}/tmdb/genres`);
}

export function searchMovies(query: string) {
  return fetch(`${API_URL}/tmdb/search?query=${query}`).then((res) => res.json());
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
  return get(`${API_URL}/users/id/${id}`);
}

export function getUserTastes() {
  return get(`${API_URL}/users/tastes`);
}
export function addTasteToUser(genreID: number) {
  return put(`${API_URL}/users/tastes/${genreID}`, {});
}

export function deleteTasteFromUser(genreID: number) {
  return fetch(`${API_URL}/users/tastes/${genreID}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${getStore(kc)?.token as string}`,
    },
    credentials: 'same-origin',
  }).then((res) => res.json());
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
  return post(`${API_URL}/groups`, { name: name });
}

export function addUserToGroup(group_id: number, user_id: number) {
  return put(`${API_URL}/groups/${group_id}/users/${user_id}`, {});
}

export function deleteUserFromGroup(group_id: number, user_id: number) {
  return fetch(`${API_URL}/groups/${group_id}/users/${user_id}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${getStore(kc)?.token as string}`,
    },
    credentials: 'same-origin',
  }).then((res) => res.json());
}

export function getGroupRecommendations(group_id: number) {
  return get(`${API_URL}/groups/${group_id}/recommendations`);
}
