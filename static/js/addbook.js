import {token,DecodedToken,userStatus} from "./userstatus.js";
userStatus(token);
if(DecodedToken(token).payload["usertype"]==="admin"){
    window.location="http://127.0.0.1:3000/admin";
}
