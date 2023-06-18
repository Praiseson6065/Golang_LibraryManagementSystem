import {token,DecodedToken,userStatus} from './userstatus.js';
var decoded = DecodedToken(token);
userStatus(token);
if(decoded.payload["usertype"]==="admin"){
  window.location="/admin.html"
}
var url="/profile";
fetch(url)
    .then(response => response.text())
    .then(data=>{
      
      if(token!=undefined) {
        document.querySelector("title").innerText="Profile";
        document.getElementById("accname").innerText="Name : " + decoded.payload["name"];
        document.getElementById("accemail").innerText="Email : "+decoded.payload["email"];
      }
      else{
        document.querySelector("title").innerText="Not Logged In";
      }

    });

    