import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
var decoded = DecodedToken(token);
userStatus(token);
if(decoded.payload["usertype"]==="user" || decoded.payload["usertype"]==="admin")
{
    window.location="/profile.html";
}