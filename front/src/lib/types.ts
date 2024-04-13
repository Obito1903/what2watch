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

export type Group = {
	id: number;
	name: string;
	members: string[];
};
