<?php
$query_params = $_GET;
?>

<!DOCTYPE html>
<html>
<head>
    <title>GET Echo</title>
</head>
<body>
    <h1>GET Request Parameters:</h1>
    <ul>
        <?php foreach ($query_params as $key => $value) : ?>
            <li><?= htmlspecialchars($key) ?>: <?= htmlspecialchars($value) ?></li>
        <?php endforeach; ?>
    </ul>
</body>
</html>
