import { writable } from 'svelte/store';
import Keycloak from 'keycloak-js';

export const activePage = writable('');

export const kc = writable<Keycloak | undefined>(undefined);
