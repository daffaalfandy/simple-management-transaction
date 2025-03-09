<?php

use App\Http\Controllers\AuthController;
use App\Http\Middleware\Authenticate;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});

Route::post('/register', [AuthController::class, 'register']);
Route::get('/register', function () {
    return view('auth.register');
});


Route::post('/login', [AuthController::class, 'login']);
Route::get('/login', function () {
    return view('auth.login');
});

Route::get('/logout', [AuthController::class, 'logout'])->middleware(Authenticate::class);

Route::get('/dashboard', function () {
    return view('user.dashboard');
})->middleware(Authenticate::class);