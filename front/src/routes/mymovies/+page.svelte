<script lang="ts">
	import { activePage } from '$lib/stores';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input';
	import * as Card from '$lib/components/ui/card/index';
	import * as Popover from '$lib/components/ui/popover';
	activePage.set('mymovies');

	import * as api from '$lib/api';
	import type { Movie, Movie_Details, Movie_Review } from '$lib/types';
	import { onMount } from 'svelte';
	import { addToMyMovies } from '$lib/api';
	import { Badge } from '$lib/components/ui/badge';
	import { Label } from '$lib/components/ui/label';
	import * as RadioGroup from '$lib/components/ui/radio-group';
	$: isPopoverVisible = false;

	let review: string;

	let myReviews: Array<Movie_Review> = [];
	let myMovies: Movie_Review[] = [];
	let searchQuery = '';
	let selectedMovies: Movie_Review[] = [];

	type Movie_Review = {
		movie_id: number;
		movie_title: string;
		rating: number;
	};


	onMount(async () => {
		const movies = await api.getMyMovies();
		myMovies = await Promise.all(
			movies.map(async (movie: any) => {
				const movieDetails = await api.getMovieDetails(movie.movie_id);
				return {
					movie_id: movie.movie_id,
					movie_title: movieDetails.title,
					rating: movie.rating
				};
			})
		);
	});

	$: filteredMovies = myMovies.filter((movie) =>
		movie.movie_title.toLowerCase().includes(searchQuery.toLowerCase())
	);

	function deleteMovie(id: number) {
		api.deleteMovie(id).then(() => {
			myMovies = myMovies.filter((m) => m.movie_id !== id);
			selectedMovies = selectedMovies.filter((m) => m.movie_id !== id);
		});
	}

	function togglePopover() {
		isPopoverVisible = !isPopoverVisible;
	}
</script>

<Card.Root>
	<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
		<Card.Title class="text-5xl font-medium">Movies</Card.Title>
	</Card.Header>
	<Card.Content>
		<form class="flex w-full max-w-sm items-center space-x-2">
			<Input placeholder="Search for a movie" bind:value={searchQuery} />
			<Button type="submit">Search</Button>
		</form>

		<div class="mt-4 grid grid-cols-4 gap-4">
			{#each filteredMovies as movie, i}
				<Card.Root class="flex flex-col items-center">
					<Card.Title class="mt-2 text-lg font-medium">{movie.movie_title}</Card.Title>
					<Card.Description class="mt-2 text-center text-sm"></Card.Description>
					<Card.Content>
						My rating : {movie.rating}
					</Card.Content>

					<Button
						variant="destructive"
						on:click={() => {
							deleteMovie(movie.movie_id);
						}}>Delete</Button
					>
				</Card.Root>
			{/each}
		</div>
	</Card.Content>
</Card.Root>
