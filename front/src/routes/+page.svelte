<script lang="ts">
	import { activePage } from '$lib/stores';
	import * as Card from '$lib/components/ui/card/index';
	import { signIn } from "@auth/sveltekit/client"
	activePage.set('');
	export let data;
</script>

<Card.Root>
	<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
		<Card.Title class="text-5xl font-medium">Welcome to what2watch</Card.Title>
	</Card.Header>
	<Card.Content>
		Lorem ipsum dolor sit amet consectetur adipisicing elit. Rem similique sint explicabo eius illum nulla quasi et mollitia quod natus impedit odit dignissimos hic pariatur quas animi, dolorum doloribus commodi?
		{data.sessionToken}
	</Card.Content>
</Card.Root>
