<script lang="ts">
	import { activePage } from '$lib/stores';
	import * as Card from '$lib/components/ui/card/index';
	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import type { Group, Movie_Details } from '$lib/types';
	import * as Avatar from '$lib/components/ui/avatar';
	import { onMount } from 'svelte';
	import * as api from '$lib/api';
	import { Badge } from '$lib/components/ui/badge';
	import * as Popover from '$lib/components/ui/popover';

	activePage.set('recommendations');
	let myGroups: Group[] = [];

	type Recommendation = {
		movie_id: number;
		accuracy: number;
	};

	let recommendations: Map<number, Movie_Details[]> = new Map();
	let isPopoverVisible = false;
	let selectedGroupId = -1;
	async function getRecommendations(groupId: number) {
		const recs = await api.getGroupRecommendations(groupId);
		console.log(recs);
		const moviePromises = recs.map((reco) => api.getMovieDetails(reco.movie_id));
		const movies = await Promise.all(moviePromises);

		recommendations.set(groupId, movies);
		selectedGroupId = groupId;
		isPopoverVisible = true;
	}

	onMount(async () => {
		const groups = await api.getGroups();
		console.log(groups);

		const updatedGroups = await Promise.all(
			groups.map(async (group) => {
				const members = await api.getGroupMembers(group.group_id);
				// console.log(members);

				const gpmembers = await Promise.all(
					members.map(async (member) => {
						const user = await api.getUserByID(member);
						// console.log("got user : " + JSON.stringify(user));
						return user.name;
					})
				);

				return { id: group.group_id, name: group.group_name, members: gpmembers };
			})
		);

		myGroups = [...myGroups, ...updatedGroups];
	});
</script>

<Card.Root>
	<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
		<Card.Title class="text-5xl font-medium">My Groups</Card.Title>
	</Card.Header>
	<Card.Content>
		<ScrollArea orientation="horizontal" class="w-auto whitespace-nowrap rounded-md border">
			<div class="flex flex-row gap-1">
				{#if myGroups.length === 0}
					<p class="text-muted-foreground text-2xl">You don't have any groups yet.</p>
				{/if}
				{#each myGroups as group}
					<Card.Root class="w-96">
						<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
							<Card.Title>{group.name}</Card.Title>
						</Card.Header>
						<Card.Content class="grid gap-3">
							<Card.Root>
								<Card.Content class="grid gap-6 p-6">
									<ScrollArea class="h-40">
										<div class="flex flex-col items-center justify-between space-y-4">
											{#each group.members as member}
												<div class="flex items-center space-x-4">
													<Avatar.Root>
														<Avatar.Image src="/avatars/01.png" alt="Sofia Davis" />
														<Avatar.Fallback>SD</Avatar.Fallback>
													</Avatar.Root>
													<div>
														<p class="text-sm font-medium leading-none">{member}</p>
														<p class="text-muted-foreground text-sm">{member}</p>
													</div>
												</div>
											{/each}
										</div>
									</ScrollArea>
								</Card.Content>
							</Card.Root>

							
							<Button on:click={() => {
								api.refreshGroupRecommendations(group.id).then(() => {
									console.log('refreshed recommendations');
								});
							}}>
								Refresh recommendations
							</Button>

							<Button on:click={() => getRecommendations(group.id)} variant="default"
								>Get recommendations</Button
							>
						</Card.Content>
					</Card.Root>
				{/each}
			</div>
		</ScrollArea>

		
	</Card.Content>
</Card.Root>

<div class="PopoverAnchor">
	<Popover.Root bind:open={isPopoverVisible}>
		<Popover.Trigger />
		<Popover.Content class="w-100">
			{#if recommendations.get(selectedGroupId) != undefined}
				<p>No recommendations yet, press refresh to get recommendations</p>
			{/if}
			{#each recommendations.get(selectedGroupId) ?? [] as movie}
					<div class="flex p-4">
						<img src={movie.poster} alt={movie.title} class="mr-4" />
						<div>
							<h2 class="text-2xl font-bold">{movie.title}</h2>
							<p class="mt-2">
								Release Date: <span class="text-muted-foreground">
									{movie.release_date}</span
								>
							</p>
							<p class="mt-2">
								Popularity: <span class="text-muted-foreground">{movie.popularity}</span>
							</p>
							<p class="mt-2">
								Overview: <span class="text-muted-foreground">{movie.overview}</span>
							</p>
							<p class="mt-2">
								Genres:
								{#each movie.genres as genre}
									<Badge>{genre.name}</Badge>
								{/each}
							</p>
						</div>
					</div>
			{/each}
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
