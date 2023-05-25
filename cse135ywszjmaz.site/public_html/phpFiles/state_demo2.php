<?php
session_start();

if (isset($_SESSION['name'])) {
    $name = $_SESSION['name'];
} else {
    header('Location: state_demo.php');
    exit;
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>State Demo - Page 2</title>
</head>
<body>
    <h1>Welcome, <?= htmlspecialchars($name) ?>!</h1>
    <p>This is the second page of the state demo.</p>
</body>
</html>