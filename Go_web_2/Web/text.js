var storage=window.sessionStorage
function formSubmit()//登录时跳转页面函数
{
    var obj=new Object();
    obj.name=document.getElementById("username").value;
    obj.password=document.getElementById("password").value;//获取username与password
    var id = document.getElementsByTagName("input");//获取id
    for(var i=0; i<id.length; i ++){//获取它的身份
        if(id[i].checked){
            obj.id=id[i].value
        }
    }
    var httpRequest =new XMLHttpRequest();
    httpRequest.open("POST","http://localhost:8080/order",true);//处理登录网址
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    httpRequest.send(JSON.stringify(obj));//发送信息name与password与id给后端
    httpRequest.onreadystatechange =()=>{
    if (httpRequest.readyState == 4 && httpRequest.status==200){//响应成功
        var data=JSON.parse(httpRequest.responseText);//获取后端发过来的数据
        if (data.statuscode==200){
            alert("登录成功，欢迎使用记事本!");
            //使用localstorage保存姓名
            if(!window.sessionStorage){
                //错误日志处理
                console.log("错误")
            }else{
                storage["uuid"]=data.information;//存储cookie，发送给浏览器
                //alert(data.information);
            }
            if(data.id=="users"){//看看登录成功后转为哪个界面
                window.location.href="user_sign_in_page.html";
            }else{
                window.location.href="ad_sign_in_page.html";
            }
        }else{
            alert("登录不成功,请再次检查账号密码");
            location.reload();//刷新
        }
    }
 }
}

function registermessage(){//注册普通用户跳转界面
    window.location.href="user_register.html";
}

function userformSubmit()//用户注册响应函数
{
    //判断两次密码是否相同
    var password1=document.getElementById("passworda").value;
    var password2=document.getElementById("passwordb").value;//从id获取内容
    if(password1==password2){//两次密码输入相同
        var obj =new Object;
        obj.password=password1;
        obj.name=document.getElementById("username").value;
       // obj.id="user"
        var httpRequest =new XMLHttpRequest();
        httpRequest.open("POST","http://localhost:8080/handleregist",true);//处理用户注册网址
        httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        httpRequest.send(JSON.stringify(obj));//发送信息name与password
        httpRequest.onreadystatechange =()=>{
        if (httpRequest.readyState == 4 && httpRequest.status==200){//响应成功
            var data=JSON.parse(httpRequest.responseText);//获取后端发过来的数据
            if(data.statuscode==200){//注册成功
                alert(data.information);
                window.location.href="mainpage.html";
            }else{
                alert("用户名重复");
                location.reload();//刷新
            }
        }
     }
    }else{//两次密码输入不相同，刷新页面，提示信息
        alert("密码两次不相同");
        location.reload();//刷新
    }
}

function addusernamemessage(){//添加用户需要保存的信息
    var obj=new Object();
    obj.information= prompt("请输入新的信息：", "请输入需要添加的信息");
    obj.uuid=storage["uuid"];//获取uuid
    if(obj.information==null)
    {
        alert("没有输入");
    }else{
        alert("已经输入");
        httpRequest.open("POST","http://localhost:8080/adduserinformation",true);//处理用户注册网址
        httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        httpRequest.send(JSON.stringify(obj));//发给后端  姓名 信息
        httpRequest.onreadystatechange =()=>{
            if (httpRequest.readyState == 4 && httpRequest.status==200){//响应成功
                var data=JSON.parse(httpRequest.responseText);//获取后端发过来的数据
                if(data.statuscode==200){
                    alert(data.information);
                    location.reload();//刷新页面
                }else{
                    alert(data.information)
                    location.reload();//刷新页面
                }
            }
        }
    }
}

function deleteusermessage(r)//用户删除记载信息
{
    //获取需要删除的地方
    tr = r.parentNode.parentNode;//前端获取表格里边内容的 得传一个this
    var obj =new Object();
    obj.id=tr.cells[0].innerText;//序号
    obj.uuid=tr.cells[1].innerText;//姓名
    obj.time=tr.cells[2].innerText;//时间
    obj.information=tr.cells[3].innerText;//信息
    //传给后端
    httpRequest.open("POST","http://localhost:8080/deleteuserinformation",true);//处理用户注册网址
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    httpRequest.send(JSON.stringify(obj));
    httpRequest.onreadystatechange =()=>{
        location.reload();//成功删除，刷新页面
    }
}

function revusermessage(r)//用户修改内容
{
    tr = r.parentNode.parentNode;
    var obj=new Object();
    var aaa=tr.cells[3].innerText;//信息
    alert(aaa);
    var person = prompt("请输入新的信息：", aaa);
    if(person==null)//没有输入任何东西
    {
        alert("没有输入哦");
    }else{
        if(person==aaa)
        {
            alert("修改内容相同");
        }else{
            obj.information=person;
            obj.id=tr.cells[0].innerText;
            obj.uuid=tr.cells[1].innerText;
            obj.time=tr.cells[2].innerText;
            httpRequest.open("POST","http://localhost:8080/reviseuserinformation",true);//处理用户注册网址
            httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
            httpRequest.send(JSON.stringify(obj));
            httpRequest.onreadystatechange =()=>{
                if (httpRequest.readyState == 4 && httpRequest.status==200){//响应成功
                    var data=JSON.parse(httpRequest.responseText);//获取后端发过来的数据
                    if(data.statuscode==200){
                        alert(data.information);
                        location.reload();//刷新页面
                    }else{
                        alert(data.information)
                        location.reload();//刷新页面
                    }
                }
            }
        }
    }
}

function findmessage()//根据id查询
{
    //获取id 发给后端 请求看看是否有然后打印
    var obj=new Object();
    obj.uuid=storage["uuid"];
    obj.id=document.getElementById("id").value;
    httpRequest.open("POST","http://localhost:8080/findinformation",true);
    httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");//Content-Type（内容类型），一般是指网页中存在的 Content-Type，Content-Type（内容类型），一般是指网页中存在的 Content-Type，用于定义网络文件的类型和网页的编码，决定浏览器将以什么形式、什么编码读取这个文件，                                                                               
    httpRequest.send(JSON.stringify(obj));//发送id和身份
        httpRequest.onreadystatechange =()=>{
            if (httpRequest.readyState == 4 && httpRequest.status==200){//响应成功
            var data=JSON.parse(httpRequest.responseText);//获取后端发过来的数据
            alert(data.information);  
        }
    }
}