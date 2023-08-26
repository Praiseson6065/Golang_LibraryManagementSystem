import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
userStatus(token);
UserPage(token);
var BooksPrice=0;
var Purchasing=false;
var BookLending = false;
var RedirectUrl="";
var error=false;
var error_books=[];    
var lend=false;
var pur=false;
if(token!=undefined)
{   
    var Decoded=DecodedToken(token);
    document.getElementById("CheckoutCart").hidden=true;    

}
function status(purc,lend){
    console.log(purc,lend)
    var modalcontent=document.querySelector(".modalcontent")
    console.log("purchlend")
    if(purc && lend){
        modalcontent.innerHTML="Books Purchased && Issued"
        
    }
    else if(purc ){
        modalcontent.innerHTML="Books Purchased"
        
    }
    else if(lend){
        modalcontent.innerHTML="Books Issued"
        
    }
    else{
        modalcontent.innerHTML="Error"
    }
    setTimeout(() => {
        console.log(RedirectUrl);
        console.log("redirect");    
        window.location.href = RedirectUrl;
        },5000);
}
async function GetUserCart(){
    var url ="/user/getusercart/" + Decoded.payload["ID"];

    await  fetch(url)
        .then(response=> response.json())
        .then(data=>  {   
            
            
            if(data.length!=0){
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
        
        if(error){

           var dataer=document.querySelectorAll("div[data-booker]") 
           for(let i=0;i<dataer.length;i++){
                dataer[i].classList.add("error");
           }
           alert("Please Check the Quantity of the books.")
        }else{
        modal.style.display="block";  
        document.querySelector(".navbar").style.zIndex=-999;
        document.querySelector(".footer").style.zIndex=-990;
        }  
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
                        
                        if(data!=null){
                            
                            Purchasing=true;
                            for (let i in data){
                                if(data[i]['PurchaseDetails']['Quantity']>=1){ 
                                    if(data[i]['PurchaseDetails']['Quantity']>data[i]['Book']['Quantity']){
                                        error=true;
                                        error_books.push(data[i]['Book']['BookId'])
                                        var erval=`data-booker=${data[i]['Book']['BookId']}`
                                    }
                            var BookDetails=`
                            <div class="PurbookHolder" ${erval}>
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
                                                window.location.reload();   
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
    
    if(!Purchasing && BookLending){
        document.querySelector(".purhsecrthead").style.display="none";

    }
    else if(Purchasing && !BookLending){
        document.querySelector(".lendcrthead").style.display="none";
        

    }
    else if(!Purchasing && !BookLending){
        document.getElementById("CheckoutBtn").style.display="none";
        document.querySelector(".purhsecrthead").style.display="none";
        document.querySelector(".lendcrthead").style.display="none";
        document.getElementById("cart-bookswrap").innerHTML="Empty Cart"


    }
    console.log(Purchasing,BookLending)
    var Confirm=document.getElementById("ModalConfirm");
        var modalcontent=document.querySelector(".modalcontent");
        
    await Confirm.addEventListener("click",async function(){
        if(BookLending && !Purchasing){
            fetch(`/user/checkoutcart/${DecodedToken(token).payload["ID"]}`,{method:"POST"})
            .then(response=> response.json())
            .then(data=>{
                if(data===true)
                {
                    lend=true;
                }   
                else{
                    modalcontent.innerHTML="<h3>Books Are Not Issued</h3>";
                }
                

            })
            .then(res=>{
                status(pur,lend)
            })
            
            // .then(res=>{
            //     status(purc,lend)
            // })

        }
        
        else if(Purchasing && !BookLending){
            modalcontent.innerHTML=`
            <div class="paymentMode">   
                <div>Amount : ${BooksPrice}</div>
                <div class="paymentdiv">
                <label for="cname">Name on Card</label>
                <input type="text" id="cname" name="cardname" placeholder="John More Doe">
                <label for="ccnum">Credit card number</label>
                <input type="text" id="ccnum" name="cardnumber" placeholder="1111-2222-3333-4444">
                <label for="expmonth">Exp Month</label>
                <input type="text" id="expmonth" name="expmonth" placeholder="September">
                </div>
                <div><button id="paynow">Pay Now</button><button id="cancel">cancel</button></div>
            </div>
            `
            document.getElementById("cancel").addEventListener("click",()=>{
                window.location.reload()
            })
            document.getElementById("paynow").addEventListener("click",async ()=>{

                fetch(`/user/paymentforpurchasing/${Decoded.payload['ID']}`,{method:"POST"})
                .then(response=>response.json())
                .then(async data=>{
                    const headers = {
                                "Content-Type": "application/json",
                      };
                      
                        await fetch(`/user/confirmpayment/${Decoded.payload["ID"]}`,{method:"POST",headers,body: JSON.stringify(data)})
                         .then(response=>response.json())
                         .then(async ans=>{
                             if(ans){
                                RedirectUrl=ans;
                                console.log(RedirectUrl)
                                  await fetch(`/user/purchasebook/${Decoded.payload["ID"]}`,{method:"PUT"})
                                     .then(response=>response.json())
                                     .then( data=>{
                                             if(data==true){
                                                        pur=true;
                                                }
                                                
                                                    
                                                })
                                    .then(res=>{
                                                    status(pur,lend)
                                                })
                                        }

                                    })

                                
                            })


                        
                    })
                        
                    }
            else if(Purchasing && BookLending){
                fetch(`/user/checkoutcart/${DecodedToken(token).payload["ID"]}`,{method:"POST"})
                .then(response=> response.json())
                .then(data=>{
                    if(data===true)
                    {
                        lend=true;
                    }   
                    else{
                        modalcontent.innerHTML="<h3>Books Are Not Issued</h3>";
                    }
                    

                })
                .then(res=>{
                    modalcontent.innerHTML=`
                    <div class="paymentMode">
                        <div>${lend===true ? "BookIssued PayNow" : "" }</div>
                        <div>Amount : ${BooksPrice}</div>
                        <div class="paymentdiv">
                        <label for="cname">Name on Card</label>
                        <input type="text" id="cname" name="cardname" placeholder="John More Doe">
                        <label for="ccnum">Credit card number</label>
                        <input type="text" id="ccnum" name="cardnumber" placeholder="1111-2222-3333-4444">
                        <label for="expmonth">Exp Month</label>
                        <input type="text" id="expmonth" name="expmonth" placeholder="September">
                        </div>
                        <div><button id="paynow">Pay Now</button><button id="cancel">cancel</button></div>
                    </div>
                    `
                    document.getElementById("cancel").addEventListener("click",()=>{
                        window.location.reload()
                    })
                    document.getElementById("paynow").addEventListener("click",async ()=>{
    
                        fetch(`/user/paymentforpurchasing/${Decoded.payload['ID']}`,{method:"POST"})
                        .then(response=>response.json())
                        .then(async data=>{
                            const headers = {
                                        "Content-Type": "application/json",
                              };
                              
                                await fetch(`/user/confirmpayment/${Decoded.payload["ID"]}`,{method:"POST",headers,body: JSON.stringify(data)})
                                 .then(response=>response.json())
                                 .then(async ans=>{
                                     if(ans){
                                        RedirectUrl=ans;
                                        console.log(RedirectUrl)
                                          await fetch(`/user/purchasebook/${Decoded.payload["ID"]}`,{method:"PUT"})
                                             .then(response=>response.json())
                                             .then( data=>{
                                                     if(data==true){
                                                                pur=true;
                                                        }
                                                        
                                                            
                                                        })
                                            .then(res=>{
                                                        console.log(res)
                                                            status(pur,lend)
                                                        })
                                                }
    
                                            })
    
                                        
                                    })
    
    
                                
                            })

                })


            }
                    
                    
                    
    
                    
                })

    
    if(lend || pur){
        console.log("hello")
        return "Both"
    }         
    
    else{
        return "none"
    }
                
                


}

  await ModalHandler();
  await GetUserCart();
  await GetUserPurchaseCart();
  var ans= await handleBoth();
  await console.log(error,"Error",error_books)
  if(ans=="Both"){
    status(pur, lend);
    setTimeout(1000, () => {
    console.log(RedirectUrl);
    window.location = RedirectUrl;
    });
  }
  

   


         



