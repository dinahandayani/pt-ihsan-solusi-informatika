#!/bin/bash

# Set permission untuk Laravel
chmod -R 775 storage bootstrap/cache

# Clear semua cache Laravel
php artisan config:clear
php artisan cache:clear
php artisan route:clear
php artisan view:clear

# Install dependencies kalau belum
npm install

# Jalankan Laravel server (di port 8000, host 0.0.0.0)
php artisan serve --host=0.0.0.0 --port=8000 &

# Jalankan Vite server di host 0.0.0.0 supaya bisa diakses dari luar container
npm run dev -- --host 0.0.0.0
