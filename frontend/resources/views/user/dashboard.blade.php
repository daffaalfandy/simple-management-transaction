<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>User Dashboard</title>
    <link href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-gray-100 text-gray-900">
    <div class="container mx-auto p-6">
        <h1 class="text-3xl font-bold mb-4">User Dashboard</h1>
        
        <div class="bg-white p-4 rounded shadow-md">
            <p class="text-xl">Balance: $1000</p>
        </div>

        <div class="mt-6">
            <a href="topup.php" class="block bg-blue-500 text-white text-center py-2 rounded mb-2">Top Up Balance</a>
            <a href="transaction.php" class="block bg-green-500 text-white text-center py-2 rounded mb-2">Make a Transaction</a>
            <a href="history.php" class="block bg-gray-500 text-white text-center py-2 rounded">View Transaction History</a>
        </div>
    </div>
</body>
</html>
