<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use App\Models\Matakuliah;
use App\Models\User;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    public function run(): void
    {
        $this->call([
            MahasiswaSeeder::class,
            MatakuliahSeeder::class,
        ]);

        User::updateOrCreate(
            ['email' => 'admin@unsoed.ac.id'],
            [
                'name' => 'Administrator',
                'password' => 'rahasia123',
                'peran' => 'admin',
            ]
        );

        User::updateOrCreate(
            ['email' => 'mahasiswa@unsoed.ac.id'],
            [
                'name' => 'Pengguna Mahasiswa',
                'password' => 'rahasia123',
                'peran' => 'mahasiswa',
            ]
        );

        $mahasiswas = Mahasiswa::all();
        $matakuliahs = Matakuliah::all();

        foreach ($mahasiswas as $mahasiswa) {
            $nilaiKuliah = $matakuliahs->shuffle()->take(rand(3, 5));

            foreach ($nilaiKuliah as $matakuliah) {
                $mahasiswa->matakuliahs()->attach($matakuliah->id, [
                    'nilai' => rand(70, 100),
                ]);
            }
        }
    }
}
