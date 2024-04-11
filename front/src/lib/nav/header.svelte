<script lang="ts">
	import { default as Nav } from '$lib/nav/navbar.svelte';
    import { default as UserNav } from '$lib/nav/user-nav.svelte';

	export let sessionToken : string | undefined;
</script>


<div class="hidden flex-col md:flex">
	<div class="border-b">
		<div class="flex h-16 items-center px-4">
			<Nav class="mx-6"/>
            <div class="ml-auto flex items-center space-x-4">
				<!-- <Search /> -->
				<UserNav sessionToken={sessionToken}/>
			</div>
		</div>
	</div>
</div>
