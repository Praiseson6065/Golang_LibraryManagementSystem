import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
userStatus(token);
UserPage(token);
var BooksPrice=0;
var Purchasing=false;
var BookLending = false;
if(token!=undefined)
{   
    var Decoded=DecodedToken(token);
    document.getElementById("CheckoutCart").hidden=true;    

}

async function GetUserCart(){
    var url ="/user/getusercart/" + Decoded.payload["ID"];

    await  fetch(url)
        .then(response=> response.json())
        .then(data=>  {   
            
            if(data.length===0){
                document.getElementById("cart-bookswrap").innerHTML="Empty Cart";
            }
            else{
                BookLending=true;
                for (let i in data){
                    var BookDetails=`
                <div class="bookHolder">
                <div class="bookImg"><a href="/book.html?BookC=${data[i]["BookCode"]}"><img class="bookImgCP" src="/img/books/${data[i]['ImgPath']}"></a></div>
                <div class="bookDetails">
                <div class="bookName">${data[i]['BookName']}</div>
                <div class="bookAuthor">by ${data[i]['Author']}</div>
               
                <div> <button class="rmBookbtn" data-bookvalue="${data[i]['BookId']}">Remove From Cart</button></div>
                </div>`;
                    document.getElementById("cart-bookswrap").innerHTML+=BookDetails;
                }
                var bookLend="";


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
                bookLend+=book;

                    }
                
                document.getElementById("BookLend").innerHTML=bookLend;
                let RmBtns = document.querySelectorAll(".rmBookbtn")
            
                for (var k = 0; k < RmBtns.length; k++) {
                    
                    RmBtns[k].addEventListener('click',function (event) {
                        
                        let bookId = event.target.dataset.bookvalue;
                        fetch(`/user/cart/${Decoded.payload["ID"]}/${bookId}`, { method: 'DELETE' })
                            .then(response => response.json())
                            .then(data => {
                                    window.location.reload();
                            });
                    } );
                }
                
                
            }
            
            
        });
    
            
}
async function ModalHandler(){
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
            <div id="BookLend">
            
                
            </div>
            <h3>Purchase Books</h3>
            <div id="BookPurchase">

            </div>
             </div>
            <div class="ModalConfirmHolder"><button id="ModalConfirm">Confirm</button></div>
                </div>
                </div>`;
    var modal =document.querySelector(".modal");
    var span = document.querySelector(".close");
    var CheckoutCartbtn = document.getElementById("CheckoutBtn");
    CheckoutCartbtn.addEventListener("click",function(){
        modal.style.display="block";  
        document.querySelector(".navbar").style.zIndex=-999;
        document.querySelector(".footer").style.zIndex=-990;  
    })
    span.onclick =  function (){
        modal.style.display="none";
    }
    window.onclick = function(event) {
        if (event.target == modal) {
        modal.style.display = "none";
        }
    }
    
        
}
async function GetUserPurchaseCart(){
    
    await fetch(`/user/purchasecart/${Decoded.payload['ID']}`)
                    .then(response=>response.json())
                    .then(data=>{
                        document.getElementById("purchasecart").innerHTML="";
                        if(data.length===0){
                            document.getElementById("purchasecart").innerHTML="Empty Cart";
                        }
                        else{
                            Purchasing=true;
                            for (let i in data){
                            if(data[i]['PurchaseDetails']['Quantity']>=1){ 
                            
                            var BookDetails=`
                            <div class="PurbookHolder">
                            <div class="bookImg"><a href="/book.html?BookC=${data[i]['Book']["BookCode"]}"><img class="bookImgCP" src="/img/books/${data[i]['Book']['ImgPath']}"></a></div>
                            <div class="bookDetails">
                            <div class="bookName">${data[i]['Book']['BookName']}</div>
                            <div class="bookAuthor">by ${data[i]['Book']['Author']}</div>
                            
                        
                            <div><button class="PurrmBookbtn" data-bookvalue="${data[i]['Book']['BookId']}">Remove From Cart</button></div>
                            
                            </div>
                            <div class="bookQuantityHandler"><button data-bookvalue="${data[i]['Book']['BookId']}" class="quantityDecrement">-</button><span class="bookQuantity">${data[i]['PurchaseDetails']['Quantity']}</span><button  class="quantityIncrement" data-book='${data[i]['Book']['BookId']},${data[i]['Book']['Quantity']},${data[i]['PurchaseDetails']['Quantity']}'>+</button></div>
                            </div>
                            `;

                                document.getElementById("purchasecart").innerHTML+=BookDetails;
                            }
                        }
                            var decrementQuantityBtns = document.querySelectorAll(".quantityDecrement")
                            var IncrementQuantityBtns = document.querySelectorAll(".quantityIncrement")
                            decrementQuantityBtns.forEach(btn =>{
                                btn.addEventListener("click",(e)=>{

                                    var bId = e.target.dataset.bookvalue
                                    console.log(bId)
                                    fetch(`/user/cartpurchasebook/${Decoded.payload['ID']}/${bId}/-1`,{method:"POST"})
                                        .then(response=>response.json())
                                        .then(data=>{
                                            if(data===true){
                                                GetUserPurchaseCart();
                                            }
                                        })
                                        

                                })
                            })
                            IncrementQuantityBtns.forEach(btn =>{
                                btn.addEventListener("click",(e)=>{
                                    var book= e.target.dataset.book.split(",")
                                    var bId = parseInt(book[0])
                                    var CurQuantity= parseInt(book[2])
                                    var bookAvailble= parseInt(book[1])

                                    if((CurQuantity+1)<=bookAvailble){
                                        fetch(`/user/cartpurchasebook/${Decoded.payload['ID']}/${bId}/1`,{method:"POST"})
                                            .then(response=>response.json())
                                            .then(data=>{
                                                if(data===true){
                                                    GetUserPurchaseCart();
                                                }
                                            })
                                    }
                                    else{
                                        alert("Max Quantity Reached")
                                    }
                                        

                                })
                            })
                            var RemovePurBtns = document.querySelectorAll(".PurrmBookbtn")
                            RemovePurBtns.forEach((btn)=>{
                                btn.addEventListener("click",(e)=>{
                                    fetch(`/user/rmPurchasecart/${Decoded.payload['ID']}/${e.target.dataset.bookvalue}`,{method:'DELETE'})
                                        .then(response=>response.json())
                                        .then(data=>{
                                            if(data===true){
                                                GetUserPurchaseCart();
                                            }
                                        })
                                })
                            })
                            
                            var BookPur="";
                            var Price=0;
                            var modalPurchase = document.getElementById("BookPurchase")
                            data.forEach((item)=>{
                                console.log(item);
                                BookPur+=`
                                <div class="ModalBookPurHolder">
                                <div class="ModalBookPurImgHolder"><img class="ModalBookImg" src="/img/books/${item['Book']["ImgPath"]}" alt=""></div>
                                <div class="ModalBookPurDetails">
                                    <div>${item['Book']["BookName"]}</div>  
                                    <div>by ${item['Book']["Author"]}</div>
                                    <div>Quantity : ${item['PurchaseDetails']["Quantity"]}</div>
                                </div>
                                <div>Price : &#8377; ${item['PurchaseDetails']["Quantity"]*item['Book']["Price"]} </div>
        
                                </div>

                                `;
                                Price+=(item['PurchaseDetails']["Quantity"]*item['Book']["Price"])
                                
                                modalPurchase.innerHTML=""; 
                                modalPurchase.innerHTML+=BookPur;
                            })
                            modalPurchase.innerHTML+=`
                                <div class="PurchaseTotal">
                                <div>Total Price : </div>
                                <div> &#8377; ${Price} </div>

                               
                                </div>
                            `;
                            BooksPrice=0;
                            BooksPrice=Price;
                            

                        }
                    });

}
async function handleBoth(){
    if(Purchasing || BookLending){
        document.getElementById("CheckoutBtn").hidden=false;
    }
    console.log(Purchasing,BookLending)
    var Confirm=document.getElementById("ModalConfirm");
        var modalcontent=document.querySelector(".modalcontent");
        var lend=false;
        var pur=false;
    await Confirm.addEventListener("click",async function(){
            if(Purchasing){
                modalcontent.innerHTML=`
                <div class="paymentMode">
                    <div>Amount : ${BooksPrice}</div>
                    <div>
                    <label for="cname">Name on Card</label>
                    <input type="text" id="cname" name="cardname" placeholder="John More Doe">
                    <label for="ccnum">Credit card number</label>
                    <input type="text" id="ccnum" name="cardnumber" placeholder="1111-2222-3333-4444">
                    <label for="expmonth">Exp Month</label>
                    <input type="text" id="expmonth" name="expmonth" placeholder="September">
                    </div>
                    <div><button id="paynow">Pay Now</button></div>
                </div>
                `

                await document.getElementById("paynow").addEventListener("click",async ()=>{

                    await fetch(`/user/paymentforpurchasing/${Decoded.payload['ID']}`,{method:"POST"})
                    .then(response=>response.json())
                    .then(async data=>{
                        const headers = {
                                    "Content-Type": "application/json",
                          };
                        //   var json = {
                        //     "id":data.id,
                        //     "client_secret":data.client_secret,
                        //           }
                            // console.log(json)
                        await fetch(`/user/confirmpayment/${Decoded.payload["ID"]}`,{method:"POST",headers,body: JSON.stringify(data)})
                             .then(response=>response.json())
                             .then(async ans=>{
                                console.log(ans)
                                 if(ans==="http://localhost:3000/payment/success"){
                                     await fetch(`/user/purchasebook/${Decoded.payload["ID"]}`,{method:"PUT"})
                                         .then(response=>response.json())
                                         .then(data=>{
                                                 if(data==true){
                                                            pur=true;
                                                        }
                                                    })
                                            }

                                        })
                                })


                            
                        })
                        
                    }
                    if(BookLending){
                        await fetch(`/user/checkoutcart/${DecodedToken(token).payload["ID"]}`,{method:"POST"})
                        .then(response=> response.json())
                        .then(data=>{
                            if(data===true)
                            {
                                lend=true;
                            }   
                            else{
                                modalcontent.innerHTML="<h3>Books Are Not Issued</h3>";
                            }
                        if(pur && lend)
                        {
            
                             modalcontent.innerHTML="<h3>Books Issued and Purchased</h3>"  ;
                                        
                        }
                    else if(pur){
                        modalcontent.innerHTML="<h3>Books Purchased</h3>"  ;

                    }
                    else if(lend){
                        modalcontent.innerHTML="<h3>Books Issued</h3>"
                        
                    }
                    else{
                        modalcontent.innerHTML="<h3>Books Are Not Issued</h3>";
                        
                    }
                    setTimeout(function(){
                        window.location.reload();
                     }, 5000);

                        });
                    }
                    
                    
                })
                
                


}
await ModalHandler();
await GetUserCart();
await GetUserPurchaseCart();
await handleBoth();

         



