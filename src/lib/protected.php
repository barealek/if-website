<?php

   session_start();
   if (!isset($_SESSION['user'])){
      echo "<strong>Fejl:</strong> du er ikke authoriseret";
      exit();
   }
?>
