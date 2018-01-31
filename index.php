<!DOCTYPE html>
<html>
  <head>
    <title>CodeRockIt</title>
    <!--Main Styles-->
    <style>
    body{
      margin:0px;
    }
    #content{
      position: absolute;
      height: 100%;
      width: 100%;
      background-color: black;
      color: white;
      z-index: 1;
    }
      #main-title{
        color:white;
        position: relative;
        margin: auto;
        top: 15%;
        width: 570px;
        text-align: center;
      }
      #main-title h{
        font-size: 50px;
      }
      #main-title-p1{
        margin-top: 40px;
        font-size: 30px;
      }
      #logo{
        position: absolute;
        top: -60px;
        left: 0px;
      }
      #get-email{
        width: 317px;
        margin: auto;
        position: relative;
        top: 30%;
      }
      #get-email-text{
        width: 620px;
        margin: auto;
        position: relative;
        top: 30%;
      }
      #get-email-p1{
        font-size: 25px;
        position: relative;
        top: 0px;
      }
      #get-email-p2{
        font-size: 12px;
        position: relative;
        margin-top: 20px;
      }
      #get-email-in{
        width: 250px;
      }
    </style>
    <script>
      function inputText(isFocused){
        var myinput = document.getElementById('get-email-in');
        if (isFocused && myinput.value =='Your Email Here'){
          myinput.value = '';
        }else if (myinput.value == ''){
          myinput.value = 'Your Email Here';
        }
      }
    </script>
    <?php
      if(isset($_GET['msg'])){
        echo "<script type='text/javascript'>alert('Please enter a valid e-mail.');</script>";
      }
    ?>
  </head>
  <body>
    <div id='content'>
      <div id='main-title'>
        <h>CodeRockit</h>
        <br>
        <p id='main-title-p1'>A better way to micro-manage shared code.</p>
        <img src='logo.svg' id='logo'/>
      </div>
      <div id='get-email-text'>
        <p id='get-email-p1'>Want to know when we're ready for you? Join our e-mail list!</p>
      </div>
      <div id='get-email'>
        <form action='get_emails.php' method='post'>
          <input name='email' id='get-email-in' onfocus="inputText(true)" onfocusout="inputText()" type="text" value="Your Email Here"/>
          <input type='submit'/>
        </form>
        <p id='get-email-p2'>We promise to only e-mail you once when CodeRockit is ready.</p>
      </div>
    </div>
  </body>
</html>
