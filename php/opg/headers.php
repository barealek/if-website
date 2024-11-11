<?php
   $count = $_COOKIE['count'] ?? 0;
   setcookie('count', ++$count, time() + 60*60);
   echo "Du har været inde på denne side $count gang".($count == 1 ? "" : "e")." den sidste time";
?>
