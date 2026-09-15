<?php

namespace App\Http\Controllers;

use App\Models\Matakuliah;
use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $kataKunci = trim($request->query('q', ''));

        $daftarMatakuliah = Matakuliah::query()
            ->when($kataKunci !== '', function ($query) use ($kataKunci) {
                $query->where('kode', 'like', "%{$kataKunci}%")
                    ->orWhere('nama', 'like', "%{$kataKunci}%");
            })
            ->orderBy('kode', 'asc')
            ->get();

        return view('matakuliah.index', compact('daftarMatakuliah', 'kataKunci'));
    }

    public function show(string $kode)
    {
        $matakuliah = Matakuliah::where('kode', strtoupper($kode))->firstOrFail();

        return view('matakuliah.show', compact('matakuliah'));
    }
}
