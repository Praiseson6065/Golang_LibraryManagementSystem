import {token,DecodedToken,userStatus} from './userstatus.js'
userStatus(token);
if(DecodedToken(token).payload["usertype"]!=="admin")
{
  window.location="http://127.0.0.1:3000/admin";
}
fetch("http://127.0.0.1:3000/admin")
    .then(response => response.json())
    .then(data => {
        if(data['error']==="UnAuthorized"){
            document.getElementById("main-wrap").innerText=data['error'];
        }
        else{
             
        }
    });