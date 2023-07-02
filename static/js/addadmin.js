import {token,DecodedToken,userStatus} from "./userstatus.js";

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



var submit = document.getElementById("addadmin");
submit.addEventListener("click",function (){
    var json={
        "Name": document.getElementById("name").value,
        "Email":document.getElementById("email").value,
        "Password":document.getElementById("password").value,
    };
    const headers = {
        "Content-Type": "application/json",
      };
    fetch(`/admin/addadmin`,{method : 'post',headers,body: JSON.stringify(json)})
    .then(response=>response.json())
    .then(data=>{
        if(data === true){
            document.querySelector(".add-admin-wrap").innerHTML="<p>User Added</p>";
            setTimeout(function(){
                window.location.reload(true);
            },5000)
        }
        else{
            document.querySelector(".add-admin-wrap").innerHTML=`<p>${data}</p>`;
        }
    })

})
