<?php

namespace App\Services;

use Illuminate\Http\Client\Request;

class AuthService extends BaseApiService
{
    public function __construct()
    {
        parent::__construct();
    }

    public function login(array $params)
    {
        return $this->post('/auth/login', $params);
    }

    public function register(array $params)
    {
        return $this->post('/auth/register', $params);
    }
}
