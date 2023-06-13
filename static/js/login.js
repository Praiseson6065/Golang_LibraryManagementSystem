import {token,DecodedToken} from "./userstatus.js";
var decoded = DecodedToken(token);
if(decoded.payload["usertype"]==="user" || decoded.payload["usertype"]==="admin")
{
    window.location="http://127.0.0.1:3000/profile.html";
}