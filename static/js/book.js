import {token,DecodedToken,userStatus} from "./userstatus.js"
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
        
        for(let i=0;i<book['Books'].length;i++)
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
        <div>Votes : ${data['votes']}</div>
        <div>Tags : ${data['Taglines']}</div>
        </div>
        <div class="BookCheckout">
            <label class="container">
            <input id="like" type="checkbox">
            <svg id="Glyph" version="1.1" viewBox="0 0 32 32" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><path d="M29.845,17.099l-2.489,8.725C26.989,27.105,25.804,28,24.473,28H11c-0.553,0-1-0.448-1-1V13  c0-0.215,0.069-0.425,0.198-0.597l5.392-7.24C16.188,4.414,17.05,4,17.974,4C19.643,4,21,5.357,21,7.026V12h5.002  c1.265,0,2.427,0.579,3.188,1.589C29.954,14.601,30.192,15.88,29.845,17.099z" id="XMLID_254_"></path><path d="M7,12H3c-0.553,0-1,0.448-1,1v14c0,0.552,0.447,1,1,1h4c0.553,0,1-0.448,1-1V13C8,12.448,7.553,12,7,12z   M5,25.5c-0.828,0-1.5-0.672-1.5-1.5c0-0.828,0.672-1.5,1.5-1.5c0.828,0,1.5,0.672,1.5,1.5C6.5,24.828,5.828,25.5,5,25.5z" id="XMLID_256_"></path></svg>
            </label>
            <button id="addtocart" >Add to Cart</button>
        </div>`;
        document.getElementById("book-wrap").innerHTML=book;

        var tags = data['Taglines'].split(",");
    
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

                        });
                    
                }
                    
            });
        }
        

});



