@extends('layouts.app')

@section('judul', 'Detail Matakuliah')

@section('konten')
<h1 class="h3 mb-4">Detail Matakuliah</h1>

<div class="card">
    <div class="card-body">
        <dl class="row mb-0">
            <dt class="col-sm-3">Kode</dt>
            <dd class="col-sm-9">{{ $matakuliah['kode'] }}</dd>

            <dt class="col-sm-3">Nama</dt>
            <dd class="col-sm-9">{{ $matakuliah['nama'] }}</dd>

            <dt class="col-sm-3">SKS</dt>
            <dd class="col-sm-9"><x-badge-sks :sks="$matakuliah['sks']" /></dd>
        </dl>
    </div>
</div>

<a href="{{ route('matakuliah.index') }}" class="btn btn-secondary mt-3">Kembali</a>
@endsection
