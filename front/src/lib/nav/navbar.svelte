<script lang="ts">
	import { cn } from '$lib/utils.js';

	let className: string | undefined | null = undefined;
	export { className as class };

	import { activePage } from '$lib/stores';
	import { page } from '$app/stores';

	$: loggedIn = $page.data.session?.user;
</script>

<nav class={cn('flex items-center space-x-4 lg:space-x-6', className)}>
	<h1 class="">What2Watch</h1>
	<a
		href="/"
		class="hover:text-primary text-sm font-medium transition-colors {$activePage === ''
			? 'text-primary'
			: 'text-muted-foreground'}"
	>
		Overview
	</a>

	<a
		href="/movies"
		class="hover:text-primary text-sm font-medium transition-colors {$activePage === 'movies'
			? 'text-primary'
			: 'text-muted-foreground'}"
	>
		Movies
	</a>

	<a
		href="/mymovies"
		class="hover:text-primary text-sm font-medium transition-colors {$activePage === 'mymovies'
			? 'text-primary'
			: 'text-muted-foreground'}"
	>
		My Movies
	</a>

	<a
		href="/recommendations"
		class="hover:text-primary text-sm font-medium transition-colors {$activePage ===
		'recommendations'
			? 'text-primary'
			: 'text-muted-foreground'}"
		class:disabled={!loggedIn}
		on:click={(e) => {
			if (!loggedIn) {
				e.preventDefault();
			}
		}}
	>
		Recommendations
	</a>

	<a
		href="/groups"
		class="hover:text-primary text-sm font-medium transition-colors {$activePage === 'groups'
			? 'text-primary'
			: 'text-muted-foreground'}"
		class:disabled={!loggedIn}
		on:click={(e) => {
			if (!loggedIn) {
				e.preventDefault();
			}
		}}
	>
		Groups
	</a>

	<a
		href="/terms"
		class="hover:text-primary text-sm font-medium transition-colors {$activePage === 'terms'
			? 'text-primary'
			: 'text-muted-foreground'}"
	>
		Terms & Conditions
	</a>
</nav>

<style>
	.disabled {
		pointer-events: none;
		opacity: 0.5;
	}
</style>
