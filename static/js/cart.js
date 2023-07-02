import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
userStatus(token);
UserPage(token);
if(token!=undefined)
{   
    var Decoded=DecodedToken(token);
    document.getElementById("CheckoutCart").hidden=true;    

}

function GetUserCart(){
    var url ="/api/getusercart/" + Decoded.payload["ID"];

    fetch(url)
        .then(response=> response.json())
        .then(data=>{   
            
            if(data.length===0){
                document.getElementById("CheckoutBtn").hidden=true;
                document.getElementById("cart-bookswrap").innerHTML="Empty Cart";
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
                <div> <button class="rmBookbtn" data-bookvalue="${data[i]['BookId']}">Remove From Cart</button></div>
    
                
            </div>`;
                    document.getElementById("cart-bookswrap").innerHTML+=BookDetails;
                }
                const RmBtns = document.querySelectorAll(".rmBookbtn")
                
                for (var k = 0; k < RmBtns.length; k++) {
                    RmBtns[k].addEventListener('click', function(event) {
                      console.log(event);
                      let bookId = event.target.dataset.bookvalue;
                      fetch(`/api/cart/${Decoded.payload["ID"]}/${bookId}`, { method: 'DELETE' })
                        .then(response => response.json())
                        .then(data => {
                          console.log(data);
                          window.location.reload();
                        });
                    });
                  }
                
                var books="";
                    for(let i in data){
                        var book=`
                        <div class="ModalBookHolder">
                        <div class="ModalBookImgHolder"><img class="ModalBookImg" src="/img/books/${data[i]["ImgPath"]}" alt=""></div>
                        <div class="ModalBookDetails">
                            <div>Title:${data[i]["BookName"]}</div>
                            <div>Pages:${data[i]["Pages"]}</div>
                            <div>Publisher:${data[i]["Publisher"]}</div>
                            
                        </div>

                        </div>
                        `;
                        books+=book;

                    }
                    var maintag=document.querySelector("main");
                
                    maintag.innerHTML+=`
                    <div class="modal">
                    
                    <div class="modalcontent">
                    <span class="close">&times;</span>
                        <div class="ModalUserWrap">
                            <div>UserDetails</div>
                            <div>Name: ${DecodedToken(token).payload["name"]}</div>
                            <div>Email: ${DecodedToken(token).payload["email"]}</div>
                            <div>Date: ${new Date()}</div>
                        </div>
                        <h3>Books</h3>
                        <div class="ModalBookWrap">
                            
                            ${books}
                        </div>
                        <div class="ModalConfirmHolder"><button id="ModalConfirm">Confirm</button></div>
                    </div>
        
        
                </div>`;
                var modal =document.querySelector(".modal");
                var span = document.querySelector(".close");
                var CheckoutCartbtn = document.getElementById("CheckoutBtn");
                CheckoutCartbtn.addEventListener("click",function(){
                    modal.style.display="block";    
                })
                span.onclick =  function (){
                    modal.style.display="none";
                }
                // window.onclick = function(event) {
                //     if (event.target == modal) {
                //     modal.style.display = "none";
                //     }
                // }
                var Confirm=document.getElementById("ModalConfirm");
                var modalcontent=document.querySelector(".modalcontent");
                Confirm.addEventListener("click",function(){
                    fetch(`/api/checkoutcart/${DecodedToken(token).payload["ID"]}`,{method:"POST"})
                        .then(response=> response.json())
                        .then(data=>{
                            if(data===true)
                            {
                                modalcontent.innerHTML="<h3>Books Issued</h3>"  ;
                                setTimeout(function(){
                                    window.location.reload();
                                 }, 5000);
                            }   
                            else{
                                modalcontent.innerHTML="<h3>Books Are Not Issued</h3>";
                            }

                        });
                })
                
            }
            
            
        });
            
}
GetUserCart();

         



