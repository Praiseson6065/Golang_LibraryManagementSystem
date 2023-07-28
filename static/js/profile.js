import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
UserPage(token);
var decoded = DecodedToken(token);
userStatus(token);
if(decoded.payload["usertype"]==="admin"){
  window.location="/admin.html"
}
var url="/user/profile";
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
fetch(`/user/approvbooks/${decoded.payload["ID"]}`)
    .then(response=> response.json())
    .then(data=> {
      if(data.length!=0)
      {
        for (let i=0;i<data.length;i++){
          var Approvedbook=`
          <div class="bookHolder">
        <div class="bookImg"><a href="/book.html?BookC=${data[i]["BookCode"]}"><img class="bookImgCP" src="/img/books/${data[i]['ImgPath']}"></a></div>
        <div class="bookDetails">
        <div class="bookName">${data[i]['BookName']}</div>
        <div class="bookAuthor">by ${(data[i]['Author']).replace(/[{}"]/g,'')}</div>
      
        <div><button class="returnbook" data-bookid=${data[i]["BookId"]}>Return Book</button></div>

        
    </div>`;
          document.getElementById("approvedbooks").innerHTML+=Approvedbook;
         
        }
  
        var ReturnBtns = document.querySelectorAll(".returnbook");
        
        for (let i=0;i<ReturnBtns.length;i++){
              ReturnBtns[i].addEventListener("click",function(event){
                  fetch(`/user/returnbook/${decoded.payload["ID"]}/${event.target.dataset.bookid}`,{method:"POST"})
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
    fetch(`/user/reqbooks/`,{method:"post",headers,body: JSON.stringify(json)})
    .then(response=> response.json())
    .then(data=>{
      if(data===true);
      alert("Book Requested");
    })
  }


})
fetch(`/user/userreqbook/${decoded.payload["ID"]}`)
  .then(response=>response.json())
  .then(data=>{
    var tbody="";
    for(let i=0;i<data.length;i++){
      var row =`
        <div class="rowReqBook">
          <div>${data[i].RequestedBooks}</div>
          <div>${data[i].ISBN}</div>
          <div>${data[i].Status}</div>

        </div>
        <div class="gradLine"></div>
        `;


        tbody+=row;
    
    }
    document.getElementById("bodyReqBook").innerHTML=tbody;
    
  })

fetch(`/user/issuedbooks/${decoded.payload["ID"]}`)
  .then(response=>response.json())
  .then(data=>{
    if(data.length!=0){
      for (let i=0;i<data.length;i++){
        var Issuedbook=`
        <div class="bookHolder">
        <div class="bookImg"><a href="/book.html?BookC=${data[i]["BookCode"]}"><img class="bookImgCP" src="/img/books/${data[i]['ImgPath']}"></a></div>
        <div class="bookDetails">
        <div class="bookName">${data[i]['BookName']}</div>
        <div class="bookAuthor">by ${(data[i]['Author']).replace(/[{}"]/g,'')}</div>
        <div class="bookCritics">
            <div class="bookDes"><div class="bookD">${data[i]['Pages']}</div><div  class="bookDName">Pages</div></div>
            <div class="bookDes"><div class="bookD">${data[i]['votes']}</div ><div class="bookDName">Likes</div></div>
        </div>
        <button data-bookid=${data[i]['BookId']} class="removeissuereqbook">Remove</button>
        
    </div>`;
        
        document.getElementById("issuedbooks").innerHTML+=Issuedbook;
       
      }

      const removeRequestBooks=document.querySelectorAll(".removeissuereqbook")
      removeRequestBooks.forEach( btn => { btn.addEventListener("click",function (eve){
          fetch(`/user/issuebook/${decoded.payload["ID"]}/${eve.target.dataset.bookid}`,{method:"Delete"})
            .then(response => response.json())
            .then(data=>{
              if(data===true)
              {
                alert("Book Removed");
                window.location.reload();
              }
            })
      })})

    }
    else{
      document.getElementById("issuedbooks").innerHTML="No Books Are Request To Issue"
    }

  });


    