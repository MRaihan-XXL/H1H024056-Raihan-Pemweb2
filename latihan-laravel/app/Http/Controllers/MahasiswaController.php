<?php
namespace App\Http\Controllers;

use App\Models\Mahasiswa;
use Illuminate\Http\Request;

class MahasiswaController extends Controller
{
    public function index()
    {
        $daftarMahasiswa = Mahasiswa::orderBy('angkatan', 'desc')
            ->orderBy('nama', 'asc')
            ->get();

        return view('mahasiswa.index', compact('daftarMahasiswa'));
    }

    public function show(string $nim)
    {
        $mahasiswa = Mahasiswa::with('matakuliahs')
            ->where('nim', $nim)
            ->firstOrFail();

        $ipk = $mahasiswa->matakuliahs->isEmpty()
            ? 0
            : round($mahasiswa->matakuliahs->avg(fn ($matakuliah) => $matakuliah->pivot->nilai) / 25, 2);

        return view('mahasiswa.show', compact('mahasiswa', 'ipk'));
    }

    public function cari(Request $request)
    {
        $kataKunci = $request->query('q', '');

        return response()->json([
            'kata_kunci' => $kataKunci,
            'metode' => $request->method(),
            'path' => $request->path(),
        ]);
    }

    public function topIpk()
    {
        $mahasiswaTerbaik = Mahasiswa::query()
            ->where('program_studi', 'Informatika')
            ->with(['matakuliahs' => function ($query) {
                $query->select('matakuliahs.id', 'matakuliahs.nama', 'matakuliahs.sks')
                    ->withPivot('nilai');
            }])
            ->get()
            ->map(function ($mahasiswa) {
                $mahasiswa->ipk = $mahasiswa->matakuliahs->isEmpty()
                    ? 0
                    : round($mahasiswa->matakuliahs->avg(fn ($matakuliah) => $matakuliah->pivot->nilai) / 25, 2);

                return $mahasiswa;
            })
            ->sortByDesc('ipk')
            ->take(10)
            ->values();

        return view('mahasiswa.top-ipk', compact('mahasiswaTerbaik'));
    }
}