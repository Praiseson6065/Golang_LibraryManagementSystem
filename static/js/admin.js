const token = document.cookie.split("=")[1];
if(token!=undefined){
  const header = token.split('.')[0];
  const payload = token.split('.')[1];
  const signature = token.split('.')[2];
  
 var decoded = {
    header: JSON.parse(atob(header)),
    payload: JSON.parse(atob(payload)),
    signature: signature
  };
  
}

fetch("http://127.0.0.1:3000/admin")
    .then(response => response.json())
    .then(data => {
        if(data['error']==="UnAuthorized"){
            document.getElementById("main-wrap").innerText=data['error'];
        }
        else{
             
        }
    });