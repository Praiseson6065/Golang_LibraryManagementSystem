import{token,DecodedToken,userStatus} from "./userstatus.js";
userStatus(token);
var RmBtns;
if(token!=undefined)
{
    var Decoded=DecodedToken(token);

}

function GetUserCart(){
    var url ="/api/getusercart/" + Decoded.payload["ID"];

    fetch(url)
        .then(response=> response.json())
        .then(data=>{
            if(data===null){
                document.getElementById("cart-bookswrap").innerHTML="Empty Cart"
            }
            else{
                for (let i in data){
                    var BookDetails=`
                <div class="bookHolder">
                <div class="bookImg"><a href="/book.html?BookC=${data[i]["BookCode"]}"><img class="bookImgCP" src="/img/books/${data[i]['ImgPath']}"></a></div>
                <div class="bookDetails">
                <div class="bookName">Title : ${data[i]['BookName']}</div>
                <div class="bookPublisher">Publisher : ${data[i]['Publisher']}</div>
                <div class="bookAuthor">Author : ${data[i]['Author']}</div>
                <div class="bookISBN">ISBN : ${data[i]['ISBN']}</div>
                <div class="bookPages">Pages : ${data[i]['Pages']}</div>
                <div class="bookTag">Tags : ${(data[i]['Taglines']).replace(/[{}"]/g,'')}</div>
    
                </div>
                <div><button data-bookvalue="${data[i]['BookId']}">RemoveFromCart</button></div>
    
                
            </div>`
                    document.getElementById("cart-bookswrap").innerHTML+=BookDetails;
                }
                window.RmBtns = document.querySelectorAll("button");
                
                
            }
            
        });
}
GetUserCart();
console.log(RmBtns);
if(RmBtns!=undefined){
        for (let i=0;i<RmBtns.length;i++){
        RmBtns[i].addEventListener("click",function (event){
            console.log(event["dataset"]);
            // fetch(`api/cart/${Decoded.payload["ID"]}/${}`)
        })
    } 
}


