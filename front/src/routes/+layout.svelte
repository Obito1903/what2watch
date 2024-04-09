<script lang="ts">
	import '../app.pcss';
	import { ModeWatcher } from "mode-watcher";
	import { default as Header } from '$lib/nav/header.svelte';
	import { writable } from 'svelte/store';

</script>

<Header/>
<ModeWatcher />
<slot />
