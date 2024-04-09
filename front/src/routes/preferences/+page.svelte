<script lang="ts">
	import { activePage } from '$lib/stores';
	import * as Card from '$lib/components/ui/card/index';
	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { default as Search } from '$lib/nav/search.svelte';
	import * as Command from '$lib/components/ui/command';
	import * as Popover from '$lib/components/ui/popover';

	activePage.set('preferences');

	const genres = [
		'Action',
		'Comedy',
		'Drama',
		'Thriller',
		'Sci-Fi',
		'Fantasy',
		'Horror',
		'Romance',
		'Mystery',
		'Crime',
		'Adventure',
		'Animation',
		'Family',
		'War',
		'History',
		'Documentary',
		'Music',
		'Western'
	];

	let myGenres = ['Documentary'];

	let addGenrePopoverOpen = false;

	function submit(genre: string) {
		if (myGenres.includes(genre)) {
			return;
		}

		myGenres = [...myGenres, genre];
		addGenrePopoverOpen = false;
	}

	function remove(genre: string) {
		myGenres = myGenres.filter((g) => g !== genre);
	}
</script>

<Card.Root>
	<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
		<Card.Title class="text-5xl font-medium">Genres</Card.Title>
	</Card.Header>
	<Card.Content>
		<h1 class="my-2 py-2 text-4xl">My genres</h1>
		<Popover.Root bind:open={addGenrePopoverOpen}>
			<Popover.Trigger asChild let:builder>
				<Button builders={[builder]} variant="default" class="ml-auto py-2">Add a genre</Button>
			</Popover.Trigger>
			<Popover.Content class="p-0" align="end">
				<Command.Root loop>
					<Command.Input placeholder="Select genre..." />
					<Command.List>
						{#each genres as genre}
							<Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
								<Button class="m-0 w-full p-0" variant="ghost" on:click={() => submit(genre)}
									>{genre}</Button
								>
							</Command.Item>
						{/each}
					</Command.List>
				</Command.Root>
			</Popover.Content>
		</Popover.Root>

		<ScrollArea orientation="horizontal" class="w-auto whitespace-nowrap rounded-md border">
			<div class="flex flex-row gap-1">
				{#if myGenres.length === 0}
					<p class="text-muted-foreground text-2xl">No genres selected</p>
				{/if}
				{#each myGenres as genre}
					<Card.Root class="w-96">
						<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
							<Card.Title class="text-2xl text-sm font-medium">{genre}</Card.Title>
						</Card.Header>
						<Card.Content>
							<Button on:click={() => remove(genre)} variant="destructive">Remove</Button>
						</Card.Content>
					</Card.Root>
				{/each}
			</div>
		</ScrollArea>
	</Card.Content>
</Card.Root>
