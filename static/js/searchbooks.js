import {token,DecodedToken,userStatus} from './userstatus.js';
var decoded = DecodedToken(token);
userStatus(token);
const queryString = window.location.search;
var url  = new URLSearchParams(queryString);
var Search = url.get('Search');
var SearchValue = document.getElementById("SearchValue");
var SearchColumn= document.getElementById("SearchColumn");
var SearchBtn = document.getElementById("Search");

    

function DisplayBook(data){
    document.getElementById("main-bookswrap").innerHTML="";
    var booksCount= data['Books'].length;
    if(booksCount===0){
        document.getElementById("main-bookswrap").innerHTML="Books Not Found"
    }
    else{
        for (let i=0;i<booksCount;i++)
    {
        
        var bookDesc = `<div class="bookHolder">
        <div class="bookImg"><a href="/book.html?BookC=${data['Books'][i]["BookCode"]}"><img class="bookImgCP" src="/img/books/${data['Books'][i]['ImgPath']}"></a></div>
        <div class="bookDetails">
        <div class="bookName">${data['Books'][i]['BookName']}</div>
        <div class="bookAuthor">by ${(data['Books'][i]['Author']).replace(/[{}"]/g,'')}</div>
        <div class="bookCritics">
            <div class="bookDes"><div class="bookD">${data['Books'][i]['Pages']}</div><div  class="bookDName">Pages</div></div>
            <div class="bookDes"><div class="bookD">${data['Books'][i]['votes']}</div ><div class="bookDName">Likes</div></div>
        </div>
        
    </div>`;
        
    document.getElementById("main-bookswrap").innerHTML=document.getElementById("main-bookswrap").innerHTML + bookDesc;
    }   
    }
    
}



fetch("/api/GetBooks")
            .then(response=>response.json())
            .then(data=>{
                DisplayBook(data);
});


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

if(Search!=null || Search!=""){
    
    SearchValue.value=Search;
    SearchColumn.value="book_name";
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
}
