<html>
<body>

<?php
  function handleEmailInput(){
    $inputString = $_POST['email'];
    $count = strlen($inputString);
    $at = false;
    $dot = false;
    $blankSpace = false;
    function writeEmail(){
      $newemail = $_POST['email'];
      $emails = fopen( "emails.txt", 'a') or die ('Unable to open file.');
      fwrite($emails, $newemail . ',');
      fclose($emails);
      header("Location: index.php");
      exit;
    }
    for ($x = 1; $x < $count; $x++){
      switch ($inputString[$x]){
        case ' ':
          $blankSpace = true;
          $x = $count;
          break;
        case '@':
          $at = true;
          break;
        case '.':
          $dot = true;
          break;
      }
    }
    if ($at && $dot && $blankSpace != true && $count < 51){
      writeEmail();
    }else {
      header("Location: index.php?msg=set");
    }
  }
  handleEmailInput();
?>

</body>
</html>
