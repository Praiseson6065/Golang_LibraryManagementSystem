import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
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
fetch(`/admin/reqbooks`)
    .then(response=>response.json())
    .then(data=>{
        var books="";
        for(let i=0;i<data.length;i++)
        {
            var row=`
                <div class="RowReqBook">
                    <div>${data[i].UserId}</div>
                    <div>${data[i].RequestedBooks}</div>
                    <div>${data[i].ISBN}</div>
                    <div>${data[i].Status}</div>
                </div>
            `;
            books+=row;


        }
        document.getElementById("BodyReqBooks").innerHTML=books;
    })