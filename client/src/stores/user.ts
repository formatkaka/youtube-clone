import { setTokenCookies, clearTokenCookies } from "../utils/cookie";
import { writable } from "svelte/store";
import { session } from '$app/stores'
// import { attemptLogin } from "../services/auth";



function User(){
    console.log("user ts : ", session)
    // const initialUserState = await attemptLogin()

    const { subscribe, update } = writable({
        isLoggedIn: false,
        user: {}
    })

    return {

        loginUser: function(user: any){
            setTokenCookies(user.token, user.refreshToken)
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