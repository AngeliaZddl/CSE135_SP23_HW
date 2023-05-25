<?php
header('Content-Type: application/json');

date_default_timezone_set('UTC');
$current_date_time = date('Y-m-d H:i:s');
$user_ip = $_SERVER['REMOTE_ADDR'];

$response = [
    'message' => 'Hello World!',
    'current_date_time' => $current_date_time,
    'user_ip' => $user_ip,
];

echo json_encode($response);
