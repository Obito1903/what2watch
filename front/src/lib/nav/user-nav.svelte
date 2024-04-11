<script lang="ts">
	import * as DropdownMenu from '$lib/registry/new-york/ui/dropdown-menu/index.js';
	import * as Avatar from '$lib/registry/new-york/ui/avatar/index.js';
	import { Button } from '$lib/registry/new-york/ui/button/index.js';

	import { signIn, signOut } from '@auth/sveltekit/client';
	import { page } from '$app/stores';
	import { SignOut } from '@auth/sveltekit/components';
</script>
{#if !$page.data.session?.user}
<Button on:click={() => signIn()}>Sign in</Button>
{:else}
<DropdownMenu.Root>
	<DropdownMenu.Trigger asChild let:builder>
		<Button variant="ghost" builders={[builder]} class="relative h-8 w-8 rounded-full">
			<Avatar.Root class="h-8 w-8">
				<Avatar.Image src="/avatars/01.png" alt="@shadcn" />
				<Avatar.Fallback>SC</Avatar.Fallback>
			</Avatar.Root>
		</Button>
	</DropdownMenu.Trigger>
	<DropdownMenu.Content class="w-56" align="end">
		<DropdownMenu.Label class="font-normal">
			<div class="flex flex-col space-y-1">
				<p class="text-sm font-medium leading-none">{$page.data.session?.user?.name}</p>
				<p class="text-muted-foreground text-xs leading-none">{$page.data.session?.user?.email}</p>
			</div>
		</DropdownMenu.Label>
		<DropdownMenu.Separator />
		<DropdownMenu.Group>
			<DropdownMenu.Item>
				Profile
				<DropdownMenu.Shortcut>⇧⌘P</DropdownMenu.Shortcut>
			</DropdownMenu.Item>
			<DropdownMenu.Item href="/preferences">
				Preferences
				<DropdownMenu.Shortcut>⌘S</DropdownMenu.Shortcut>
			</DropdownMenu.Item>
		</DropdownMenu.Group>
		<DropdownMenu.Separator />
		<DropdownMenu.Item on:click={() => signOut()}>
			Sign out
			<DropdownMenu.Shortcut>⇧⌘Q</DropdownMenu.Shortcut>
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>

{/if}
