<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;
use Illuminate\Validation\Rules\Password;

class AuthController extends Controller
{
    public function register(Request $request): JsonResponse
    {
        $data = $request->validate([
            'name' => ['required', 'string', 'max:100'],
            'email' => ['required', 'email', 'max:100', 'unique:users,email'],
            'password' => ['required', 'confirmed', Password::min(8)->letters()->numbers()],
        ]);

        $data['password'] = Hash::make($data['password']);
        $data['peran'] = 'mahasiswa';

        $pengguna = User::create($data);
        $token = $pengguna->createToken('token-perangkat', ['mahasiswa:baca'])->plainTextToken;

        return response()->json([
            'sukses' => true,
            'pesan' => 'Registrasi berhasil',
            'data' => $this->dataPengguna($pengguna, $token),
        ], 201);
    }

    public function login(Request $request): JsonResponse
    {
        $data = $request->validate([
            'email' => ['required', 'email'],
            'password' => ['required', 'string'],
        ]);

        $pengguna = User::where('email', $data['email'])->first();

        if ($pengguna === null || ! Hash::check($data['password'], $pengguna->password)) {
            return response()->json([
                'sukses' => false,
                'pesan' => 'Email atau kata sandi tidak sesuai',
            ], 401);
        }

        $pengguna->terakhir_login = now();
        $pengguna->save();

        $kemampuan = $pengguna->peran === 'admin'
            ? ['mahasiswa:baca', 'mahasiswa:tulis']
            : ['mahasiswa:baca'];

        $token = $pengguna->createToken('token-perangkat', $kemampuan)->plainTextToken;

        return response()->json([
            'sukses' => true,
            'pesan' => 'Login berhasil',
            'data' => $this->dataPengguna($pengguna, $token),
        ]);
    }

    public function profil(Request $request): JsonResponse
    {
        $pengguna = $request->user();
        $token = $pengguna->currentAccessToken();

        return response()->json([
            'sukses' => true,
            'data' => [
                'id' => $pengguna->id,
                'nama' => $pengguna->name,
                'email' => $pengguna->email,
                'peran' => $pengguna->peran,
                'terakhir_login' => $pengguna->terakhir_login,
                'kemampuan' => $token?->abilities ?? [],
            ],
        ]);
    }

    public function ubahPassword(Request $request): JsonResponse
    {
        $data = $request->validate([
            'kata_sandi_lama' => ['required', 'string'],
            'kata_sandi_baru' => ['required', 'confirmed', Password::min(8)->letters()->numbers()],
        ]);

        $pengguna = $request->user();

        if (! Hash::check($data['kata_sandi_lama'], $pengguna->password)) {
            return response()->json([
                'sukses' => false,
                'pesan' => 'Kata sandi lama tidak sesuai.',
            ], 422);
        }

        $pengguna->update(['password' => Hash::make($data['kata_sandi_baru'])]);

        return response()->json([
            'sukses' => true,
            'pesan' => 'Kata sandi berhasil diubah.',
        ]);
    }

    public function logout(Request $request): JsonResponse
    {
        $request->user()->currentAccessToken()?->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Logout berhasil',
        ]);
    }

    private function dataPengguna(User $pengguna, ?string $token = null): array
    {
        return [
            'pengguna' => [
                'id' => $pengguna->id,
                'nama' => $pengguna->name,
                'email' => $pengguna->email,
                'peran' => $pengguna->peran,
                'terakhir_login' => $pengguna->terakhir_login,
            ],
            ...($token === null ? [] : ['token' => $token]),
        ];
    }
}
