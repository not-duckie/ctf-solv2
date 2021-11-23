<!doctype html>
<html lang="en">
  <head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css">

	<title>Secure Coding</title>
  </head>
  <body>
  <div class="container py-2">
	<div class="jumbotron">
	  <h1>Login</h1>
	  <form  method="POST" action="index.php">
		<div class="form-group">
		  <label>Username</label>
		  <input type="text" class="form-control" name="username" placeholder="Enter Username">
		</div>
		<div class="form-group">
		  <label>Password</label>
		  <input type="password" class="form-control" name="password" placeholder="Enter Password">
		</div>
		<div class="form-group">
		  <button class="btn btn-success" type="submit">Login</button>
		</div>
	  </form>
	</div>
  </div>
  <?php
include "connection.php";
session_start();

if (isset($_POST['username']) && isset($_POST['password']))
{
    $uname = $_POST['username'];
    $password = $_POST['password'];

    //$query=mysqli_query($connection,"SELECT * FROM users WHERE username ='$uname' AND password ='$password'") or die("Query Unsuccessfull:".mysqli_error($connection));
    
	
	//using prepare sql queries
	
	$stmt = $connection->prepare("SELECT * FROM users WHERE username = ? AND password = ?");
    $stmt->bind_param("ss", $uname, $password);

    /*
    if($query_run)
    {
    echo "Login Successfull";
    }
    */
    $stmt->execute();

    ($stmt_result = $stmt->get_result()) or trigger_error($stmt->error, E_USER_ERROR);

    $num_rows = $stmt_result->num_rows;

    while ($row = $stmt_result->fetch_assoc())
    {
        if ($num_rows > 0)
        {
            echo "Login Sucessfull";
            $_SESSION["id"] = $row['id'];
            header("Location: profile.php?id=" . $_SESSION["id"]);
        } else {
			echo "Login Unsucessfull";
		}
    }
}
?>

	<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js"></script>
  </body>
</html>
