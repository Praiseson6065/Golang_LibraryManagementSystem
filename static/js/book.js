import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
userStatus(token);
console.log(token);
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
            <div><input type="checkbox" id="purchasebkchk"><span>Purchase Book</span></div>
            <div id="purchasediv"><div><button id="quantityDecrement">-</button><span id="quantityValue">1</span><button id="quantityIncrement">+</button></div><div><button id="purchasebtn"><i class='bx bx-cart-add'></i></button></div></div>
            <div><button id="addtocart" >Lend</button></div>
           
            
        
        </div>
        
        `;
        
        document.getElementById("book-wrap").innerHTML=book;
        document.getElementById("purchasebkchk").addEventListener("change",(event)=>{
            if(event.target.checked){
                document.getElementById("addtocart").style.display="none"
                document.getElementById("purchasediv").style.display="flex";
                var btnDecrement=document.getElementById("quantityDecrement");
                var btnIncrement=document.getElementById("quantityIncrement");
                var quantityValue=document.getElementById("quantityValue");
                if(parseInt(quantityValue.innerHTML)===1){
                    btnDecrement.disabled=true;
                    btnDecrement.classList.add("disabled");
                }
                btnDecrement.addEventListener("click",()=>{
                    if(parseInt(quantityValue.innerHTML)===1){
                        btnDecrement.disabled=true;
                        btnDecrement.classList.add("disabled");
                    }
                    else{
                        btnDecrement.disabled=false;
                        quantityValue.innerHTML=parseInt(quantityValue.innerHTML)-1;
                        btnIncrement.disabled=false;
                        btnDecrement.classList.remove("disabled");
                        btnIncrement.classList.remove("disabled");
                        if(parseInt(quantityValue.innerHTML)===1){
                            btnDecrement.disabled=true;
                            btnDecrement.classList.add("disabled");
                        }
                    }
                });
                btnIncrement.addEventListener("click",()=>{
                    if(parseInt(quantityValue.innerHTML)===data['Quantity']){
                        btnIncrement.disabled=true;
                        btnIncrement.classList.add("disabled");

                    }
                    else{
                        btnDecrement.disabled=false;
                        btnDecrement.classList.remove("disabled");
                        quantityValue.innerHTML=parseInt(quantityValue.innerHTML)+1;
                        btnIncrement.classList.remove("disabled");
                        if(parseInt(quantityValue.innerHTML)===data['Quantity']){
                            btnIncrement.disabled=true;
                            btnIncrement.classList.add("disabled");
                        }
                    }
                });
                var purchasebtn = document.getElementById("purchasebtn");
                purchasebtn.addEventListener("click",()=>{
                    fetch(`/user/cartpurchasebook/${decoded.payload['ID']}/${bId}/${parseInt(quantityValue.innerHTML)}`,{method:"POST"})
                        .then(response=>response.json())
                        .then(data=>{
                            if(data===true){
                            quantityValue.innerHTML=1;
                            btnDecrement.disabled=true;
                            btnDecrement.classList.add("disabled");
                            }
                        })

                })

            }
            else{
                document.getElementById("purchasediv").style.display="none";
                document.getElementById("addtocart").style.display="block";
                

            }
        })
        
        
        
        if(token!=undefined && decoded.payload["usertype"]==="admin"){
            document.getElementById("Bookchk").innerHTML+=`<a href="/updatebook.html?BId=${data['BookId']}"><button id="editbook" >Edit Book</button></a>`;
        }   
        if(data['Quantity']===0){
            document.getElementById("addtocart").innerText="Out of Stock";
            document.getElementById("addtocart").classList.add("bookoutofstock");
            document.getElementById("addtocart").disabled=true;

        }
        
        async function CartLimit(){
            await fetch(`/user/getusercart/+${decoded.payload["ID"]}`)
                .then(response=>response.json())
                .then(data=>{
                    if(data.length>=5){
                        document.getElementById("addtocart").innerText="Cart Limit Reached";
                        document.getElementById("addtocart").classList.add("bookoutofstock");
                        document.getElementById("addtocart").disabled=true;
                    }
                })
        }
        async function UserBookRelation(){
            
            if(token!=undefined){
                CartLimit();    
                await   fetch(`/user/userbookdetails/${decoded.payload["ID"]}/${bId}`)
                .then(response=>response.json())
                .then(data=>{
                    console.log(data)
                    if(data.Approve==true){
                        document.getElementById("addtocart").innerText="Already Book Taken";
                        document.getElementById("addtocart").classList.add("bookoutofstock");
                        document.getElementById("addtocart").disabled=true;
                    }
                    else if(data.Issued==true){
                        document.getElementById("addtocart").innerText="Already Requested";
                        document.getElementById("addtocart").classList.add("bookoutofstock");
                        document.getElementById("addtocart").disabled=true;
                    }
                    else if(data.Cart==true){
                        document.getElementById("addtocart").innerText="Added to Cart";
                        document.getElementById("addtocart").classList.add("bookoutofstock");
                        document.getElementById("addtocart").disabled=true;
                    }
                    
                })
                
            }
        }
        UserBookRelation();
        
    
        var tags = data['Taglines'].split(",");
        var likebtn=document.getElementById("likebtn");
        function isliked(){
            if(token!=undefined){
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
            
        }
        isliked();
        likebtn.addEventListener("click",function(event){
            if(token!=undefined){
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
                })
            }
            else{
                alert("Please Login");
            }
            
        })
        suggestedBooks(tags);
        var AddtoCart=document.getElementById("addtocart");
        if(AddtoCart!=null){
            AddtoCart.addEventListener("click",function(){
                console.log(token)
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





