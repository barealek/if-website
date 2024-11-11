<?php

if (isset($_POST['name'])) {
    setcookie('name', $_POST['name'], time() + 3600, '/');
    echo "Dit navn er sat!";
} else {
   header('Location: navn.php');
}
?>
