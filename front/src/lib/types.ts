// id	238
// title	"The Godfather"
// poster	"https://image.tmdb.org/t/p/w300_and_h450_bestv2/3bhkrj58Vtu7enYsRolD1fZdja1.jpg"
// release_date	"1972-03-14"
// genres	[]


export type Movie = {
  id: number;
  title: string;
  poster: string;
  release_date: string;
  genres: string[];
}

export type Group = {
    id: number;
    name: string;
    members: string[];
};