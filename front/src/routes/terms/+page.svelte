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
		Very serious terms and conditions.
	</Card.Content>
</Card.Root>
