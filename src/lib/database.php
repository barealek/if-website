<?php

/* mineDatabaser.php
* indeholder alle funktioner der benyttes til at få adgang
* til databasen, som angives som variabel i $databasenavn
*/

function DB_Connect($databasenavn)
{
	$servername="mysql-db";
	$username="myuser";
	$password="mypassword";
	$dbname=$databasenavn;

	global $forbindelse;
	$forbindelse = mysqli_connect($servername,$username,$password,$dbname);

	if(!$forbindelse)
		die("Ingen forbindelse: ".mysqli_connect_error());

	// Change character set to UTF-8
	mysqli_set_charset($forbindelse,"utf8");

}
?>
