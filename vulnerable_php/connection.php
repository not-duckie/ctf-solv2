<?php
    $dbuser = "root";
    $dbpass = "password";
    $dbname = "secure_coding";
    $host = "127.0.0.1";
    error_reporting(0);
    $connection = mysqli_connect($host, $dbuser, $dbpass, $dbname) or die("Connection Failed: ".mysqli_connect_errno());
?>
