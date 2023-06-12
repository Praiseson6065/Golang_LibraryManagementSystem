nameInput = document.getElementById("name");
emailInput = document.getElementById("email");
password = document.getElementById("password");
submit = document.getElementById("submit");
const isName= str => /^[a-zA-Z .]*$/.test(str);
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
});



    