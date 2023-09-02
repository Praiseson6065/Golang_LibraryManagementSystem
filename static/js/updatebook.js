import{token,DecodedToken,userStatus,UserPage} from "./userstatus.js";
UserPage(token);
if(token === undefined){
    window.location=`/home.html`;
    if(DecodedToken(token).payload["usertype"]!="admin")
    {
    window.location="/profile.html";
    }
}
else{
    userStatus(token);
}

function ExistingData(){
    var url = "/api/book/";
    if(IdValue.value==='')
    {
        var divFields =["BookId","BookName","ISBN","Author","Price","Publisher","Pages","Taglines","Quantity"];
                for (let i in divFields){
                    document.getElementById(divFields[i].toString()).innerText = "";
                
                }
        document.getElementById("CoverPage").src="";
    }
    else{
        fetch(url+IdValue.value)
        .then(response => {
            if(response.status===405)
            {
                console.log('405 Method Not Allowed\n Not Found');
                return 405;
                
            }
            else if(response.status===500)
            {
                console.log('500 Method Not Allowed\n Sql error');
                return 500;
                
            }
            else{
                return response.text();
            }
            
        })
        .then(data => {
            if(data===405)
            {
                document.getElementById("BookId").innerText="NotFound";
            }
            else if(data===500)
            {
                
                document.getElementById("BookId").innerText="Error";
            }
            else{
                data = JSON.parse(data);
                
                var divFields =["BookId","BookName","ISBN","Author","Publisher","Pages","Taglines","Quantity"];
                for (let i in divFields){
                    document.getElementById(divFields[i].toString()).innerText = divFields[i] + " : "+ data[divFields[i]] ;
                
                }
                document.getElementById("CoverPage").setAttribute("src","/img/books/"+data["ImgPath"]);
            }
                
        });
    }
}

var IdValue = document.getElementById("id");
const queryString = window.location.search;
var url  = new URLSearchParams(queryString);
var bookId = url.get('BId');
if(bookId!=null){
    document.getElementById("id").value=bookId; 
    ExistingData();
}
IdValue.addEventListener("input",ExistingData);
const form = document.querySelector('form');

form.addEventListener('submit', (event) => {
  event.preventDefault();
  const bookId = document.querySelector('#id').value;
  const bookName = document.querySelector('input[name="BookName"]').value;
  const isbn = document.querySelector('input[name="ISBN"]').value;
  const pages = document.querySelector('input[name="Pages"]').value;
  const publisher = document.querySelector('input[name="Publisher"]').value;
 
  const author = document.querySelector('input[name="Author"]').value;
  const price = document.querySelector('input[name="Price"]').value;
  const taglines = document.querySelector('input[name="Taglines"]').value;
  const quantity = document.querySelector('input[name="Quantity"]').value
  const imageFile = document.querySelector('input[name="ImgPath"]').files[0];
  const formData = new FormData();
  formData.append('BookId', bookId);
  formData.append('BookName', bookName);
  formData.append('ISBN', isbn);
  formData.append('Pages', pages);
  formData.append('Publisher', publisher);
  formData.append('Author', author);
  formData.append('Taglines', taglines);
  formData.append('Quantity', quantity);
  formData.append('Price', price)
  formData.append('ImgPath', imageFile);
  fetch(`/admin/updatebook/${bookId}`, {
    method: 'PUT',
    body: formData,
  })
    .then(response => response.json())
    .then(data => {
        
      console.log(data["msg"]);
      
    })
    .catch(error => {
      console.error('Error:', error);
    });
});
var ImgPath=document.getElementById("ImgPath");
ImgPath.addEventListener("change",function (){
    const file = ImgPath.files[0];
    document.getElementById("updatedimg").src=URL.createObjectURL(file);
})   
