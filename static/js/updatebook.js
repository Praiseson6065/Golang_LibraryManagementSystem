import {token,userStatus,DecodedToken} from "./userstatus.js";
userStatus(token);
var decoded = DecodedToken(token);
if(decoded.payload["usertype"]!="admin")
{
    window.location="/profile.html";
}
function ExistingData(){
    var url = "http://127.0.0.1:3000/api/book/";
    if(IdValue.value==='')
    {
        var divFields =["Id","BookName","ISBN","Author","Publisher","Pages","Taglines"];
                for (let i in divFields){
                    document.getElementById(divFields[i].toString()).innerText = "";
                
                }
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
                document.getElementById("Id").innerText="NotFound";
            }
            else if(data===500)
            {
                
                document.getElementById("Id").innerText="Error";
            }
            else{
                data = JSON.parse(data);
                var divFields =["Id","BookName","ISBN","Author","Publisher","Pages","Taglines"];
                for (let i in divFields){
                    document.getElementById(divFields[i].toString()).innerText = divFields[i] + " : "+ data[divFields[i]] ;
                
                }
                document.getElementById("CoverPage").setAttribute("src","http://127.0.0.1:3000/img/"+data["ImgPath"]);
            }
                
        });
    }
}

var IdValue = document.getElementById("id");
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
  const taglines = document.querySelector('input[name="Taglines"]').value;
  const imageFile = document.querySelector('input[name="ImgPath"]').files[0];
  const formData = new FormData();
  formData.append('BookId', bookId);
  formData.append('BookName', bookName);
  formData.append('ISBN', isbn);
  formData.append('Pages', pages);
  formData.append('Publisher', publisher);
  formData.append('Author', author);
  formData.append('Taglines', taglines);
  formData.append('ImgPath', imageFile);
  fetch(`/api/updatebook/${bookId}`, {
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
