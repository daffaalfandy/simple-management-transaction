<?php 
namespace App\Services;

use Illuminate\Support\Facades\App;
use Illuminate\Support\Facades\Http;

class BaseApiService 
{
    protected $baseUrl;
    protected $apiUrl;
    
    public function __construct()
    {
        $this->baseUrl = getenv("BACKEND_HOST") . ":" . getenv("BACKEND_PORT");
        $this->apiUrl = $this->baseUrl . "/api";
    }

    public function post($path, $data)
    {
        $response = Http::post($this->apiUrl . $path, $data);

        if ($response->failed()) {
            throw new \Exception($response->body());
        }

        return $response->json()['data'];
    }

    public function get($path)
    {
        $response = Http::get($this->apiUrl . $path);

        if ($response->failed()) {
            throw new \Exception($response->body());
        }

        return $response->json();
    }
}