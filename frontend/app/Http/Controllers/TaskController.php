<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class TaskController extends Controller
{
    public function index()
    {
        // Ambil data dari API
        $response_o = Http::get(env('BACKEND_URL', 'http://localhost:8080').'/tasks-ongoing');
        $response_c = Http::get(env('BACKEND_URL', 'http://localhost:8080').'/tasks-completed');

        // Cek jika responsenya sukses
        $tasks_o = $response_o->successful() && isset($response_o['data']) ? $response_o['data'] : [];
        $tasks_c = $response_c->successful() && isset($response_c['data']) ? $response_c['data'] : [];

        return view('todo', [
            'tasks_o' => $tasks_o,
            'tasks_c' => $tasks_c,
        ]);
    }

    public function store(Request $request)
    {
        // Validasi input
        $request->validate([
            'title' => 'required|string|max:255',
        ]);

        // Kirim data ke backend Go
        $response = Http::post(env('BACKEND_URL', 'http://localhost:8080').'/tasks', [
            'title' => $request->title,
        ]);

        // Cek apakah berhasil
        if ($response->successful() && $response['status'] === 'success') {
            return redirect()->back()->with('success', $response['message'] ?? 'Task berhasil ditambahkan');
        }

        // Jika gagal
        return redirect()->back()->with('error', $response['message'] ?? 'Gagal menambahkan task');
    }

    public function destroy($id)
    {
        // Kirim permintaan DELETE ke backend Go dengan ID di URL
        $response = Http::delete(env('BACKEND_URL', 'http://localhost:8080')."/tasks?id={$id}");

        // Cek apakah permintaan DELETE berhasil
        if ($response->successful() && $response['status'] === 'success') {
            return redirect()->route('tasks.index')->with('success', 'Task berhasil dihapus');
        }

        // Jika gagal
        return redirect()->route('tasks.index')->with('error', 'Gagal menghapus task');
    }

    public function signtocompleted($id)
    {
        // Kirim permintaan PUT ke backend Go dengan ID di URL sebagai query parameter
        $response = Http::put(env('BACKEND_URL', 'http://localhost:8080')."/sign-to-completed?id={$id}", [
            'completed' => true,  // Body raw untuk menandai task sebagai completed
        ]);

        // Cek apakah permintaan PUT berhasil
        if ($response->successful() && $response['status'] === 'success') {
            return redirect()->route('tasks.index')->with('success', 'Task berhasil ditandai sebagai selesai');
        }

        // Jika gagal
        return redirect()->route('tasks.index')->with('error', 'Gagal menandai task sebagai selesai');
    }

    public function edit($id)
    {
        // Mengambil data task berdasarkan ID dari API
        $response = Http::get(env('BACKEND_URL', 'http://localhost:8080')."/tasks-id?id={$id}");

        // Ambil data dari API
        $response_o = Http::get(env('BACKEND_URL', 'http://localhost:8080').'/tasks-ongoing');
        $response_c = Http::get(env('BACKEND_URL', 'http://localhost:8080').'/tasks-completed');

        // Cek jika responsenya sukses
        $tasks_o = $response_o->successful() && isset($response_o['data']) ? $response_o['data'] : [];
        $tasks_c = $response_c->successful() && isset($response_c['data']) ? $response_c['data'] : [];

        // Mengecek apakah task ditemukan
        if ($response->successful()) {
            $task = $response->json()['data']; // Mengambil data task dari response API
            return view('edit', [
                'task' => $task,
                'tasks_o' => $tasks_o,
                'tasks_c' => $tasks_c,
            ]);
        } else {
            return redirect()->route('tasks.index')->with('error', 'Task tidak ditemukan');
        }
    }

    public function update(Request $request, $id)
    {
        // Validasi input
        $request->validate([
            'title' => 'required|string|max:255',
        ]);

        // Ambil data "title" dari input
        $title = $request->input('title');

        // Kirim PUT request ke API eksternal
        $response = Http::put(env('BACKEND_URL', 'http://localhost:8080')."/tasks?id={$id}", [
            'title' => $title
        ]);

        // Cek apakah request API berhasil
        if ($response->successful()) {
            $data = $response->json();

            if ($data['status'] === 'success') {
                // Redirect dengan pesan sukses jika task berhasil diupdate
                return redirect()->route('tasks.edit', $id)->with('success', 'Task berhasil diupdate');
            } else {
                // Redirect dengan pesan error jika task gagal diupdate
                return redirect()->route('tasks.edit', $id)->with('error', 'Task gagal diupdate: ' . $data['message']);
            }
        } else {
            // Redirect dengan pesan error jika API gagal dijangkau
            return redirect()->route('tasks.edit', $id)->with('error', 'Terjadi kesalahan saat menghubungi API.');
        }
    }

}
