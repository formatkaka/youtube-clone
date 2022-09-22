function secondsToHms(seconds: number) {
    seconds = Number(seconds);
    const h = Math.floor(seconds % (3600*24) / 3600);
    const m = Math.floor(seconds % 3600 / 60);
    const s = Math.floor(seconds % 60);
    let result="";
    result += h > 0 ? h+":"  : "";
    result += m > 0 ? convertToTimeNotation(m)+":"  : "00:";
    result += convertToTimeNotation(s)
    

    return  result;
    }


function convertToTimeNotation(num: number){
    return num < 9 ? "0"+num : num;
}

export { secondsToHms}