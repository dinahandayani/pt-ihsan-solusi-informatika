@extends('layouts.app')

@section('content')
    <div class="rounded-lg">

        @if(session('success'))
            <div class="bg-green-100 text-green-800 text-sm p-3 rounded mb-4">
                {{ session('success') }}
            </div>
        @endif

        @if(session('error'))
            <div class="bg-red-100 text-red-800 text-sm p-3 rounded mb-4">
                {{ session('error') }}
            </div>
        @endif

        <form action="{{ route('tasks.update', $task['id']) }}" method="POST" class="text-center">
            @csrf
            @method('PUT') <!-- Menggunakan method PUT untuk update -->

            <div class="mb-4 text-left">
                <label for="title" class="block text-sm font-medium text-gray-700 mb-1">Title</label>
                <input
                    type="text"
                    name="title"
                    id="title"
                    value="{{ $task['title'] }}"
                    class="border p-2 rounded w-full bg-white"
                >
            </div>

            <button type="submit" class="bg-orange-400 text-black px-4 py-1 rounded text-sm cursor-pointer hover:bg-orange-300 transition">
                Update Task
            </button>
            <a href="{{ route('tasks.index') }}" class="bg-red-400 text-black px-4 py-1 rounded text-sm cursor-pointer hover:bg-red-300 transition">
                Cancel
            </a>            
        </form>

        <h2 class="font-semibold text-gray-700 mb-2 mt-2">Ongoing Task</h2>
        @if(empty($tasks_o) || count($tasks_o) === 0)
            <div class="bg-yellow-100 text-yellow-800 text-sm p-3 rounded mb-4">
                Belum ada data Ongoing Task
            </div>
        @else
            @foreach($tasks_o as $task_o)
            <div class="bg-gray-200 p-4 rounded-lg shadow-md mb-4 mt-4">
                <div class="flex items-center justify-between">
                    <!-- Kiri: Tugas + Waktu -->
                    <div class="flex flex-col">
                        <div class="flex items-center">
                            <p class="text-sm font-medium text-gray-800 mr-3 text-[16px]">{{ $task_o['title'] }}</p>
                            <a href="#" class="text-gray-600 hover:text-gray-900 cursor-pointer" data-tooltip="Edit">
                                <i class="fas fa-pencil-alt text-sm"></i>
                            </a>
                        </div>
                        <div class="text-sm text-gray-600 mt-1 text-[12px]">
                            {{ \Carbon\Carbon::parse($task_o['created_at'])->translatedFormat('d F Y H:i') }}
                        </div>
                    </div>

                    <!-- Kanan: Tombol aksi -->
                    <div class="flex items-center">
                        <!-- Tombol Hapus -->
                        <a href="#" onclick="event.preventDefault(); document.getElementById('delete-task-{{ $task_o['id'] }}').submit();"
                           class="border border-gray-400 rounded-full w-6 h-6 items-center flex justify-center mr-2 hover:bg-red-200 cursor-pointer" 
                           data-tooltip="Delete">
                           <i class="fas fa-times text-[12px]"></i>
                        </a>
                        <!-- Tombol Completed -->
                        <a href="#" onclick="event.preventDefault(); document.getElementById('complete-task-{{ $task_o['id'] }}').submit();"
                            class="bg-white border border-gray-400 rounded-full w-6 h-6 items-center flex justify-center mr-2 hover:bg-green-200 cursor-pointer" 
                            data-tooltip="Completed">
                            <!-- Ganti dengan ikon sesuai keinginan -->
                        </a>
                    </div>
                </div>
            </div>

            <!-- Form untuk hapus task -->
            <form id="delete-task-{{ $task_o['id'] }}" action="{{ route('tasks.destroy', $task_o['id']) }}" method="POST" class="hidden">
                @csrf
                @method('DELETE')
            </form>
            <!-- Form untuk tanda task selesai -->
            <form id="complete-task-{{ $task_o['id'] }}" action="{{ route('tasks.signtocompleted', $task_o['id']) }}" method="POST" class="hidden">
                @csrf
                @method('PUT')
            </form>
            @endforeach
        @endif

        <h2 class="font-semibold text-gray-700 mb-2 mt-2">Completed Task</h2>
        @if(empty($tasks_c) || count($tasks_c) === 0)
            <div class="bg-yellow-100 text-yellow-800 text-sm p-3 rounded mb-4">
                Belum ada data Completed Task
            </div>
        @else
            @foreach($tasks_c as $task_c)
            <div class="bg-gray-200 p-4 rounded-lg shadow-md mb-4 mt-4">
                <div class="flex items-center justify-between">
                    <!-- Kiri: Tugas + Waktu -->
                    <div class="flex flex-col">
                        <div class="flex items-center">
                            <p class="text-sm font-medium text-gray-800 mr-3 text-[16px] line-through">{{ $task_c['title'] }}</p>
                            <a href="#" class="text-gray-600 hover:text-gray-900 cursor-pointer" data-tooltip="Edit">
                                <i class="fas fa-pencil-alt text-sm"></i>
                            </a>
                        </div>
                        <div class="text-sm text-gray-600 mt-1 text-[12px]">
                            {{ \Carbon\Carbon::parse($task_c['created_at'])->translatedFormat('d F Y H:i') }}
                        </div>
                    </div>

                    <!-- Kanan: Tombol aksi -->
                    <div class="flex items-center">
                        <button class="border border-gray-400 rounded-full w-6 h-6 items-center flex justify-center mr-2 hover:bg-red-200 cursor-pointer" data-tooltip="Delete">
                            <i class="fas fa-times text-[12px]"></i>
                        </button>
                        <button class="border border-gray-400 rounded-full w-6 h-6 items-center flex justify-center hover:bg-green-200 cursor-pointer" data-tooltip="Completed">
                            <i class="fas fa-check text-[12px]"></i>
                        </button>
                    </div>
                </div>
            </div>
            @endforeach
        @endif

    </div>
@endsection
