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

document.querySelector("form").addEventListener("submit", function(event){
    event.preventDefault();
    fetch("/api/book", {
        method: "POST",
        body: new FormData(event.target)})
    .then(response => response.json())
    .then(data=>{
        if(data===true){
            window.location="/home.html";
        }
        else{
            alert(data);
        }
    })

});

