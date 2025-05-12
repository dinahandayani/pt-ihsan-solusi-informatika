#!/bin/sh

echo "Menjalankan migrasi database..."
go run cmd/migrate.go

echo "Menjalankan air (live reload)..."
air
