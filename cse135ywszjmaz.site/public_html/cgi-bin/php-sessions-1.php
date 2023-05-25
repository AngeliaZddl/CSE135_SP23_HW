<?php
// Start PHP session
session_start();

// Check if the username is set in the session or in the request
if (isset($_SESSION['username'])) {
    $name = $_SESSION['username'];
} elseif (isset($_REQUEST['username'])) {
    $name = $_REQUEST['username'];
    $_SESSION['username'] = $name;
} else {
    $name = null;
}

$sessionId = session_id();

// Generate HTML output
?>
<!DOCTYPE html>
<html>
<head>
    <title>PHP Sessions</title>
</head>
<body>
    <h1>PHP Sessions Page 1</h1>
    <?php if ($name): ?>
        <p><b>Name:</b> <?= htmlspecialchars($name) ?></p>
    <?php else: ?>
        <p><b>Name:</b> You do not have a name set</p>
    <?php endif; ?>
    <br/><br/>
    <a href="php-sessions-2.php?PHPSESSID=<?= urlencode($sessionId) ?>">Session Page 2</a><br/>
    <a href="/php-cgiform.html">PHP CGI Form</a><br />
    <form style="margin-top:30px" action="php-destroy-session.php?PHPSESSID=<?= urlencode($sessionId) ?>" method="get">
        <button type="submit">Destroy Session</button>
    </form>
</body>
</html>