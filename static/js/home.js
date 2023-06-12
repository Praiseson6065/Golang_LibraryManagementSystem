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
  fetch('http://127.0.0.1:3000/home.html')
      .then(response=>response.text())
      .then(data=>{
        document.body.innerHTML=data;
          if(token!=undefined) {
              
              document.getElementById("accstatus").innerText="Logout";
              document.getElementById("accstatus").setAttribute("href","/logout");    
              document.getElementById("accsignup").hidden=true;
              
          }
          else{
              document.getElementById("profile").hidden=true;
      }
  });