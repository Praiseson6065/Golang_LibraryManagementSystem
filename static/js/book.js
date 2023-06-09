import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
userStatus(token);
var decoded=DecodedToken(token);
const queryString = window.location.search;
var url  = new URLSearchParams(queryString);
var bookId = url.get('BookC');
function suggestedBooks(taglines){
    var tag = taglines[0];
    
    var json={
        "SearchValue": tag,
        "SearchColumn": "taglines",
    };
    const headers = {
        "Content-Type": "application/json",
      };
    const SearchData = fetch("/api/searchbook", {method: "POST",headers,body: JSON.stringify(json),})
      .then((response) => response.json())
      .then((books) => { 
        return books;
      });
    const PrintBook = async () => {
        const data = await SearchData;
        const noOfBooks =data['Books'].length<=3 ? data['Books'].length : 4;
        for(let i=0;i<noOfBooks;i++)
        {   if(data['Books'][i]["BookCode"]===bookId){
            continue;
            }
            
            var suggested=`<div class="bookHolder-Suggested">
            <div class="bookImg"><a href="/book.html?BookC=${data['Books'][i]["BookCode"]}"><img class="bookImgCPS" src="/img/books/${data['Books'][i]['ImgPath']}"></a></div>
            <div class="bookDetails">
            <div class="bookName">${data['Books'][i]['BookName']}</div>
            <div class="bookAuthor">by ${(data['Books'][i]['Author']).replace(/[{}"]/g,'')}</div>
            <div class="bookCritics">
                <div class="bookDes"><div class="bookD">${data['Books'][i]['Pages']}</div><div  class="bookDName">Pages</div></div>
                <div class="bookDes"><div class="bookD">${data['Books'][i]['votes']}</div ><div class="bookDName">Likes</div></div>
            </div>
            
        </div>`;
            
            document.getElementById("book-suggestions").innerHTML =document.getElementById("book-suggestions").innerHTML+suggested;
        }
        
    }
    PrintBook();
    
}

