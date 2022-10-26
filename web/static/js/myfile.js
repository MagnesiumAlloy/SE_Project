function user1(){
    $.ajax({
        //几个参数需要注意一下
        dataType: "json",
        url: "/User1" ,
        type: "POST", 
        data: $('#user1').serialize(),
        success: function (result) {
            console.log(result);//打印服务端返回的数据(调试用)
            if (result.status == 200) {
               location.href="/myfile";
            } else {
                alert("错误");
            }
        },
        error : function() {
            console.log("failed");
            alert("异常！");
        }
    });
}


function user4(){
    $.ajax({
        //几个参数需要注意一下
        dataType: "json",
        url: "/User4" ,
        type: "POST", 
        data: $('#user4').serialize(),
        success: function (result) {
            console.log(result);//打印服务端返回的数据(调试用)
            if (result.status == 200) {
               location.href="recycle";
            } else {
                alert("错误");
            }
        },
        error : function() {
            console.log("failed");
            alert("异常！");
        }
    });
}

function user5(){
    $.ajax({
        //几个参数需要注意一下
        dataType: "json",
        url: "/User5" ,
        type: "POST", 
        data: $('#user5').serialize(),
        success: function (result) {
            console.log(result);//打印服务端返回的数据(调试用)
            if (result.status == 200) {
               location.href="share";
            } else {
                alert("错误");
            }
        },
        error : function() {
            console.log("failed");
            alert("异常！");
        }
    });
}

function user6(){
    $.ajax({
        //几个参数需要注意一下
        dataType: "json",
        url: "/User6" ,
        type: "POST", 
        data: $('#user6').serialize(),
        success: function (result) {
            console.log(result);//打印服务端返回的数据(调试用)
            if (result.status == 200) {
               location.href="/secret";
            } else {
                alert("错误");
            }
        },
        error : function() {
            console.log("failed");
            alert("异常！");
        }
    });
}

function user2(){ //私密上传 函数要重新写
    console.log(123)
    const fle = document.getElementById('fileUpload')
    console.log('111',fle)
     fle.click()
    console.log(123)
    console.log($("#fileUpload")[0].files);
    upload()
}

function user3(){ //普通上传 函数要重新写
  console.log(123)
  const fle = document.getElementById('fileUpload')
  console.log('111',fle)
   fle.click()
  console.log(123)
  console.log($("#fileUpload")[0].files);
  upload()
}


function upload(){
  $.ajax({
        //几个参数需要注意一下
        dataType: "json",
        url: "/Upload" ,
        type: "POST", 
        data: $('#fileUpload').serialize(),
        success: function (result) {
            console.log(result);//打印服务端返回的数据(调试用)
            if (result.status == 200) {
               location.href="/zhuye";
            } else {
                alert("错误");
            }
        },
        error : function() {
            console.log("failed");
            alert("异常！");
        }
    });
    console.log(11);
  console.log($("#fileUpload")[0].files);
}


function getFile(value){
  // 获取文本框dom
 // var doc = document.getElementById('doc');
  // 获取上传控件dom
  console.log(this.location.pathname);
  var upload = document.getElementById('fileUpload');
  // 获取文件名
  console.log(upload);
  console.log($("#fileUpload")[0].files);
  // 获取文件路径
  var fileName = upload.files[0].name;
  var filePath = upload.value;
  // 将文件名载入文本框
 // doc.value = fileName;

  console.log(fileName);
  console.log(filePath);

}

  function my(){
       $.ajax({
           url: "/fileData",
           type:"get",
           dataType:"json",
           success:function (data) {
               console.log(data);
               myFunction(data.data);
           }
       })};

  function myFunction(arr) {
      var out = "";
      for(let i = 0; i < arr.length; i++) {
          out += "<tr  onclick="+"edit(this)"+">"
                   +"<td>"+arr[i].ID+"</td>"
                   +"<td>"+arr[i].Name+"</td>"
                   +"<td>"+arr[i].Name+"</td>"
                   +"<td>"+arr[i].Path+"</td>"
                   +"<td>"+arr[i].Type+"</td>"
                +"</tr>";
      }
      document.getElementById("content").innerHTML=out;
  }

  function myy(filepath){
       $.ajax({
           url: "/innerData",
           type:"get",
           dataType:"json",
           data: {
            filepath:filepath
           },
           success:function (res) {
               console.log(res);
               myFunction(res.data);
           }
       })};
    

    function edit(obj){
            let node=obj;
            let id=node.children[0].innerHTML;
            let filetype=node.children[4].innerHTML;
            let filepath=node.children[3].innerHTML;
            //alert(id,filepath,filetype);
            if(filetype == "t1"){ //文件夹类型,要你指定
              //更进一步
              myy(filepath);
            } else {
              let f=document.forms[0];
              f.action="/search?id="+id;
              f.submit();
            }
     }

