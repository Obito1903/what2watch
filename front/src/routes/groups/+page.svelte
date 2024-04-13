<script lang="ts">
	import { activePage } from '$lib/stores';
	import * as Card from '$lib/components/ui/card/index';
	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Command from '$lib/components/ui/command';
	import * as Popover from '$lib/components/ui/popover';
	import type { Group } from '$lib/types';
	import * as Avatar from '$lib/components/ui/avatar';
	import { onMount, tick } from 'svelte';
	import * as api from '$lib/api';
	import Input from '$lib/components/ui/input/input.svelte';

	activePage.set('groups');
	let myGroups: Group[] = [
		{ id: 1, name: 'les foufous de socheau', members: ['Titou', 'Paul'] },
		{ id: 2, name: 'Eistiens', members: ['Samuel', 'Quentin', 'Melody', 'Nathan'] }
	];

	let selectedGroupIdToAddMember: number = -1;
	let groupName = '';
	let emailInput = '';
	function leaveGroup(groupID: number) {
		myGroups = myGroups.filter((g) => g.id != groupID);
	}

	function deleteMember(member: string, groupId: number) {
		console.log(member);
		const groupIndex = myGroups.findIndex((group) => group.id === groupId);
		if (groupIndex !== -1) {
			const updatedMembers = myGroups[groupIndex].members.filter((m) => m !== member);
			myGroups = myGroups.map((group, index) =>
				index === groupIndex ? { ...group, members: updatedMembers } : group
			);
		}
	}

	function addMember(member: string, groupId: number) {
		console.log(member);
		const groupIndex = myGroups.findIndex((group) => group.id === groupId);
		if (groupIndex !== -1) {
			const updatedMembers = [...myGroups[groupIndex].members, member];
			myGroups = myGroups.map((group, index) =>
				index === groupIndex ? { ...group, members: updatedMembers } : group
			);
			selectedGroupIdToAddMember = -1;
		}
	}

	function addMemberByEmail(email: string) {
		api.getUserByEmail(email).then((user) => {
			if (user) {
				api.addUserToGroup(selectedGroupIdToAddMember, user.user_id).then(() => {
					addMember(user.name, selectedGroupIdToAddMember);
				});
			}
		});
	}

	function createGroup(name: string) {
		console.log(name);
		api.createGroup(name).then((group) => {
			let groupId = Number(group.message);
			myGroups = [...myGroups, { id: groupId, name: name, members: [] }];

			api.getMe().then((me) => {
				api.addUserToGroup(groupId, me.user_id).then(() => {
					myGroups = myGroups.map((group) =>
						group.id === groupId ? group = { ...group, members: [me.name] } : group
					);
					console.log('User added to group');
				});
			});
		});
	}
	onMount(() => {
		api.getGroups().then((groups) => {
	
			// for (let group of groups) {
			// 	let gpmembers: string[] = [];
			// 	api.getGroupMembers(group.group_id).then((members) => {
			// 		members.forEach(member => {
			// 			api.getUserByID(member.user_id).then((user) => {
			// 				gpmembers.push(user.name);
			// 			});
			// 		});
			// 		myGroups = [...myGroups, { id: group.group_id, name: group.name, members: gpmembers }];
			// 	});
			// }
		});
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
													<Button
														on:click={() => deleteMember(member, group.id)}
														variant="destructive"
													>
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
								<Popover.Content class="p-5 m-5" align="end">
									<Command.Root loop>
										<Command.Input bind:value={emailInput} placeholder="Enter the mail address of the user to add to the group..." />
										<Button on:click={() => {addMemberByEmail(emailInput)}}>Add to the group</Button>
									</Command.Root>
								</Popover.Content>
							</Popover.Root>

							<Button on:click={() => leaveGroup(group.id)} variant="destructive">Leave</Button>
						</Card.Content>
					</Card.Root>
				{/each}
				<Card.Root class="w-96">
					<Card.Header>
						<Card.Title>Create a new group</Card.Title>
					</Card.Header>
					<Card.Content>
						<Input placeholder="Group name..." class="w-64" bind:value={groupName} />
						<Button
							on:click={() => {
								createGroup(groupName);
							}}
							variant="default">Create Group</Button
						>
					</Card.Content>
				</Card.Root>
			</div>
		</ScrollArea>
	</Card.Content>
</Card.Root>
