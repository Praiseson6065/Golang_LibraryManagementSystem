import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
UserPage(token);
var decoded = DecodedToken(token);
userStatus(token);
if(decoded.payload["usertype"]==="admin"){
  window.location="/admin.html"
}
var url="/profile";
fetch(url)
    .then(response => response.text())
    .then(data=>{
      
      if(token!=undefined) {
        document.querySelector("title").innerText="Profile";
        document.getElementById("accname").innerText="Name : " + decoded.payload["name"];
        document.getElementById("accemail").innerText="Email : "+decoded.payload["email"];
      }
      else{
        document.querySelector("title").innerText="Not Logged In";
      }

    });
fetch(`/api/issuedbooks/${decoded.payload["ID"]}`)
    .then(response=> response.json())
    .then(data=> {
      if(data!=[])
      {
        for (let i=0;i<data.length;i++){
          var Issuedbook=`
          <div class="bookholder">
              <div><img class="bookImg" src="/img/books/${data[i]['ImgPath']}" alt="${data[i]["BookName"]}"></div>
              <div class="bookDetails">
              <div class="bookDesc">
                <div>Title : ${data[i]["BookName"]}</div>
                <div>Pages : ${data[i]["Pages"]}</div>
                <div>ISBN : ${data[i]["ISBN"]}</div>
                <div>Author : ${data[i]["Author"]}</div>
                <div>Publisher : ${data[i]["Publisher"]}</div>
              </div>
              <button class="returnbook" data-bookid=${data[i]["BookId"]}>Return Book</button> 
              </div>
              
          </div>
          `
          document.getElementById("issuedbooks").innerHTML+=Issuedbook;
         
        }
  
        var ReturnBtns = document.querySelectorAll(".returnbook");
        
        for (let i=0;i<ReturnBtns.length;i++){
              ReturnBtns[i].addEventListener("click",function(event){
                  fetch(`/api/returnbook/${decoded.payload["ID"]}/${event.target.dataset.bookid}`,{method:"POST"})
                    .then(response=> response.json())
                    .then(data=> {
                        if(data===true)
                        {
                          alert("Book is Returned");
                          window.location.reload();
                        }
                    });
                }
              )
        }
      }
      
    });

var RequestBook = document.getElementById("RSubmit");
RequestBook.addEventListener("click",function (){
  const json={
    "UserId":decoded.payload["ID"],
    "RequestedBooks": document.getElementById("RbookName").value,
    "ISBN":document.getElementById("RISBN").value,
    "Status":false,

};
const headers = {
  "Content-Type": "application/json",
};
  if(json.RequestedBooks==="" && json.ISBN===""){
    alert("Empty Details of Book");
  }
  else{
    fetch(`/api/reqbooks/`,{method:"post",headers,body: JSON.stringify(json)})
    .then(response=> response.json())
    .then(data=>{
      if(data===true);
      alert("Book Requested");
    })
  }


})
fetch(`/api/userreqbook/${decoded.payload["ID"]}`)
  .then(response=>response.json())
  .then(data=>{
    var tbody="";
    for(let i=0;i<data.length;i++){
      var row =`
        <div class="rowReqBook">
          <div>${data[i].RequestedBooks}</div>
          <div>${data[i].ISBN}</div>
          <div>${data[i].Status}</div>
        </div>`;
        tbody+=row;
    
    }
    document.getElementById("bodyReqBook").innerHTML=tbody;
    
  })


    