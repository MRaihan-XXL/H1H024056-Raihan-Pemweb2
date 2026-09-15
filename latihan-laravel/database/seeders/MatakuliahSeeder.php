<?php

namespace Database\Seeders;

use App\Models\Matakuliah;
use Illuminate\Database\Seeder;

class MatakuliahSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        $matakuliahs = [
            ['kode' => 'TK101', 'nama' => 'Pemrograman Web', 'sks' => 3, 'semester' => 1],
            ['kode' => 'TK102', 'nama' => 'Dasar Elektronika', 'sks' => 2, 'semester' => 1],
            ['kode' => 'TK201', 'nama' => 'Basis Data', 'sks' => 3, 'semester' => 2],
            ['kode' => 'TK202', 'nama' => 'Jaringan Komputer', 'sks' => 3, 'semester' => 2],
            ['kode' => 'TK203', 'nama' => 'Struktur Data', 'sks' => 3, 'semester' => 2],
            ['kode' => 'TK301', 'nama' => 'Sistem Digital', 'sks' => 3, 'semester' => 3],
            ['kode' => 'TK302', 'nama' => 'Pemrograman Berorientasi Objek', 'sks' => 3, 'semester' => 3],
            ['kode' => 'TK401', 'nama' => 'Sistem Operasi', 'sks' => 3, 'semester' => 4],
            ['kode' => 'TK402', 'nama' => 'Kecerdasan Buatan', 'sks' => 3, 'semester' => 4],
            ['kode' => 'TK501', 'nama' => 'Tugas Akhir', 'sks' => 4, 'semester' => 5],
        ];

        foreach ($matakuliahs as $matakuliah) {
            Matakuliah::create($matakuliah);
        }
    }
}