fetch(`/api/bookc/${bookId}`)
    .then(response => response.json())
    .then(data => {
        var bId=data['BookId'];
        document.querySelector("title").innerText=data['BookName'];
        var book =`
        <div class="BookImg"><img  class="bookCoverPage" src="/img/books/${data['ImgPath']}" alt="${data['BookName']}"></div>
        <div class="BookDetails">
        <div>${data['BookName']}</div>
        <div>${data['Pages']} pages,Paperback</div>
        <div>by ${data['Author']}</div>
        <div>Published by ${data['Publisher']}</div>
        <div>Available  ${data['Quantity']} Books</div>
        <div>ISBN : ${data['ISBN']}</div>
    
        <div>Tags : ${data['Taglines']}</div>
        </div>
        <div class="BookCheckout" id="Bookchk">
        
            <div class="likediv"><button id="likebtn"><i class='bx bx-like'id="like"></i></button><span id="likecnt">${data['votes']}</span></div>
            <button id="addtocart" >Add to Cart</button>
        </div>`;
        
        document.getElementById("book-wrap").innerHTML=book;
        if(token!=undefined && decoded.payload["usertype"]==="admin"){
            document.getElementById("Bookchk").innerHTML+=`<a href="/updatebook.html?BId=${data['BookId']}"><button id="editbook" >Edit Book</button></a>`;
        }
        if(data['Quantity']===0){
            document.getElementById("addtocart").innerText="Out of Stock";
            document.getElementById("addtocart").classList.add("bookoutofstock");
            document.getElementById("addtocart").disabled=true;

        }
        function cartstatus(){
            fetch(`/user/cartbkchk/${decoded.payload["ID"]}/${bId}`)
                .then(response=>response.json())
                .then(data=>{
                    if(data===true)
                    {
                        document.getElementById("addtocart").innerText="AddedToCart";
                        document.getElementById("addtocart").classList.add("bookoutofstock");
                        document.getElementById("addtocart").disabled=true;
                        
                    }
                });
        }
        function bookIssueChk(){
            if(token!=undefined){
                fetch(`/user/isbookissued/${decoded.payload["ID"]}/${bId}`)
                    .then(response=>response.json())
                    .then(data=>{
                        if(data===true)
                        {
                            document.getElementById("addtocart").innerText="Already Book Issued   ";
                            document.getElementById("addtocart").classList.add("bookoutofstock");
                            document.getElementById("addtocart").disabled=true;
                        }
                    });
                
            }
        }
        function bookApproved(){
            if(token!=undefined){
                fetch(`/user/approvbooks/${decoded.payload["ID"]}`)
                    .then(response=>response.json())
                    .then(data=>{
                        for(let i=0;i<data.length;i++){
                            if (data[i].BookId===bId){

                                document.getElementById("addtocart").innerText="Already Book Taken";
                                document.getElementById("addtocart").classList.add("bookoutofstock");
                                document.getElementById("addtocart").disabled=true;
                                return;
                            }
                            continue;
                        }
                    })
            }
        }
        var tags = data['Taglines'].split(",");
        var likebtn=document.getElementById("likebtn");
        function isliked(){
            fetch(`/user/isliked/${decoded.payload["ID"]}/${data['BookId']}`)
                    .then(response=> response.json())
                    .then(data=>{
                        if(data===true)
                        {
                            document.getElementById("like").classList.add("bxs-like");
                            document.getElementById("like").classList.remove("bx-like");
                        }
                        else{
                            document.getElementById("like").classList.remove("bxs-like");
                            document.getElementById("like").classList.add("bx-like");
                        }

                    });
        }
        isliked();
        likebtn.addEventListener("click",function(event){
            fetch(`/user/like/${decoded.payload["ID"]}/${data['BookId']}`,{method:"POST"})
                .then(response=> response.json())
                .then(data=> {
                        console.log(data);
                        isliked();
                        fetch(`/api/bookc/${bookId}`)
                            .then(response => response.json())
                            .then(data=>{
                               document.getElementById("likecnt").innerText=`${data['votes']}`;
                            })
                });
        })
        suggestedBooks(tags);
        var AddtoCart=document.getElementById("addtocart");
        if(AddtoCart!=null){
            AddtoCart.addEventListener("click",function(){
                if(token===undefined){
                    alert("Please Login");
                }
                else{
                    
                    fetch(`/user/cart/${decoded.payload["ID"]}/${bId}`,{method:`Post`})
                        .then(response=>response.json())
                        .then(data=>{
                                document.getElementById("cartstatus").innerText=data["msg"];
                                cartstatus();

                        });
                    
                }
                    
            });
        }

        cartstatus();   
        bookIssueChk();
        bookApproved();
        async function BookReviewed(){
            
            var flag=0;
            await fetch(`/api/reviews/${bId}`)
            .then(response=>response.json())
            .then(data=>{
                var ans =0;
                document.getElementById("BookReviews").innerHTML="";
                for(let i=0;i<data.length;i++){

                    ans = data[i].UserId === decoded.payload["ID"] ? 1 :0
                    if(ans===1){
                        flag =1;
                    }
                    var review=`<div class="BookReviewBody">
                    <div>Reviewed by <span class="UserName">${data[i].UserName}</span></div>
                    <div>Review :    <div ${ans===1 ? `data-user=${data[i].UserId}` :""} class="UserReview" ${ans===1 ? `id="UserRevwed"`:""} >${data[i].Review}</div></div>
                    ${ans===1 ? `<div id="BtnDiv"><button id="EditRvw">Edit Review</button><button id="deletervw">Delete Review</button></div>`:`` }
                    </div>`
                    document.getElementById("BookReviews").innerHTML+=review;
                }
                console.log("BookReviewed")
                if(flag===1){
                    var EditRvw = document.getElementById("EditRvw")
                    var uRvw= document.getElementById("UserRevwed")
                    var deletervw = document.getElementById("deletervw")
                    deletervw.addEventListener("click",function(){
                        fetch(`/user/delbookreview/${decoded.payload["ID"]}/${bId}`,{method:"delete"})
                        .then(response=>response.json())
                        .then(data=>{
                            if(data===true)
                            {
                                BookReviewed();
                            }
                        })
                        

                    });
                    EditRvw.addEventListener("click",function(){
                           
                        var preRvw = document.getElementById("UserRevwed").innerText;
                        uRvw.contentEditable=true;
                        uRvw.classList.add("reviewEntryOutline");
                        document.getElementById("BtnDiv").innerHTML+=`<button id="UpdateRvw">Update Review</button>`
                        var updateRvw= document.getElementById("UpdateRvw")
                        function windowclick(event) {
                            if(event.target!=uRvw && event.target!=updateRvw && event.target !=EditRvw && event.target!=deletervw &&event.target!=UsRv)
                            {
                                console.log(event);
                                uRvw.contentEditable=false;
                                uRvw.classList.remove("reviewEntryOutline");
                                updateRvw.remove();
                                BookReviewed(); 
                            }
    
                        }
                        
                        // function hasEventListeners() {
                        //     const listeners = window.eventListeners;
                        //     return listeners !== undefined && listeners.length > 0;
                        //   }
                          
                        
                        window.onclick = windowclick
                        
                        updateRvw.addEventListener("click",async function(){
                            if(preRvw===uRvw.innerText){
                                alert("Review Not Changed")
                            }
                            else{
                                const headers = {
                                    "Content-Type": "application/json",
                                  };
                                var json={
                                    "BookId" : bId,
                                    "Review" :uRvw.innerText,
                                }
                                await fetch(`/user/updatereview/${decoded.payload["ID"]}`,{method:"put",headers,body:JSON.stringify(json)})
                                    .then(response=>response.json())
                                    .then(data=>{
                                        alert("Review Updated");
                                        BookReviewed()
                                    })
                            }
                        })
                    })
                    

                }
                
            })
            
            BookRviewEntry(flag)
            console.log("Bokj");
        }
        BookReviewed()
        
        function BookRviewEntry(ans){
            if (ans===0){
                document.getElementById("BookReviewEntry").innerHTML=`<textarea id="UsRv" placeholder="Write a Review"></textarea><div><button id="submitRv">Submit</button></div>`
                var SubmitRv = document.getElementById("submitRv")
                SubmitRv.addEventListener("click",function(){
                    var Review = document.getElementById("UsRv")
                    if(Review.value===""){
                        alert("Empty Review")
                    }
                    else{
                        var json={
                            "BookId" : bId,
                            "Review" :Review.value,
                        }
                        
                        const headers = {
                            "Content-Type": "application/json",
                          };
                        fetch(`/user/bookreview/${decoded.payload['ID']}`,{method:"post",headers,body:JSON.stringify(json)})
                            .then(response=>response.json())
                            .then(data=>{
                                if(data===true)
                                {
                                    BookReviewed();
                                    
                                }
                            })
                    }
                })
            }
            else{
                document.getElementById("BookReviewEntry").innerHTML="";
            }
            
            
        }
        
        
        
        

});





