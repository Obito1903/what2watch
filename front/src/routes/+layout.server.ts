let sessionToken : string | null = null;
let loggedIn = false;


export function load({cookies}) {
    const sessionCookie = cookies.get('authjs.session-token');

    return {
        sessionToken: sessionCookie
    }
}