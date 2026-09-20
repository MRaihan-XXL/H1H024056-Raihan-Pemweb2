@extends('layouts.app')

@section('judul', 'Detail Mata Kuliah')

@section('konten')
<div class="card">
    <p class="muted">Mata Kuliah</p>
    <h1>Detail Mata Kuliah</h1>
    <div style="display:grid; gap:12px; max-width:520px;">
        <div><span class="label">Kode</span><span class="value">{{ $matakuliah->kode }}</span></div>
        <div><span class="label">Nama</span><span class="value">{{ $matakuliah->nama }}</span></div>
        <div><span class="label">SKS</span><span class="value"><x-badge-sks :sks="$matakuliah->sks" /></span></div>
        <div><span class="label">Semester</span><span class="value">{{ $matakuliah->semester }}</span></div>
    </div>
    <a href="{{ route('matakuliah.index') }}" class="btn" style="margin-top:20px;">Kembali</a>
</div>
@endsection
