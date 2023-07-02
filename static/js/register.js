import{token,DecodedToken,userStatus} from "./userstatus.js";
userStatus(token);
if(token!=undefined){
    var decoded=DecodedToken(token);
    if(decoded.payload["usertype"]==="user"){
        window.location="/profile.html";    
    }
}
    
var submit=document.getElementById("submit");
submit.addEventListener("click",function(){
    var json={
        "Name": document.getElementById("name").value,
        "Email":document.getElementById("email").value,
        "Password":document.getElementById("password").value,
    };
    const headers = {
        "Content-Type": "application/json",
      };
    if(json.Name==="" || json.Email==="" || json.Password===""){
        alert("Empty Fields");
    }
    else{
        fetch("/register",{method:"POST",headers,body: JSON.stringify(json)})
        .then(response=>response.json())
        .then(data=>{
            console.log(data);
            if(data===true)
            {
                document.querySelector("main").innerHTML="<h2>Registered Successfully</h2>";
                setTimeout(function(){
                    window.location="/register.html";
                },5000)
            }
            else{
                document.querySelector("main").innerText="Error";
            }
        });
    }
    
})



// nameInput = document.getElementById("name");
// emailInput = document.getElementById("email");
// password = document.getElementById("password");
// submit = document.getElementById("submit");
/*const isName= str => /^[a-zA-Z .]*$/.test(str);
const isEmail= str=> /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/.test(str);
const isPass = str => /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$/.test(str);
nameInput.addEventListener("input",function (){
    nameInput.classList.remove("error");
    if(nameInput.value!="" && isName(nameInput.value)){
        nameInput.classList.remove("error");
        
    }
    else{
        nameInput.classList.add("error");
    }
});
password.addEventListener("input",function(){
    password.classList.remove("error");

    if(password.value!="" && isPass(password.value) && nameInput.value!="" && isName(nameInput.value) && emailInput.value!=""){
        password.classList.remove("error");
        submit.disabled= false;
        submit.classList.add("active-btn")
    }
    else{
        password.classList.add("error");
    }
});*/



    