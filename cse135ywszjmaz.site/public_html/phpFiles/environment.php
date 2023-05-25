<?php
function get_request_headers() {
    $headers = [];
    foreach ($_SERVER as $key => $value) {
        if (substr($key, 0, 5) == 'HTTP_') {
            $header_key = str_replace(' ', '-', ucwords(str_replace('_', ' ', strtolower(substr($key, 5)))));
            $headers[$header_key] = $value;
        }
    }
    return $headers;
}

$request_headers = get_request_headers();
$server_variables = $_SERVER;
?>

<!DOCTYPE html>
<html>
<head>
    <title>Environment</title>
</head>
<body>
    <h1>Request Headers:</h1>
    <ul>
        <?php foreach ($request_headers as $key => $value) : ?>
            <li><?= htmlspecialchars($key) ?>: <?= htmlspecialchars($value) ?></li>
        <?php endforeach; ?>
    </ul>
    
    <h1>Server Variables:</h1>
    <ul>
        <?php foreach ($server_variables as $key => $value) : ?>
            <li><?= htmlspecialchars($key) ?>: <?= htmlspecialchars($value) ?></li>
        <?php endforeach; ?>
    </ul>
</body>
</html>
