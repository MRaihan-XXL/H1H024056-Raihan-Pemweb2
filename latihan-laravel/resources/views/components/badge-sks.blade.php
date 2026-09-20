@php
    $warna = $sks < 3 ? 'warning' : 'success';
@endphp

<span class="badge {{ $warna }}">{{ $sks }} SKS</span>
