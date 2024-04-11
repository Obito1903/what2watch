<script lang="ts">
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
	import { default as Header } from '$lib/nav/header.svelte';

	export let data;
</script>

<Header sessionToken={data.sessionToken}/>
<ModeWatcher />
<slot />
