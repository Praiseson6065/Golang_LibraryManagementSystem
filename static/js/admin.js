
import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
console.log(token)
UserPage(token);
userStatus(token);
if(token === undefined){
  window.location=`/home.html`;
  
}
else{
  console.log("HEllo")
  userStatus(token);
  if(DecodedToken(token).payload["usertype"]!="admin")
  {
  window.location="/profile.html";
  }
}
