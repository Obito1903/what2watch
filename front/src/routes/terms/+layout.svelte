<script>
import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	// Correctly dispatch the event with the event name as the first argument
	dispatch('updatePage', {page: 'terms'});
</script>

<slot />