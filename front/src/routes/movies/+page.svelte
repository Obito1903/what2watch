<script lang="ts">
	import { activePage } from '$lib/stores';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input';
	import * as Card from '$lib/components/ui/card/index';
	activePage.set('movies');
</script>


<Card.Root>
	<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
		<Card.Title class="text-5xl font-medium">Movies</Card.Title>
	</Card.Header>
	<Card.Content>
		<form class="flex w-full max-w-sm items-center space-x-2">
			<Input placeholder="Search for a movie" />
			<Button type="submit">Search</Button>
		</form>

		<div class="gap-4 mt-4">
			{#each Array.from({ length: 9 }) as _, i}
				<Card.Root class="flex flex-col items-center">
					<!-- <Card.Image src={`https://picsum.photos/200/300?random=${i}`} alt="Movie poster" /> -->
					<Card.Title class="text-lg font-medium mt-2">Movie {i + 1}</Card.Title>
					<Card.Description class="text-sm text-center mt-2">
						Description of the movie {i + 1}
					</Card.Description>
					<Button class="mt-4">View details</Button>
				</Card.Root>
			{/each}
		</div>
	</Card.Content>
</Card.Root>
