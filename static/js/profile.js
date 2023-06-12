const token = document.cookie.split("=")[1];
if(token!=undefined){
  const header = token.split('.')[0];
  const payload = token.split('.')[1];
  const signature = token.split('.')[2];
  
 var decoded = {
    header: JSON.parse(atob(header)),
    payload: JSON.parse(atob(payload)),
    signature: signature
  };
  
}
url="http://127.0.0.1:3000/profile" ;
fetch(url)
    .then(response => response.text())
    .then(data=>{
      
      if(token!=undefined) {
        document.querySelector("title").innerText="Profile";
        document.getElementById("accname").innerText="Name : " + decoded.payload["name"];
        document.getElementById("accemail").innerText="Email : "+decoded.payload["email"];
        document.getElementById("accstatus").innerText="Logout";
        document.getElementById("accstatus").setAttribute("href","/logout");

      }
      else{
        document.querySelector("title").innerText="Not Logged In";
        document.getElementById("profile").innerText="Login"; 
        document.getElementById("profile").setAttribute("href","/login")
        document.getElementById("accstatus").innerText="Sign Up";
        document.getElementById("accstatus").setAttribute("href","/register");
      }

    });

    