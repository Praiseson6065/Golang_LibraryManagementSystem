import {token,DecodedToken,userStatus} from './userstatus.js';
var decoded = DecodedToken(token);
if(decoded.payload["usertype"]==="admin"){
  window.location="http://127.0.0.1:3000/admin.html"
}
var url="http://127.0.0.1:3000/profile";
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

    