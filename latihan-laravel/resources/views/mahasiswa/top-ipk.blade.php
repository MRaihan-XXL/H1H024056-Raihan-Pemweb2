@extends('layouts.app')

@section('judul', '10 IPK Tertinggi Informatika')

@section('konten')
<div class="card">
    <div style="display:flex; justify-content:space-between; align-items:center; gap:16px; margin-bottom:20px; flex-wrap:wrap;">
        <div>
            <p class="muted">Peringkat</p>
            <h1>10 IPK Tertinggi Informatika</h1>
        </div>
        <a href="{{ route('mahasiswa.index') }}" class="btn">Kembali</a>
    </div>

    <table>
        <thead>
            <tr>
                <th>#</th>
                <th>NIM</th>
                <th>Nama</th>
                <th>Prodi</th>
                <th>IPK</th>
            </tr>
        </thead>
        <tbody>
            @forelse ($mahasiswaTerbaik as $index => $mahasiswa)
                <tr>
                    <td>{{ $index + 1 }}</td>
                    <td>{{ $mahasiswa->nim }}</td>
                    <td>{{ $mahasiswa->nama }}</td>
                    <td>{{ $mahasiswa->program_studi }}</td>
                    <td>{{ $mahasiswa->ipk }}</td>
                </tr>
            @empty
                <tr>
                    <td colspan="5">Data IPK belum tersedia.</td>
                </tr>
            @endforelse
        </tbody>
    </table>
</div>
@endsection
