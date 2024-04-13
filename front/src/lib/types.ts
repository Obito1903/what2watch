export type Movie = {
	id: number;
	title: string;
	poster: string;
	release_date: string;
	genres: string[];
};

export type Movie_Details = {
	id: number;
	title: string;
	genres: Array<{ id: number; name: string }>;
	overview: string;
	popularity: number;
	release_date: string;
  poster: string; // This field is not present in the API response, adding it manually
};

export type Movie_Review = {
  movie_id: number;
  rating: number;
  review_id: number;
  viewed: boolean;
};

export type Group = {
	id: number;
	name: string;
	members: string[];
};
