<?php

use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    $currentDateTime = now()->format('Y-m-d H:i:s');
    $hostName = gethostname();

    return response()->json([
        'message' => 'Hello Devops World!!!',
        'date_time' => $currentDateTime,
        'host_name' => $hostName,
    ]);
});
