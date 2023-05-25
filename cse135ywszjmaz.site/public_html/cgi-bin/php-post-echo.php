<?php
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $post_data = $_POST;
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>POST Echo</title>
</head>
<body>
    <?php if (isset($post_data)) : ?>
        <h1>POST Request Data:</h1>
        <ul>
            <?php foreach ($post_data as $key => $value) : ?>
                <li><?= htmlspecialchars($key) ?>: <?= htmlspecialchars($value) ?></li>
            <?php endforeach; ?>
        </ul>
    <?php else : ?>
        <h1>Submit a POST request:</h1>
        <form method="POST">
            <label for="param1">Param1:</label>
            <input type="text" id="param1" name="param1">
            <br>
            <label for="param2">Param2:</label>
            <input type="text" id="param2" name="param2">
            <br>
            <button type="submit">Submit</button>
        </form>
    <?php endif; ?>
</body>
</html>
