<?php

namespace App\Http\Controllers;

use App\Services\AuthService;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Session;

class AuthController extends Controller
{
    protected $authService;

    public function __construct(AuthService $authService) 
    {
        $this->authService = $authService;
    }

    public function register(Request $request)
    {
        try {
            $user = $this->authService->register($request->all());
            Session::put('token', $user->token);
    
            return redirect('/login')->with('success', 'Registration successful! Please login.');
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage());
        }
    }

    public function login(Request $request)
    {
        try {
            $token = $this->authService->login($request->all());
            Session::put('token', $token['token']);
    
            return redirect('/');
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage());
        }
    }

    public function logout()
    {
        Session::forget('token');
        return redirect('/');
    }
}
