<?php
date_default_timezone_set('UTC');
$current_date_time = date('Y-m-d H:i:s');
$user_ip = $_SERVER['REMOTE_ADDR'];
?>

<!DOCTYPE html>
<html>
<head>
    <title>Hello HTML World</title>
</head>
<body>
    <h1>Hello World!</h1>
    <p>Current Date/Time: <?= $current_date_time ?></p>
    <p>Your IP Address: <?= $user_ip ?></p>
</body>
</html>
