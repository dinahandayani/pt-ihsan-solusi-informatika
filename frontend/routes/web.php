<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\TaskController;

Route::redirect('/', '/tasks');
Route::get('/tasks', [TaskController::class, 'index'])->name('tasks.index');
Route::post('/tasks', [TaskController::class, 'store'])->name('tasks.store');
Route::get('/tasks/{id}/edit', [TaskController::class, 'edit'])->name('tasks.edit');
Route::delete('/tasks/{id}', [TaskController::class, 'destroy'])->name('tasks.destroy');
Route::put('/sign-to-completed/{id}', [TaskController::class, 'signtocompleted'])->name('tasks.signtocompleted');
Route::put('/tasks/{id}', [TaskController::class, 'update'])->name('tasks.update');