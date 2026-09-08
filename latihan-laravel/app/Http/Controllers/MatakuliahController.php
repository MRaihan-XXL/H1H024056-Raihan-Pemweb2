<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    private function daftarMatakuliah(): array
    {
        return [
            ['kode' => 'IF101', 'nama' => 'Pemrograman Web II', 'sks' => 3],
            ['kode' => 'IF102', 'nama' => 'Struktur Data', 'sks' => 3],
            ['kode' => 'IF103', 'nama' => 'Basis Data', 'sks' => 3],
            ['kode' => 'IF104', 'nama' => 'Sistem Operasi', 'sks' => 2],
            ['kode' => 'IF105', 'nama' => 'Jaringan Komputer', 'sks' => 2],
        ];
    }

    public function index(Request $request)
    {
        $kataKunci = trim($request->query('q', ''));
        $daftarMatakuliah = collect($this->daftarMatakuliah())
            ->filter(function (array $matakuliah) use ($kataKunci) {
                if ($kataKunci === '') {
                    return true;
                }

                return str_contains(strtolower($matakuliah['kode']), strtolower($kataKunci))
                    || str_contains(strtolower($matakuliah['nama']), strtolower($kataKunci));
            })
            ->values();

        return view('matakuliah.index', compact('daftarMatakuliah', 'kataKunci'));
    }

    public function show(string $kode)
    {
        $matakuliah = collect($this->daftarMatakuliah())
            ->firstWhere('kode', strtoupper($kode));

        abort_if($matakuliah === null, 404);

        return view('matakuliah.show', compact('matakuliah'));
    }
}
