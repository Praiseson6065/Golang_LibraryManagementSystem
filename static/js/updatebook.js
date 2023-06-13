import {token,userStatus,DecodedToken} from "./userstatus.js";
userStatus(token);
var url = "http://127.0.0.1:3000/api/book/";
var IdValue = document.getElementById("id");
IdValue.addEventListener("input",function(){
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
    
});
function updateBook(id) {
    const formData = new FormData();
    formData.append('BookName', document.getElementById('BookName').value);
    formData.append('ISBN', document.getElementById('ISBN').value);
    formData.append('Pages', document.getElementById('Pages').value);
    formData.append('Publisher', document.getElementById('Publisher').value);
    formData.append('Author', document.getElementById('Author').value);
    formData.append('Taglines', document.getElementById('Taglines').value);
    const file = document.getElementById('ImgPath').files[0];
    if (file) {
      formData.append('ImgPath', file);
    }
  
    const url = `http://127.0.0.1:3000/api/updatebook/${id}`;
    console.log(formData);
    return fetch(url, {
      method: 'PUT',
      body: formData
    });
  }
  updateBook(1);
  const form = document.querySelector('form');
  form.addEventListener('submit', async (event) => {
    event.preventDefault();
    const response = await updateBook(IdValue.value);
    if (response.status === 200) {
      console.log('Book updated successfully!');
    } else {
      console.log('Error updating book:', response.status);
    }
  });

