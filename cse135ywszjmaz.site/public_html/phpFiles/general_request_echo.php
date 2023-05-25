<?php
$request_method = $_SERVER['REQUEST_METHOD'];
$request_data = [];

if ($request_method === 'POST' || $request_method === 'PUT' || $request_method === 'DELETE') {
    parse_str(file_get_contents('php://input'), $request_data);
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>General Request Echo</title>
</head>
<body>
    <h1>Request Method: <?= htmlspecialchars($request_method) ?></h1>

    <?php if (!empty($request_data)) : ?>
        <h1>Request Payload:</h1>
        <ul>
            <?php foreach ($request_data as $key => $value) : ?>
                <li><?= htmlspecialchars($key) ?>: <?= htmlspecialchars($value) ?></li>
            <?php endforeach; ?>
        </ul>
    <?php endif; ?>
</body>
</html>
