export const token = document.cookie.split("=")[1];
export function UserPage(token){
    if(token===undefined){
        window.location="/home.html";
    }
}
export function DecodedToken(token){
    if(token!=undefined){
        const header = token.split('.')[0];
        const payload = token.split('.')[1];
        const signature = token.split('.')[2];
        
        var decoded = {
            header: JSON.parse(atob(header)),
            payload: JSON.parse(atob(payload)),
            signature: signature
        };
        return decoded;
    }
    
}
export function userStatus(t){
    
    if(t!=undefined) {
        document.getElementById("accstatus").innerText="Logout";
        document.getElementById("accstatus").setAttribute("href","/logout");
        document.getElementById("accsignup").hidden=true;
        if(DecodedToken(t).payload["usertype"]==="admin"){
            document.querySelector("#profile a").innerText="Admin"; 
            document.querySelector("#profile a").setAttribute("href","/admin.html");  
        }

      }
      else{
        document.getElementById("cart").hidden=true;
        document.getElementById("profile").hidden=true; 
        document.getElementById("accstatus").innerText="Login";
        document.getElementById("accstatus").setAttribute("href","/login.html");
      }
}

