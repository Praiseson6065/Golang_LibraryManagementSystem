import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
userStatus(token);
UserPage(token);
if(token === undefined){
    window.location=`/home.html`;
    if(DecodedToken(token).payload["usertype"]!="admin")
    {
    window.location="/profile.html";
    }
}
else{
    userStatus(token);
}
fetch(`/admin/approvalbookslist`)
    .then(response=>response.json())
    .then(data=>{
         console.log(data)
        var ans=0;
        for (let i=0;i<data.length;i++){

            if(data[i].IssuedBooks.length===0) continue;
            else{ans=1}
            var IssuBooks=``;
            for(let j=0;j<data[i].IssuedBooks.length;j++){
                let IssuBook =`<div class="ApprovBook"><div class="t-align-center">${data[i].IssuedBooks[j].BookId}</div><div class="t-align-center apprvbtndiv">${data[i].IssuedBooks[j].BookName}</div><div class="t-align-center"><input class="apprvChk" data-user="${data[i].User.Id}" data-book="${data[i].IssuedBooks[j].BookId}" type="checkbox"></div></div>`
                IssuBooks+=IssuBook;
            }
            var ApprovRow =`
            <div class="ApprovRow ">
                        <div class="ApprovUserId t-align-center ">${data[i].User.Id}</div>
                        <div class="ApprovUserName t-align-center ">${data[i].User.Name}</div>
                        <div class="ApprovBooks">
                            ${IssuBooks}
                        </div>
                    </div>
            `;
            document.getElementById("ApprovBody").innerHTML+=ApprovRow;
            
        }

        var approvalChk = document.querySelectorAll(".apprvChk");
        const ApprovalUserBooks=[];
        class ApprovalUserData{
            constructor(UserId,BookId=[]){
                this.UserId = UserId
                this.BookId=BookId
            }
            AddBook(BookId){
                this.BookId.push(BookId);
            }
            RemoveBook(BookId){
                let index = this.BookId.indexOf(BookId);
                if (index > -1) {
                    this.BookId.splice(index, 1); 
                }
            }
        }
        function AddBookWithUser(UsId,Bkid)
        {
            if(ApprovalUserBooks.length===0){
                let arr=[];
                arr.push(Bkid)
                let UsrData = new ApprovalUserData(UsId,arr)
                ApprovalUserBooks.push(UsrData);
            }
            else{
                for(let UsrD of ApprovalUserBooks){
                    if(UsrD.UserId===UsId){
                        UsrD.AddBook(Bkid);
                        return;
                    }
                }
                let arr=[];
                arr.push(Bkid)
                let UsrData = new ApprovalUserData(UsId,arr)
                ApprovalUserBooks.push(UsrData);
            }
        }
        function RemoveBookWithUser(UsId,Bkid){
            for(let UsrD of ApprovalUserBooks){
                if(UsrD.UserId===UsId){
                    UsrD.RemoveBook(Bkid);
                    if(UsrD.BookId.length===0){
                        let index = ApprovalUserBooks.indexOf(UsrD);
                        ApprovalUserBooks.splice(index, 1); 
                    }
                    return;
                }
            }
        }
        
        for(let k =0;k<approvalChk.length;k++){
            approvalChk[k].addEventListener("change",function(event){
                if(event.target.checked===true){
                    AddBookWithUser(parseInt(event.target.dataset.user,10),parseInt(event.target.dataset.book,10))
                }
                else{
                    RemoveBookWithUser(parseInt(event.target.dataset.user,10),parseInt(event.target.dataset.book,10))
                }
               
                
            })}
            if(ans===1  ){
                document.getElementById("ApprvBtnDiv").innerHTML=`<button id="ApprovBtn">Approve</button>`;
                var ApprovalBtn = document.getElementById("ApprovBtn")
                ApprovalBtn.addEventListener("click",function(){
                    console.log(ApprovalUserBooks)
                    if(ApprovalUserBooks.length!=0){
                        fetch(`/admin/approvbooks`,{method:"Post",headers: {"Content-Type": "application/json",},body: JSON.stringify(ApprovalUserBooks)})
                        .then(response=>response.json())
                        .then(data=>{
                            if(data===true){
                                window.location.reload();
                            }
                            else{
                                alert(data);
                            }
                        })

                    }
                    else{
                        alert("Select atleast One book to Approve")
                    }
                    
                })

            }
        
        
    });
    