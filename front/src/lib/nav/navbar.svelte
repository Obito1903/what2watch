<script lang="ts">
	import { cn } from "$lib/utils.js";

	let className: string | undefined | null = undefined;
	export { className as class };

	import { activePage } from "$lib/stores";

	
</script>

<nav class={cn("flex items-center space-x-4 lg:space-x-6", className)}>
	<a href="/" class="text-sm font-medium transition-colors hover:text-primary {$activePage === '' ? 'text-primary' : 'text-muted-foreground'}">
		Overview
	</a>

	<a
		href="/preferences"
		class="text-sm font-medium transition-colors hover:text-primary {$activePage === 'preferences' ? 'text-primary' : 'text-muted-foreground'}"
	>
		Preferences
	</a>
	<a
		href="/groups"
		class="text-sm font-medium transition-colors hover:text-primary {$activePage === 'groups' ? 'text-primary' : 'text-muted-foreground'}"
	>
		Groups
	</a>
	<a
		href="/recommendations"
		class="text-sm font-medium transition-colors hover:text-primary {$activePage === 'recommendations' ? 'text-primary' : 'text-muted-foreground'}"
	>
		Recommendations
	</a>

	<a
		href="/settings"
		class="text-sm font-medium transition-colors hover:text-primary {$activePage === 'settings' ? 'text-primary' : 'text-muted-foreground'}"
	>
		Settings
	</a>

	<a
		href="/terms"
		class="text-sm font-medium transition-colors hover:text-primary {$activePage === 'terms' ? 'text-primary' : 'text-muted-foreground'}"
	>
		Terms & Conditions
	</a>
</nav>