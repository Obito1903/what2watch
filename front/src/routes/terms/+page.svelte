<script lang="ts">
	import { activePage } from '$lib/stores';
	import * as Card from '$lib/components/ui/card/index';

	activePage.set('terms');
</script>

<Card.Root>
	<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
		<Card.Title class="text-5xl font-medium">Terms and conditions</Card.Title>
	</Card.Header>
	<Card.Content>
		Lorem ipsum dolor sit amet consectetur adipisicing elit. Rem similique sint explicabo eius illum nulla quasi et mollitia quod natus impedit odit dignissimos hic pariatur quas animi, dolorum doloribus commodi?
	</Card.Content>
</Card.Root>
