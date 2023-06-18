import {token,DecodedToken,userStatus} from './userstatus.js';
var decoded = DecodedToken(token);
userStatus(token);
function DisplayBook(data){
    document.getElementById("main-bookswrap").innerHTML="";
    var booksCount= data['Books'].length;
    for (let i=0;i<booksCount;i++)
    {
        
        var bookDesc = `<div class="bookHolder">
        <div class="bookDetails">
        <div class="bookName">Title : ${data['Books'][i]['BookName']}</div>
        <div class="bookPublisher">Publisher : ${data['Books'][i]['Publisher']}</div>
        <div class="bookAuthor">Author : ${data['Books'][i]['Author']}</div>
        <div class="bookISBN">ISBN : ${data['Books'][i]['ISBN']}</div>
        <div class="bookPages">Pages : ${data['Books'][i]['Pages']}</div>
        <div class="bookTag">Tags : ${(data['Books'][i]['Taglines']).replace(/[{}"]/g,'')}</div>
        </div>
        <div class="bookImg"><a href="/book.html?BookC=${data['Books'][i]["BookCode"]}"><img class="bookImgCP" src="/img/books/${data['Books'][i]['ImgPath']}"></a></div>
    </div>`;
        
    document.getElementById("main-bookswrap").innerHTML=document.getElementById("main-bookswrap").innerHTML + bookDesc;
    }
}



fetch("/api/GetBooks")
            .then(response=>response.json())
            .then(data=>{
                DisplayBook(data);
});

var SearchValue = document.getElementById("SearchValue");
var SearchColumn= document.getElementById("SearchColumn");
var SearchBtn = document.getElementById("Search");
SearchValue.addEventListener("input",function(){
    if(SearchValue.value===""){
        fetch("http://127.0.0.1:3000/api/GetBooks")
            .then(response=>response.json())
            .then(data=>{
                DisplayBook(data);
        });
    }
    else{
        SearchBtn.addEventListener("click",function(){
            var json={
                "SearchValue": SearchValue.value,
                "SearchColumn": SearchColumn.value,
            };
            const headers = {
                "Content-Type": "application/json",
              };
            const SearchData = fetch("http://127.0.0.1:3000/api/searchbook", {method: "POST",headers,body: JSON.stringify(json),})
              .then((response) => response.json())
              .then((books) => { 
                return books;

              });
            const PrintBook = async () => {
                const book = await SearchData;
                DisplayBook(book);
            }
            PrintBook();
              

        });

    }
})

