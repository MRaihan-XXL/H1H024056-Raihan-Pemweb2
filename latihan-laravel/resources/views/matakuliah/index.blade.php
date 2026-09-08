@extends('layouts.app')

@section('judul', 'Daftar Matakuliah')

@section('konten')
<h1 class="h3 mb-4">Daftar Matakuliah</h1>

<form method="GET" action="{{ route('matakuliah.index') }}" class="row g-2 mb-4">
    <div class="col-md-8">
        <label for="q" class="visually-hidden">Cari matakuliah</label>
        <input type="search" id="q" name="q" value="{{ $kataKunci }}" class="form-control" placeholder="Cari kode atau nama matakuliah">
    </div>
    <div class="col-auto">
        <button type="submit" class="btn btn-primary">Cari</button>
    </div>
    <div class="col-auto">
        <a href="{{ route('matakuliah.index') }}" class="btn btn-outline-secondary">Reset</a>
    </div>
</form>

<div class="table-responsive">
    <table class="table table-bordered bg-white align-middle">
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
                <td>{{ $matakuliah['kode'] }}</td>
                <td>{{ $matakuliah['nama'] }}</td>
                <td><x-badge-sks :sks="$matakuliah['sks']" /></td>
                <td>
                    <a href="{{ route('matakuliah.show', $matakuliah['kode']) }}" class="btn btn-sm btn-primary">
                        Detail
                    </a>
                </td>
            </tr>
            @empty
            <tr>
                <td colspan="4">Matakuliah tidak ditemukan.</td>
            </tr>
            @endforelse
        </tbody>
    </table>
</div>
@endsection
