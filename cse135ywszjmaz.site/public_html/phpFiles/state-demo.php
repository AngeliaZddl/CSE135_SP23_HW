<?php
session_start();

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $_SESSION['name'] = $_POST['name'];
    header('Location: state_demo2.php');
    exit;
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>State Demo - Page 1</title>
</head>
<body>
    <h1>Enter your name:</h1>
    <form method="POST">
        <label for="name">Name:</label>
        <input type="text" id="name" name="name">
        <br>
        <button type="submit">Submit</button>
    </form>
</body>
</html>
