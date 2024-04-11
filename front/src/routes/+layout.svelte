<script lang="ts">
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
	import { default as Header } from '$lib/nav/header.svelte';

</script>

<Header sessionToken={undefined}/>
<ModeWatcher />
<slot />
