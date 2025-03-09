<?php

namespace App\Services;

use Illuminate\Support\Facades\App;
use Illuminate\Support\Facades\Http;

class ApiService
{
    protected $baseUrl;

    public function __construct()
    {
        $this->baseUrl = env("BACKEND_HOST") . ":" . env("BACKEND_PORT");
    }

    public function getUsers()
    {
        $url = "{$this->baseUrl}/users";
        $response = Http::get($url);

        if ($response->successful()) {
            return $response->json() ?? [];
        }

        return [];
    }
}
