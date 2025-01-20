<?php
   if(isset($_POST['user'], $_POST['pass'])){
      $user = array(
         "user" => "alek",
         "pass" => "sander"
      );

      $ussrname = $_POST['user'];
      $password = $_POST['pass'];

      if ($ussrname == $user['user'] && $password == $user['pass']) {
         session_start();
         $_SESSION['user'] = $ussrname;
         echo "<script>top.window.location = '/admin/'</script>";
         exit();
      } else {
         header("Location: /login/bad.php");
      }
   } else {
      echo 'post data not found';
   }
?>
