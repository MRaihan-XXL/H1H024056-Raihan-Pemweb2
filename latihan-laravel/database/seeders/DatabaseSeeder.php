<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use App\Models\Matakuliah;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    public function run(): void
    {
        $this->call([
            MahasiswaSeeder::class,
            MatakuliahSeeder::class,
        ]);

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
