<script lang="ts">
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
	import { default as Header } from '$lib/nav/header.svelte';
	import { onMount } from 'svelte';
	import Keycloak from 'keycloak-js';
	import { kc } from '$lib/stores';

	async function fetchUsers() {
		const response = await fetch('/api/users', {
			headers: {
				accept: 'application/json',
				authorization: `Bearer ${keycloak.token}`
			}
		});

		return response.json();
	}

	let keycloak: Keycloak;

	
	onMount(async () => {
		keycloak = new Keycloak({
			url: 'http://auth.localhost/',
			realm: 'what2watch',
			clientId: 'what2watch'
		});

		try {
			const authenticated = await keycloak.init({
				onLoad: 'login-required',
			});
			console.log(`User is ${authenticated ? 'authenticated' : 'not authenticated'}`);

			if (authenticated) {
				keycloak.updateToken(30).then(() => {
					console.log('token is ' + keycloak.token);
				});
				kc.set(keycloak);
			}



		} catch (error) {
			console.error('Failed to initialize adapter:', error);
		}
	});
</script>

<Header />
<ModeWatcher />
<slot />
