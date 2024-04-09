<script lang="ts">
	import { activePage } from '$lib/stores';
	import * as Card from '$lib/components/ui/card/index';
	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Command from '$lib/components/ui/command';
	import * as Popover from '$lib/components/ui/popover';
	import type { Group } from '$lib/global/group';
	import * as Avatar from '$lib/components/ui/avatar';
	import { tick } from 'svelte';

	activePage.set('groups');
	let users = ['Samuel', 'Quentin', 'Melody', 'Nathan', 'Titou', 'Paul'];
	let myGroups: Group[] = [
		{ id: 1, name: 'les foufous de socheau', members: ['Titou', 'Paul'] },
		{ id: 2, name: 'Eistiens', members: ['Samuel', 'Quentin', 'Melody', 'Nathan'] }
	];

	let selectedGroup: Group | null = null;
	let selectedGroupIdToAddMember: number = -1;

	function leaveGroup(groupID: number) {
		myGroups = myGroups.filter((g) => g.id != groupID);
	}

	function deleteMember(member: string, groupId: number) {
		console.log(member);
		// Find the index of the group from which you're removing the member
		const groupIndex = myGroups.findIndex((group) => group.id === groupId);
		if (groupIndex !== -1) {
			// Create a new array without the member
			const updatedMembers = myGroups[groupIndex].members.filter((m) => m !== member);
			// Update the group with the new members array
			myGroups = myGroups.map((group, index) =>
				index === groupIndex ? { ...group, members: updatedMembers } : group
			);
		}
	}

	function addMember(member: string, groupId: number) {
		console.log(member);
		// Find the index of the group to which you're adding the member
		const groupIndex = myGroups.findIndex((group) => group.id === groupId);
		if (groupIndex !== -1) {
			// Create a new array with the updated members
			const updatedMembers = [...myGroups[groupIndex].members, member];
			// Update the group with the new members array
			myGroups = myGroups.map((group, index) =>
				index === groupIndex ? { ...group, members: updatedMembers } : group
			);
			selectedGroupIdToAddMember = -1;
		}
	}
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
													<Button on:click={() => deleteMember(member, group.id)} variant="destructive">
														KICK</Button
													>
												</div>
											{/each}
										</div>
									</ScrollArea>
								</Card.Content>
							</Card.Root>

							<Popover.Root open={selectedGroupIdToAddMember === group.id}>
								<Popover.Trigger asChild let:builder>
									<Button
										builders={[builder]}
										variant="default"
										on:click={async () => {
											selectedGroupIdToAddMember = group.id;
											await tick();
										}}>Add someone to this group</Button
									>
								</Popover.Trigger>
								<Popover.Content class="p-0" align="end">
									<Command.Root loop>
										<Command.Input placeholder="Select user to add to the group..." />
										<Command.List>
											{#each users as user}
												<Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
													<Button
														on:click={(event) => {
															addMember(user, group.id);
														}}
														class="m-0 w-full p-0"
														variant="ghost">{user}</Button
													>
												</Command.Item>
											{/each}
										</Command.List>
									</Command.Root>
								</Popover.Content>
							</Popover.Root>

							<Button on:click={() => leaveGroup(group.id)} variant="destructive">Leave</Button>
						</Card.Content>
					</Card.Root>
				{/each}
			</div>
		</ScrollArea>
	</Card.Content>
</Card.Root>
