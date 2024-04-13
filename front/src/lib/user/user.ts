
import * as api from '../api';
import { page } from '$app/stores';


export function createUserInDB() {
    // Create user in db if it doesn't exist
    api.getMe().then((user) => {
        console.log(user);
    });
}
