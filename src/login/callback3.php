<?php
	include_once($_SERVER['DOCUMENT_ROOT'] . '/lib/database.php');
	DB_Connect("btghold");

	if(isset($_POST['user'],$_POST['pass'])){

		$UserName = mysqli_real_escape_string($forbindelse, $_REQUEST['user']);
		$PassWord = mysqli_real_escape_string($forbindelse, $_REQUEST['pass']);

		$PassWord=$PassWord;
		$sql= "SELECT userID, username FROM users WHERE username='$UserName' AND password='$PassWord'";

		// Debug stuff
		/*echo $UserName;
		echo $PassWord;
		echo $sql;*/

		$resultat=mysqli_query($forbindelse,$sql);

		if(mysqli_num_rows($resultat)>0){
			$post=mysqli_fetch_assoc($resultat);
			$Brugernr=$post['userID'];
			$BrugerNavn=$post['username'];

			session_start();
			$_SESSION['user']=$Brugernr;
         header("Location: /panel/");
			exit();
		}else{
         echo '<strong>Fejl</strong>: forkerte credentials';
         exit();
		}

	}

?>
