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
        const book = await SearchData;
        const noOfBooks =book['Books'].length<=3 ? book['Books'].length : 4;
        for(let i=0;i<noOfBooks;i++)
        {   if(book['Books'][i]["BookCode"]===bookId){
            continue;
            }
            
            var suggested=`<div class="bookHolder-Suggested">
            <div class="bookImg"><a href="/book.html?BookC=${book['Books'][i]["BookCode"]}"><img class="bookImgCP" src="/img/books/${book['Books'][i]['ImgPath']}"></a></div>
            <div class="bookDetails">
            <div class="bookName">Title : ${book['Books'][i]['BookName']}</div>
            <div class="bookPublisher">Publisher : ${book['Books'][i]['Publisher']}</div>
            <div class="bookAuthor">Author : ${book['Books'][i]['Author']}</div>
            <div class="bookISBN">ISBN : ${book['Books'][i]['ISBN']}</div>
            <div class="bookPages">Pages : ${book['Books'][i]['Pages']}</div>
            <div class="bookTag">Tags : ${book['Books'][i]['Taglines']}</div>
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
        <div>Title : ${data['BookName']}</div>
        <div>Pages : ${data['Pages']}</div>
        <div>Author : ${data['Author']}</div>
        <div>Publisher : ${data['Publisher']}</div>
        <div>Quantity : ${data['Quantity']}</div>
        <div>ISBN : ${data['ISBN']}</div>
        <div id="votes">Votes : ${data['votes']}</div>
        <div>Tags : ${data['Taglines']}</div>
        </div>
        <div class="BookCheckout" id="Bookchk">
            <button id="likebtn"><i class='bx bx-like'id="like"></i></button>
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
            fetch(`/api/cartbkchk/${decoded.payload["ID"]}/${bId}`)
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
                fetch(`/api/isbookissued/${decoded.payload["ID"]}/${bId}`)
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
        var tags = data['Taglines'].split(",");
        var likebtn=document.getElementById("likebtn");
        function isliked(){
            fetch(`/api/isliked/${decoded.payload["ID"]}/${data['BookId']}`)
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
            fetch(`/api/like/${decoded.payload["ID"]}/${data['BookId']}`,{method:"POST"})
                .then(response=> response.json())
                .then(data=> {
                        console.log(data);
                        isliked();
                        fetch(`/api/bookc/${bookId}`)
                            .then(response => response.json())
                            .then(data=>{
                               document.getElementById("votes").innerText=`Votes : ${data['votes']}`;
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
                    
                    fetch(`/api/cart/${decoded.payload["ID"]}/${bId}`,{method:`Post`})
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
        
        
        

});



