function SearchBook(data){
    document.getElementById("main-bookswrap").innerHTML="";
    booksCount= data['Books'].length;
    for (let i=0;i<booksCount;i++)
    {
        bookDesc = `<div class="bookHolder">
        <div class="bookName">Title : ${data['Books'][i]['BookName']}</div>
        <div class="bookPublisher">Publisher : ${data['Books'][i]['Publisher']}</div>
        <div class="bookAuthor">Author : ${data['Books'][i]['Author']}</div>
        <div class="bookISBN">ISBN : ${data['Books'][i]['ISBN']}</div>
        <div class="bookPages">Pages : ${data['Books'][i]['Pages']}</div>
        <div class="bookTag">Tags : ${data['Books'][i]['Taglines']}</div>
    </div>`;
        
    document.getElementById("main-bookswrap").innerHTML=document.getElementById("main-bookswrap").innerHTML + bookDesc;
    }
}



fetch("http://127.0.0.1:3000/api/GetBooks")
            .then(response=>response.json())
            .then(data=>{
                SearchBook(data);
});

SearchValue = document.getElementById("SearchValue");
SearchColumn= document.getElementById("SearchColumn");
SearchBtn = document.getElementById("Search");
SearchValue.addEventListener("input",function(){
    if(SearchValue.value===""){
        fetch("http://127.0.0.1:3000/api/GetBooks")
            .then(response=>response.json())
            .then(data=>{
                SearchBook(data);
        });
    }
    else{
        SearchBtn.addEventListener("click",function(){
            json={
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
                SearchBook(book);
            }
            PrintBook();
              

        });

    }
})
