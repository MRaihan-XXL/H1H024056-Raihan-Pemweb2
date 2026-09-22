<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Requests\StoreMatakuliahRequest;
use App\Http\Requests\UpdateMatakuliahRequest;
use App\Http\Resources\MatakuliahResource;
use App\Models\Matakuliah;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $query = Matakuliah::query();

        if ($request->filled('cari')) {
            $kataKunci = $request->query('cari');

            $query->where(function ($subQuery) use ($kataKunci) {
                $subQuery->where('kode', 'like', '%' . $kataKunci . '%')
                    ->orWhere('nama', 'like', '%' . $kataKunci . '%');
            });
        }

        if ($request->filled('semester')) {
            $query->where('semester', $request->integer('semester'));
        }

        $kolomDiizinkan = ['kode', 'nama', 'sks', 'semester'];
        $urut = $request->query('urut', 'kode');
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

        return MatakuliahResource::collection($query->paginate($perHalaman));
    }

    public function store(StoreMatakuliahRequest $request): JsonResponse
    {
        $matakuliah = Matakuliah::create($request->validated());

        return response()->json([
            'sukses' => true,
            'pesan' => 'Mata kuliah berhasil dibuat',
            'data' => new MatakuliahResource($matakuliah),
        ], 201);
    }

    public function show(Matakuliah $matakuliah): JsonResponse
    {
        return response()->json([
            'sukses' => true,
            'data' => new MatakuliahResource($matakuliah),
        ]);
    }

    public function update(
        UpdateMatakuliahRequest $request,
        Matakuliah $matakuliah
    ): JsonResponse
    {
        $matakuliah->update($request->validated());

        return response()->json([
            'sukses' => true,
            'pesan' => 'Mata kuliah berhasil diperbarui',
            'data' => new MatakuliahResource($matakuliah->fresh()),
        ]);
    }

    public function destroy(Matakuliah $matakuliah): JsonResponse
    {
        $matakuliah->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Mata kuliah berhasil dihapus',
        ]);
    }
}