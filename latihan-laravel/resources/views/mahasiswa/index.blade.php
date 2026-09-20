@extends('layouts.app')

@section('judul', 'Daftar Mahasiswa')

@section('konten')
<div class="card">
    <div style="display:flex; justify-content:space-between; align-items:center; gap:16px; margin-bottom:20px; flex-wrap:wrap;">
        <div>
            <p class="muted">Informatika</p>
            <h1>Daftar Mahasiswa</h1>
        </div>
        <a href="{{ route('mahasiswa.top-ipk') }}" class="btn primary">10 IPK Tertinggi</a>
    </div>

    <table>
        <thead>
            <tr>
                <th>NIM</th>
                <th>Nama</th>
                <th>Angkatan</th>
                <th>Prodi</th>
                <th>Aksi</th>
            </tr>
        </thead>
        <tbody>
            @forelse ($daftarMahasiswa as $mahasiswa)
                <tr>
                    <td>{{ $mahasiswa->nim }}</td>
                    <td>{{ $mahasiswa->nama }}</td>
                    <td>{{ $mahasiswa->angkatan }}</td>
                    <td>{{ $mahasiswa->program_studi }}</td>
                    <td>
                        <a href="{{ route('mahasiswa.show', $mahasiswa->nim) }}" class="btn">Detail</a>
                    </td>
                </tr>
            @empty
                <tr>
                    <td colspan="5">Data belum tersedia.</td>
                </tr>
            @endforelse
        </tbody>
    </table>
</div>
@endsection