
import { SvelteKitAuth } from "@auth/sveltekit"
import Keycloak from "@auth/sveltekit/providers/keycloak"
 
export const { handle } = SvelteKitAuth({
  providers: [Keycloak],
}) 