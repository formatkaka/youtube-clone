// import  from '@sveltejs/kit'

function setTokenCookies(token: string, refreshToken: string){

    document.cookie = `token=${token}; SameSite=None; Secure`;
    document.cookie = `refreshToken=${refreshToken}; SameSite=None; Secure`;
    //
}

function getTokenCookies(){
    let token;
    let refreshToken;
    // if(browser){
    //     token = getCookie("token")
    //     refreshToken = getCookie("refreshToken")
    // }else{
    //     token = page.
    // }

    return {token, refreshToken}
}

function clearTokenCookies(){
    //
}

function getCookie(name: string) {
    const nameEQ = name + "=";
    const ca = document.cookie.split(';');
    for(let i=0;i < ca.length;i++) {
        let c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1,c.length);
        if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
    }
    return null;
}

export {
    setTokenCookies,
    getTokenCookies,
    clearTokenCookies
}