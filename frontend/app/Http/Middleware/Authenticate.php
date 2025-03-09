<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Session;
use Illuminate\Support\Facades\Http;

class Authenticate
{
    public function handle(Request $request, Closure $next)
    {
        if (!Session::has('token')) {
            return redirect('/login');
        }

        // Validate JWT token with the backend
        $token = Session::get('token');
        $backendUrl = env('BACKEND_HOST') . ":" . env('BACKEND_PORT') . '/api/auth/validate';

        $response = Http::withHeaders([
            'Authorization' => 'Bearer ' . $token,
        ])->get($backendUrl);
        
        if ($response->failed()) {
            Session::forget('token');
            return redirect('/login')->withErrors(['error' => 'Session expired. Please log in again.']);
        }

        return $next($request);
    }
}
