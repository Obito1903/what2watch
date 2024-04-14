<script lang="ts">
	import { activePage } from '$lib/stores';
	import * as Card from '$lib/components/ui/card/index';
	import { signIn } from "@auth/sveltekit/client"
	import { page } from "$app/stores"
	activePage.set('');
</script>

<Card.Root>
	<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
		<Card.Title class="text-5xl font-medium">Welcome to what2watch</Card.Title>
	</Card.Header>
	<Card.Content>
		In this web app you can create groups with your friends and get movie recommendations based on your tastes.
	</Card.Content>
</Card.Root>
