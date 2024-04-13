<script lang="ts">
	import { activePage } from '$lib/stores';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input';
	import * as Card from '$lib/components/ui/card/index';
	import * as Popover from '$lib/components/ui/popover';
	activePage.set('mymovies');

	import * as api from '$lib/api';
	import type { Movie, Movie_Details } from '$lib/types';
	import { onMount } from 'svelte';
	import { addToMyMovies } from '$lib/api';
	import { Badge } from '$lib/components/ui/badge';
	import { Label } from '$lib/components/ui/label';
	import * as RadioGroup from '$lib/components/ui/radio-group';
	$: isPopoverVisible = false;

	let review: string;

	let myMovies: Movie[] = [];
	let searchQuery = '';
	let selectedMovies: Movie[] = [];
	let selectedMovie: Movie_Details = {
		id: 0,
		title: 'title',
		genres: [],
		popularity: 0,
		overview: '',
		release_date: '',
		poster: '' // This field is not present in the API response, adding it manually
	};
	$: {
		onMount(() => {
			api.getMyMovies().then((movies) => {
				myMovies = movies;
				selectedMovies = movies;
			});
		});
	}

	$: filteredMovies = selectedMovies.filter((movie) =>
		movie.title.toLowerCase().includes(searchQuery.toLowerCase())
	);

	function addMovie(movie: Movie_Details, note: number) {
		addToMyMovies(movie.id, note).then(() => {
			console.log('Movie added to my movies');
		});
	}

	function togglePopover() {
		isPopoverVisible = !isPopoverVisible;
	}

	function showMovieDetails(movie: Movie) {
		api.getMovieDetails(movie.id).then((movieDetails: Movie_Details) => {
			console.log(movieDetails);
			selectedMovie = movieDetails;
			movieDetails.poster = movie.poster;
			togglePopover();
		});
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
					<img src={movie.poster} alt={movie.title} />
					<Card.Title class="mt-2 text-lg font-medium">{movie.title}</Card.Title>
					<Card.Description class="mt-2 text-center text-sm">
						{movie.release_date}
					</Card.Description>

					<Button
						on:click={() => {
							showMovieDetails(movie);
						}}>View details</Button
					>
				</Card.Root>
			{/each}
		</div>
	</Card.Content>
</Card.Root>

<div class="PopoverAnchor">
	<Popover.Root bind:open={isPopoverVisible}>
		<Popover.Trigger />
		<Popover.Content class="w-100">
			{#if selectedMovie}
				<div class="flex p-4">
					<img src={selectedMovie.poster} alt={selectedMovie.title} class="mr-4" />
					<div>
						<h2 class="text-2xl font-bold">{selectedMovie.title}</h2>
						<p class="mt-2">
							Release Date: <span class="text-muted-foreground"> {selectedMovie.release_date}</span>
						</p>
						<p class="mt-2">
							Popularity: <span class="text-muted-foreground">{selectedMovie.popularity}</span>
						</p>
						<p class="mt-2">
							Overview: <span class="text-muted-foreground">{selectedMovie.overview}</span>
						</p>
						<p class="mt-2">
							Genres:
							{#each selectedMovie.genres as genre}
								<Badge>{genre.name}</Badge>
							{/each}
						</p>
						<p>Review:</p>
						<div class="flex items-center space-x-2">
							<RadioGroup.Root bind:value={review}>
								<div class="flex items-center space-x-2">
									<RadioGroup.Item value="1" id="option-one" />
									<Label for="option-one">1</Label>
								</div>
								<div class="flex items-center space-x-2">
									<RadioGroup.Item value="2" id="option-two" />
									<Label for="option-two">2</Label>
								</div>
								<div class="flex items-center space-x-2">
									<RadioGroup.Item value="3" id="option-three" />
									<Label for="option-three">3</Label>
								</div>
								<div class="flex items-center space-x-2">
									<RadioGroup.Item value="4" id="option-four" />
									<Label for="option-four">4</Label>
								</div>
								<div class="flex items-center space-x-2">
									<RadioGroup.Item value="5" id="option-five" />
									<Label for="option-five">5</Label>
								</div>
							</RadioGroup.Root>
						</div>

						<Button
							class="mt-4"
							on:click={() => {
								addMovie(selectedMovie, Number(review));
							}}>Add to my movies</Button
						>
					</div>
				</div>
			{/if}
		</Popover.Content>
	</Popover.Root>
</div>

<style>
	.PopoverAnchor {
		position: fixed;
		top: 10%;
		left: 10%;
		width: 100%;
		height: 100%;
		z-index: 100;
		pointer-events: none;
	}
</style>
