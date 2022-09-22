import { getTokenCookies } from "../utils/cookie";
import { BASE_URL } from "../utils/consts"

async function SignIn(email: string, password: string){

    // return new Promise(function(resolve, reject){
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

        const rawData = JSON.stringify({
        email,
        password
        });

        const requestOptions: RequestInit = {
        method: 'POST',
        headers: myHeaders,
        body: rawData,
        redirect: 'follow'
        };

        // fetch(`${BASE_URL}/login`, requestOptions)
        // .then(response => {
        //     if(!response.ok) return Promise.reject(response.json())
        //     return response.json()
        // })
        // .then(result => {
        //     resolve(result)
        // })
        // .catch(error => {
        //     console.log("🚀 ~ file: auth.ts ~ line 33 ~ returnnewPromise ~ error", error)
        //     reject(error)
        // });

        const response = await fetch(`${BASE_URL}/login`, requestOptions);
        
        if(!response.ok){
            const errRes = await response.json()
            console.log("🚀 ~ file: auth.ts ~ line 38 ~ //returnnewPromise ~ errRes", errRes)
            throw new Error(errRes.error)
        }

        const data = response.json()
        return data;

    // })
        
}

function SignUp(email: string, password: string, name: string){

    // return new Promise(function(resolve, reject){
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

        const rawData = JSON.stringify({
        name, 
        email,
        password
        });

        const requestOptions: RequestInit = {
        method: 'POST',
        headers: myHeaders,
        body: rawData,
        redirect: 'follow'
        };

        return fetch(`${BASE_URL}/signup`, requestOptions)
        .then(response => response.text())
        // .then(result => console.log(result))
        // .catch(error => console.log('error', error));
    // })
}

async function getUserDetails(userToken: string){
    const {token = userToken}: any = getTokenCookies()

    if(!token) return null

    const myHeaders = new Headers();
    myHeaders.append("token", token);


    const requestOptions: any = {
        method: 'GET',
        headers: myHeaders,
        redirect: 'follow'
    };

    const response = await fetch(`${BASE_URL}/user`, requestOptions)
    
    if(!response.ok) return null

    const data = response.json()
    return data;

}

export {
    SignIn,
    SignUp,
    getUserDetails
}