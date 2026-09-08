@php
    $warna = $sks < 3 ? 'bg-warning text-dark' : 'bg-success';
@endphp

<span class="badge {{ $warna }}">{{ $sks }} SKS</span>
