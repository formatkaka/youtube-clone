function validateEmail (email: string) {
    return String(email)
      .toLowerCase()
      .match(
        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
      );
  };

  function validatePassword(password: string){
    return password.length > 8
  }
  function validateName(name: string){
    return name.length > 2
  }

export  {
  validateName,
  validateEmail,
  validatePassword
}