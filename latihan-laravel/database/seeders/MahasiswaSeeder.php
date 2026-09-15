<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use Illuminate\Database\Seeder;

class MahasiswaSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $mahasiswas = [
            ['nim' => 'H1D004001', 'nama' => 'Rina Putri', 'angkatan' => 2024, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004002', 'nama' => 'Fajar Wijaya', 'angkatan' => 2024, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004003', 'nama' => 'Dewi Lestari', 'angkatan' => 2024, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004004', 'nama' => 'Aldi Saputra', 'angkatan' => 2023, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004005', 'nama' => 'Siti Aisyah', 'angkatan' => 2023, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004006', 'nama' => 'Budi Hartono', 'angkatan' => 2022, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004007', 'nama' => 'Maya Sari', 'angkatan' => 2024, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004008', 'nama' => 'Reza Pratama', 'angkatan' => 2024, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004009', 'nama' => 'Nadia Rahma', 'angkatan' => 2022, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004010', 'nama' => 'Ilham Kurnia', 'angkatan' => 2024, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004011', 'nama' => 'Ayu Maulida', 'angkatan' => 2023, 'program_studi' => 'Informatika'],
            ['nim' => 'H1D004012', 'nama' => 'Arif Rahman', 'angkatan' => 2023, 'program_studi' => 'Informatika'],
        ];

        foreach ($mahasiswas as $mahasiswa) {
            Mahasiswa::create($mahasiswa);
        }
    }
}
