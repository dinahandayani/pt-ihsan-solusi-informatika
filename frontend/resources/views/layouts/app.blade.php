<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Task Management</title>
    @vite('resources/css/app.css') <!-- Menggunakan Vite untuk memuat TailwindCSS -->

    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css" rel="stylesheet">
</head>
<body class="bg-gray-100">

    <div class="w-full md:w-[40%] mx-auto mt-10 p-6 rounded-lg">
        <h1 class="text-2xl font-bold mb-6 text-center">Task Management</h1>
        @yield('content') <!-- Tempat untuk content halaman yang ditampilkan -->
    </div>

    @vite('resources/js/app.js') <!-- Menambahkan JS jika diperlukan -->
</body>
</html>
