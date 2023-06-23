import {token,userStatus} from './userstatus.js';

userStatus(token);
fetch(`/api/Getbooks/`) 
    .then(response=> response.json())
    .then(data=> {
        var booksCount= data['Books'].length;
    for (let i=0;i<=3;i++)
    {
            
            var bookDesc = `<div class="bookholder">
            <div class="BookCountHolder"><p class="BookCount">${i+1}</p></div>
            <a target="blank" class="bookimga" href="/book.html?BookC=${data['Books'][i]["BookCode"]}">
            <img class="bookimg" src="/img/books/${data['Books'][i]['ImgPath']}" alt="">
            </a>

        </div>`;
            
        document.getElementById("wrap-addedbooks").innerHTML += bookDesc;
    }
    })
