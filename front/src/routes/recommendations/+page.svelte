<script lang="ts">
	import { activePage } from '$lib/stores';
	import * as Card from '$lib/components/ui/card/index';
	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import type { Group } from '$lib/types';
	import * as Avatar from '$lib/components/ui/avatar';
	import { onMount } from 'svelte';
	import * as api from '$lib/api';

	activePage.set('recommendations');
	let myGroups: Group[] = [];

	function getRecommendations(groupId: number) {

		console.log('Getting recommendations');

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
