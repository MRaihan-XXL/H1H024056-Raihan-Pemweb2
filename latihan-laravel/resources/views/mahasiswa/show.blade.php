@extends('layouts.app')

@section('judul', 'Detail Mahasiswa')

@section('konten')
<div class="card">
    <div style="display:flex; justify-content:space-between; align-items:center; gap:16px; margin-bottom:20px; flex-wrap:wrap;">
        <div>
            <p class="muted">Mahasiswa</p>
            <h1>Detail Mahasiswa</h1>
        </div>
        <a href="{{ route('mahasiswa.index') }}" class="btn">Kembali</a>
    </div>

    <div class="row">
        <div class="box">
            <p class="muted">Profil</p>
            <h2 style="margin-bottom:16px;">{{ $mahasiswa->nama }}</h2>
            <div style="display:grid; gap:12px;">
                <div><span class="label">NIM</span><span class="value">{{ $mahasiswa->nim }}</span></div>
                <div><span class="label">Program Studi</span><span class="value">{{ $mahasiswa->program_studi }}</span></div>
                <div><span class="label">Angkatan</span><span class="value">{{ $mahasiswa->angkatan }}</span></div>
                <div><span class="label">IPK</span><span class="value">{{ $ipk }}</span></div>
            </div>
        </div>

        <div class="box">
            <p class="muted">Mata Kuliah</p>
            <table>
                <thead>
                    <tr>
                        <th>Kode</th>
                        <th>Nama</th>
                        <th>SKS</th>
                        <th>Nilai</th>
                    </tr>
                </thead>
                <tbody>
                    @forelse ($mahasiswa->matakuliahs as $matakuliah)
                        <tr>
                            <td>{{ $matakuliah->kode }}</td>
                            <td>{{ $matakuliah->nama }}</td>
                            <td>{{ $matakuliah->sks }}</td>
                            <td>{{ $matakuliah->pivot->nilai }}</td>
                        </tr>
                    @empty
                        <tr>
                            <td colspan="4">Belum ada mata kuliah yang diambil.</td>
                        </tr>
                    @endforelse
                </tbody>
            </table>
        </div>
    </div>
</div>
@endsection