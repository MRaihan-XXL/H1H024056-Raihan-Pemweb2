<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\MahasiswaResource;
use App\Models\Mahasiswa;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class MahasiswaController extends Controller
{
    public function index(Request $request)
    {
        $query = Mahasiswa::query();

        if ($request->filled('cari')) {
            $kataKunci = $request->query('cari');

            $query->where(function ($subQuery) use ($kataKunci) {
                $subQuery->where('nama', 'like', '%' . $kataKunci . '%')
                    ->orWhere('nim', 'like', '%' . $kataKunci . '%');
            });
        }

        if ($request->filled('angkatan')) {
            $query->where('angkatan', $request->integer('angkatan'));
        }

        if ($request->filled('program_studi')) {
            $query->where(
                'program_studi',
                $request->query('program_studi')
            );
        }

        $kolomDiizinkan = [
            'nim',
            'nama',
            'angkatan',
            'program_studi',
        ];

        $fields = $request->query('fields');

        if ($fields) {
            $fields = array_values(array_intersect(
                explode(',', $fields),
                $kolomDiizinkan
            ));

            if ($fields !== []) {
                $query->select($fields);
            }
        }

        $urut = $request->query('urut', 'nama');
        $arah = $request->query('arah', 'asc');

        if (in_array($urut, $kolomDiizinkan, true)) {
            $query->orderBy(
                $urut,
                $arah === 'desc' ? 'desc' : 'asc'
            );
        }

        $perHalaman = min(
            $request->integer('per_halaman', 10),
            100
        );

        return MahasiswaResource::collection(
            $query->paginate($perHalaman)
        );
    }

    public function berdasarkanProgramStudi(Request $request, string $programStudi)
    {
        $perHalaman = min($request->integer('per_halaman', 10), 100);

        return MahasiswaResource::collection(
            Mahasiswa::query()
                ->where('program_studi', $programStudi)
                ->orderBy('nama')
                ->paginate($perHalaman)
        );
    }

    public function store(Request $request): JsonResponse
    {
        $data = $request->validate([
            'nim' => ['required', 'string', 'max:20', 'unique:mahasiswas,nim'],
            'nama' => ['required', 'string', 'max:100'],
            'angkatan' => ['required', 'integer', 'min:2000', 'max:2100'],
            'program_studi' => ['required', 'string', 'max:100'],
        ]);

        $mahasiswa = Mahasiswa::create($data);

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mahasiswa berhasil dibuat',
            'data' => new MahasiswaResource($mahasiswa),
        ], 201);
    }

    public function show(Mahasiswa $mahasiswa): JsonResponse
    {
        return response()->json([
            'sukses' => true,
            'data' => new MahasiswaResource($mahasiswa),
        ]);
    }

    public function update(
        Request $request,
        Mahasiswa $mahasiswa
    ): JsonResponse {
        $data = $request->validate([
            'nim' => [
                'sometimes',
                'string',
                'max:20',
                'unique:mahasiswas,nim,' . $mahasiswa->id,
            ],
            'nama' => ['sometimes', 'string', 'max:100'],
            'angkatan' => ['sometimes', 'integer', 'min:2000', 'max:2100'],
            'program_studi' => ['sometimes', 'string', 'max:100'],
        ]);

        $mahasiswa->update($data);

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mahasiswa berhasil diperbarui',
            'data' => new MahasiswaResource($mahasiswa->fresh()),
        ]);
    }

    public function destroy(Mahasiswa $mahasiswa): JsonResponse
    {
        $mahasiswa->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mahasiswa berhasil dihapus',
        ]);
    }
}