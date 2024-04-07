<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { default as Nav } from '$lib/nav/header.svelte';
    import { default as UserNav } from '$lib/nav/user-nav.svelte';
</script>



<div class="hidden flex-col md:flex">
	<div class="border-b">
		<div class="flex h-16 items-center px-4">
			<Nav class="mx-6" />
            <div class="ml-auto flex items-center space-x-4">
				<!-- <Search /> -->
				<UserNav />
			</div>
		</div>
	</div>
</div>
