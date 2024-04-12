<script lang="ts">
	import { activePage } from '$lib/stores';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input';
	import * as Card from '$lib/components/ui/card/index';
	activePage.set('movies');

	import * as api from '$lib/api';
	import type { Movie } from '$lib/types';
	import { onMount } from 'svelte';

	let topMovies: Movie[] = [];
	let popularMovies: Movie[] = [];
	let searchQuery = '';
	let selectedMovies: Movie[] = []; 

	$: {
		api.getTopRatedMovies().then((movies) => {
			topMovies = movies;
			selectedMovies = movies;
		});

		api.getPopularMovies().then((movies) => {
			popularMovies = movies;
		});
	}

	function switchMovies(type: 'popular' | 'topRated') {
		if (type === 'popular') {
			selectedMovies = [];
			selectedMovies = popularMovies;
		} else {
			selectedMovies = [];
			selectedMovies = topMovies;
		}
	}

	$: filteredMovies = selectedMovies.filter((movie) =>
		movie.title.toLowerCase().includes(searchQuery.toLowerCase())
	);
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

		<!-- Switch to toggle between popular and top-rated movies -->
		<div class="mt-4 flex justify-center space-x-4">
			<Button on:click={() => switchMovies('topRated')}>
				<span>Top Rated</span>
			</Button>
			<Button on:click={() => switchMovies('popular')}>
				<span>Popular</span>
			</Button>
		</div>

		<div class="mt-4 grid grid-cols-4 gap-4">
			{#each filteredMovies as movie, i}
				<Card.Root class="flex flex-col items-center">
					<img src={movie.poster} alt={movie.title} />
					<Card.Title class="mt-2 text-lg font-medium">{movie.title}</Card.Title>
					<Card.Description class="mt-2 text-center text-sm">
						{movie.release_date}
					</Card.Description>
					<Button class="mt-4">View details</Button>
				</Card.Root>
			{/each}
		</div>
	</Card.Content>
</Card.Root>
