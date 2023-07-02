import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
userStatus(token);
UserPage(token);
fetch("/admin/users")
    .then(response=>response.json())
    .then(data=>{
        
        var UserHolder=document.querySelector("tbody");
        for(let i=0;i<data.length;i++){
            var user = `
            <tr>
                <td>${data[i]["Id"]}</td>
                <td>${data[i]["UserId"]}</td>
                <td>${data[i]["Name"]}</td>
                <td>${data[i]["Email"]}</td>
            </tr>
            `;
            UserHolder.innerHTML+=user;

        }
    })