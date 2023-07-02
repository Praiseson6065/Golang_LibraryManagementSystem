import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
userStatus(token);
UserPage(token);
if(token === undefined){
    window.location=`/home.html`;
    if(DecodedToken(token).payload["usertype"]!="admin")
    {
    window.location="http://127.0.0.1:3000/profile.html";
    }
}
else{
    userStatus(token);
}