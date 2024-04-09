<script lang="ts">
	import { Input } from "$lib/registry/new-york/ui/input/index.js";

    export let placeholder = 'Search...';
</script>

<div>
	<Input type="search" placeholder={placeholder} class="h-9 md:w-[100px] lg:w-[300px]" />
</div>