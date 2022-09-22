import { setTokenCookies, clearTokenCookies } from "../utils/cookie";
import { writable } from "svelte/store";
import { browser } from '$app/environment'



function User(){

    const { subscribe, update } = writable({
        isLoggedIn: false,
        user: {}
    })

    return {

        loginUser: function(user: any){
            if(browser) setTokenCookies(user.token, user.refreshToken)
            update(() => (
                 {
                    isLoggedIn: true,
                    user
                }
            ))
        },
        logoutUser: function(){
            clearTokenCookies()
            update(() => (
                 {
                    isLoggedIn: false,
                    user: {}
                }
            ))
        },
        subscribe

    }
}

export const user = User()