
export function get(url: string) {
  return fetch(url).then(res => res.json())
}


export function getTopRatedMovies() {
  return get('http://api.localhost/tmdb/movies/toprated');
}

export function getPopularMovies() {
  return get('http://api.localhost/tmdb/movies/popular');
}
export function getMovieDetails(id: number) {
  return get(`http://api.localhost/tmdb/movies/${id}`);
}