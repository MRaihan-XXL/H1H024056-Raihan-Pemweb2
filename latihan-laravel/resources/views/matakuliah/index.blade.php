@extends('layouts.app')

@section('judul', 'Daftar Mata Kuliah')

@section('konten')
<div class="card">
    <h1>Daftar Mata Kuliah</h1>

    <form method="GET" action="{{ route('matakuliah.index') }}" class="toolbar">
        <div class="field">
            <label class="label" for="q">Cari mata kuliah</label>
            <input type="search" id="q" name="q" value="{{ $kataKunci }}" placeholder="Masukkan kode atau nama mata kuliah">
        </div>
        <button type="submit" class="btn primary">Cari</button>
        <a href="{{ route('matakuliah.index') }}" class="btn">Reset</a>
    </form>

    <div class="table-wrap">
        <table>
            <thead>
                <tr>
                    <th>Kode</th>
                    <th>Nama</th>
                    <th>SKS</th>
                    <th>Aksi</th>
                </tr>
            </thead>
            <tbody>
                @forelse ($daftarMatakuliah as $matakuliah)
                    <tr>
                        <td>{{ $matakuliah->kode }}</td>
                        <td>{{ $matakuliah->nama }}</td>
                        <td><x-badge-sks :sks="$matakuliah->sks" /></td>
                        <td>
                            <a href="{{ route('matakuliah.show', $matakuliah->kode) }}" class="btn">Lihat Detail</a>
                        </td>
                    </tr>
                @empty
                    <tr>
                        <td colspan="4">Mata kuliah tidak ditemukan.</td>
                    </tr>
                @endforelse
            </tbody>
        </table>
    </div>
</div>
@endsection
